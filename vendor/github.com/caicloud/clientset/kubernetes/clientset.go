/*
Copyright 2018 caicloud authors. All rights reserved.
*/

// Code generated by client-gen. DO NOT EDIT.

package kubernetes

import (
	apiextensionsv1beta1 "github.com/caicloud/clientset/kubernetes/typed/apiextensions/v1beta1"
	apiregistrationv1 "github.com/caicloud/clientset/kubernetes/typed/apiregistration/v1"
	cleverv1alpha1 "github.com/caicloud/clientset/kubernetes/typed/clever/v1alpha1"
	cnetworkingv1alpha1 "github.com/caicloud/clientset/kubernetes/typed/cnetworking/v1alpha1"
	configv1alpha1 "github.com/caicloud/clientset/kubernetes/typed/config/v1alpha1"
	datasetv1alpha1 "github.com/caicloud/clientset/kubernetes/typed/dataset/v1alpha1"
	devopsv1 "github.com/caicloud/clientset/kubernetes/typed/devops/v1"
	loadbalancev1alpha2 "github.com/caicloud/clientset/kubernetes/typed/loadbalance/v1alpha2"
	loggingv1alpha1 "github.com/caicloud/clientset/kubernetes/typed/logging/v1alpha1"
	microservicev1alpha1 "github.com/caicloud/clientset/kubernetes/typed/microservice/v1alpha1"
	modelv1alpha1 "github.com/caicloud/clientset/kubernetes/typed/model/v1alpha1"
	orchestrationv1alpha1 "github.com/caicloud/clientset/kubernetes/typed/orchestration/v1alpha1"
	releasev1alpha1 "github.com/caicloud/clientset/kubernetes/typed/release/v1alpha1"
	resourcev1alpha1 "github.com/caicloud/clientset/kubernetes/typed/resource/v1alpha1"
	resourcev1beta1 "github.com/caicloud/clientset/kubernetes/typed/resource/v1beta1"
	tenantv1alpha1 "github.com/caicloud/clientset/kubernetes/typed/tenant/v1alpha1"
	kubernetes "k8s.io/client-go/kubernetes"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	kubernetes.Interface
	ApiextensionsV1beta1() apiextensionsv1beta1.ApiextensionsV1beta1Interface
	// Deprecated: please explicitly pick a version if possible.
	Apiextensions() apiextensionsv1beta1.ApiextensionsV1beta1Interface
	ApiregistrationV1() apiregistrationv1.ApiregistrationV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Apiregistration() apiregistrationv1.ApiregistrationV1Interface
	CleverV1alpha1() cleverv1alpha1.CleverV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Clever() cleverv1alpha1.CleverV1alpha1Interface
	CnetworkingV1alpha1() cnetworkingv1alpha1.CnetworkingV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Cnetworking() cnetworkingv1alpha1.CnetworkingV1alpha1Interface
	ConfigV1alpha1() configv1alpha1.ConfigV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Config() configv1alpha1.ConfigV1alpha1Interface
	DatasetV1alpha1() datasetv1alpha1.DatasetV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Dataset() datasetv1alpha1.DatasetV1alpha1Interface
	DevopsV1() devopsv1.DevopsV1Interface
	// Deprecated: please explicitly pick a version if possible.
	Devops() devopsv1.DevopsV1Interface
	LoadbalanceV1alpha2() loadbalancev1alpha2.LoadbalanceV1alpha2Interface
	// Deprecated: please explicitly pick a version if possible.
	Loadbalance() loadbalancev1alpha2.LoadbalanceV1alpha2Interface
	LoggingV1alpha1() loggingv1alpha1.LoggingV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Logging() loggingv1alpha1.LoggingV1alpha1Interface
	MicroserviceV1alpha1() microservicev1alpha1.MicroserviceV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Microservice() microservicev1alpha1.MicroserviceV1alpha1Interface
	ModelV1alpha1() modelv1alpha1.ModelV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Model() modelv1alpha1.ModelV1alpha1Interface
	OrchestrationV1alpha1() orchestrationv1alpha1.OrchestrationV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Orchestration() orchestrationv1alpha1.OrchestrationV1alpha1Interface
	ReleaseV1alpha1() releasev1alpha1.ReleaseV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Release() releasev1alpha1.ReleaseV1alpha1Interface
	ResourceV1alpha1() resourcev1alpha1.ResourceV1alpha1Interface
	ResourceV1beta1() resourcev1beta1.ResourceV1beta1Interface
	// Deprecated: please explicitly pick a version if possible.
	Resource() resourcev1beta1.ResourceV1beta1Interface
	TenantV1alpha1() tenantv1alpha1.TenantV1alpha1Interface
	// Deprecated: please explicitly pick a version if possible.
	Tenant() tenantv1alpha1.TenantV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*kubernetes.Clientset
	apiextensionsV1beta1  *apiextensionsv1beta1.ApiextensionsV1beta1Client
	apiregistrationV1     *apiregistrationv1.ApiregistrationV1Client
	cleverV1alpha1        *cleverv1alpha1.CleverV1alpha1Client
	cnetworkingV1alpha1   *cnetworkingv1alpha1.CnetworkingV1alpha1Client
	configV1alpha1        *configv1alpha1.ConfigV1alpha1Client
	datasetV1alpha1       *datasetv1alpha1.DatasetV1alpha1Client
	devopsV1              *devopsv1.DevopsV1Client
	loadbalanceV1alpha2   *loadbalancev1alpha2.LoadbalanceV1alpha2Client
	loggingV1alpha1       *loggingv1alpha1.LoggingV1alpha1Client
	microserviceV1alpha1  *microservicev1alpha1.MicroserviceV1alpha1Client
	modelV1alpha1         *modelv1alpha1.ModelV1alpha1Client
	orchestrationV1alpha1 *orchestrationv1alpha1.OrchestrationV1alpha1Client
	releaseV1alpha1       *releasev1alpha1.ReleaseV1alpha1Client
	resourceV1alpha1      *resourcev1alpha1.ResourceV1alpha1Client
	resourceV1beta1       *resourcev1beta1.ResourceV1beta1Client
	tenantV1alpha1        *tenantv1alpha1.TenantV1alpha1Client
}

