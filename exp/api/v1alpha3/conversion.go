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

package v1alpha3

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	v1alpha3 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha3"
	v1alpha4 "sigs.k8s.io/cluster-api-provider-azure/api/v1alpha4"
	clusterapiapiv1alpha3 "sigs.k8s.io/cluster-api/api/v1alpha3"
	clusterapiapiv1alpha4 "sigs.k8s.io/cluster-api/api/v1alpha4"
)

// Convert_v1alpha3_OSDisk_To_v1alpha4_OSDisk is a conversion function.
func Convert_v1alpha3_OSDisk_To_v1alpha4_OSDisk(in *v1alpha3.OSDisk, out *v1alpha4.OSDisk, s conversion.Scope) error {
	return v1alpha3.Convert_v1alpha3_OSDisk_To_v1alpha4_OSDisk(in, out, s)
}

// Convert_v1alpha4_OSDisk_To_v1alpha3_OSDisk is a conversion function.
func Convert_v1alpha4_OSDisk_To_v1alpha3_OSDisk(in *v1alpha4.OSDisk, out *v1alpha3.OSDisk, s conversion.Scope) error {
	return v1alpha3.Convert_v1alpha4_OSDisk_To_v1alpha3_OSDisk(in, out, s)
}

// Convert_v1alpha3_Image_To_v1alpha4_Image is a conversion function.
func Convert_v1alpha3_Image_To_v1alpha4_Image(in *v1alpha3.Image, out *v1alpha4.Image, s conversion.Scope) error {
	return v1alpha3.Convert_v1alpha3_Image_To_v1alpha4_Image(in, out, s)
}

// Convert_v1alpha4_Image_To_v1alpha3_Image is a conversion function.
func Convert_v1alpha4_Image_To_v1alpha3_Image(in *v1alpha4.Image, out *v1alpha3.Image, s conversion.Scope) error {
	return v1alpha3.Convert_v1alpha4_Image_To_v1alpha3_Image(in, out, s)
}

// Convert_v1alpha3_APIEndpoint_To_v1alpha4_APIEndpoint is an autogenerated conversion function.
func Convert_v1alpha3_APIEndpoint_To_v1alpha4_APIEndpoint(in *clusterapiapiv1alpha3.APIEndpoint, out *clusterapiapiv1alpha4.APIEndpoint, s conversion.Scope) error {
	return clusterapiapiv1alpha3.Convert_v1alpha3_APIEndpoint_To_v1alpha4_APIEndpoint(in, out, s)
}

// Convert_v1alpha4_APIEndpoint_To_v1alpha3_APIEndpoint is an autogenerated conversion function.
func Convert_v1alpha4_APIEndpoint_To_v1alpha3_APIEndpoint(in *clusterapiapiv1alpha4.APIEndpoint, out *clusterapiapiv1alpha3.APIEndpoint, s conversion.Scope) error {
	return clusterapiapiv1alpha3.Convert_v1alpha4_APIEndpoint_To_v1alpha3_APIEndpoint(in, out, s)
}