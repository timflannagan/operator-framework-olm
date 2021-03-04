/*
Copyright Red Hat, Inc.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	operatorsv1 "github.com/openshift/operator-framework-olm/staging/api/pkg/operators/v1"
	versioned "github.com/openshift/operator-framework-olm/staging/operator-lifecycle-manager/pkg/api/client/clientset/versioned"
	internalinterfaces "github.com/openshift/operator-framework-olm/staging/operator-lifecycle-manager/pkg/api/client/informers/externalversions/internalinterfaces"
	v1 "github.com/openshift/operator-framework-olm/staging/operator-lifecycle-manager/pkg/api/client/listers/operators/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// OperatorInformer provides access to a shared informer and lister for
// Operators.
type OperatorInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.OperatorLister
}

type operatorInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewOperatorInformer constructs a new informer for Operator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewOperatorInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredOperatorInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredOperatorInformer constructs a new informer for Operator type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredOperatorInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorsV1().Operators().List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OperatorsV1().Operators().Watch(context.TODO(), options)
			},
		},
		&operatorsv1.Operator{},
		resyncPeriod,
		indexers,
	)
}

func (f *operatorInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredOperatorInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *operatorInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operatorsv1.Operator{}, f.defaultInformer)
}

func (f *operatorInformer) Lister() v1.OperatorLister {
	return v1.NewOperatorLister(f.Informer().GetIndexer())
}
