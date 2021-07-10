// +build e2e

/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package e2e

import (
	"context"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	. "github.com/onsi/gomega"
	"github.com/pkg/errors"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/types"
	e2e_pod "sigs.k8s.io/cluster-api-provider-azure/test/e2e/kubernetes/pod"
	"sigs.k8s.io/cluster-api/test/framework"
	"sigs.k8s.io/controller-runtime/pkg/client"
	kinderrors "sigs.k8s.io/kind/pkg/errors"
	"sigs.k8s.io/yaml"
)

// AzureTimeSyncSpecInput is the input for AzureTimeSyncSpec.
type AzureTimeSyncSpecInput struct {
	BootstrapClusterProxy framework.ClusterProxy
	Namespace             *corev1.Namespace
	ClusterName           string
}

// AzureTimeSyncSpec implements a test that verifies time synchronization is healthy for
// the nodes in a cluster.
func AzureTimeSyncSpec(ctx context.Context, inputGetter func() AzureTimeSyncSpecInput) {
	var (
		specName = "azure-timesync"
		input    AzureTimeSyncSpecInput
		thirty   = 30 * time.Second
	)

	input = inputGetter()
	Expect(input.BootstrapClusterProxy).NotTo(BeNil(), "Invalid argument. input.BootstrapClusterProxy can't be nil when calling %s spec", specName)
	namespace, clusterName := input.Namespace.Name, input.ClusterName
	Eventually(func() error {
		sshInfo, err := getClusterSSHInfo(ctx, input.BootstrapClusterProxy, namespace, clusterName)
		if err != nil {
			return err
		}

		if len(sshInfo) <= 0 {
			return errors.New("sshInfo did not contain any machines")
		}

		var testFuncs []func() error
		for _, s := range sshInfo {
			Byf("checking that time synchronization is healthy on %s", s.Hostname)

			execToStringFn := func(expected, command string, args ...string) func() error {
				// don't assert in this test func, just return errors
				return func() error {
					f := &strings.Builder{}
					if err := execOnHost(s.Endpoint, s.Hostname, s.Port, f, command, args...); err != nil {
						return err
					}
					if !strings.Contains(f.String(), expected) {
						return fmt.Errorf("expected \"%s\" in command output:\n%s", expected, f.String())
					}
					return nil
				}
			}

			testFuncs = append(testFuncs,
				execToStringFn(
					"✓ chronyd is active",
					"systemctl", "is-active", "chronyd", "&&",
					"echo", "✓ chronyd is active",
				),
				execToStringFn(
					"Reference ID",
					"chronyc", "tracking",
				),
			)
		}

		return kinderrors.AggregateConcurrent(testFuncs)
	}, thirty, thirty).Should(Succeed())
}

const (
	nsenterWorkloadFile = "workloads/nsenter/daemonset.yaml"
)

