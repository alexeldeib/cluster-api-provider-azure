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

package scalesetvms

import (
	"context"
	"errors"
)

// Spec contains properties to create a managed cluster.
type Spec struct {
	Name          string
	ResourceGroup string
}

// ListInstances returns a list of provider IDs for the given VM scale set.
func (s *Service) ListInstances(ctx context.Context, spec interface{}) ([]string, error) {
	scaleSetVMsSpec, ok := spec.(*Spec)
	if !ok {
		return nil, errors.New("expected scale set vms specification")
	}
	return s.Client.ListInstances(ctx, scaleSetVMsSpec.ResourceGroup, scaleSetVMsSpec.Name)
}
