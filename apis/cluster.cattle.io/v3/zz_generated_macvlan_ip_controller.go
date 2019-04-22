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
	MacvlanIPGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "MacvlanIP",
	}
	MacvlanIPResource = metav1.APIResource{
		Name:         "macvlanips",
		SingularName: "macvlanip",
		Namespaced:   true,

		Kind: MacvlanIPGroupVersionKind.Kind,
	}
)

func NewMacvlanIP(namespace, name string, obj MacvlanIP) *MacvlanIP {
	obj.APIVersion, obj.Kind = MacvlanIPGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type MacvlanIPList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MacvlanIP
}

type MacvlanIPHandlerFunc func(key string, obj *MacvlanIP) (runtime.Object, error)

type MacvlanIPChangeHandlerFunc func(obj *MacvlanIP) (runtime.Object, error)

type MacvlanIPLister interface {
	List(namespace string, selector labels.Selector) (ret []*MacvlanIP, err error)
	Get(namespace, name string) (*MacvlanIP, error)
}

type MacvlanIPController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() MacvlanIPLister
	AddHandler(ctx context.Context, name string, handler MacvlanIPHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler MacvlanIPHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type MacvlanIPInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*MacvlanIP) (*MacvlanIP, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*MacvlanIP, error)
	Get(name string, opts metav1.GetOptions) (*MacvlanIP, error)
	Update(*MacvlanIP) (*MacvlanIP, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*MacvlanIPList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() MacvlanIPController
	AddHandler(ctx context.Context, name string, sync MacvlanIPHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle MacvlanIPLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync MacvlanIPHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle MacvlanIPLifecycle)
}

type macvlanIPLister struct {
	controller *macvlanIPController
}

func (l *macvlanIPLister) List(namespace string, selector labels.Selector) (ret []*MacvlanIP, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*MacvlanIP))
	})
	return
}

func (l *macvlanIPLister) Get(namespace, name string) (*MacvlanIP, error) {
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
			Group:    MacvlanIPGroupVersionKind.Group,
			Resource: "macvlanIP",
		}, key)
	}
	return obj.(*MacvlanIP), nil
}

type macvlanIPController struct {
	controller.GenericController
}

func (c *macvlanIPController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *macvlanIPController) Lister() MacvlanIPLister {
	return &macvlanIPLister{
		controller: c,
	}
}

func (c *macvlanIPController) AddHandler(ctx context.Context, name string, handler MacvlanIPHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*MacvlanIP); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *macvlanIPController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler MacvlanIPHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*MacvlanIP); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type macvlanIPFactory struct {
}

func (c macvlanIPFactory) Object() runtime.Object {
	return &MacvlanIP{}
}

func (c macvlanIPFactory) List() runtime.Object {
	return &MacvlanIPList{}
}

func (s *macvlanIPClient) Controller() MacvlanIPController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.macvlanIPControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(MacvlanIPGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &macvlanIPController{
		GenericController: genericController,
	}

	s.client.macvlanIPControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type macvlanIPClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   MacvlanIPController
}

func (s *macvlanIPClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *macvlanIPClient) Create(o *MacvlanIP) (*MacvlanIP, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*MacvlanIP), err
}

func (s *macvlanIPClient) Get(name string, opts metav1.GetOptions) (*MacvlanIP, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*MacvlanIP), err
}

func (s *macvlanIPClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*MacvlanIP, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*MacvlanIP), err
}

func (s *macvlanIPClient) Update(o *MacvlanIP) (*MacvlanIP, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*MacvlanIP), err
}

func (s *macvlanIPClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *macvlanIPClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *macvlanIPClient) List(opts metav1.ListOptions) (*MacvlanIPList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*MacvlanIPList), err
}

func (s *macvlanIPClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *macvlanIPClient) Patch(o *MacvlanIP, patchType types.PatchType, data []byte, subresources ...string) (*MacvlanIP, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*MacvlanIP), err
}

