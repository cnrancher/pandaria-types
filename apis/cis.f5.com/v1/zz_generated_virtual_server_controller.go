package v1

import (
	"context"
	"time"

	"github.com/F5Networks/k8s-bigip-ctlr/config/apis/cis/v1"
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
	VirtualServerGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "VirtualServer",
	}
	VirtualServerResource = metav1.APIResource{
		Name:         "virtualservers",
		SingularName: "virtualserver",
		Namespaced:   true,

		Kind: VirtualServerGroupVersionKind.Kind,
	}

	VirtualServerGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "virtualservers",
	}
)

func init() {
	resource.Put(VirtualServerGroupVersionResource)
}

func NewVirtualServer(namespace, name string, obj v1.VirtualServer) *v1.VirtualServer {
	obj.APIVersion, obj.Kind = VirtualServerGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type VirtualServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1.VirtualServer `json:"items"`
}

type VirtualServerHandlerFunc func(key string, obj *v1.VirtualServer) (runtime.Object, error)

type VirtualServerChangeHandlerFunc func(obj *v1.VirtualServer) (runtime.Object, error)

type VirtualServerLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1.VirtualServer, err error)
	Get(namespace, name string) (*v1.VirtualServer, error)
}

type VirtualServerController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() VirtualServerLister
	AddHandler(ctx context.Context, name string, handler VirtualServerHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync VirtualServerHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler VirtualServerHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler VirtualServerHandlerFunc)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, after time.Duration)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type VirtualServerInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v1.VirtualServer) (*v1.VirtualServer, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.VirtualServer, error)
	Get(name string, opts metav1.GetOptions) (*v1.VirtualServer, error)
	Update(*v1.VirtualServer) (*v1.VirtualServer, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*VirtualServerList, error)
	ListNamespaced(namespace string, opts metav1.ListOptions) (*VirtualServerList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() VirtualServerController
	AddHandler(ctx context.Context, name string, sync VirtualServerHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync VirtualServerHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle VirtualServerLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle VirtualServerLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync VirtualServerHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync VirtualServerHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle VirtualServerLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle VirtualServerLifecycle)
}

type virtualServerLister struct {
	controller *virtualServerController
}

func (l *virtualServerLister) List(namespace string, selector labels.Selector) (ret []*v1.VirtualServer, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1.VirtualServer))
	})
	return
}

func (l *virtualServerLister) Get(namespace, name string) (*v1.VirtualServer, error) {
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
			Group:    VirtualServerGroupVersionKind.Group,
			Resource: "virtualServer",
		}, key)
	}
	return obj.(*v1.VirtualServer), nil
}

type virtualServerController struct {
	controller.GenericController
}

func (c *virtualServerController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *virtualServerController) Lister() VirtualServerLister {
	return &virtualServerLister{
		controller: c,
	}
}

func (c *virtualServerController) AddHandler(ctx context.Context, name string, handler VirtualServerHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.VirtualServer); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *virtualServerController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler VirtualServerHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.VirtualServer); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *virtualServerController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler VirtualServerHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.VirtualServer); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *virtualServerController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler VirtualServerHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.VirtualServer); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type virtualServerFactory struct {
}

func (c virtualServerFactory) Object() runtime.Object {
	return &v1.VirtualServer{}
}

func (c virtualServerFactory) List() runtime.Object {
	return &VirtualServerList{}
}

func (s *virtualServerClient) Controller() VirtualServerController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.virtualServerControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(VirtualServerGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &virtualServerController{
		GenericController: genericController,
	}

	s.client.virtualServerControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type virtualServerClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   VirtualServerController
}

func (s *virtualServerClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *virtualServerClient) Create(o *v1.VirtualServer) (*v1.VirtualServer, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1.VirtualServer), err
}

func (s *virtualServerClient) Get(name string, opts metav1.GetOptions) (*v1.VirtualServer, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1.VirtualServer), err
}

func (s *virtualServerClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.VirtualServer, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v1.VirtualServer), err
}

func (s *virtualServerClient) Update(o *v1.VirtualServer) (*v1.VirtualServer, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1.VirtualServer), err
}

func (s *virtualServerClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *virtualServerClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *virtualServerClient) List(opts metav1.ListOptions) (*VirtualServerList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*VirtualServerList), err
}

func (s *virtualServerClient) ListNamespaced(namespace string, opts metav1.ListOptions) (*VirtualServerList, error) {
	obj, err := s.objectClient.ListNamespaced(namespace, opts)
	return obj.(*VirtualServerList), err
}

func (s *virtualServerClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *virtualServerClient) Patch(o *v1.VirtualServer, patchType types.PatchType, data []byte, subresources ...string) (*v1.VirtualServer, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v1.VirtualServer), err
}

func (s *virtualServerClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *virtualServerClient) AddHandler(ctx context.Context, name string, sync VirtualServerHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *virtualServerClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync VirtualServerHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *virtualServerClient) AddLifecycle(ctx context.Context, name string, lifecycle VirtualServerLifecycle) {
	sync := NewVirtualServerLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *virtualServerClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle VirtualServerLifecycle) {
	sync := NewVirtualServerLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *virtualServerClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync VirtualServerHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *virtualServerClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync VirtualServerHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *virtualServerClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle VirtualServerLifecycle) {
	sync := NewVirtualServerLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *virtualServerClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle VirtualServerLifecycle) {
	sync := NewVirtualServerLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}
