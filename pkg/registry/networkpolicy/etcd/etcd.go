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

package etcd

import (
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/fields"
	"k8s.io/kubernetes/pkg/labels"
	"k8s.io/kubernetes/pkg/registry/cachesize"
	"k8s.io/kubernetes/pkg/registry/generic"
	etcdgeneric "k8s.io/kubernetes/pkg/registry/generic/etcd"
	"k8s.io/kubernetes/pkg/registry/networkpolicy"
	"k8s.io/kubernetes/pkg/runtime"
)

// NetworkPolicyStorage includes dummy storage for NetworkPolicys.
type NetworkPolicyStorage struct {
	NetworkPolicy *REST
}

// REST implements a RESTStorage for PodSecurityPolicies against etcd.
type REST struct {
	*etcdgeneric.Etcd
}

func NewStorage(opts generic.RESTOptions) NetworkPolicyStorage {
	networkPolicyRest := NewREST(opts)
	//networkPolicyRegistry := networkpolicy.NewRegistry(networkPolicyRest)

	return NetworkPolicyStorage{
		NetworkPolicy: networkPolicyRest,
	}
}

type REST struct {
	*etcdgeneric.Etcd
}

// NewREST returns a RESTStorage object that will work against NetworkPolicy.
func NewREST(opts generic.RESTOptions) *REST {
	prefix := "/networkpolicys"

	newListFunc := func() runtime.Object { return &extensions.NetworkPolicyList{} }
	storageInterface := opts.Decorator(
		opts.Storage, cachesize.GetWatchCacheSizeByResource(cachesize.NetworkPolicys), &extensions.NetworkPolicy{}, prefix, networkpolicy.Strategy, newListFunc)

	store := &etcdgeneric.Etcd{
		NewFunc: func() runtime.Object { return &extensions.NetworkPolicy{} },

		// NewListFunc returns an object capable of storing results of an etcd list.
		NewListFunc: newListFunc,
		// Produces a path that etcd understands, to the root of the resource
		// by combining the namespace in the context with the given prefix
		KeyRootFunc: func(ctx api.Context) string {
			return etcdgeneric.NamespaceKeyRootFunc(ctx, prefix)
		},
		// Produces a path that etcd understands, to the resource by combining
		// the namespace in the context with the given prefix
		KeyFunc: func(ctx api.Context, name string) (string, error) {
			return etcdgeneric.NamespaceKeyFunc(ctx, prefix, name)
		},
		// Retrieve the name field of a NetworkPolicy
		ObjectNameFunc: func(obj runtime.Object) (string, error) {
			return obj.(*extensions.NetworkPolicy).Name, nil
		},
		// Used to match objects based on labels/fields for list and watch
		PredicateFunc: func(label labels.Selector, field fields.Selector) generic.Matcher {
			return networkpolicy.MatchNetworkPolicy(label, field)
		},
		QualifiedResource:       api.Resource("networkpolicys"),
		DeleteCollectionWorkers: opts.DeleteCollectionWorkers,

		// Used to validate NetworkPolicy creation
		CreateStrategy: networkpolicy.Strategy,

		// Used to validate NetworkPolicy updates
		UpdateStrategy: networkpolicy.Strategy,
		DeleteStrategy: networkpolicy.Strategy,

		Storage: storageInterface,
	}

	return &REST{store}
}
