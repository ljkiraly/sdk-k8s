// Copyright (c) 2020-2022 Doc.ai and/or its affiliates.
//
// Copyright (c) 2020-2022 Cisco and/or its affiliates.
//
// Copyright (c) 2020-2022 VMware, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at:
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"

	v1 "github.com/ljkiraly/sdk-k8s/pkg/tools/k8s/apis/networkservicemesh.io/v1"
)

// NetworkServiceLister helps list NetworkServices.
// All objects returned here must be treated as read-only.
type NetworkServiceLister interface {
	// List lists all NetworkServices in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.NetworkService, err error)
	// NetworkServices returns an object that can list and get NetworkServices.
	NetworkServices(namespace string) NetworkServiceNamespaceLister
	NetworkServiceListerExpansion
}

// networkServiceLister implements the NetworkServiceLister interface.
type networkServiceLister struct {
	indexer cache.Indexer
}

// NewNetworkServiceLister returns a new NetworkServiceLister.
func NewNetworkServiceLister(indexer cache.Indexer) NetworkServiceLister {
	return &networkServiceLister{indexer: indexer}
}

// List lists all NetworkServices in the indexer.
func (s *networkServiceLister) List(selector labels.Selector) (ret []*v1.NetworkService, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NetworkService))
	})
	return ret, err
}

// NetworkServices returns an object that can list and get NetworkServices.
func (s *networkServiceLister) NetworkServices(namespace string) NetworkServiceNamespaceLister {
	return networkServiceNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// NetworkServiceNamespaceLister helps list and get NetworkServices.
// All objects returned here must be treated as read-only.
type NetworkServiceNamespaceLister interface {
	// List lists all NetworkServices in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.NetworkService, err error)
	// Get retrieves the NetworkService from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.NetworkService, error)
	NetworkServiceNamespaceListerExpansion
}

// networkServiceNamespaceLister implements the NetworkServiceNamespaceLister
// interface.
type networkServiceNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all NetworkServices in the indexer for a given namespace.
func (s networkServiceNamespaceLister) List(selector labels.Selector) (ret []*v1.NetworkService, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.NetworkService))
	})
	return ret, err
}

// Get retrieves the NetworkService from the indexer for a given namespace and name.
func (s networkServiceNamespaceLister) Get(name string) (*v1.NetworkService, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("networkservice"), name)
	}
	return obj.(*v1.NetworkService), nil
}
