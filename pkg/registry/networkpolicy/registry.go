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
	"fmt"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/rest"
	"k8s.io/kubernetes/pkg/apis/extensions"
	"k8s.io/kubernetes/pkg/watch"
)

// Registry is an interface for things that know how to store NetworkPolicys.
type Registry interface {
	ListNetworkPolicys(ctx api.Context, options *api.ListOptions) (*extensions.NetworkPolicyList, error)
	WatchNetworkPolicys(ctx api.Context, options *api.ListOptions) (watch.Interface, error)
	GetNetworkPolicy(ctx api.Context, networkPolicyID string) (*extensions.NetworkPolicy, error)
	CreateNetworkPolicy(ctx api.Context, networkPolicy *extensions.NetworkPolicy) (*extensions.NetworkPolicy, error)
	UpdateNetworkPolicy(ctx api.Context, networkPolicy *extensions.NetworkPolicy) (*extensions.NetworkPolicy, error)
	DeleteNetworkPolicy(ctx api.Context, networkPolicyID string) error
}

// storage puts strong typing around storage calls
type storage struct {
	rest.StandardStorage
}

// NewRegistry returns a new Registry interface for the given Storage. Any mismatched
// types will panic.
func NewRegistry(s rest.StandardStorage) Registry {
	return &storage{s}
}

func (s *storage) ListNetworkPolicys(ctx api.Context, options *api.ListOptions) (*extensions.NetworkPolicyList, error) {
	if options != nil && options.FieldSelector != nil && !options.FieldSelector.Empty() {
		return nil, fmt.Errorf("field selector not supported yet")
	}
	obj, err := s.List(ctx, options)
	if err != nil {
		return nil, err
	}
	return obj.(*extensions.NetworkPolicyList), err
}

func (s *storage) WatchNetworkPolicys(ctx api.Context, options *api.ListOptions) (watch.Interface, error) {
	return s.Watch(ctx, options)
}

func (s *storage) GetNetworkPolicy(ctx api.Context, networkPolicyID string) (*extensions.NetworkPolicy, error) {
	obj, err := s.Get(ctx, networkPolicyID)
	if err != nil {
		return nil, err
	}
	return obj.(*extensions.NetworkPolicy), nil
}

func (s *storage) CreateNetworkPolicy(ctx api.Context, networkPolicy *extensions.NetworkPolicy) (*extensions.NetworkPolicy, error) {
	obj, err := s.Create(ctx, networkPolicy)
	if err != nil {
		return nil, err
	}
	return obj.(*extensions.NetworkPolicy), nil
}

func (s *storage) UpdateNetworkPolicy(ctx api.Context, networkPolicy *extensions.NetworkPolicy) (*extensions.NetworkPolicy, error) {
	obj, _, err := s.Update(ctx, networkPolicy)
	if err != nil {
		return nil, err
	}
	return obj.(*extensions.NetworkPolicy), nil
}

func (s *storage) DeleteNetworkPolicy(ctx api.Context, networkPolicyID string) error {
	_, err := s.Delete(ctx, networkPolicyID, nil)
	return err
}
