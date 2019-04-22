package v3

import (
	"context"

	"github.com/rancher/norman/controller"
	"github.com/rancher/norman/objectclient"
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
	MacvlanSubnetGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "MacvlanSubnet",
	}
	MacvlanSubnetResource = metav1.APIResource{
		Name:         "macvlansubnets",
		SingularName: "macvlansubnet",
		Namespaced:   true,

		Kind: MacvlanSubnetGroupVersionKind.Kind,
	}
)

func NewMacvlanSubnet(namespace, name string, obj MacvlanSubnet) *MacvlanSubnet {
	obj.APIVersion, obj.Kind = MacvlanSubnetGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type MacvlanSubnetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MacvlanSubnet
}

type MacvlanSubnetHandlerFunc func(key string, obj *MacvlanSubnet) (runtime.Object, error)

type MacvlanSubnetChangeHandlerFunc func(obj *MacvlanSubnet) (runtime.Object, error)

type MacvlanSubnetLister interface {
	List(namespace string, selector labels.Selector) (ret []*MacvlanSubnet, err error)
	Get(namespace, name string) (*MacvlanSubnet, error)
}

type MacvlanSubnetController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() MacvlanSubnetLister
	AddHandler(ctx context.Context, name string, handler MacvlanSubnetHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler MacvlanSubnetHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type MacvlanSubnetInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*MacvlanSubnet) (*MacvlanSubnet, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*MacvlanSubnet, error)
	Get(name string, opts metav1.GetOptions) (*MacvlanSubnet, error)
	Update(*MacvlanSubnet) (*MacvlanSubnet, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*MacvlanSubnetList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() MacvlanSubnetController
	AddHandler(ctx context.Context, name string, sync MacvlanSubnetHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle MacvlanSubnetLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync MacvlanSubnetHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle MacvlanSubnetLifecycle)
}

type macvlanSubnetLister struct {
	controller *macvlanSubnetController
}

func (l *macvlanSubnetLister) List(namespace string, selector labels.Selector) (ret []*MacvlanSubnet, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*MacvlanSubnet))
	})
	return
}

func (l *macvlanSubnetLister) Get(namespace, name string) (*MacvlanSubnet, error) {
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
			Group:    MacvlanSubnetGroupVersionKind.Group,
			Resource: "macvlanSubnet",
		}, key)
	}
	return obj.(*MacvlanSubnet), nil
}

type macvlanSubnetController struct {
	controller.GenericController
}

func (c *macvlanSubnetController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *macvlanSubnetController) Lister() MacvlanSubnetLister {
	return &macvlanSubnetLister{
		controller: c,
	}
}

func (c *macvlanSubnetController) AddHandler(ctx context.Context, name string, handler MacvlanSubnetHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*MacvlanSubnet); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *macvlanSubnetController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler MacvlanSubnetHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*MacvlanSubnet); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type macvlanSubnetFactory struct {
}

func (c macvlanSubnetFactory) Object() runtime.Object {
	return &MacvlanSubnet{}
}

func (c macvlanSubnetFactory) List() runtime.Object {
	return &MacvlanSubnetList{}
}

func (s *macvlanSubnetClient) Controller() MacvlanSubnetController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.macvlanSubnetControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(MacvlanSubnetGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &macvlanSubnetController{
		GenericController: genericController,
	}

	s.client.macvlanSubnetControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type macvlanSubnetClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   MacvlanSubnetController
}

func (s *macvlanSubnetClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *macvlanSubnetClient) Create(o *MacvlanSubnet) (*MacvlanSubnet, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*MacvlanSubnet), err
}

func (s *macvlanSubnetClient) Get(name string, opts metav1.GetOptions) (*MacvlanSubnet, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*MacvlanSubnet), err
}

func (s *macvlanSubnetClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*MacvlanSubnet, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*MacvlanSubnet), err
}

func (s *macvlanSubnetClient) Update(o *MacvlanSubnet) (*MacvlanSubnet, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*MacvlanSubnet), err
}

func (s *macvlanSubnetClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *macvlanSubnetClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *macvlanSubnetClient) List(opts metav1.ListOptions) (*MacvlanSubnetList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*MacvlanSubnetList), err
}

func (s *macvlanSubnetClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *macvlanSubnetClient) Patch(o *MacvlanSubnet, patchType types.PatchType, data []byte, subresources ...string) (*MacvlanSubnet, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*MacvlanSubnet), err
}

