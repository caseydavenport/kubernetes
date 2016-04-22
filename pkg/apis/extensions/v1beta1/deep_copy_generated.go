// +build !ignore_autogenerated

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1beta1

import (
	api "k8s.io/kubernetes/pkg/api"
	resource "k8s.io/kubernetes/pkg/api/resource"
	unversioned "k8s.io/kubernetes/pkg/api/unversioned"
	v1 "k8s.io/kubernetes/pkg/api/v1"
	conversion "k8s.io/kubernetes/pkg/conversion"
	intstr "k8s.io/kubernetes/pkg/util/intstr"
)

func init() {
	if err := api.Scheme.AddGeneratedDeepCopyFuncs(
		DeepCopy_v1beta1_APIVersion,
		DeepCopy_v1beta1_CPUTargetUtilization,
		DeepCopy_v1beta1_CustomMetricCurrentStatus,
		DeepCopy_v1beta1_CustomMetricCurrentStatusList,
		DeepCopy_v1beta1_CustomMetricTarget,
		DeepCopy_v1beta1_CustomMetricTargetList,
		DeepCopy_v1beta1_DaemonSet,
		DeepCopy_v1beta1_DaemonSetList,
		DeepCopy_v1beta1_DaemonSetSpec,
		DeepCopy_v1beta1_DaemonSetStatus,
		DeepCopy_v1beta1_Deployment,
		DeepCopy_v1beta1_DeploymentList,
		DeepCopy_v1beta1_DeploymentRollback,
		DeepCopy_v1beta1_DeploymentSpec,
		DeepCopy_v1beta1_DeploymentStatus,
		DeepCopy_v1beta1_DeploymentStrategy,
		DeepCopy_v1beta1_ExportOptions,
		DeepCopy_v1beta1_HTTPIngressPath,
		DeepCopy_v1beta1_HTTPIngressRuleValue,
		DeepCopy_v1beta1_HorizontalPodAutoscaler,
		DeepCopy_v1beta1_HorizontalPodAutoscalerList,
		DeepCopy_v1beta1_HorizontalPodAutoscalerSpec,
		DeepCopy_v1beta1_HorizontalPodAutoscalerStatus,
		DeepCopy_v1beta1_HostPortRange,
		DeepCopy_v1beta1_IDRange,
		DeepCopy_v1beta1_Ingress,
		DeepCopy_v1beta1_IngressBackend,
		DeepCopy_v1beta1_IngressList,
		DeepCopy_v1beta1_IngressRule,
		DeepCopy_v1beta1_IngressRuleValue,
		DeepCopy_v1beta1_IngressSpec,
		DeepCopy_v1beta1_IngressStatus,
		DeepCopy_v1beta1_IngressTLS,
		DeepCopy_v1beta1_Job,
		DeepCopy_v1beta1_JobCondition,
		DeepCopy_v1beta1_JobList,
		DeepCopy_v1beta1_JobSpec,
		DeepCopy_v1beta1_JobStatus,
		DeepCopy_v1beta1_LabelSelector,
		DeepCopy_v1beta1_LabelSelectorRequirement,
		DeepCopy_v1beta1_ListOptions,
		DeepCopy_v1beta1_NetworkPolicy,
		DeepCopy_v1beta1_NetworkPolicyIngressRule,
		DeepCopy_v1beta1_NetworkPolicyList,
		DeepCopy_v1beta1_NetworkPolicyPort,
		DeepCopy_v1beta1_NetworkPolicySource,
		DeepCopy_v1beta1_NetworkPolicySpec,
		DeepCopy_v1beta1_PodSecurityPolicy,
		DeepCopy_v1beta1_PodSecurityPolicyList,
		DeepCopy_v1beta1_PodSecurityPolicySpec,
		DeepCopy_v1beta1_ReplicaSet,
		DeepCopy_v1beta1_ReplicaSetList,
		DeepCopy_v1beta1_ReplicaSetSpec,
		DeepCopy_v1beta1_ReplicaSetStatus,
		DeepCopy_v1beta1_ReplicationControllerDummy,
		DeepCopy_v1beta1_RollbackConfig,
		DeepCopy_v1beta1_RollingUpdateDeployment,
		DeepCopy_v1beta1_RunAsUserStrategyOptions,
		DeepCopy_v1beta1_SELinuxStrategyOptions,
		DeepCopy_v1beta1_Scale,
		DeepCopy_v1beta1_ScaleSpec,
		DeepCopy_v1beta1_ScaleStatus,
		DeepCopy_v1beta1_SubresourceReference,
		DeepCopy_v1beta1_ThirdPartyResource,
		DeepCopy_v1beta1_ThirdPartyResourceData,
		DeepCopy_v1beta1_ThirdPartyResourceDataList,
		DeepCopy_v1beta1_ThirdPartyResourceList,
	); err != nil {
		// if one of the deep copy functions is malformed, detect it immediately.
		panic(err)
	}
}

