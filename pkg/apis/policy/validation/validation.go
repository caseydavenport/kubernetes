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
	"reflect"

	"github.com/golang/glog"
	unversionedvalidation "k8s.io/kubernetes/pkg/api/unversioned/validation"
	apivalidation "k8s.io/kubernetes/pkg/api/validation"
	extensionsvalidation "k8s.io/kubernetes/pkg/apis/extensions/validation"
	"k8s.io/kubernetes/pkg/apis/policy"
	"k8s.io/kubernetes/pkg/util/validation/field"
)

func ValidatePodDisruptionBudget(pdb *policy.PodDisruptionBudget) field.ErrorList {
	allErrs := ValidatePodDisruptionBudgetSpec(pdb.Spec, field.NewPath("spec"))
	return allErrs
}

func ValidatePodDisruptionBudgetUpdate(pdb, oldPdb *policy.PodDisruptionBudget) field.ErrorList {
	allErrs := field.ErrorList{}

	restoreGeneration := pdb.Generation
	pdb.Generation = oldPdb.Generation

	if !reflect.DeepEqual(pdb, oldPdb) {
		allErrs = append(allErrs, field.Forbidden(field.NewPath("spec"), "updates to poddisruptionbudget spec are forbidden."))
	}

	pdb.Generation = restoreGeneration
	return allErrs
}

func ValidatePodDisruptionBudgetSpec(spec policy.PodDisruptionBudgetSpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}

	allErrs = append(allErrs, extensionsvalidation.ValidatePositiveIntOrPercent(spec.MinAvailable, fldPath.Child("minAvailable"))...)
	allErrs = append(allErrs, extensionsvalidation.IsNotMoreThan100Percent(spec.MinAvailable, fldPath.Child("minAvailable"))...)
	allErrs = append(allErrs, unversionedvalidation.ValidateLabelSelector(spec.Selector, fldPath.Child("selector"))...)

	return allErrs
}

// ValidateNetworkPolicyName can be used to check whether the given networkpolicy
// name is valid.
func ValidateNetworkPolicyName(name string, prefix bool) (bool, string) {
	return apivalidation.NameIsDNSSubdomain(name, prefix)
}

// ValidateNetworkPolicySpec tests if required fields in the networkpolicy spec are set.
func ValidateNetworkPolicySpec(spec *policy.NetworkPolicySpec, fldPath *field.Path) field.ErrorList {
	allErrs := field.ErrorList{}
	// TODO CD4
	allErrs = append(allErrs, unversionedvalidation.ValidateLabelSelector(&spec.PodSelector, fldPath.Child("podSelector"))...)
	return allErrs
}

// ValidateNetworkPolicy validates a networkpolicy.
func ValidateNetworkPolicy(networkPolicy *policy.NetworkPolicy) field.ErrorList {
	glog.V(1).Infof("CD4 Validating network policy: %s", networkPolicy)
	allErrs := apivalidation.ValidateObjectMeta(&networkPolicy.ObjectMeta, true, ValidateNetworkPolicyName, field.NewPath("metadata"))
	allErrs = append(allErrs, ValidateNetworkPolicySpec(&networkPolicy.Spec, field.NewPath("spec"))...)
	return allErrs
}

// ValidateNetworkPolicyUpdate tests if required fields in the networkpolicy are set.
func ValidateNetworkPolicyUpdate(networkPolicy, oldNetworkPolicy *policy.NetworkPolicy) field.ErrorList {
	allErrs := field.ErrorList{}
	// TODO
	return allErrs
}

// ValidateNetworkPolicyStatusUpdate tests if required fields in the networkpolicy are set.
func ValidateNetworkPolicyStatusUpdate(networkPolicy, oldNetworkPolicy *policy.NetworkPolicy) field.ErrorList {
	allErrs := field.ErrorList{}
	allErrs = append(allErrs, apivalidation.ValidateObjectMetaUpdate(&networkPolicy.ObjectMeta, &oldNetworkPolicy.ObjectMeta, field.NewPath("metadata"))...)
	// TODO: Validate status.
	return allErrs
}