func (s *macvlanSubnetClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *macvlanSubnetClient) AddHandler(ctx context.Context, name string, sync MacvlanSubnetHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *macvlanSubnetClient) AddLifecycle(ctx context.Context, name string, lifecycle MacvlanSubnetLifecycle) {
	sync := NewMacvlanSubnetLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *macvlanSubnetClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync MacvlanSubnetHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *macvlanSubnetClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle MacvlanSubnetLifecycle) {
	sync := NewMacvlanSubnetLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

type MacvlanSubnetIndexer func(obj *MacvlanSubnet) ([]string, error)

type MacvlanSubnetClientCache interface {
	Get(namespace, name string) (*MacvlanSubnet, error)
	List(namespace string, selector labels.Selector) ([]*MacvlanSubnet, error)

	Index(name string, indexer MacvlanSubnetIndexer)
	GetIndexed(name, key string) ([]*MacvlanSubnet, error)
}

type MacvlanSubnetClient interface {
	Create(*MacvlanSubnet) (*MacvlanSubnet, error)
	Get(namespace, name string, opts metav1.GetOptions) (*MacvlanSubnet, error)
	Update(*MacvlanSubnet) (*MacvlanSubnet, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*MacvlanSubnetList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() MacvlanSubnetClientCache

	OnCreate(ctx context.Context, name string, sync MacvlanSubnetChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync MacvlanSubnetChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync MacvlanSubnetChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() MacvlanSubnetInterface
}

type macvlanSubnetClientCache struct {
	client *macvlanSubnetClient2
}

type macvlanSubnetClient2 struct {
	iface      MacvlanSubnetInterface
	controller MacvlanSubnetController
}

func (n *macvlanSubnetClient2) Interface() MacvlanSubnetInterface {
	return n.iface
}

func (n *macvlanSubnetClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *macvlanSubnetClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *macvlanSubnetClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *macvlanSubnetClient2) Create(obj *MacvlanSubnet) (*MacvlanSubnet, error) {
	return n.iface.Create(obj)
}

func (n *macvlanSubnetClient2) Get(namespace, name string, opts metav1.GetOptions) (*MacvlanSubnet, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *macvlanSubnetClient2) Update(obj *MacvlanSubnet) (*MacvlanSubnet, error) {
	return n.iface.Update(obj)
}

func (n *macvlanSubnetClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *macvlanSubnetClient2) List(namespace string, opts metav1.ListOptions) (*MacvlanSubnetList, error) {
	return n.iface.List(opts)
}

func (n *macvlanSubnetClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *macvlanSubnetClientCache) Get(namespace, name string) (*MacvlanSubnet, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *macvlanSubnetClientCache) List(namespace string, selector labels.Selector) ([]*MacvlanSubnet, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *macvlanSubnetClient2) Cache() MacvlanSubnetClientCache {
	n.loadController()
	return &macvlanSubnetClientCache{
		client: n,
	}
}

func (n *macvlanSubnetClient2) OnCreate(ctx context.Context, name string, sync MacvlanSubnetChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &macvlanSubnetLifecycleDelegate{create: sync})
}

func (n *macvlanSubnetClient2) OnChange(ctx context.Context, name string, sync MacvlanSubnetChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &macvlanSubnetLifecycleDelegate{update: sync})
}

func (n *macvlanSubnetClient2) OnRemove(ctx context.Context, name string, sync MacvlanSubnetChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &macvlanSubnetLifecycleDelegate{remove: sync})
}

func (n *macvlanSubnetClientCache) Index(name string, indexer MacvlanSubnetIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*MacvlanSubnet); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *macvlanSubnetClientCache) GetIndexed(name, key string) ([]*MacvlanSubnet, error) {
	var result []*MacvlanSubnet
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*MacvlanSubnet); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *macvlanSubnetClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type macvlanSubnetLifecycleDelegate struct {
	create MacvlanSubnetChangeHandlerFunc
	update MacvlanSubnetChangeHandlerFunc
	remove MacvlanSubnetChangeHandlerFunc
}

func (n *macvlanSubnetLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *macvlanSubnetLifecycleDelegate) Create(obj *MacvlanSubnet) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *macvlanSubnetLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *macvlanSubnetLifecycleDelegate) Remove(obj *MacvlanSubnet) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *macvlanSubnetLifecycleDelegate) Updated(obj *MacvlanSubnet) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
