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

package scalesets

import (
	"context"
	"errors"
)

// Spec contains properties to identify VMSS in a resource group.
type Spec struct {
	ResourceGroup string
}

// Listreturns a list of provider IDs for the given VM scale set.
func (s *Service) List(ctx context.Context, spec interface{}) ([]interface{}, error) {
	scaleSetsSpec, ok := spec.(*Spec)
	if !ok {
		return nil, errors.New("expected scale set specification")
	}
	return s.Client.List(ctx, scaleSetsSpec.ResourceGroup)
}
