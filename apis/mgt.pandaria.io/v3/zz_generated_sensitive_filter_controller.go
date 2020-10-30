package v3

import (
	"context"
	"time"

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
	SensitiveFilterGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "SensitiveFilter",
	}
	SensitiveFilterResource = metav1.APIResource{
		Name:         "sensitivefilters",
		SingularName: "sensitivefilter",
		Namespaced:   false,
		Kind:         SensitiveFilterGroupVersionKind.Kind,
	}

	SensitiveFilterGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "sensitivefilters",
	}
)

func init() {
	resource.Put(SensitiveFilterGroupVersionResource)
}

func NewSensitiveFilter(namespace, name string, obj SensitiveFilter) *SensitiveFilter {
	obj.APIVersion, obj.Kind = SensitiveFilterGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type SensitiveFilterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SensitiveFilter `json:"items"`
}

type SensitiveFilterHandlerFunc func(key string, obj *SensitiveFilter) (runtime.Object, error)

type SensitiveFilterChangeHandlerFunc func(obj *SensitiveFilter) (runtime.Object, error)

type SensitiveFilterLister interface {
	List(namespace string, selector labels.Selector) (ret []*SensitiveFilter, err error)
	Get(namespace, name string) (*SensitiveFilter, error)
}

type SensitiveFilterController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() SensitiveFilterLister
	AddHandler(ctx context.Context, name string, handler SensitiveFilterHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync SensitiveFilterHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler SensitiveFilterHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler SensitiveFilterHandlerFunc)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, after time.Duration)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type SensitiveFilterInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*SensitiveFilter) (*SensitiveFilter, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*SensitiveFilter, error)
	Get(name string, opts metav1.GetOptions) (*SensitiveFilter, error)
	Update(*SensitiveFilter) (*SensitiveFilter, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*SensitiveFilterList, error)
	ListNamespaced(namespace string, opts metav1.ListOptions) (*SensitiveFilterList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() SensitiveFilterController
	AddHandler(ctx context.Context, name string, sync SensitiveFilterHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync SensitiveFilterHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle SensitiveFilterLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle SensitiveFilterLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync SensitiveFilterHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync SensitiveFilterHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle SensitiveFilterLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle SensitiveFilterLifecycle)
}

type sensitiveFilterLister struct {
	controller *sensitiveFilterController
}

func (l *sensitiveFilterLister) List(namespace string, selector labels.Selector) (ret []*SensitiveFilter, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*SensitiveFilter))
	})
	return
}

func (l *sensitiveFilterLister) Get(namespace, name string) (*SensitiveFilter, error) {
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
			Group:    SensitiveFilterGroupVersionKind.Group,
			Resource: "sensitiveFilter",
		}, key)
	}
	return obj.(*SensitiveFilter), nil
}

type sensitiveFilterController struct {
	controller.GenericController
}

func (c *sensitiveFilterController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *sensitiveFilterController) Lister() SensitiveFilterLister {
	return &sensitiveFilterLister{
		controller: c,
	}
}

func (c *sensitiveFilterController) AddHandler(ctx context.Context, name string, handler SensitiveFilterHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*SensitiveFilter); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *sensitiveFilterController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler SensitiveFilterHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*SensitiveFilter); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *sensitiveFilterController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler SensitiveFilterHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*SensitiveFilter); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *sensitiveFilterController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler SensitiveFilterHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*SensitiveFilter); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type sensitiveFilterFactory struct {
}

func (c sensitiveFilterFactory) Object() runtime.Object {
	return &SensitiveFilter{}
}

func (c sensitiveFilterFactory) List() runtime.Object {
	return &SensitiveFilterList{}
}

func (s *sensitiveFilterClient) Controller() SensitiveFilterController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.sensitiveFilterControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(SensitiveFilterGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &sensitiveFilterController{
		GenericController: genericController,
	}

	s.client.sensitiveFilterControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type sensitiveFilterClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   SensitiveFilterController
}

func (s *sensitiveFilterClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *sensitiveFilterClient) Create(o *SensitiveFilter) (*SensitiveFilter, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*SensitiveFilter), err
}

func (s *sensitiveFilterClient) Get(name string, opts metav1.GetOptions) (*SensitiveFilter, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*SensitiveFilter), err
}

func (s *sensitiveFilterClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*SensitiveFilter, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*SensitiveFilter), err
}

func (s *sensitiveFilterClient) Update(o *SensitiveFilter) (*SensitiveFilter, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*SensitiveFilter), err
}

func (s *sensitiveFilterClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *sensitiveFilterClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *sensitiveFilterClient) List(opts metav1.ListOptions) (*SensitiveFilterList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*SensitiveFilterList), err
}

func (s *sensitiveFilterClient) ListNamespaced(namespace string, opts metav1.ListOptions) (*SensitiveFilterList, error) {
	obj, err := s.objectClient.ListNamespaced(namespace, opts)
	return obj.(*SensitiveFilterList), err
}

func (s *sensitiveFilterClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *sensitiveFilterClient) Patch(o *SensitiveFilter, patchType types.PatchType, data []byte, subresources ...string) (*SensitiveFilter, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*SensitiveFilter), err
}

func (s *sensitiveFilterClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *sensitiveFilterClient) AddHandler(ctx context.Context, name string, sync SensitiveFilterHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *sensitiveFilterClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync SensitiveFilterHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *sensitiveFilterClient) AddLifecycle(ctx context.Context, name string, lifecycle SensitiveFilterLifecycle) {
	sync := NewSensitiveFilterLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *sensitiveFilterClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle SensitiveFilterLifecycle) {
	sync := NewSensitiveFilterLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *sensitiveFilterClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync SensitiveFilterHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *sensitiveFilterClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync SensitiveFilterHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *sensitiveFilterClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle SensitiveFilterLifecycle) {
	sync := NewSensitiveFilterLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *sensitiveFilterClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle SensitiveFilterLifecycle) {
	sync := NewSensitiveFilterLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}