// ApiextensionsV1beta1 retrieves the ApiextensionsV1beta1Client
func (c *Clientset) ApiextensionsV1beta1() apiextensionsv1beta1.ApiextensionsV1beta1Interface {
	return c.apiextensionsV1beta1
}

// Deprecated: Apiextensions retrieves the default version of ApiextensionsClient.
// Please explicitly pick a version.
func (c *Clientset) Apiextensions() apiextensionsv1beta1.ApiextensionsV1beta1Interface {
	return c.apiextensionsV1beta1
}

// ApiregistrationV1 retrieves the ApiregistrationV1Client
func (c *Clientset) ApiregistrationV1() apiregistrationv1.ApiregistrationV1Interface {
	return c.apiregistrationV1
}

// Deprecated: Apiregistration retrieves the default version of ApiregistrationClient.
// Please explicitly pick a version.
func (c *Clientset) Apiregistration() apiregistrationv1.ApiregistrationV1Interface {
	return c.apiregistrationV1
}

// CleverV1alpha1 retrieves the CleverV1alpha1Client
func (c *Clientset) CleverV1alpha1() cleverv1alpha1.CleverV1alpha1Interface {
	return c.cleverV1alpha1
}

// Deprecated: Clever retrieves the default version of CleverClient.
// Please explicitly pick a version.
func (c *Clientset) Clever() cleverv1alpha1.CleverV1alpha1Interface {
	return c.cleverV1alpha1
}

// CnetworkingV1alpha1 retrieves the CnetworkingV1alpha1Client
func (c *Clientset) CnetworkingV1alpha1() cnetworkingv1alpha1.CnetworkingV1alpha1Interface {
	return c.cnetworkingV1alpha1
}

// Deprecated: Cnetworking retrieves the default version of CnetworkingClient.
// Please explicitly pick a version.
func (c *Clientset) Cnetworking() cnetworkingv1alpha1.CnetworkingV1alpha1Interface {
	return c.cnetworkingV1alpha1
}

// ConfigV1alpha1 retrieves the ConfigV1alpha1Client
func (c *Clientset) ConfigV1alpha1() configv1alpha1.ConfigV1alpha1Interface {
	return c.configV1alpha1
}

// Deprecated: Config retrieves the default version of ConfigClient.
// Please explicitly pick a version.
func (c *Clientset) Config() configv1alpha1.ConfigV1alpha1Interface {
	return c.configV1alpha1
}

