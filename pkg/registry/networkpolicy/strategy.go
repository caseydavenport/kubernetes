/*
Copyright 2016 The Kubernetes Authors All rights reserved.

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

package networkpolicy

import (
	"fmt"
	"reflect"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/apis/extensions/validation"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/validation/field"
)

// npStrategy implements verification logic for NetworkPolicys.
type npStrategy struct {
	runtime.ObjectTyper
	api.NameGenerator
}

// Strategy is the default logic that applies when creating and updating NetworkPolicy objects.
var Strategy = npStrategy{api.Scheme, api.SimpleNameGenerator}

// NamespaceScoped returns true because all NetworkPolicys need to be within a namespace.
func (npStrategy) NamespaceScoped() bool {
	return true
}

// PrepareForCreate clears the status of a NetworkPolicy before creation.
func (npStrategy) PrepareForCreate(obj runtime.Object) {
	np := obj.(*extensions.NetworkPolicy)

	np.Generation = 1
}

// PrepareForUpdate clears fields that are not allowed to be set by end users on update.
func (npStrategy) PrepareForUpdate(obj, old runtime.Object) {
	newNP := obj.(*extensions.NetworkPolicy)
	oldNP := old.(*extensions.NetworkPolicy)

	// Any changes to the spec increment the generation number, any changes to the
	// status should reflect the generation number of the corresponding object. We push
	// the burden of managing the status onto the clients because we can't (in general)
	// know here what version of spec the writer of the status has seen. It may seem like
	// we can at first -- since obj contains spec -- but in the future we will probably make
	// status its own object, and even if we don't, writes may be the result of a
	// read-update-write loop, so the contents of spec may not actually be the spec that
	// the NetworkPolicy has *seen*.
	if !reflect.DeepEqual(oldNP.Spec, newNP.Spec) {
		newNP.Generation = oldNP.Generation + 1
	}
}

// Validate validates a new NetworkPolicy.
func (npStrategy) Validate(ctx api.Context, obj runtime.Object) field.ErrorList {
	np := obj.(*extensions.NetworkPolicy)
	return validation.ValidateNetworkPolicy(np)
}

// Canonicalize normalizes the object after validation.
func (npStrategy) Canonicalize(obj runtime.Object) {
}

// AllowCreateOnUpdate is false for NetworkPolicys; this means a POST is
// needed to create one.
func (npStrategy) AllowCreateOnUpdate() bool {
	return false
}

// ValidateUpdate is the default update validation for an end user.
func (npStrategy) ValidateUpdate(ctx api.Context, obj, old runtime.Object) field.ErrorList {
	validationErrorList := validation.ValidateNetworkPolicy(obj.(*extensions.NetworkPolicy))
	updateErrorList := validation.ValidateNetworkPolicyUpdate(obj.(*extensions.NetworkPolicy), old.(*extensions.NetworkPolicy))
	return append(validationErrorList, updateErrorList...)
}

func (npStrategy) AllowUnconditionalUpdate() bool {
	return true
}

// NetworkPolicyToSelectableFields returns a field set that represents the object.
func NetworkPolicyToSelectableFields(np *extensions.NetworkPolicy) fields.Set {
	objectMetaFieldsSet := generic.ObjectMetaFieldsSet(np.ObjectMeta, true)
	npSpecificFieldsSet := fields.Set{}
	return generic.MergeFieldsSets(objectMetaFieldsSet, npSpecificFieldsSet)
}

// MatchNetworkPolicy is the filter used by the generic etcd backend to route
// watch events from etcd to clients of the apiserver only interested in specific
// labels/fields.
func MatchNetworkPolicy(label labels.Selector, field fields.Selector) generic.Matcher {
	return &generic.SelectionPredicate{
		Label: label,
		Field: field,
		GetAttrs: func(obj runtime.Object) (labels.Set, fields.Set, error) {
			np, ok := obj.(*extensions.NetworkPolicy)
			if !ok {
				return nil, nil, fmt.Errorf("Given object is not a NetworkPolicy.")
			}
			return labels.Set(np.ObjectMeta.Labels), NetworkPolicyToSelectableFields(np), nil
		},
	}
}
