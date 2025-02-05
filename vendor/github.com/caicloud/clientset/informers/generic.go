/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by informer-gen. DO NOT EDIT.

package informers

import (
	v1beta1 "github.com/caicloud/clientset/pkg/apis/apiextensions/v1beta1"
	v1 "github.com/caicloud/clientset/pkg/apis/apiregistration/v1"
	v1alpha1 "github.com/caicloud/clientset/pkg/apis/clever/v1alpha1"
	cnetworkingv1alpha1 "github.com/caicloud/clientset/pkg/apis/cnetworking/v1alpha1"
	configv1alpha1 "github.com/caicloud/clientset/pkg/apis/config/v1alpha1"
	datasetv1alpha1 "github.com/caicloud/clientset/pkg/apis/dataset/v1alpha1"
	devopsv1 "github.com/caicloud/clientset/pkg/apis/devops/v1"
	v1alpha2 "github.com/caicloud/clientset/pkg/apis/loadbalance/v1alpha2"
	loggingv1alpha1 "github.com/caicloud/clientset/pkg/apis/logging/v1alpha1"
	microservicev1alpha1 "github.com/caicloud/clientset/pkg/apis/microservice/v1alpha1"
	modelv1alpha1 "github.com/caicloud/clientset/pkg/apis/model/v1alpha1"
	orchestrationv1alpha1 "github.com/caicloud/clientset/pkg/apis/orchestration/v1alpha1"
	releasev1alpha1 "github.com/caicloud/clientset/pkg/apis/release/v1alpha1"
	resourcev1alpha1 "github.com/caicloud/clientset/pkg/apis/resource/v1alpha1"
	resourcev1beta1 "github.com/caicloud/clientset/pkg/apis/resource/v1beta1"
	tenantv1alpha1 "github.com/caicloud/clientset/pkg/apis/tenant/v1alpha1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	informers "k8s.io/client-go/informers"
	cache "k8s.io/client-go/tools/cache"
)

type genericInformer struct {
	informer cache.SharedIndexInformer
	resource schema.GroupResource
}

// Informer returns the SharedIndexInformer.
func (f *genericInformer) Informer() cache.SharedIndexInformer {
	return f.informer
}

// Lister returns the GenericLister.
func (f *genericInformer) Lister() cache.GenericLister {
	return cache.NewGenericLister(f.Informer().GetIndexer(), f.resource)
}

