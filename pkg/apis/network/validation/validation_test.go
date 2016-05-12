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
	"strings"
	"testing"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/apis/network"
)

func TestValidateNetworkPolicy(t *testing.T) {
	validLabels := map[string]string{"a": "b"}
	invalidLabels := map[string]string{"NoUppercaseOrSpecialCharsLike=Equals": "b"}
	successCases := []network.NetworkPolicy{
		{
			ObjectMeta: api.ObjectMeta{Name: "abc", Namespace: api.NamespaceDefault},
			Spec: network.NetworkPolicySpec{
				PodSelector: &unversioned.LabelSelector{MatchLabels: validLabels},
				Ingress:     []network.NetworkPolicyIngressRule{},
			},
		},
		{
			ObjectMeta: api.ObjectMeta{Name: "abc-123", Namespace: api.NamespaceDefault},
			Spec: network.NetworkPolicySpec{
				PodSelector: &unversioned.LabelSelector{MatchLabels: validLabels},
				Ingress:     []network.NetworkPolicyIngressRule{},
			},
		},
	}
	for _, successCase := range successCases {
		if errs := ValidateNetworkPolicy(&successCase); len(errs) != 0 {
			t.Errorf("expected success: %v", errs)
		}
	}

	errorCases := map[string]network.NetworkPolicy{
		"zero-length ID": {
			ObjectMeta: api.ObjectMeta{Name: "", Namespace: api.NamespaceDefault},
			Spec: network.NetworkPolicySpec{
				PodSelector: &unversioned.LabelSelector{MatchLabels: validLabels},
				Ingress:     []network.NetworkPolicyIngressRule{},
			},
		},
		"missing-namespace": {
			ObjectMeta: api.ObjectMeta{Name: "abc-123"},
			Spec: network.NetworkPolicySpec{
				PodSelector: &unversioned.LabelSelector{MatchLabels: validLabels},
				Ingress:     []network.NetworkPolicyIngressRule{},
			},
		},
		"invalid_label": {
			ObjectMeta: api.ObjectMeta{
				Name:      "abc-123",
				Namespace: api.NamespaceDefault,
				Labels: map[string]string{
					"NoUppercaseOrSpecialCharsLike=Equals": "bar",
				},
			},
			Spec: network.NetworkPolicySpec{
				PodSelector: &unversioned.LabelSelector{MatchLabels: validLabels},
				Ingress:     []network.NetworkPolicyIngressRule{},
			},
		},
		"invalid_label 2": {
			ObjectMeta: api.ObjectMeta{
				Name:      "abc-123",
				Namespace: api.NamespaceDefault,
				Labels: map[string]string{
					"NoUppercaseOrSpecialCharsLike=Equals": "bar",
				},
			},
			Spec: network.NetworkPolicySpec{
				PodSelector: &unversioned.LabelSelector{MatchLabels: validLabels},
				Ingress:     []network.NetworkPolicyIngressRule{},
			},
		},
		"invalid_annotation": {
			ObjectMeta: api.ObjectMeta{
				Name:      "abc-123",
				Namespace: api.NamespaceDefault,
				Annotations: map[string]string{
					"NoUppercaseOrSpecialCharsLike=Equals": "bar",
				},
			},
			Spec: network.NetworkPolicySpec{
				PodSelector: &unversioned.LabelSelector{MatchLabels: validLabels},
				Ingress:     []network.NetworkPolicyIngressRule{},
			},
		},
	}
	for k, v := range errorCases {
		errs := ValidateNetworkPolicy(&v)
		if len(errs) == 0 {
			t.Errorf("expected failure for %s", k)
		}
		for i := range errs {
			field := errs[i].Field
			// TODO
			if !strings.HasPrefix(field, "spec.template.") &&
				field != "metadata.name" &&
				field != "metadata.namespace" &&
				field != "spec.selector" &&
				field != "spec.template" &&
				field != "GCEPersistentDisk.ReadOnly" &&
				field != "spec.replicas" &&
				field != "spec.template.labels" &&
				field != "metadata.annotations" &&
				field != "metadata.labels" &&
				field != "status.replicas" {
				t.Errorf("%s: missing prefix for: %v", k, errs[i])
			}
		}
	}
}