// AzureDaemonsetTimeSyncSpec implements a test that verifies time synchronization is healthy for
// the nodes in a cluster.
func AzureDaemonsetTimeSyncSpec(ctx context.Context, inputGetter func() AzureTimeSyncSpecInput) {
	var (
		specName = "azure-timesync"
		input    AzureTimeSyncSpecInput
		ninety   = 90 * time.Second
		five     = 5 * time.Second
	)

	input = inputGetter()
	Expect(input.BootstrapClusterProxy).NotTo(BeNil(), "Invalid argument. input.BootstrapClusterProxy can't be nil when calling %s spec", specName)
	namespace, clusterName := input.Namespace.Name, input.ClusterName
	workloadCluster := input.BootstrapClusterProxy.GetWorkloadCluster(ctx, namespace, clusterName)
	kubeclient := workloadCluster.GetClient()
	clientset := workloadCluster.GetClientSet()
	config := workloadCluster.GetRESTConfig()
	var nsenterDs unstructured.Unstructured

	yamlData, err := ioutil.ReadFile(nsenterWorkloadFile)
	if err != nil {
		Logf("failed daemonset time sync: %v", err)
	}

	jsonData, err := yaml.YAMLToJSON(yamlData)
	if err != nil {
		Logf("failed to convert nsenter yaml to json: %v", err)
	}

	if err := nsenterDs.UnmarshalJSON(jsonData); err != nil {
		Logf("failed daemonset time synx: %v", err)
	}

	nsenterDs.SetNamespace("default")

	if err := kubeclient.Create(ctx, &nsenterDs); err != nil {
		Logf("failed to create daemonset for time sync check: %v", err)
		return
	}

	matchingLabels := client.MatchingLabels(map[string]string{
		"app": "nsenter",
	})

	Eventually(func() error {
		var nodes corev1.PodList
		if err := kubeclient.List(ctx, &nodes); err != nil {
			Logf("failed to list nodes for daemonset timesync check: %v", err)
			return err
		}

		if len(nodes.Items) < 1 {
			msg := "expected to find >= 1 node for timesync check"
			Logf(msg)
			return fmt.Errorf(msg)
		}
		desired := len(nodes.Items)

		var ds appsv1.DaemonSet
		if err := kubeclient.Get(ctx, types.NamespacedName{"default", "nsenter"}, &ds); err != nil {
			Logf("failed to get nsenter ds: %s", err)
			return err
		}
		ready := ds.Status.NumberReady
		available := ds.Status.NumberAvailable
		allReadyAndAvailable := desired == ready && desired == available
		generationOk := ds.Metadata.Generation == ds.Status.ObservedGeneration

		Logf("want %d instances, found %d ready and %d available. generation: %s, observedGeneration: %s", desired, ready, available, ds.Metadata.Generation, ds.Status.ObservedGeneration)
		return allReadyAndAvailable && generationOk
	}).Should(Succeed())

	var podList corev1.PodList
	if err := kubeclient.List(ctx, &podList, matchingLabels); err != nil {
		Logf("failed to list pods for daemonset timesync check: %v", err)
		return
	}

	Logf("mapping pods to hostnames")
	podMap := map[string]corev1.Pod{}
	for _, pod := range podList.Items {
		podMap[pod.Spec.NodeName] = pod
	}

	for k, v := range podMap {
		Logf("found host %s with pod %s", k, v.Name)
	}

	Eventually(func() error {
		execInfo, err := getClusterSSHInfo(ctx, input.BootstrapClusterProxy, namespace, clusterName)
		if err != nil {
			return err
		}

		if len(execInfo) <= 0 {
			return errors.New("execInfo did not contain any machines")
		}

		// 	var testFuncs []func() error
		for _, s := range execInfo {
			Byf("checking that time synchronization is healthy on %s", s.Hostname)

			pod, exists := podMap[s.Hostname]
			if !exists {
				Logf("failed to find pod matching host %s", s.Hostname)
				return err
			}

			cmd := []string{"nsenter", "-t", "1", "-a", "--", "bash", "-c", "systemctl is-active chronyd && echo chronyd is active"}
			// command := []string{"systemctl", "is-active", "chronyd", "&&", "echo", "✓ chronyd is active"}
			stdout, stderr, err := e2e_pod.ExecWithOutput(clientset, config, pod, cmd)
			if err != nil {
				Logf("failed to nsenter host %s, error: '%s'", s.Hostname, err)
				Logf("stderr: %s", stderr.String())
				return err
			}

			Logf("stdout: '%s'", stdout.String())
			Logf("stderr: %s", stderr.String())

			if !strings.Contains(stdout.String(), "chronyd is active") {
				return fmt.Errorf("expected \"%s\" in command output:\n%s", "chronyd is active", stdout.String())
			}
			// 		execToStringFn := func(expected, command string, args ...string) func() error {
			// 			// don't assert in this test func, just return errors
			// 			return func() error {
			// 				f := &strings.Builder{}
			// 				if err := execOnHost(s.Endpoint, s.Hostname, s.Port, f, command, args...); err != nil {
			// 					return err
			// 				}
			// 				if !strings.Contains(f.String(), expected) {
			// 					return fmt.Errorf("expected \"%s\" in command output:\n%s", expected, f.String())
			// 				}
			// 				return nil
			// 			}
			// 		}

			// 		testFuncs = append(testFuncs,
			// 			execToStringFn(
			// 				"✓ chronyd is active",
			// 				"systemctl", "is-active", "chronyd", "&&",
			// 				"echo", "✓ chronyd is active",
			// 			),
			// 			execToStringFn(
			// 				"Reference ID",
			// 				"chronyc", "tracking",
			// 			),
			// 		)
		}

		return nil
	}, ninety, five).Should(Succeed())
}
