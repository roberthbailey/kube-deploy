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

// This file was automatically generated by informer-gen

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
	cluster_v1alpha1 "k8s.io/kube-deploy/cluster-api/pkg/apis/cluster/v1alpha1"
	clientset "k8s.io/kube-deploy/cluster-api/pkg/client/clientset_generated/clientset"
	internalinterfaces "k8s.io/kube-deploy/cluster-api/pkg/client/informers_generated/externalversions/internalinterfaces"
	v1alpha1 "k8s.io/kube-deploy/cluster-api/pkg/client/listers_generated/cluster/v1alpha1"
	time "time"
)

// MachineClassInformer provides access to a shared informer and lister for
// MachineClasses.
type MachineClassInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.MachineClassLister
}

type machineClassInformer struct {
	factory internalinterfaces.SharedInformerFactory
}

// NewMachineClassInformer constructs a new informer for MachineClass type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewMachineClassInformer(client clientset.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				return client.ClusterV1alpha1().MachineClasses(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				return client.ClusterV1alpha1().MachineClasses(namespace).Watch(options)
			},
		},
		&cluster_v1alpha1.MachineClass{},
		resyncPeriod,
		indexers,
	)
}

func defaultMachineClassInformer(client clientset.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewMachineClassInformer(client, v1.NamespaceAll, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc})
}

func (f *machineClassInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cluster_v1alpha1.MachineClass{}, defaultMachineClassInformer)
}

func (f *machineClassInformer) Lister() v1alpha1.MachineClassLister {
	return v1alpha1.NewMachineClassLister(f.Informer().GetIndexer())
}
