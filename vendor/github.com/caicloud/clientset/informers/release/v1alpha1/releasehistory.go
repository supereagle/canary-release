/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	kubernetes "github.com/caicloud/clientset/kubernetes"
	v1alpha1 "github.com/caicloud/clientset/listers/release/v1alpha1"
	releasev1alpha1 "github.com/caicloud/clientset/pkg/apis/release/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	internalinterfaces "k8s.io/client-go/informers/internalinterfaces"
	clientgokubernetes "k8s.io/client-go/kubernetes"
	cache "k8s.io/client-go/tools/cache"
)

// ReleaseHistoryInformer provides access to a shared informer and lister for
// ReleaseHistories.
type ReleaseHistoryInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ReleaseHistoryLister
}

type releaseHistoryInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewReleaseHistoryInformer constructs a new informer for ReleaseHistory type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewReleaseHistoryInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredReleaseHistoryInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredReleaseHistoryInformer constructs a new informer for ReleaseHistory type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredReleaseHistoryInformer(client kubernetes.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ReleaseV1alpha1().ReleaseHistories(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.ReleaseV1alpha1().ReleaseHistories(namespace).Watch(options)
			},
		},
		&releasev1alpha1.ReleaseHistory{},
		resyncPeriod,
		indexers,
	)
}

func (f *releaseHistoryInformer) defaultInformer(client clientgokubernetes.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredReleaseHistoryInformer(client.(kubernetes.Interface), f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *releaseHistoryInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&releasev1alpha1.ReleaseHistory{}, f.defaultInformer)
}

func (f *releaseHistoryInformer) Lister() v1alpha1.ReleaseHistoryLister {
	return v1alpha1.NewReleaseHistoryLister(f.Informer().GetIndexer())
}
