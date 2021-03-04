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

package internalversion

import (
	"context"
	time "time"

	operators "github.com/openshift/operator-framework-olm/staging/operator-lifecycle-manager/pkg/package-server/apis/operators"
	clientsetinternalversion "github.com/openshift/operator-framework-olm/staging/operator-lifecycle-manager/pkg/package-server/client/clientset/internalversion"
	internalinterfaces "github.com/openshift/operator-framework-olm/staging/operator-lifecycle-manager/pkg/package-server/client/informers/internalversion/internalinterfaces"
	internalversion "github.com/openshift/operator-framework-olm/staging/operator-lifecycle-manager/pkg/package-server/client/listers/operators/internalversion"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// PackageManifestInformer provides access to a shared informer and lister for
// PackageManifests.
type PackageManifestInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() internalversion.PackageManifestLister
}

type packageManifestInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewPackageManifestInformer constructs a new informer for PackageManifest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewPackageManifestInformer(client clientsetinternalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredPackageManifestInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredPackageManifestInformer constructs a new informer for PackageManifest type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredPackageManifestInformer(client clientsetinternalversion.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Operators().PackageManifests(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.Operators().PackageManifests(namespace).Watch(context.TODO(), options)
			},
		},
		&operators.PackageManifest{},
		resyncPeriod,
		indexers,
	)
}

func (f *packageManifestInformer) defaultInformer(client clientsetinternalversion.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredPackageManifestInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *packageManifestInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&operators.PackageManifest{}, f.defaultInformer)
}

func (f *packageManifestInformer) Lister() internalversion.PackageManifestLister {
	return internalversion.NewPackageManifestLister(f.Informer().GetIndexer())
}
