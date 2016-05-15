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

package policy

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/util/intstr"
)

// PodDisruptionBudgetSpec is a description of a PodDisruptionBudget.
type PodDisruptionBudgetSpec struct {
	// The minimum number of pods that must be available simultaneously.  This
	// can be either an integer or a string specifying a percentage, e.g. "28%".
	MinAvailable intstr.IntOrString `json:"minAvailable,omitempty"`

	// Label query over pods whose evictions are managed by the disruption
	// budget.
	Selector *unversioned.LabelSelector `json:"selector,omitempty"`
}

// PodDisruptionBudgetStatus represents information about the status of a
// PodDisruptionBudget. Status may trail the actual state of a system.
type PodDisruptionBudgetStatus struct {
	// Whether or not a disruption is currently allowed.
	PodDisruptionAllowed bool `json:"disruptionAllowed"`

	// current number of healthy pods
	CurrentHealthy int32 `json:"currentHealthy"`

	// minimum desired number of healthy pods
	DesiredHealthy int32 `json:"desiredHealthy"`

	// total number of pods counted by this disruption budget
	ExpectedPods int32 `json:"expectedPods"`
}

// +genclient=true,noMethods=true

// PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods
type PodDisruptionBudget struct {
	unversioned.TypeMeta `json:",inline"`
	api.ObjectMeta       `json:"metadata,omitempty"`

	// Specification of the desired behavior of the PodDisruptionBudget.
	Spec PodDisruptionBudgetSpec `json:"spec,omitempty"`
	// Most recently observed status of the PodDisruptionBudget.
	Status PodDisruptionBudgetStatus `json:"status,omitempty"`
}

// PodDisruptionBudgetList is a collection of PodDisruptionBudgets.
type PodDisruptionBudgetList struct {
	unversioned.TypeMeta `json:",inline"`
	unversioned.ListMeta `json:"metadata,omitempty"`
	Items                []PodDisruptionBudget `json:"items"`
}

// NetworkPolicy is an object to define network policies that can be applied to pods.
type NetworkPolicy struct {
	unversioned.TypeMeta `json:",inline"`
	// Standard object metadata; More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata.
	api.ObjectMeta `json:"metadata,omitempty"`

	// Specification of the desired behavior for this NetworkPolicy.
	Spec NetworkPolicySpec `json:"spec,omitempty"`
}

type NetworkPolicySpec struct {
	// Selects the pods to which this NetworkPolicy object applies.
	PodSelector unversioned.LabelSelector `json:"podSelector"`

	// List of ingress rules to be applied to the selected pods.
	Ingress []NetworkPolicyIngressRule `json:"ingress,omitempty"`
}

// This NetworkPolicyIngressRule matches traffic if and only if the traffic matches both Ports AND From.
type NetworkPolicyIngressRule struct {
	// List of ports which should be made accessible on the pods selected by PodSelector.
	Ports []NetworkPolicyPort `json:"ports,omitempty"`

	// List of sources which should be able to access the pods selected by PodSelector.
	From []NetworkPolicyPeer `json:"from,omitempty"`
}

type NetworkPolicyPort struct {
	// The protocol (TCP or UDP) which traffic must match.
	Protocol *api.Protocol `json:"protocol,omitempty"`

	// If specified, the port on the given protocol.
	Port *intstr.IntOrString `json:"port,omitempty"`
}

type NetworkPolicyPeer struct {
	// This is a label selector which selects Pods in this namespace.
	Pods *unversioned.LabelSelector `json:"pods,omitempty"`

	// If 'Pods' is defined, 'Namespaces' must not be.
	Namespaces *unversioned.LabelSelector `json:"namespaces,omitempty"`
}

// NetworkPolicyList is a collection of NetworkPolicys.
type NetworkPolicyList struct {
	unversioned.TypeMeta `json:",inline"`
	unversioned.ListMeta `json:"metadata,omitempty"`
	Items                []NetworkPolicy `json:"items"`
}
