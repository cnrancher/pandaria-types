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
	TransportServerGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "TransportServer",
	}
	TransportServerResource = metav1.APIResource{
		Name:         "transportservers",
		SingularName: "transportserver",
		Namespaced:   true,

		Kind: TransportServerGroupVersionKind.Kind,
	}

	TransportServerGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "transportservers",
	}
)

func init() {
	resource.Put(TransportServerGroupVersionResource)
}

func NewTransportServer(namespace, name string, obj v1.TransportServer) *v1.TransportServer {
	obj.APIVersion, obj.Kind = TransportServerGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type TransportServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1.TransportServer `json:"items"`
}

type TransportServerHandlerFunc func(key string, obj *v1.TransportServer) (runtime.Object, error)

type TransportServerChangeHandlerFunc func(obj *v1.TransportServer) (runtime.Object, error)

type TransportServerLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1.TransportServer, err error)
	Get(namespace, name string) (*v1.TransportServer, error)
}

type TransportServerController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() TransportServerLister
	AddHandler(ctx context.Context, name string, handler TransportServerHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync TransportServerHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler TransportServerHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler TransportServerHandlerFunc)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, after time.Duration)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type TransportServerInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v1.TransportServer) (*v1.TransportServer, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.TransportServer, error)
	Get(name string, opts metav1.GetOptions) (*v1.TransportServer, error)
	Update(*v1.TransportServer) (*v1.TransportServer, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*TransportServerList, error)
	ListNamespaced(namespace string, opts metav1.ListOptions) (*TransportServerList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() TransportServerController
	AddHandler(ctx context.Context, name string, sync TransportServerHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync TransportServerHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle TransportServerLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle TransportServerLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync TransportServerHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync TransportServerHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle TransportServerLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle TransportServerLifecycle)
}

type transportServerLister struct {
	controller *transportServerController
}

func (l *transportServerLister) List(namespace string, selector labels.Selector) (ret []*v1.TransportServer, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1.TransportServer))
	})
	return
}

func (l *transportServerLister) Get(namespace, name string) (*v1.TransportServer, error) {
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
			Group:    TransportServerGroupVersionKind.Group,
			Resource: "transportServer",
		}, key)
	}
	return obj.(*v1.TransportServer), nil
}

type transportServerController struct {
	controller.GenericController
}

func (c *transportServerController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *transportServerController) Lister() TransportServerLister {
	return &transportServerLister{
		controller: c,
	}
}

func (c *transportServerController) AddHandler(ctx context.Context, name string, handler TransportServerHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.TransportServer); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *transportServerController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler TransportServerHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.TransportServer); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *transportServerController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler TransportServerHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.TransportServer); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *transportServerController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler TransportServerHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.TransportServer); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type transportServerFactory struct {
}

func (c transportServerFactory) Object() runtime.Object {
	return &v1.TransportServer{}
}

func (c transportServerFactory) List() runtime.Object {
	return &TransportServerList{}
}

func (s *transportServerClient) Controller() TransportServerController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.transportServerControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(TransportServerGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &transportServerController{
		GenericController: genericController,
	}

	s.client.transportServerControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type transportServerClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   TransportServerController
}

func (s *transportServerClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *transportServerClient) Create(o *v1.TransportServer) (*v1.TransportServer, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1.TransportServer), err
}

func (s *transportServerClient) Get(name string, opts metav1.GetOptions) (*v1.TransportServer, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1.TransportServer), err
}

func (s *transportServerClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.TransportServer, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v1.TransportServer), err
}

func (s *transportServerClient) Update(o *v1.TransportServer) (*v1.TransportServer, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1.TransportServer), err
}

func (s *transportServerClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *transportServerClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *transportServerClient) List(opts metav1.ListOptions) (*TransportServerList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*TransportServerList), err
}

func (s *transportServerClient) ListNamespaced(namespace string, opts metav1.ListOptions) (*TransportServerList, error) {
	obj, err := s.objectClient.ListNamespaced(namespace, opts)
	return obj.(*TransportServerList), err
}

func (s *transportServerClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *transportServerClient) Patch(o *v1.TransportServer, patchType types.PatchType, data []byte, subresources ...string) (*v1.TransportServer, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v1.TransportServer), err
}

func (s *transportServerClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *transportServerClient) AddHandler(ctx context.Context, name string, sync TransportServerHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *transportServerClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync TransportServerHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *transportServerClient) AddLifecycle(ctx context.Context, name string, lifecycle TransportServerLifecycle) {
	sync := NewTransportServerLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *transportServerClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle TransportServerLifecycle) {
	sync := NewTransportServerLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *transportServerClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync TransportServerHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *transportServerClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync TransportServerHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *transportServerClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle TransportServerLifecycle) {
	sync := NewTransportServerLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *transportServerClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle TransportServerLifecycle) {
	sync := NewTransportServerLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}
