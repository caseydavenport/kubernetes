/*
Copyright 2015 The Kubernetes Authors All rights reserved.

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

package unversioned

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/network"
	"k8s.io/kubernetes/pkg/watch"
)

// NetworkPolicyNamespacer has methods to work with NetworkPolicy resources in a namespace
type NetworkPolicyNamespacer interface {
	NetworkPolicys(namespace string) NetworkPolicyInterface
}

// NetworkPolicyInterface exposes methods to work on NetworkPolicy resources.
type NetworkPolicyInterface interface {
	List(opts api.ListOptions) (*network.NetworkPolicyList, error)
	Get(name string) (*network.NetworkPolicy, error)
	Create(networkPolicy *network.NetworkPolicy) (*network.NetworkPolicy, error)
	Update(networkPolicy *network.NetworkPolicy) (*network.NetworkPolicy, error)
	Delete(name string, options *api.DeleteOptions) error
	Watch(opts api.ListOptions) (watch.Interface, error)
	UpdateStatus(networkPolicy *network.NetworkPolicy) (*network.NetworkPolicy, error)
}

// networkPolicy implements NetworkPolicyNamespacer interface
type networkPolicy struct {
	r  *NetworkClient
	ns string
}

// newNetworkPolicy returns a networkPolicy
func newNetworkPolicy(c *NetworkClient, namespace string) *networkPolicy {
	return &networkPolicy{c, namespace}
}

// List returns a list of networkPolicy that match the label and field selectors.
func (c *networkPolicy) List(opts api.ListOptions) (result *network.NetworkPolicyList, err error) {
	result = &network.NetworkPolicyList{}
	err = c.r.Get().Namespace(c.ns).Resource("networkpolicys").VersionedParams(&opts, api.ParameterCodec).Do().Into(result)
	return
}

// Get returns information about a particular networkPolicy.
func (c *networkPolicy) Get(name string) (result *network.NetworkPolicy, err error) {
	result = &network.NetworkPolicy{}
	err = c.r.Get().Namespace(c.ns).Resource("networkpolicys").Name(name).Do().Into(result)
	return
}

// Create creates a new networkPolicy.
func (c *networkPolicy) Create(networkPolicy *network.NetworkPolicy) (result *network.NetworkPolicy, err error) {
	result = &network.NetworkPolicy{}
	err = c.r.Post().Namespace(c.ns).Resource("networkpolicys").Body(networkPolicy).Do().Into(result)
	return
}

// Update updates an existing networkPolicy.
func (c *networkPolicy) Update(networkPolicy *network.NetworkPolicy) (result *network.NetworkPolicy, err error) {
	result = &network.NetworkPolicy{}
	err = c.r.Put().Namespace(c.ns).Resource("networkpolicys").Name(networkPolicy.Name).Body(networkPolicy).Do().Into(result)
	return
}

// Delete deletes a networkPolicy, returns error if one occurs.
func (c *networkPolicy) Delete(name string, options *api.DeleteOptions) (err error) {
	return c.r.Delete().Namespace(c.ns).Resource("networkpolicys").Name(name).Body(options).Do().Error()
}

// Watch returns a watch.Interface that watches the requested networkPolicy.
func (c *networkPolicy) Watch(opts api.ListOptions) (watch.Interface, error) {
	return c.r.Get().
		Prefix("watch").
		Namespace(c.ns).
		Resource("networkpolicys").
		VersionedParams(&opts, api.ParameterCodec).
		Watch()
}

// UpdateStatus takes the name of the networkPolicy and the new status.  Returns the server's representation of the networkPolicy, and an error, if it occurs.
func (c *networkPolicy) UpdateStatus(networkPolicy *network.NetworkPolicy) (result *network.NetworkPolicy, err error) {
	result = &network.NetworkPolicy{}
	err = c.r.Put().Namespace(c.ns).Resource("networkpolicys").Name(networkPolicy.Name).SubResource("status").Body(networkPolicy).Do().Into(result)
	return
}
