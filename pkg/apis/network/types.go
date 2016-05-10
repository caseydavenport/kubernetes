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

package network

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
)

type NetworkPolicy struct {
	TypeMeta
	ObjectMeta

	// Specification of the desired behavior for this NetworkPolicy.
	Spec NetworkPolicySpec
}

type NetworkPolicySpec struct {
	// Selects the pods to which this NetworkPolicy object applies.  The array of NetworkPolicyIngressRules below
	// is applied to any pods selected by this field. Multiple NetworkPolicy objects can select the
	// same set of pods.  In this case, the NetworkPolicyRules for each are combined additively.
	// This field is NOT optional and follows standard unversioned.LabelSelector semantics.
	// An empty PodSelector matches all pods in this namespace.
	PodSelector unversioned.LabelSelector `json:"podSelector"`

	// List of ingress rules to be applied to the selected pods.
	// Traffic is allowed to a pod if Namespace.NetworkPolicy.Ingress.Isolation is undefined,
	// OR if the traffic source is the pod's local kubelet (for health checks),
	// OR if the traffic matches at least one NetworkPolicyIngressRule across all of the NetworkPolicy
	// objects whose podSelector matches the pod.
	// If this field is nil, this NetworkPolicy does not affect ingress to the selected pods.
	// If this field is non-nil but contains no rules, this NetworkPolicy allows no traffic.
	// If this field is non-nil and contains at least one rule, this NetworkPolicy allows any traffic
	// which matches at least one of the NetworkPolicyIngressRules in this list.
	Ingress []NetworkPolicyIngressRule `json:"ingress,omitempty"`
}

// This NetworkPolicyIngressRule matches traffic if and only if the traffic matches both Ports AND From.
type NetworkPolicyIngressRule struct {
	// List of ports which should be made accessible on the pods selected by PodSelector.
	// Each item in this list is combined using a logical OR.
	// If this field is nil, this NetworkPolicyIngressRule matches all ports (traffic not restricted by Port).
	// If this field is non-nil but contains no items, this NetworkPolicyIngressRule matches no ports (no traffic matches).
	// If this field is non-nil and contains at least one item, then this NetworkPolicyIngressRule allows traffic
	// only if the traffic matches at least one NetworkPolicyPort in the Ports list.
	Ports []NetworkPolicyPort `json:"ports,omitempty"`

	// List of sources which should be able to access the pods selected by PodSelector.
	// Items in this list are combined using a logical OR operation.
	// If this field nil, this NetworkPolicyIngressRule matches all sources (traffic not restricted by source).
	// If this field is non-nil but contains no items, this NetworkPolicyIngressRule matches no sources (no traffic matches).
	// If this field is non-nil and contains at least on item, this NetworkPolicyIngressRule allows traffic only if the
	// traffic matches at least one NetworkPolicyPeer in the From list.
	From []NetworkPolicyPeer `json:"from,omitempty"`
}

type NetworkPolicyPort struct {
	// The protocol (TCP or UDP) which traffic must match.
	// If not defined, this field defaults to TCP.
	Protocol api.Protocol `json:"protocol"`

	// If specified, the port on the given protocol.  This can
	// either be a numerical or named port.  If this field is nil,
	// this NetworkPolicyPort matches all port names and numbers.
	// If non-nil, only traffic on the specified protocol AND port
	// will be matched by this NetworkPolicyPort.
	Port *intstr.IntOrString `json:"port,omitempty"`
}

type NetworkPolicyPeer struct {
	// If 'Namespaces' is defined, 'Pods' must not be.
	// This is a label selector which selects Pods in this namespace.
	// This NetworkPolicyPeer matches any pods selected by this selector.
	// This field follows standard unversioned.LabelSelector semantics.
	// If nil, this selector selects no pods.
	// If non-nil but empty, this selector selects all pods in this namespace.
	Pods *unversioned.LabelSelector `json:"pods,omitempty"`

	// If 'Pods' is defined, 'Namespaces' must not be.
	// Selects Kubernetes Namespaces.  This NetworkPolicyPeer matches
	// all pods in all namespaces selected by this label selector.
	// This field follows standard unversioned.LabelSelector semantics.
	// If nil, this selector selects no namespaces.
	// If non-nil but empty, this selector selects all namespaces.
	Namespaces *unversioned.LabelSelector `json:"namespaces,omitempty"`
}

// NetworkPolicyList is a collection of NetworkPolicys.
type NetworkPolicyList struct {
	unversioned.TypeMeta `json:",inline"`
	unversioned.ListMeta `json:"metadata,omitempty"`
	Items                []NetworkPolicy `json:"items"`
}
