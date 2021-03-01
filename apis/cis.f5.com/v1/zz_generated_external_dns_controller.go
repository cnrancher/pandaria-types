package v1

import (
	"context"
	"time"

	v1 "github.com/F5Networks/k8s-bigip-ctlr/config/apis/cis/v1"
	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/tools/cache"
)

var (
	ExternalDNSGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "ExternalDNS",
	}
	ExternalDNSResource = metav1.APIResource{
		Name:         "externaldnss",
		SingularName: "externaldns",
		Namespaced:   true,

		Kind: ExternalDNSGroupVersionKind.Kind,
	}

	ExternalDNSGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "externaldnss",
	}
)

func init() {
	resource.Put(ExternalDNSGroupVersionResource)
}

func NewExternalDNS(namespace, name string, obj v1.ExternalDNS) *v1.ExternalDNS {
	obj.APIVersion, obj.Kind = ExternalDNSGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type ExternalDNSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1.ExternalDNS `json:"items"`
}

type ExternalDNSHandlerFunc func(key string, obj *v1.ExternalDNS) (runtime.Object, error)

type ExternalDNSChangeHandlerFunc func(obj *v1.ExternalDNS) (runtime.Object, error)

type ExternalDNSLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1.ExternalDNS, err error)
	Get(namespace, name string) (*v1.ExternalDNS, error)
}

type ExternalDNSController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() ExternalDNSLister
	AddHandler(ctx context.Context, name string, handler ExternalDNSHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync ExternalDNSHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler ExternalDNSHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler ExternalDNSHandlerFunc)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, after time.Duration)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type ExternalDNSInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v1.ExternalDNS) (*v1.ExternalDNS, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.ExternalDNS, error)
	Get(name string, opts metav1.GetOptions) (*v1.ExternalDNS, error)
	Update(*v1.ExternalDNS) (*v1.ExternalDNS, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*ExternalDNSList, error)
	ListNamespaced(namespace string, opts metav1.ListOptions) (*ExternalDNSList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() ExternalDNSController
	AddHandler(ctx context.Context, name string, sync ExternalDNSHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync ExternalDNSHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle ExternalDNSLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle ExternalDNSLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync ExternalDNSHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync ExternalDNSHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle ExternalDNSLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle ExternalDNSLifecycle)
}

type externalDNSLister struct {
	controller *externalDNSController
}

func (l *externalDNSLister) List(namespace string, selector labels.Selector) (ret []*v1.ExternalDNS, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1.ExternalDNS))
	})
	return
}

func (l *externalDNSLister) Get(namespace, name string) (*v1.ExternalDNS, error) {
	var key string
	if namespace != "" {
		key = namespace + "/" + name
	} else {
		key = name
	}
	obj, exists, err := l.controller.Informer().GetIndexer().GetByKey(key)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(schema.GroupResource{
			Group:    ExternalDNSGroupVersionKind.Group,
			Resource: "externalDNS",
		}, key)
	}
	return obj.(*v1.ExternalDNS), nil
}

type externalDNSController struct {
	controller.GenericController
}

func (c *externalDNSController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *externalDNSController) Lister() ExternalDNSLister {
	return &externalDNSLister{
		controller: c,
	}
}

func (c *externalDNSController) AddHandler(ctx context.Context, name string, handler ExternalDNSHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.ExternalDNS); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *externalDNSController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler ExternalDNSHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.ExternalDNS); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *externalDNSController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler ExternalDNSHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.ExternalDNS); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *externalDNSController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler ExternalDNSHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.ExternalDNS); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type externalDNSFactory struct {
}

func (c externalDNSFactory) Object() runtime.Object {
	return &v1.ExternalDNS{}
}

func (c externalDNSFactory) List() runtime.Object {
	return &ExternalDNSList{}
}

func (s *externalDNSClient) Controller() ExternalDNSController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.externalDNSControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(ExternalDNSGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &externalDNSController{
		GenericController: genericController,
	}

	s.client.externalDNSControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type externalDNSClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   ExternalDNSController
}

func (s *externalDNSClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *externalDNSClient) Create(o *v1.ExternalDNS) (*v1.ExternalDNS, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1.ExternalDNS), err
}

func (s *externalDNSClient) Get(name string, opts metav1.GetOptions) (*v1.ExternalDNS, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1.ExternalDNS), err
}

func (s *externalDNSClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.ExternalDNS, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v1.ExternalDNS), err
}

func (s *externalDNSClient) Update(o *v1.ExternalDNS) (*v1.ExternalDNS, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1.ExternalDNS), err
}

func (s *externalDNSClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *externalDNSClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *externalDNSClient) List(opts metav1.ListOptions) (*ExternalDNSList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*ExternalDNSList), err
}

func (s *externalDNSClient) ListNamespaced(namespace string, opts metav1.ListOptions) (*ExternalDNSList, error) {
	obj, err := s.objectClient.ListNamespaced(namespace, opts)
	return obj.(*ExternalDNSList), err
}

func (s *externalDNSClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *externalDNSClient) Patch(o *v1.ExternalDNS, patchType types.PatchType, data []byte, subresources ...string) (*v1.ExternalDNS, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v1.ExternalDNS), err
}

func (s *externalDNSClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *externalDNSClient) AddHandler(ctx context.Context, name string, sync ExternalDNSHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *externalDNSClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync ExternalDNSHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *externalDNSClient) AddLifecycle(ctx context.Context, name string, lifecycle ExternalDNSLifecycle) {
	sync := NewExternalDNSLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *externalDNSClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle ExternalDNSLifecycle) {
	sync := NewExternalDNSLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *externalDNSClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync ExternalDNSHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *externalDNSClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync ExternalDNSHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *externalDNSClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle ExternalDNSLifecycle) {
	sync := NewExternalDNSLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *externalDNSClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle ExternalDNSLifecycle) {
	sync := NewExternalDNSLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}
