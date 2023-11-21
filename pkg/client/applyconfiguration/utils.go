/*
Copyright The Kubernetes Authors.

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

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfiguration

import (
	v2 "Persona_CRD/pkg/apis/compute/v2"
	computev2 "Persona_CRD/pkg/client/applyconfiguration/compute/v2"

	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=compute.genesis.msys.com, Version=v2
	case v2.SchemeGroupVersion.WithKind("Persona"):
		return &computev2.PersonaApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("PersonaSpec"):
		return &computev2.PersonaSpecApplyConfiguration{}
	case v2.SchemeGroupVersion.WithKind("PersonaStatus"):
		return &computev2.PersonaStatusApplyConfiguration{}

	}
	return nil
}