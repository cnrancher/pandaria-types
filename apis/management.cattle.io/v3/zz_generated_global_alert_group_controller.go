package v3

import (
	"context"

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
	GlobalAlertGroupGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "GlobalAlertGroup",
	}
	GlobalAlertGroupResource = metav1.APIResource{
		Name:         "globalalertgroups",
		SingularName: "globalalertgroup",
		Namespaced:   true,

		Kind: GlobalAlertGroupGroupVersionKind.Kind,
	}

	GlobalAlertGroupGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "globalalertgroups",
	}
)

func init() {
	resource.Put(GlobalAlertGroupGroupVersionResource)
}

func NewGlobalAlertGroup(namespace, name string, obj GlobalAlertGroup) *GlobalAlertGroup {
	obj.APIVersion, obj.Kind = GlobalAlertGroupGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type GlobalAlertGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalAlertGroup `json:"items"`
}

type GlobalAlertGroupHandlerFunc func(key string, obj *GlobalAlertGroup) (runtime.Object, error)

type GlobalAlertGroupChangeHandlerFunc func(obj *GlobalAlertGroup) (runtime.Object, error)

type GlobalAlertGroupLister interface {
	List(namespace string, selector labels.Selector) (ret []*GlobalAlertGroup, err error)
	Get(namespace, name string) (*GlobalAlertGroup, error)
}

type GlobalAlertGroupController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() GlobalAlertGroupLister
	AddHandler(ctx context.Context, name string, handler GlobalAlertGroupHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync GlobalAlertGroupHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler GlobalAlertGroupHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler GlobalAlertGroupHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type GlobalAlertGroupInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*GlobalAlertGroup) (*GlobalAlertGroup, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*GlobalAlertGroup, error)
	Get(name string, opts metav1.GetOptions) (*GlobalAlertGroup, error)
	Update(*GlobalAlertGroup) (*GlobalAlertGroup, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*GlobalAlertGroupList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() GlobalAlertGroupController
	AddHandler(ctx context.Context, name string, sync GlobalAlertGroupHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync GlobalAlertGroupHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle GlobalAlertGroupLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle GlobalAlertGroupLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync GlobalAlertGroupHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync GlobalAlertGroupHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle GlobalAlertGroupLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle GlobalAlertGroupLifecycle)
}

type globalAlertGroupLister struct {
	controller *globalAlertGroupController
}

func (l *globalAlertGroupLister) List(namespace string, selector labels.Selector) (ret []*GlobalAlertGroup, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*GlobalAlertGroup))
	})
	return
}

func (l *globalAlertGroupLister) Get(namespace, name string) (*GlobalAlertGroup, error) {
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
			Group:    GlobalAlertGroupGroupVersionKind.Group,
			Resource: "globalAlertGroup",
		}, key)
	}
	return obj.(*GlobalAlertGroup), nil
}

type globalAlertGroupController struct {
	controller.GenericController
}

func (c *globalAlertGroupController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *globalAlertGroupController) Lister() GlobalAlertGroupLister {
	return &globalAlertGroupLister{
		controller: c,
	}
}

func (c *globalAlertGroupController) AddHandler(ctx context.Context, name string, handler GlobalAlertGroupHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*GlobalAlertGroup); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *globalAlertGroupController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler GlobalAlertGroupHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*GlobalAlertGroup); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *globalAlertGroupController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler GlobalAlertGroupHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*GlobalAlertGroup); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *globalAlertGroupController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler GlobalAlertGroupHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*GlobalAlertGroup); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type globalAlertGroupFactory struct {
}

func (c globalAlertGroupFactory) Object() runtime.Object {
	return &GlobalAlertGroup{}
}

func (c globalAlertGroupFactory) List() runtime.Object {
	return &GlobalAlertGroupList{}
}

func (s *globalAlertGroupClient) Controller() GlobalAlertGroupController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.globalAlertGroupControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(GlobalAlertGroupGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &globalAlertGroupController{
		GenericController: genericController,
	}

	s.client.globalAlertGroupControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type globalAlertGroupClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   GlobalAlertGroupController
}

func (s *globalAlertGroupClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *globalAlertGroupClient) Create(o *GlobalAlertGroup) (*GlobalAlertGroup, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*GlobalAlertGroup), err
}

func (s *globalAlertGroupClient) Get(name string, opts metav1.GetOptions) (*GlobalAlertGroup, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*GlobalAlertGroup), err
}

func (s *globalAlertGroupClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*GlobalAlertGroup, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*GlobalAlertGroup), err
}

func (s *globalAlertGroupClient) Update(o *GlobalAlertGroup) (*GlobalAlertGroup, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*GlobalAlertGroup), err
}

func (s *globalAlertGroupClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *globalAlertGroupClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *globalAlertGroupClient) List(opts metav1.ListOptions) (*GlobalAlertGroupList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*GlobalAlertGroupList), err
}