func DeepCopy_v1beta1_APIVersion(in APIVersion, out *APIVersion, c *conversion.Cloner) error {
	out.Name = in.Name
	return nil
}

func DeepCopy_v1beta1_CPUTargetUtilization(in CPUTargetUtilization, out *CPUTargetUtilization, c *conversion.Cloner) error {
	out.TargetPercentage = in.TargetPercentage
	return nil
}

func DeepCopy_v1beta1_CustomMetricCurrentStatus(in CustomMetricCurrentStatus, out *CustomMetricCurrentStatus, c *conversion.Cloner) error {
	out.Name = in.Name
	if err := resource.DeepCopy_resource_Quantity(in.CurrentValue, &out.CurrentValue, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_CustomMetricCurrentStatusList(in CustomMetricCurrentStatusList, out *CustomMetricCurrentStatusList, c *conversion.Cloner) error {
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]CustomMetricCurrentStatus, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_CustomMetricCurrentStatus(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_v1beta1_CustomMetricTarget(in CustomMetricTarget, out *CustomMetricTarget, c *conversion.Cloner) error {
	out.Name = in.Name
	if err := resource.DeepCopy_resource_Quantity(in.TargetValue, &out.TargetValue, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_CustomMetricTargetList(in CustomMetricTargetList, out *CustomMetricTargetList, c *conversion.Cloner) error {
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]CustomMetricTarget, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_CustomMetricTarget(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_v1beta1_DaemonSet(in DaemonSet, out *DaemonSet, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_DaemonSetSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_DaemonSetStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_DaemonSetList(in DaemonSetList, out *DaemonSetList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]DaemonSet, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_DaemonSet(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_v1beta1_DaemonSetSpec(in DaemonSetSpec, out *DaemonSetSpec, c *conversion.Cloner) error {
	if in.Selector != nil {
		in, out := in.Selector, &out.Selector
		*out = new(LabelSelector)
		if err := DeepCopy_v1beta1_LabelSelector(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Selector = nil
	}
	if err := v1.DeepCopy_v1_PodTemplateSpec(in.Template, &out.Template, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_DaemonSetStatus(in DaemonSetStatus, out *DaemonSetStatus, c *conversion.Cloner) error {
	out.CurrentNumberScheduled = in.CurrentNumberScheduled
	out.NumberMisscheduled = in.NumberMisscheduled
	out.DesiredNumberScheduled = in.DesiredNumberScheduled
	return nil
}

func DeepCopy_v1beta1_Deployment(in Deployment, out *Deployment, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_DeploymentSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_DeploymentStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_DeploymentList(in DeploymentList, out *DeploymentList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]Deployment, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_Deployment(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_v1beta1_DeploymentRollback(in DeploymentRollback, out *DeploymentRollback, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	out.Name = in.Name
	if in.UpdatedAnnotations != nil {
		in, out := in.UpdatedAnnotations, &out.UpdatedAnnotations
		*out = make(map[string]string)
		for key, val := range in {
			(*out)[key] = val
		}
	} else {
		out.UpdatedAnnotations = nil
	}
	if err := DeepCopy_v1beta1_RollbackConfig(in.RollbackTo, &out.RollbackTo, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_DeploymentSpec(in DeploymentSpec, out *DeploymentSpec, c *conversion.Cloner) error {
	if in.Replicas != nil {
		in, out := in.Replicas, &out.Replicas
		*out = new(int32)
		**out = *in
	} else {
		out.Replicas = nil
	}
	if in.Selector != nil {
		in, out := in.Selector, &out.Selector
		*out = new(LabelSelector)
		if err := DeepCopy_v1beta1_LabelSelector(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Selector = nil
	}
	if err := v1.DeepCopy_v1_PodTemplateSpec(in.Template, &out.Template, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_DeploymentStrategy(in.Strategy, &out.Strategy, c); err != nil {
		return err
	}
	out.MinReadySeconds = in.MinReadySeconds
	if in.RevisionHistoryLimit != nil {
		in, out := in.RevisionHistoryLimit, &out.RevisionHistoryLimit
		*out = new(int32)
		**out = *in
	} else {
		out.RevisionHistoryLimit = nil
	}
	out.Paused = in.Paused
	if in.RollbackTo != nil {
		in, out := in.RollbackTo, &out.RollbackTo
		*out = new(RollbackConfig)
		if err := DeepCopy_v1beta1_RollbackConfig(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.RollbackTo = nil
	}
	return nil
}

func DeepCopy_v1beta1_DeploymentStatus(in DeploymentStatus, out *DeploymentStatus, c *conversion.Cloner) error {
	out.ObservedGeneration = in.ObservedGeneration
	out.Replicas = in.Replicas
	out.UpdatedReplicas = in.UpdatedReplicas
	out.AvailableReplicas = in.AvailableReplicas
	out.UnavailableReplicas = in.UnavailableReplicas
	return nil
}

func DeepCopy_v1beta1_DeploymentStrategy(in DeploymentStrategy, out *DeploymentStrategy, c *conversion.Cloner) error {
	out.Type = in.Type
	if in.RollingUpdate != nil {
		in, out := in.RollingUpdate, &out.RollingUpdate
		*out = new(RollingUpdateDeployment)
		if err := DeepCopy_v1beta1_RollingUpdateDeployment(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.RollingUpdate = nil
	}
	return nil
}

func DeepCopy_v1beta1_ExportOptions(in ExportOptions, out *ExportOptions, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	out.Export = in.Export
	out.Exact = in.Exact
	return nil
}

func DeepCopy_v1beta1_HTTPIngressPath(in HTTPIngressPath, out *HTTPIngressPath, c *conversion.Cloner) error {
	out.Path = in.Path
	if err := DeepCopy_v1beta1_IngressBackend(in.Backend, &out.Backend, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_HTTPIngressRuleValue(in HTTPIngressRuleValue, out *HTTPIngressRuleValue, c *conversion.Cloner) error {
	if in.Paths != nil {
		in, out := in.Paths, &out.Paths
		*out = make([]HTTPIngressPath, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_HTTPIngressPath(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Paths = nil
	}
	return nil
}

func DeepCopy_v1beta1_HorizontalPodAutoscaler(in HorizontalPodAutoscaler, out *HorizontalPodAutoscaler, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_HorizontalPodAutoscalerSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_HorizontalPodAutoscalerStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_HorizontalPodAutoscalerList(in HorizontalPodAutoscalerList, out *HorizontalPodAutoscalerList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]HorizontalPodAutoscaler, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_HorizontalPodAutoscaler(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_v1beta1_HorizontalPodAutoscalerSpec(in HorizontalPodAutoscalerSpec, out *HorizontalPodAutoscalerSpec, c *conversion.Cloner) error {
	if err := DeepCopy_v1beta1_SubresourceReference(in.ScaleRef, &out.ScaleRef, c); err != nil {
		return err
	}
	if in.MinReplicas != nil {
		in, out := in.MinReplicas, &out.MinReplicas
		*out = new(int32)
		**out = *in
	} else {
		out.MinReplicas = nil
	}
	out.MaxReplicas = in.MaxReplicas
	if in.CPUUtilization != nil {
		in, out := in.CPUUtilization, &out.CPUUtilization
		*out = new(CPUTargetUtilization)
		if err := DeepCopy_v1beta1_CPUTargetUtilization(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.CPUUtilization = nil
	}
	return nil
}

func DeepCopy_v1beta1_HorizontalPodAutoscalerStatus(in HorizontalPodAutoscalerStatus, out *HorizontalPodAutoscalerStatus, c *conversion.Cloner) error {
	if in.ObservedGeneration != nil {
		in, out := in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = *in
	} else {
		out.ObservedGeneration = nil
	}
	if in.LastScaleTime != nil {
		in, out := in.LastScaleTime, &out.LastScaleTime
		*out = new(unversioned.Time)
		if err := unversioned.DeepCopy_unversioned_Time(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.LastScaleTime = nil
	}
	out.CurrentReplicas = in.CurrentReplicas
	out.DesiredReplicas = in.DesiredReplicas
	if in.CurrentCPUUtilizationPercentage != nil {
		in, out := in.CurrentCPUUtilizationPercentage, &out.CurrentCPUUtilizationPercentage
		*out = new(int32)
		**out = *in
	} else {
		out.CurrentCPUUtilizationPercentage = nil
	}
	return nil
}

func DeepCopy_v1beta1_HostPortRange(in HostPortRange, out *HostPortRange, c *conversion.Cloner) error {
	out.Min = in.Min
	out.Max = in.Max
	return nil
}

func DeepCopy_v1beta1_IDRange(in IDRange, out *IDRange, c *conversion.Cloner) error {
	out.Min = in.Min
	out.Max = in.Max
	return nil
}

func DeepCopy_v1beta1_Ingress(in Ingress, out *Ingress, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_IngressSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_IngressStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_IngressBackend(in IngressBackend, out *IngressBackend, c *conversion.Cloner) error {
	out.ServiceName = in.ServiceName
	if err := intstr.DeepCopy_intstr_IntOrString(in.ServicePort, &out.ServicePort, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_IngressList(in IngressList, out *IngressList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]Ingress, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_Ingress(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_v1beta1_IngressRule(in IngressRule, out *IngressRule, c *conversion.Cloner) error {
	out.Host = in.Host
	if err := DeepCopy_v1beta1_IngressRuleValue(in.IngressRuleValue, &out.IngressRuleValue, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_IngressRuleValue(in IngressRuleValue, out *IngressRuleValue, c *conversion.Cloner) error {
	if in.HTTP != nil {
		in, out := in.HTTP, &out.HTTP
		*out = new(HTTPIngressRuleValue)
		if err := DeepCopy_v1beta1_HTTPIngressRuleValue(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.HTTP = nil
	}
	return nil
}

func DeepCopy_v1beta1_IngressSpec(in IngressSpec, out *IngressSpec, c *conversion.Cloner) error {
	if in.Backend != nil {
		in, out := in.Backend, &out.Backend
		*out = new(IngressBackend)
		if err := DeepCopy_v1beta1_IngressBackend(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Backend = nil
	}
	if in.TLS != nil {
		in, out := in.TLS, &out.TLS
		*out = make([]IngressTLS, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_IngressTLS(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.TLS = nil
	}
	if in.Rules != nil {
		in, out := in.Rules, &out.Rules
		*out = make([]IngressRule, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_IngressRule(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Rules = nil
	}
	return nil
}

func DeepCopy_v1beta1_IngressStatus(in IngressStatus, out *IngressStatus, c *conversion.Cloner) error {
	if err := v1.DeepCopy_v1_LoadBalancerStatus(in.LoadBalancer, &out.LoadBalancer, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_IngressTLS(in IngressTLS, out *IngressTLS, c *conversion.Cloner) error {
	if in.Hosts != nil {
		in, out := in.Hosts, &out.Hosts
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.Hosts = nil
	}
	out.SecretName = in.SecretName
	return nil
}

func DeepCopy_v1beta1_Job(in Job, out *Job, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_JobSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_JobStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_JobCondition(in JobCondition, out *JobCondition, c *conversion.Cloner) error {
	out.Type = in.Type
	out.Status = in.Status
	if err := unversioned.DeepCopy_unversioned_Time(in.LastProbeTime, &out.LastProbeTime, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_Time(in.LastTransitionTime, &out.LastTransitionTime, c); err != nil {
		return err
	}
	out.Reason = in.Reason
	out.Message = in.Message
	return nil
}

func DeepCopy_v1beta1_JobList(in JobList, out *JobList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]Job, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_Job(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_v1beta1_JobSpec(in JobSpec, out *JobSpec, c *conversion.Cloner) error {
	if in.Parallelism != nil {
		in, out := in.Parallelism, &out.Parallelism
		*out = new(int32)
		**out = *in
	} else {
		out.Parallelism = nil
	}
	if in.Completions != nil {
		in, out := in.Completions, &out.Completions
		*out = new(int32)
		**out = *in
	} else {
		out.Completions = nil
	}
	if in.ActiveDeadlineSeconds != nil {
		in, out := in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
		*out = new(int64)
		**out = *in
	} else {
		out.ActiveDeadlineSeconds = nil
	}
	if in.Selector != nil {
		in, out := in.Selector, &out.Selector
		*out = new(LabelSelector)
		if err := DeepCopy_v1beta1_LabelSelector(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Selector = nil
	}
	if in.AutoSelector != nil {
		in, out := in.AutoSelector, &out.AutoSelector
		*out = new(bool)
		**out = *in
	} else {
		out.AutoSelector = nil
	}
	if err := v1.DeepCopy_v1_PodTemplateSpec(in.Template, &out.Template, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_JobStatus(in JobStatus, out *JobStatus, c *conversion.Cloner) error {
	if in.Conditions != nil {
		in, out := in.Conditions, &out.Conditions
		*out = make([]JobCondition, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_JobCondition(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Conditions = nil
	}
	if in.StartTime != nil {
		in, out := in.StartTime, &out.StartTime
		*out = new(unversioned.Time)
		if err := unversioned.DeepCopy_unversioned_Time(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.StartTime = nil
	}
	if in.CompletionTime != nil {
		in, out := in.CompletionTime, &out.CompletionTime
		*out = new(unversioned.Time)
		if err := unversioned.DeepCopy_unversioned_Time(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.CompletionTime = nil
	}
	out.Active = in.Active
	out.Succeeded = in.Succeeded
	out.Failed = in.Failed
	return nil
}

func DeepCopy_v1beta1_LabelSelector(in LabelSelector, out *LabelSelector, c *conversion.Cloner) error {
	if in.MatchLabels != nil {
		in, out := in.MatchLabels, &out.MatchLabels
		*out = make(map[string]string)
		for key, val := range in {
			(*out)[key] = val
		}
	} else {
		out.MatchLabels = nil
	}
	if in.MatchExpressions != nil {
		in, out := in.MatchExpressions, &out.MatchExpressions
		*out = make([]LabelSelectorRequirement, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_LabelSelectorRequirement(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.MatchExpressions = nil
	}
	return nil
}

func DeepCopy_v1beta1_LabelSelectorRequirement(in LabelSelectorRequirement, out *LabelSelectorRequirement, c *conversion.Cloner) error {
	out.Key = in.Key
	out.Operator = in.Operator
	if in.Values != nil {
		in, out := in.Values, &out.Values
		*out = make([]string, len(in))
		copy(*out, in)
	} else {
		out.Values = nil
	}
	return nil
}

func DeepCopy_v1beta1_ListOptions(in ListOptions, out *ListOptions, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	out.LabelSelector = in.LabelSelector
	out.FieldSelector = in.FieldSelector
	out.Watch = in.Watch
	out.ResourceVersion = in.ResourceVersion
	if in.TimeoutSeconds != nil {
		in, out := in.TimeoutSeconds, &out.TimeoutSeconds
		*out = new(int64)
		**out = *in
	} else {
		out.TimeoutSeconds = nil
	}
	return nil
}

func DeepCopy_v1beta1_NetworkPolicy(in NetworkPolicy, out *NetworkPolicy, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_NetworkPolicySpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_NetworkPolicyIngressRule(in NetworkPolicyIngressRule, out *NetworkPolicyIngressRule, c *conversion.Cloner) error {
	if in.Ports != nil {
		in, out := in.Ports, &out.Ports
		*out = make([]NetworkPolicyPort, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_NetworkPolicyPort(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Ports = nil
	}
	if in.From != nil {
		in, out := in.From, &out.From
		*out = make([]NetworkPolicySource, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_NetworkPolicySource(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.From = nil
	}
	return nil
}

func DeepCopy_v1beta1_NetworkPolicyList(in NetworkPolicyList, out *NetworkPolicyList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]NetworkPolicy, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_NetworkPolicy(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_v1beta1_NetworkPolicyPort(in NetworkPolicyPort, out *NetworkPolicyPort, c *conversion.Cloner) error {
	out.Protocol = in.Protocol
	if in.Port != nil {
		in, out := in.Port, &out.Port
		*out = new(intstr.IntOrString)
		if err := intstr.DeepCopy_intstr_IntOrString(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Port = nil
	}
	return nil
}

func DeepCopy_v1beta1_NetworkPolicySource(in NetworkPolicySource, out *NetworkPolicySource, c *conversion.Cloner) error {
	if in.Pods != nil {
		in, out := in.Pods, &out.Pods
		*out = new(unversioned.LabelSelector)
		if err := unversioned.DeepCopy_unversioned_LabelSelector(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Pods = nil
	}
	if in.Namespaces != nil {
		in, out := in.Namespaces, &out.Namespaces
		*out = new(unversioned.LabelSelector)
		if err := unversioned.DeepCopy_unversioned_LabelSelector(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Namespaces = nil
	}
	return nil
}

func DeepCopy_v1beta1_NetworkPolicySpec(in NetworkPolicySpec, out *NetworkPolicySpec, c *conversion.Cloner) error {
	if in.PodSelector != nil {
		in, out := in.PodSelector, &out.PodSelector
		*out = new(unversioned.LabelSelector)
		if err := unversioned.DeepCopy_unversioned_LabelSelector(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.PodSelector = nil
	}
	if in.Ingress != nil {
		in, out := in.Ingress, &out.Ingress
		*out = make([]NetworkPolicyIngressRule, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_NetworkPolicyIngressRule(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Ingress = nil
	}
	return nil
}

func DeepCopy_v1beta1_PodSecurityPolicy(in PodSecurityPolicy, out *PodSecurityPolicy, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_PodSecurityPolicySpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_PodSecurityPolicyList(in PodSecurityPolicyList, out *PodSecurityPolicyList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]PodSecurityPolicy, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_PodSecurityPolicy(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_v1beta1_PodSecurityPolicySpec(in PodSecurityPolicySpec, out *PodSecurityPolicySpec, c *conversion.Cloner) error {
	out.Privileged = in.Privileged
	if in.Capabilities != nil {
		in, out := in.Capabilities, &out.Capabilities
		*out = make([]v1.Capability, len(in))
		for i := range in {
			(*out)[i] = in[i]
		}
	} else {
		out.Capabilities = nil
	}
	if in.Volumes != nil {
		in, out := in.Volumes, &out.Volumes
		*out = make([]FSType, len(in))
		for i := range in {
			(*out)[i] = in[i]
		}
	} else {
		out.Volumes = nil
	}
	out.HostNetwork = in.HostNetwork
	if in.HostPorts != nil {
		in, out := in.HostPorts, &out.HostPorts
		*out = make([]HostPortRange, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_HostPortRange(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.HostPorts = nil
	}
	out.HostPID = in.HostPID
	out.HostIPC = in.HostIPC
	if err := DeepCopy_v1beta1_SELinuxStrategyOptions(in.SELinux, &out.SELinux, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_RunAsUserStrategyOptions(in.RunAsUser, &out.RunAsUser, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_ReplicaSet(in ReplicaSet, out *ReplicaSet, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_ReplicaSetSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_ReplicaSetStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_ReplicaSetList(in ReplicaSetList, out *ReplicaSetList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]ReplicaSet, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_ReplicaSet(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_v1beta1_ReplicaSetSpec(in ReplicaSetSpec, out *ReplicaSetSpec, c *conversion.Cloner) error {
	if in.Replicas != nil {
		in, out := in.Replicas, &out.Replicas
		*out = new(int32)
		**out = *in
	} else {
		out.Replicas = nil
	}
	if in.Selector != nil {
		in, out := in.Selector, &out.Selector
		*out = new(LabelSelector)
		if err := DeepCopy_v1beta1_LabelSelector(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.Selector = nil
	}
	if err := v1.DeepCopy_v1_PodTemplateSpec(in.Template, &out.Template, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_ReplicaSetStatus(in ReplicaSetStatus, out *ReplicaSetStatus, c *conversion.Cloner) error {
	out.Replicas = in.Replicas
	out.FullyLabeledReplicas = in.FullyLabeledReplicas
	out.ObservedGeneration = in.ObservedGeneration
	return nil
}

func DeepCopy_v1beta1_ReplicationControllerDummy(in ReplicationControllerDummy, out *ReplicationControllerDummy, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_RollbackConfig(in RollbackConfig, out *RollbackConfig, c *conversion.Cloner) error {
	out.Revision = in.Revision
	return nil
}

func DeepCopy_v1beta1_RollingUpdateDeployment(in RollingUpdateDeployment, out *RollingUpdateDeployment, c *conversion.Cloner) error {
	if in.MaxUnavailable != nil {
		in, out := in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		if err := intstr.DeepCopy_intstr_IntOrString(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.MaxUnavailable = nil
	}
	if in.MaxSurge != nil {
		in, out := in.MaxSurge, &out.MaxSurge
		*out = new(intstr.IntOrString)
		if err := intstr.DeepCopy_intstr_IntOrString(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.MaxSurge = nil
	}
	return nil
}

func DeepCopy_v1beta1_RunAsUserStrategyOptions(in RunAsUserStrategyOptions, out *RunAsUserStrategyOptions, c *conversion.Cloner) error {
	out.Rule = in.Rule
	if in.Ranges != nil {
		in, out := in.Ranges, &out.Ranges
		*out = make([]IDRange, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_IDRange(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Ranges = nil
	}
	return nil
}

func DeepCopy_v1beta1_SELinuxStrategyOptions(in SELinuxStrategyOptions, out *SELinuxStrategyOptions, c *conversion.Cloner) error {
	out.Rule = in.Rule
	if in.SELinuxOptions != nil {
		in, out := in.SELinuxOptions, &out.SELinuxOptions
		*out = new(v1.SELinuxOptions)
		if err := v1.DeepCopy_v1_SELinuxOptions(*in, *out, c); err != nil {
			return err
		}
	} else {
		out.SELinuxOptions = nil
	}
	return nil
}

func DeepCopy_v1beta1_Scale(in Scale, out *Scale, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_ScaleSpec(in.Spec, &out.Spec, c); err != nil {
		return err
	}
	if err := DeepCopy_v1beta1_ScaleStatus(in.Status, &out.Status, c); err != nil {
		return err
	}
	return nil
}

func DeepCopy_v1beta1_ScaleSpec(in ScaleSpec, out *ScaleSpec, c *conversion.Cloner) error {
	out.Replicas = in.Replicas
	return nil
}

func DeepCopy_v1beta1_ScaleStatus(in ScaleStatus, out *ScaleStatus, c *conversion.Cloner) error {
	out.Replicas = in.Replicas
	if in.Selector != nil {
		in, out := in.Selector, &out.Selector
		*out = make(map[string]string)
		for key, val := range in {
			(*out)[key] = val
		}
	} else {
		out.Selector = nil
	}
	out.TargetSelector = in.TargetSelector
	return nil
}

func DeepCopy_v1beta1_SubresourceReference(in SubresourceReference, out *SubresourceReference, c *conversion.Cloner) error {
	out.Kind = in.Kind
	out.Name = in.Name
	out.APIVersion = in.APIVersion
	out.Subresource = in.Subresource
	return nil
}

func DeepCopy_v1beta1_ThirdPartyResource(in ThirdPartyResource, out *ThirdPartyResource, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	out.Description = in.Description
	if in.Versions != nil {
		in, out := in.Versions, &out.Versions
		*out = make([]APIVersion, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_APIVersion(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Versions = nil
	}
	return nil
}

func DeepCopy_v1beta1_ThirdPartyResourceData(in ThirdPartyResourceData, out *ThirdPartyResourceData, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := v1.DeepCopy_v1_ObjectMeta(in.ObjectMeta, &out.ObjectMeta, c); err != nil {
		return err
	}
	if in.Data != nil {
		in, out := in.Data, &out.Data
		*out = make([]byte, len(in))
		copy(*out, in)
	} else {
		out.Data = nil
	}
	return nil
}

func DeepCopy_v1beta1_ThirdPartyResourceDataList(in ThirdPartyResourceDataList, out *ThirdPartyResourceDataList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]ThirdPartyResourceData, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_ThirdPartyResourceData(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}

func DeepCopy_v1beta1_ThirdPartyResourceList(in ThirdPartyResourceList, out *ThirdPartyResourceList, c *conversion.Cloner) error {
	if err := unversioned.DeepCopy_unversioned_TypeMeta(in.TypeMeta, &out.TypeMeta, c); err != nil {
		return err
	}
	if err := unversioned.DeepCopy_unversioned_ListMeta(in.ListMeta, &out.ListMeta, c); err != nil {
		return err
	}
	if in.Items != nil {
		in, out := in.Items, &out.Items
		*out = make([]ThirdPartyResource, len(in))
		for i := range in {
			if err := DeepCopy_v1beta1_ThirdPartyResource(in[i], &(*out)[i], c); err != nil {
				return err
			}
		}
	} else {
		out.Items = nil
	}
	return nil
}