func TestValidateNetworkPolicyUpdate(t *testing.T) {
	validLabels := map[string]string{"a": "b"}
	validPodTemplate := api.PodTemplate{
		Template: api.PodTemplateSpec{
			ObjectMeta: api.ObjectMeta{
				Labels: validLabels,
			},
			Spec: api.PodSpec{
				RestartPolicy: api.RestartPolicyAlways,
				DNSPolicy:     api.DNSClusterFirst,
				Containers:    []api.Container{{Name: "abc", Image: "image", ImagePullPolicy: "IfNotPresent"}},
			},
		},
	}
	readWriteVolumePodTemplate := api.PodTemplate{
		Template: api.PodTemplateSpec{
			ObjectMeta: api.ObjectMeta{
				Labels: validLabels,
			},
			Spec: api.PodSpec{
				RestartPolicy: api.RestartPolicyAlways,
				DNSPolicy:     api.DNSClusterFirst,
				Containers:    []api.Container{{Name: "abc", Image: "image", ImagePullPolicy: "IfNotPresent"}},
				Volumes:       []api.Volume{{Name: "gcepd", VolumeSource: api.VolumeSource{GCEPersistentDisk: &api.GCEPersistentDiskVolumeSource{PDName: "my-PD", FSType: "ext4", Partition: 1, ReadOnly: false}}}},
			},
		},
	}
	invalidLabels := map[string]string{"NoUppercaseOrSpecialCharsLike=Equals": "b"}
	invalidPodTemplate := api.PodTemplate{
		Template: api.PodTemplateSpec{
			Spec: api.PodSpec{
				RestartPolicy: api.RestartPolicyAlways,
				DNSPolicy:     api.DNSClusterFirst,
			},
			ObjectMeta: api.ObjectMeta{
				Labels: invalidLabels,
			},
		},
	}
	type psUpdateTest struct {
		old    network.NetworkPolicy
		update network.NetworkPolicy
	}
	successCases := []psUpdateTest{
		{
			old: network.NetworkPolicy{
				ObjectMeta: api.ObjectMeta{Name: "abc", Namespace: api.NamespaceDefault},
				Spec: network.NetworkPolicySpec{
					Selector: &unversioned.LabelSelector{MatchLabels: validLabels},
					Template: validPodTemplate.Template,
				},
			},
			update: network.NetworkPolicy{
				ObjectMeta: api.ObjectMeta{Name: "abc", Namespace: api.NamespaceDefault},
				Spec: network.NetworkPolicySpec{
					Replicas: 3,
					Selector: &unversioned.LabelSelector{MatchLabels: validLabels},
					Template: validPodTemplate.Template,
				},
			},
		},
	}
	for _, successCase := range successCases {
		successCase.old.ObjectMeta.ResourceVersion = "1"
		successCase.update.ObjectMeta.ResourceVersion = "1"
		if errs := ValidateNetworkPolicyUpdate(&successCase.update, &successCase.old); len(errs) != 0 {
			t.Errorf("expected success: %v", errs)
		}
	}
	errorCases := map[string]psUpdateTest{
		"more than one read/write": {
			old: network.NetworkPolicy{
				ObjectMeta: api.ObjectMeta{Name: "", Namespace: api.NamespaceDefault},
				Spec: network.NetworkPolicySpec{
					Selector: &unversioned.LabelSelector{MatchLabels: validLabels},
					Template: validPodTemplate.Template,
				},
			},
			update: network.NetworkPolicy{
				ObjectMeta: api.ObjectMeta{Name: "abc", Namespace: api.NamespaceDefault},
				Spec: network.NetworkPolicySpec{
					Replicas: 2,
					Selector: &unversioned.LabelSelector{MatchLabels: validLabels},
					Template: readWriteVolumePodTemplate.Template,
				},
			},
		},
		"updates to a field other than spec.Replicas": {
			old: network.NetworkPolicy{
				ObjectMeta: api.ObjectMeta{Name: "abc", Namespace: api.NamespaceDefault},
				Spec: network.NetworkPolicySpec{
					Selector: &unversioned.LabelSelector{MatchLabels: validLabels},
					Template: validPodTemplate.Template,
				},
			},
			update: network.NetworkPolicy{
				ObjectMeta: api.ObjectMeta{Name: "abc", Namespace: api.NamespaceDefault},
				Spec: network.NetworkPolicySpec{
					Replicas: 1,
					Selector: &unversioned.LabelSelector{MatchLabels: validLabels},
					Template: readWriteVolumePodTemplate.Template,
				},
			},
		},
		"invalid selector": {
			old: network.NetworkPolicy{
				ObjectMeta: api.ObjectMeta{Name: "", Namespace: api.NamespaceDefault},
				Spec: network.NetworkPolicySpec{
					Selector: &unversioned.LabelSelector{MatchLabels: validLabels},
					Template: validPodTemplate.Template,
				},
			},
			update: network.NetworkPolicy{
				ObjectMeta: api.ObjectMeta{Name: "abc", Namespace: api.NamespaceDefault},
				Spec: network.NetworkPolicySpec{
					Replicas: 2,
					Selector: &unversioned.LabelSelector{MatchLabels: invalidLabels},
					Template: validPodTemplate.Template,
				},
			},
		},
		"invalid pod": {
			old: network.NetworkPolicy{
				ObjectMeta: api.ObjectMeta{Name: "", Namespace: api.NamespaceDefault},
				Spec: network.NetworkPolicySpec{
					Selector: &unversioned.LabelSelector{MatchLabels: validLabels},
					Template: validPodTemplate.Template,
				},
			},
			update: network.NetworkPolicy{
				ObjectMeta: api.ObjectMeta{Name: "abc", Namespace: api.NamespaceDefault},
				Spec: network.NetworkPolicySpec{
					Replicas: 2,
					Selector: &unversioned.LabelSelector{MatchLabels: validLabels},
					Template: invalidPodTemplate.Template,
				},
			},
		},
		"negative replicas": {
			old: network.NetworkPolicy{
				ObjectMeta: api.ObjectMeta{Name: "abc", Namespace: api.NamespaceDefault},
				Spec: network.NetworkPolicySpec{
					Selector: &unversioned.LabelSelector{MatchLabels: validLabels},
					Template: validPodTemplate.Template,
				},
			},
			update: network.NetworkPolicy{
				ObjectMeta: api.ObjectMeta{Name: "abc", Namespace: api.NamespaceDefault},
				Spec: network.NetworkPolicySpec{
					Replicas: -1,
					Selector: &unversioned.LabelSelector{MatchLabels: validLabels},
					Template: validPodTemplate.Template,
				},
			},
		},
	}
	for testName, errorCase := range errorCases {
		if errs := ValidateNetworkPolicyUpdate(&errorCase.update, &errorCase.old); len(errs) == 0 {
			t.Errorf("expected failure: %s", testName)
		}
	}
}