// DatasetV1alpha1 retrieves the DatasetV1alpha1Client
func (c *Clientset) DatasetV1alpha1() datasetv1alpha1.DatasetV1alpha1Interface {
	return c.datasetV1alpha1
}

// Deprecated: Dataset retrieves the default version of DatasetClient.
// Please explicitly pick a version.
func (c *Clientset) Dataset() datasetv1alpha1.DatasetV1alpha1Interface {
	return c.datasetV1alpha1
}

// DevopsV1 retrieves the DevopsV1Client
func (c *Clientset) DevopsV1() devopsv1.DevopsV1Interface {
	return c.devopsV1
}

// Deprecated: Devops retrieves the default version of DevopsClient.
// Please explicitly pick a version.
func (c *Clientset) Devops() devopsv1.DevopsV1Interface {
	return c.devopsV1
}

// LoadbalanceV1alpha2 retrieves the LoadbalanceV1alpha2Client
func (c *Clientset) LoadbalanceV1alpha2() loadbalancev1alpha2.LoadbalanceV1alpha2Interface {
	return c.loadbalanceV1alpha2
}

// Deprecated: Loadbalance retrieves the default version of LoadbalanceClient.
// Please explicitly pick a version.
func (c *Clientset) Loadbalance() loadbalancev1alpha2.LoadbalanceV1alpha2Interface {
	return c.loadbalanceV1alpha2
}

// LoggingV1alpha1 retrieves the LoggingV1alpha1Client
func (c *Clientset) LoggingV1alpha1() loggingv1alpha1.LoggingV1alpha1Interface {
	return c.loggingV1alpha1
}

// Deprecated: Logging retrieves the default version of LoggingClient.
// Please explicitly pick a version.
func (c *Clientset) Logging() loggingv1alpha1.LoggingV1alpha1Interface {
	return c.loggingV1alpha1
}

// MicroserviceV1alpha1 retrieves the MicroserviceV1alpha1Client
func (c *Clientset) MicroserviceV1alpha1() microservicev1alpha1.MicroserviceV1alpha1Interface {
	return c.microserviceV1alpha1
}

// Deprecated: Microservice retrieves the default version of MicroserviceClient.
// Please explicitly pick a version.
func (c *Clientset) Microservice() microservicev1alpha1.MicroserviceV1alpha1Interface {
	return c.microserviceV1alpha1
}

// ModelV1alpha1 retrieves the ModelV1alpha1Client
func (c *Clientset) ModelV1alpha1() modelv1alpha1.ModelV1alpha1Interface {
	return c.modelV1alpha1
}

// Deprecated: Model retrieves the default version of ModelClient.
// Please explicitly pick a version.
func (c *Clientset) Model() modelv1alpha1.ModelV1alpha1Interface {
	return c.modelV1alpha1
}

// OrchestrationV1alpha1 retrieves the OrchestrationV1alpha1Client
func (c *Clientset) OrchestrationV1alpha1() orchestrationv1alpha1.OrchestrationV1alpha1Interface {
	return c.orchestrationV1alpha1
}

// Deprecated: Orchestration retrieves the default version of OrchestrationClient.
// Please explicitly pick a version.
func (c *Clientset) Orchestration() orchestrationv1alpha1.OrchestrationV1alpha1Interface {
	return c.orchestrationV1alpha1
}

// ReleaseV1alpha1 retrieves the ReleaseV1alpha1Client
func (c *Clientset) ReleaseV1alpha1() releasev1alpha1.ReleaseV1alpha1Interface {
	return c.releaseV1alpha1
}

// Deprecated: Release retrieves the default version of ReleaseClient.
// Please explicitly pick a version.
func (c *Clientset) Release() releasev1alpha1.ReleaseV1alpha1Interface {
	return c.releaseV1alpha1
}

// ResourceV1alpha1 retrieves the ResourceV1alpha1Client
func (c *Clientset) ResourceV1alpha1() resourcev1alpha1.ResourceV1alpha1Interface {
	return c.resourceV1alpha1
}

// ResourceV1beta1 retrieves the ResourceV1beta1Client
func (c *Clientset) ResourceV1beta1() resourcev1beta1.ResourceV1beta1Interface {
	return c.resourceV1beta1
}

