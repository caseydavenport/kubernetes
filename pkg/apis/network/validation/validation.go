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

package validation

import (
	apivalidation "k8s.io/kubernetes/pkg/api/validation"
	"k8s.io/kubernetes/pkg/apis/network"
	"k8s.io/kubernetes/pkg/util/validation/field"
)

// ValidateNetworkPolicyName can be used to check whether the given networkpolicy
// name is valid.
func ValidateNetworkPolicyName(name string, prefix bool) (bool, string) {
	return apivalidation.NameIsDNSSubdomain(name, prefix)
}

// ValidateNetworkPolicySpec tests if required fields in the networkpolicy spec are set.
func ValidateNetworkPolicySpec(spec *network.NetworkPolicySpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	// TODO
	return allErrs
}

// ValidateNetworkPolicy validates a networkpolicy.
func ValidateNetworkPolicy(networkpolicy *network.NetworkPolicy) field.ErrorList {
	allErrs := apivalidation.ValidateObjectMeta(&networkpolicy.ObjectMeta, true, ValidateNetworkPolicyName, field.NewPath("metadata"))
	allErrs = append(allErrs, ValidateNetworkPolicySpec(&networkpolicy.Spec, field.NewPath("spec"))...)
	return allErrs
}

// ValidateNetworkPolicyUpdate tests if required fields in the networkpolicy are set.
func ValidateNetworkPolicyUpdate(networkPolicy, oldNetworkPolicy *network.NetworkPolicy) field.ErrorList {
	allErrs := field.ErrorList{}
	// TODO
	return allErrs
}

// ValidateNetworkPolicyStatusUpdate tests if required fields in the networkpolicy are set.
func ValidateNetworkPolicyStatusUpdate(networkPolicy, oldNetworkPolicy *network.NetworkPolicy) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, apivalidation.ValidateObjectMetaUpdate(&networkPolicy.ObjectMeta, &oldNetworkPolicy.ObjectMeta, field.NewPath("metadata"))...)
	// TODO: Validate status.
	return allErrs
}
