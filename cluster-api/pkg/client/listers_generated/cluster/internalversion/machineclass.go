/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package internalversion

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	cluster "k8s.io/kube-deploy/cluster-api/pkg/apis/cluster"
)

// MachineClassLister helps list MachineClasses.
type MachineClassLister interface {
	// List lists all MachineClasses in the indexer.
	List(selector labels.Selector) (ret []*cluster.MachineClass, err error)
	// MachineClasses returns an object that can list and get MachineClasses.
	MachineClasses(namespace string) MachineClassNamespaceLister
	MachineClassListerExpansion
}

// machineClassLister implements the MachineClassLister interface.
type machineClassLister struct {
	indexer cache.Indexer
}

// NewMachineClassLister returns a new MachineClassLister.
func NewMachineClassLister(indexer cache.Indexer) MachineClassLister {
	return &machineClassLister{indexer: indexer}
}

// List lists all MachineClasses in the indexer.
func (s *machineClassLister) List(selector labels.Selector) (ret []*cluster.MachineClass, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*cluster.MachineClass))
	})
	return ret, err
}

// MachineClasses returns an object that can list and get MachineClasses.
func (s *machineClassLister) MachineClasses(namespace string) MachineClassNamespaceLister {
	return machineClassNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MachineClassNamespaceLister helps list and get MachineClasses.
type MachineClassNamespaceLister interface {
	// List lists all MachineClasses in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*cluster.MachineClass, err error)
	// Get retrieves the MachineClass from the indexer for a given namespace and name.
	Get(name string) (*cluster.MachineClass, error)
	MachineClassNamespaceListerExpansion
}

// machineClassNamespaceLister implements the MachineClassNamespaceLister
// interface.
type machineClassNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MachineClasses in the indexer for a given namespace.
func (s machineClassNamespaceLister) List(selector labels.Selector) (ret []*cluster.MachineClass, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*cluster.MachineClass))
	})
	return ret, err
}

// Get retrieves the MachineClass from the indexer for a given namespace and name.
func (s machineClassNamespaceLister) Get(name string) (*cluster.MachineClass, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(cluster.Resource("machineclass"), name)
	}
	return obj.(*cluster.MachineClass), nil
}
