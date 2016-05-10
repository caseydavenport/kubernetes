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
	"testing"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apis/networkpolicy"
)

func TestNetworkPolicyStrategy(t *testing.T) {
	ctx := api.NewDefaultContext()
	if !Strategy.NamespaceScoped() {
		t.Errorf("NetworkPolicy must be namespace scoped")
	}
	if Strategy.AllowCreateOnUpdate() {
		t.Errorf("NetworkPolicy should not allow create on update")
	}

	validSelector := map[string]string{"a": "b"}
	np := &networkpolicy.NetworkPolicy{
		ObjectMeta: api.ObjectMeta{Name: "abc", Namespace: api.NamespaceDefault},
		Spec: networkpolicy.NetworkPolicySpec{
			PodSelector: &unversioned.LabelSelector{MatchLabels: validSelector},
			Ingress:     []network.NetworkPolicyIngressRule{},
		},
	}

	Strategy.PrepareForCreate(np)
	errs := Strategy.Validate(ctx, np)
	if len(errs) != 0 {
		t.Errorf("Unexpected error validating %v", errs)
	}

	validNp.Spec.Selector = &unversioned.LabelSelector{MatchLabels: map[string]string{"a": "bar"}}
	Strategy.PrepareForUpdate(validNp, np)
	errs = Strategy.ValidateUpdate(ctx, validNp, np)
	if len(errs) == 0 {
		t.Errorf("Expected a validation error since updates are disallowed on networkpolicys.")
	}
}

func TestNetworkPolicyStatusStrategy(t *testing.T) {
	ctx := api.NewDefaultContext()
	if !StatusStrategy.NamespaceScoped() {
		t.Errorf("NetworkPolicy must be namespace scoped")
	}
	if StatusStrategy.AllowCreateOnUpdate() {
		t.Errorf("NetworkPolicy should not allow create on update")
	}
	validSelector := map[string]string{"a": "b"}
	validPodTemplate := api.PodTemplate{
		Template: api.PodTemplateSpec{
			ObjectMeta: api.ObjectMeta{
				Labels: validSelector,
			},
			Spec: api.PodSpec{
				RestartPolicy: api.RestartPolicyAlways,
				DNSPolicy:     api.DNSClusterFirst,
				Containers:    []api.Container{{Name: "abc", Image: "image", ImagePullPolicy: "IfNotPresent"}},
			},
		},
	}
	oldNP := &networkpolicy.NetworkPolicy{
		ObjectMeta: api.ObjectMeta{Name: "abc", Namespace: api.NamespaceDefault, ResourceVersion: "10"},
		Spec: networkpolicy.NetworkPolicySpec{
			Replicas: 3,
			Selector: &unversioned.LabelSelector{MatchLabels: validSelector},
			Template: validPodTemplate.Template,
		},
		Status: networkpolicy.NetworkPolicyStatus{
			Replicas: 1,
		},
	}
	newNP := &networkpolicy.NetworkPolicy{
		ObjectMeta: api.ObjectMeta{Name: "abc", Namespace: api.NamespaceDefault, ResourceVersion: "9"},
		Spec: networkpolicy.NetworkPolicySpec{
			Replicas: 1,
			Selector: &unversioned.LabelSelector{MatchLabels: validSelector},
			Template: validPodTemplate.Template,
		},
		Status: networkpolicy.NetworkPolicyStatus{
			Replicas: 2,
		},
	}
	StatusStrategy.PrepareForUpdate(newNP, oldNP)
	if newNP.Status.Replicas != 2 {
		t.Errorf("NetworkPolicy status updates should allow change of pets: %v", newNP.Status.Replicas)
	}
	if newNP.Spec.Replicas != 3 {
		t.Errorf("NetworkPolicy status updates should not clobber spec: %v", newNP.Spec)
	}
	errs := StatusStrategy.ValidateUpdate(ctx, newNP, oldNP)
	if len(errs) != 0 {
		t.Errorf("Unexpected error %v", errs)
	}
}