// ForResource gives generic access to a shared informer of the matching type
// TODO extend this to unknown resources with a client pool
func (f *sharedInformerFactory) ForResource(resource schema.GroupVersionResource) (informers.GenericInformer, error) {
	switch resource {
	// Group=apiextensions.k8s.io, Version=v1beta1
	case v1beta1.SchemeGroupVersion.WithResource("customresourcedefinitions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apiextensions().V1beta1().CustomResourceDefinitions().Informer()}, nil

		// Group=apiregistration.k8s.io, Version=v1
	case v1.SchemeGroupVersion.WithResource("apiservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Apiregistration().V1().APIServices().Informer()}, nil

		// Group=clever.caicloud.io, Version=v1alpha1
	case v1alpha1.SchemeGroupVersion.WithResource("flavors"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Clever().V1alpha1().Flavors().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("projects"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Clever().V1alpha1().Projects().Informer()}, nil
	case v1alpha1.SchemeGroupVersion.WithResource("templates"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Clever().V1alpha1().Templates().Informer()}, nil

		// Group=cnetworking.caicloud.io, Version=v1alpha1
	case cnetworkingv1alpha1.SchemeGroupVersion.WithResource("networkpolicies"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Cnetworking().V1alpha1().NetworkPolicies().Informer()}, nil

		// Group=config.caicloud.io, Version=v1alpha1
	case configv1alpha1.SchemeGroupVersion.WithResource("configclaims"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1alpha1().ConfigClaims().Informer()}, nil
	case configv1alpha1.SchemeGroupVersion.WithResource("configreferences"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Config().V1alpha1().ConfigReferences().Informer()}, nil

		// Group=dataset.caicloud.io, Version=v1alpha1
	case datasetv1alpha1.SchemeGroupVersion.WithResource("datasets"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Dataset().V1alpha1().Datasets().Informer()}, nil

		// Group=devops.caicloud.io, Version=v1
	case devopsv1.SchemeGroupVersion.WithResource("cargos"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Devops().V1().Cargos().Informer()}, nil

		// Group=loadbalance.caicloud.io, Version=v1alpha2
	case v1alpha2.SchemeGroupVersion.WithResource("loadbalancers"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Loadbalance().V1alpha2().LoadBalancers().Informer()}, nil

		// Group=logging.caicloud.io, Version=v1alpha1
	case loggingv1alpha1.SchemeGroupVersion.WithResource("logendpoints"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Logging().V1alpha1().LogEndpoints().Informer()}, nil

		// Group=microservice.caicloud.io, Version=v1alpha1
	case microservicev1alpha1.SchemeGroupVersion.WithResource("springclouds"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Microservice().V1alpha1().Springclouds().Informer()}, nil

		// Group=model.caicloud.io, Version=v1alpha1
	case modelv1alpha1.SchemeGroupVersion.WithResource("models"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Model().V1alpha1().Models().Informer()}, nil

		// Group=orchestration.caicloud.io, Version=v1alpha1
	case orchestrationv1alpha1.SchemeGroupVersion.WithResource("applications"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Orchestration().V1alpha1().Applications().Informer()}, nil
	case orchestrationv1alpha1.SchemeGroupVersion.WithResource("applicationdrafts"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Orchestration().V1alpha1().ApplicationDrafts().Informer()}, nil

		// Group=release.caicloud.io, Version=v1alpha1
	case releasev1alpha1.SchemeGroupVersion.WithResource("canaryreleases"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Release().V1alpha1().CanaryReleases().Informer()}, nil
	case releasev1alpha1.SchemeGroupVersion.WithResource("releases"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Release().V1alpha1().Releases().Informer()}, nil
	case releasev1alpha1.SchemeGroupVersion.WithResource("releasehistories"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Release().V1alpha1().ReleaseHistories().Informer()}, nil

		// Group=resource.caicloud.io, Version=v1alpha1
	case resourcev1alpha1.SchemeGroupVersion.WithResource("storageservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Resource().V1alpha1().StorageServices().Informer()}, nil
	case resourcev1alpha1.SchemeGroupVersion.WithResource("storagetypes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Resource().V1alpha1().StorageTypes().Informer()}, nil

		// Group=resource.caicloud.io, Version=v1beta1
	case resourcev1beta1.SchemeGroupVersion.WithResource("clusters"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Resource().V1beta1().Clusters().Informer()}, nil
	case resourcev1beta1.SchemeGroupVersion.WithResource("configs"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Resource().V1beta1().Configs().Informer()}, nil
	case resourcev1beta1.SchemeGroupVersion.WithResource("machines"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Resource().V1beta1().Machines().Informer()}, nil
	case resourcev1beta1.SchemeGroupVersion.WithResource("machineautoscalinggroups"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Resource().V1beta1().MachineAutoScalingGroups().Informer()}, nil
	case resourcev1beta1.SchemeGroupVersion.WithResource("requirementgaps"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Resource().V1beta1().RequirementGaps().Informer()}, nil
	case resourcev1beta1.SchemeGroupVersion.WithResource("snapshots"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Resource().V1beta1().Snapshots().Informer()}, nil
	case resourcev1beta1.SchemeGroupVersion.WithResource("storageservices"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Resource().V1beta1().StorageServices().Informer()}, nil
	case resourcev1beta1.SchemeGroupVersion.WithResource("storagetypes"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Resource().V1beta1().StorageTypes().Informer()}, nil
	case resourcev1beta1.SchemeGroupVersion.WithResource("tags"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Resource().V1beta1().Tags().Informer()}, nil

		// Group=tenant.caicloud.io, Version=v1alpha1
	case tenantv1alpha1.SchemeGroupVersion.WithResource("clusterquotas"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Tenant().V1alpha1().ClusterQuotas().Informer()}, nil
	case tenantv1alpha1.SchemeGroupVersion.WithResource("partitions"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Tenant().V1alpha1().Partitions().Informer()}, nil
	case tenantv1alpha1.SchemeGroupVersion.WithResource("tenants"):
		return &genericInformer{resource: resource.GroupResource(), informer: f.Tenant().V1alpha1().Tenants().Informer()}, nil

	}

	return f.SharedInformerFactory.ForResource(resource)
}
