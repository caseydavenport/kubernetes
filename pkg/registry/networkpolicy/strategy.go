/*
Copyright 2014 The Kubernetes Authors All rights reserved.

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
	"k8s.io/kubernetes/pkg/apis/policy"
	"k8s.io/kubernetes/pkg/apis/policy/validation"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/generic"
	"k8s.io/kubernetes/pkg/runtime"
	"k8s.io/kubernetes/pkg/util/validation/field"
)

// networkPolicyStrategy implements verification logic for NetworkPolicys.
type networkPolicyStrategy struct {
	runtime.ObjectTyper
	api.NameGenerator
}

// Strategy is the default logic that applies when creating and updating NetworkPolicy objects.
var Strategy = networkPolicyStrategy{api.Scheme, api.SimpleNameGenerator}

// NamespaceScoped returns true because all NetworkPolicys need to be within a namespace.
func (networkPolicyStrategy) NamespaceScoped() bool {
	return true
}

// PrepareForCreate clears the status of an NetworkPolicy before creation.
func (networkPolicyStrategy) PrepareForCreate(obj runtime.Object) {
	networkPolicy := obj.(*policy.NetworkPolicy)
	networkPolicy.Generation = 1
}

// PrepareForUpdate clears fields that are not allowed to be set by end users on update.
func (networkPolicyStrategy) PrepareForUpdate(obj, old runtime.Object) {
	newNetworkPolicy := obj.(*policy.NetworkPolicy)
	oldNetworkPolicy := old.(*policy.NetworkPolicy)

	// Any changes to the spec increment the generation number, any changes to the
	// status should reflect the generation number of the corresponding object.
	// See api.ObjectMeta description for more information on Generation.
	if !reflect.DeepEqual(oldNetworkPolicy.Spec, newNetworkPolicy.Spec) {
		newNetworkPolicy.Generation = oldNetworkPolicy.Generation + 1
	}
}

// Validate validates a new NetworkPolicy.
func (networkPolicyStrategy) Validate(ctx api.Context, obj runtime.Object) field.ErrorList {
	networkPolicy := obj.(*policy.NetworkPolicy)
	return validation.ValidateNetworkPolicy(networkPolicy)
}

// Canonicalize normalizes the object after validation.
func (networkPolicyStrategy) Canonicalize(obj runtime.Object) {
}

// AllowCreateOnUpdate is true for NetworkPolicy; this means you may create one with a PUT request.
func (networkPolicyStrategy) AllowCreateOnUpdate() bool {
	return false
}

// ValidateUpdate is the default update validation for an end user.
func (networkPolicyStrategy) ValidateUpdate(ctx api.Context, obj, old runtime.Object) field.ErrorList {
	validationErrorList := validation.ValidateNetworkPolicy(obj.(*policy.NetworkPolicy))
	updateErrorList := validation.ValidateNetworkPolicyUpdate(obj.(*policy.NetworkPolicy), old.(*policy.NetworkPolicy))
	return append(validationErrorList, updateErrorList...)
}

// AllowUnconditionalUpdate is the default update policy for NetworkPolicy objects.
func (networkPolicyStrategy) AllowUnconditionalUpdate() bool {
	return true
}

// NetworkPolicyToSelectableFields returns a field set that represents the object.
func NetworkPolicyToSelectableFields(networkPolicy *policy.NetworkPolicy) fields.Set {
	return generic.ObjectMetaFieldsSet(networkPolicy.ObjectMeta, true)
}

// MatchNetworkPolicy is the filter used by the generic etcd backend to watch events
// from etcd to clients of the apiserver only interested in specific labels/fields.
func MatchNetworkPolicy(label labels.Selector, field fields.Selector) generic.Matcher {
	return &generic.SelectionPredicate{
		Label: label,
		Field: field,
		GetAttrs: func(obj runtime.Object) (labels.Set, fields.Set, error) {
			networkPolicy, ok := obj.(*policy.NetworkPolicy)
			if !ok {
				return nil, nil, fmt.Errorf("given object is not a NetworkPolicy.")
			}
			return labels.Set(networkPolicy.ObjectMeta.Labels), NetworkPolicyToSelectableFields(networkPolicy), nil
		},
	}
}