func (s *globalAlertGroupClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *globalAlertGroupClient) Patch(o *GlobalAlertGroup, patchType types.PatchType, data []byte, subresources ...string) (*GlobalAlertGroup, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*GlobalAlertGroup), err
}

func (s *globalAlertGroupClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *globalAlertGroupClient) AddHandler(ctx context.Context, name string, sync GlobalAlertGroupHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *globalAlertGroupClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync GlobalAlertGroupHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *globalAlertGroupClient) AddLifecycle(ctx context.Context, name string, lifecycle GlobalAlertGroupLifecycle) {
	sync := NewGlobalAlertGroupLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *globalAlertGroupClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle GlobalAlertGroupLifecycle) {
	sync := NewGlobalAlertGroupLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *globalAlertGroupClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync GlobalAlertGroupHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *globalAlertGroupClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync GlobalAlertGroupHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *globalAlertGroupClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle GlobalAlertGroupLifecycle) {
	sync := NewGlobalAlertGroupLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *globalAlertGroupClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle GlobalAlertGroupLifecycle) {
	sync := NewGlobalAlertGroupLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

type GlobalAlertGroupIndexer func(obj *GlobalAlertGroup) ([]string, error)

type GlobalAlertGroupClientCache interface {
	Get(namespace, name string) (*GlobalAlertGroup, error)
	List(namespace string, selector labels.Selector) ([]*GlobalAlertGroup, error)

	Index(name string, indexer GlobalAlertGroupIndexer)
	GetIndexed(name, key string) ([]*GlobalAlertGroup, error)
}

type GlobalAlertGroupClient interface {
	Create(*GlobalAlertGroup) (*GlobalAlertGroup, error)
	Get(namespace, name string, opts metav1.GetOptions) (*GlobalAlertGroup, error)
	Update(*GlobalAlertGroup) (*GlobalAlertGroup, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*GlobalAlertGroupList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() GlobalAlertGroupClientCache

	OnCreate(ctx context.Context, name string, sync GlobalAlertGroupChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync GlobalAlertGroupChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync GlobalAlertGroupChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() GlobalAlertGroupInterface
}

type globalAlertGroupClientCache struct {
	client *globalAlertGroupClient2
}

type globalAlertGroupClient2 struct {
	iface      GlobalAlertGroupInterface
	controller GlobalAlertGroupController
}

func (n *globalAlertGroupClient2) Interface() GlobalAlertGroupInterface {
	return n.iface
}

func (n *globalAlertGroupClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *globalAlertGroupClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *globalAlertGroupClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *globalAlertGroupClient2) Create(obj *GlobalAlertGroup) (*GlobalAlertGroup, error) {
	return n.iface.Create(obj)
}

func (n *globalAlertGroupClient2) Get(namespace, name string, opts metav1.GetOptions) (*GlobalAlertGroup, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *globalAlertGroupClient2) Update(obj *GlobalAlertGroup) (*GlobalAlertGroup, error) {
	return n.iface.Update(obj)
}

func (n *globalAlertGroupClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *globalAlertGroupClient2) List(namespace string, opts metav1.ListOptions) (*GlobalAlertGroupList, error) {
	return n.iface.List(opts)
}

func (n *globalAlertGroupClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *globalAlertGroupClientCache) Get(namespace, name string) (*GlobalAlertGroup, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *globalAlertGroupClientCache) List(namespace string, selector labels.Selector) ([]*GlobalAlertGroup, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *globalAlertGroupClient2) Cache() GlobalAlertGroupClientCache {
	n.loadController()
	return &globalAlertGroupClientCache{
		client: n,
	}
}

func (n *globalAlertGroupClient2) OnCreate(ctx context.Context, name string, sync GlobalAlertGroupChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &globalAlertGroupLifecycleDelegate{create: sync})
}

func (n *globalAlertGroupClient2) OnChange(ctx context.Context, name string, sync GlobalAlertGroupChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &globalAlertGroupLifecycleDelegate{update: sync})
}

func (n *globalAlertGroupClient2) OnRemove(ctx context.Context, name string, sync GlobalAlertGroupChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &globalAlertGroupLifecycleDelegate{remove: sync})
}

func (n *globalAlertGroupClientCache) Index(name string, indexer GlobalAlertGroupIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*GlobalAlertGroup); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *globalAlertGroupClientCache) GetIndexed(name, key string) ([]*GlobalAlertGroup, error) {
	var result []*GlobalAlertGroup
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*GlobalAlertGroup); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *globalAlertGroupClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type globalAlertGroupLifecycleDelegate struct {
	create GlobalAlertGroupChangeHandlerFunc
	update GlobalAlertGroupChangeHandlerFunc
	remove GlobalAlertGroupChangeHandlerFunc
}

func (n *globalAlertGroupLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *globalAlertGroupLifecycleDelegate) Create(obj *GlobalAlertGroup) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *globalAlertGroupLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *globalAlertGroupLifecycleDelegate) Remove(obj *GlobalAlertGroup) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *globalAlertGroupLifecycleDelegate) Updated(obj *GlobalAlertGroup) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