// Deprecated: Resource retrieves the default version of ResourceClient.
// Please explicitly pick a version.
func (c *Clientset) Resource() resourcev1beta1.ResourceV1beta1Interface {
	return c.resourceV1beta1
}

// TenantV1alpha1 retrieves the TenantV1alpha1Client
func (c *Clientset) TenantV1alpha1() tenantv1alpha1.TenantV1alpha1Interface {
	return c.tenantV1alpha1
}

// Deprecated: Tenant retrieves the default version of TenantClient.
// Please explicitly pick a version.
func (c *Clientset) Tenant() tenantv1alpha1.TenantV1alpha1Interface {
	return c.tenantV1alpha1
}

// NewForConfig creates a new Clientset for the given config.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.apiextensionsV1beta1, err = apiextensionsv1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.apiregistrationV1, err = apiregistrationv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.cleverV1alpha1, err = cleverv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.cnetworkingV1alpha1, err = cnetworkingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.configV1alpha1, err = configv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.datasetV1alpha1, err = datasetv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.devopsV1, err = devopsv1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.loadbalanceV1alpha2, err = loadbalancev1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.loggingV1alpha1, err = loggingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.microserviceV1alpha1, err = microservicev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.modelV1alpha1, err = modelv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.orchestrationV1alpha1, err = orchestrationv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.releaseV1alpha1, err = releasev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.resourceV1alpha1, err = resourcev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.resourceV1beta1, err = resourcev1beta1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.tenantV1alpha1, err = tenantv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.Clientset, err = kubernetes.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.apiextensionsV1beta1 = apiextensionsv1beta1.NewForConfigOrDie(c)
	cs.apiregistrationV1 = apiregistrationv1.NewForConfigOrDie(c)
	cs.cleverV1alpha1 = cleverv1alpha1.NewForConfigOrDie(c)
	cs.cnetworkingV1alpha1 = cnetworkingv1alpha1.NewForConfigOrDie(c)
	cs.configV1alpha1 = configv1alpha1.NewForConfigOrDie(c)
	cs.datasetV1alpha1 = datasetv1alpha1.NewForConfigOrDie(c)
	cs.devopsV1 = devopsv1.NewForConfigOrDie(c)
	cs.loadbalanceV1alpha2 = loadbalancev1alpha2.NewForConfigOrDie(c)
	cs.loggingV1alpha1 = loggingv1alpha1.NewForConfigOrDie(c)
	cs.microserviceV1alpha1 = microservicev1alpha1.NewForConfigOrDie(c)
	cs.modelV1alpha1 = modelv1alpha1.NewForConfigOrDie(c)
	cs.orchestrationV1alpha1 = orchestrationv1alpha1.NewForConfigOrDie(c)
	cs.releaseV1alpha1 = releasev1alpha1.NewForConfigOrDie(c)
	cs.resourceV1alpha1 = resourcev1alpha1.NewForConfigOrDie(c)
	cs.resourceV1beta1 = resourcev1beta1.NewForConfigOrDie(c)
	cs.tenantV1alpha1 = tenantv1alpha1.NewForConfigOrDie(c)

	cs.Clientset = kubernetes.NewForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.apiextensionsV1beta1 = apiextensionsv1beta1.New(c)
	cs.apiregistrationV1 = apiregistrationv1.New(c)
	cs.cleverV1alpha1 = cleverv1alpha1.New(c)
	cs.cnetworkingV1alpha1 = cnetworkingv1alpha1.New(c)
	cs.configV1alpha1 = configv1alpha1.New(c)
	cs.datasetV1alpha1 = datasetv1alpha1.New(c)
	cs.devopsV1 = devopsv1.New(c)
	cs.loadbalanceV1alpha2 = loadbalancev1alpha2.New(c)
	cs.loggingV1alpha1 = loggingv1alpha1.New(c)
	cs.microserviceV1alpha1 = microservicev1alpha1.New(c)
	cs.modelV1alpha1 = modelv1alpha1.New(c)
	cs.orchestrationV1alpha1 = orchestrationv1alpha1.New(c)
	cs.releaseV1alpha1 = releasev1alpha1.New(c)
	cs.resourceV1alpha1 = resourcev1alpha1.New(c)
	cs.resourceV1beta1 = resourcev1beta1.New(c)
	cs.tenantV1alpha1 = tenantv1alpha1.New(c)

	cs.Clientset = kubernetes.New(c)
	return &cs
}