func (s *macvlanIPClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *macvlanIPClient) AddHandler(ctx context.Context, name string, sync MacvlanIPHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *macvlanIPClient) AddLifecycle(ctx context.Context, name string, lifecycle MacvlanIPLifecycle) {
	sync := NewMacvlanIPLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *macvlanIPClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync MacvlanIPHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *macvlanIPClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle MacvlanIPLifecycle) {
	sync := NewMacvlanIPLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

type MacvlanIPIndexer func(obj *MacvlanIP) ([]string, error)

type MacvlanIPClientCache interface {
	Get(namespace, name string) (*MacvlanIP, error)
	List(namespace string, selector labels.Selector) ([]*MacvlanIP, error)

	Index(name string, indexer MacvlanIPIndexer)
	GetIndexed(name, key string) ([]*MacvlanIP, error)
}

type MacvlanIPClient interface {
	Create(*MacvlanIP) (*MacvlanIP, error)
	Get(namespace, name string, opts metav1.GetOptions) (*MacvlanIP, error)
	Update(*MacvlanIP) (*MacvlanIP, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*MacvlanIPList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() MacvlanIPClientCache

	OnCreate(ctx context.Context, name string, sync MacvlanIPChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync MacvlanIPChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync MacvlanIPChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() MacvlanIPInterface
}

type macvlanIPClientCache struct {
	client *macvlanIPClient2
}

type macvlanIPClient2 struct {
	iface      MacvlanIPInterface
	controller MacvlanIPController
}

func (n *macvlanIPClient2) Interface() MacvlanIPInterface {
	return n.iface
}

func (n *macvlanIPClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *macvlanIPClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *macvlanIPClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *macvlanIPClient2) Create(obj *MacvlanIP) (*MacvlanIP, error) {
	return n.iface.Create(obj)
}

func (n *macvlanIPClient2) Get(namespace, name string, opts metav1.GetOptions) (*MacvlanIP, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *macvlanIPClient2) Update(obj *MacvlanIP) (*MacvlanIP, error) {
	return n.iface.Update(obj)
}

func (n *macvlanIPClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *macvlanIPClient2) List(namespace string, opts metav1.ListOptions) (*MacvlanIPList, error) {
	return n.iface.List(opts)
}

func (n *macvlanIPClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *macvlanIPClientCache) Get(namespace, name string) (*MacvlanIP, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *macvlanIPClientCache) List(namespace string, selector labels.Selector) ([]*MacvlanIP, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *macvlanIPClient2) Cache() MacvlanIPClientCache {
	n.loadController()
	return &macvlanIPClientCache{
		client: n,
	}
}

func (n *macvlanIPClient2) OnCreate(ctx context.Context, name string, sync MacvlanIPChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &macvlanIPLifecycleDelegate{create: sync})
}

func (n *macvlanIPClient2) OnChange(ctx context.Context, name string, sync MacvlanIPChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &macvlanIPLifecycleDelegate{update: sync})
}

func (n *macvlanIPClient2) OnRemove(ctx context.Context, name string, sync MacvlanIPChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &macvlanIPLifecycleDelegate{remove: sync})
}

func (n *macvlanIPClientCache) Index(name string, indexer MacvlanIPIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*MacvlanIP); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *macvlanIPClientCache) GetIndexed(name, key string) ([]*MacvlanIP, error) {
	var result []*MacvlanIP
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*MacvlanIP); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *macvlanIPClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type macvlanIPLifecycleDelegate struct {
	create MacvlanIPChangeHandlerFunc
	update MacvlanIPChangeHandlerFunc
	remove MacvlanIPChangeHandlerFunc
}

func (n *macvlanIPLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *macvlanIPLifecycleDelegate) Create(obj *MacvlanIP) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *macvlanIPLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *macvlanIPLifecycleDelegate) Remove(obj *MacvlanIP) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *macvlanIPLifecycleDelegate) Updated(obj *MacvlanIP) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
