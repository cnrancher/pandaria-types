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
	GlobalAlertRuleGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "GlobalAlertRule",
	}
	GlobalAlertRuleResource = metav1.APIResource{
		Name:         "globalalertrules",
		SingularName: "globalalertrule",
		Namespaced:   true,

		Kind: GlobalAlertRuleGroupVersionKind.Kind,
	}

	GlobalAlertRuleGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "globalalertrules",
	}
)

func init() {
	resource.Put(GlobalAlertRuleGroupVersionResource)
}

func NewGlobalAlertRule(namespace, name string, obj GlobalAlertRule) *GlobalAlertRule {
	obj.APIVersion, obj.Kind = GlobalAlertRuleGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type GlobalAlertRuleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []GlobalAlertRule `json:"items"`
}

type GlobalAlertRuleHandlerFunc func(key string, obj *GlobalAlertRule) (runtime.Object, error)

type GlobalAlertRuleChangeHandlerFunc func(obj *GlobalAlertRule) (runtime.Object, error)

type GlobalAlertRuleLister interface {
	List(namespace string, selector labels.Selector) (ret []*GlobalAlertRule, err error)
	Get(namespace, name string) (*GlobalAlertRule, error)
}

type GlobalAlertRuleController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() GlobalAlertRuleLister
	AddHandler(ctx context.Context, name string, handler GlobalAlertRuleHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync GlobalAlertRuleHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler GlobalAlertRuleHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler GlobalAlertRuleHandlerFunc)
	Enqueue(namespace, name string)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type GlobalAlertRuleInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*GlobalAlertRule) (*GlobalAlertRule, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*GlobalAlertRule, error)
	Get(name string, opts metav1.GetOptions) (*GlobalAlertRule, error)
	Update(*GlobalAlertRule) (*GlobalAlertRule, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*GlobalAlertRuleList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() GlobalAlertRuleController
	AddHandler(ctx context.Context, name string, sync GlobalAlertRuleHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync GlobalAlertRuleHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle GlobalAlertRuleLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle GlobalAlertRuleLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync GlobalAlertRuleHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync GlobalAlertRuleHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle GlobalAlertRuleLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle GlobalAlertRuleLifecycle)
}

type globalAlertRuleLister struct {
	controller *globalAlertRuleController
}

func (l *globalAlertRuleLister) List(namespace string, selector labels.Selector) (ret []*GlobalAlertRule, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*GlobalAlertRule))
	})
	return
}

func (l *globalAlertRuleLister) Get(namespace, name string) (*GlobalAlertRule, error) {
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
			Group:    GlobalAlertRuleGroupVersionKind.Group,
			Resource: "globalAlertRule",
		}, key)
	}
	return obj.(*GlobalAlertRule), nil
}

type globalAlertRuleController struct {
	controller.GenericController
}

func (c *globalAlertRuleController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *globalAlertRuleController) Lister() GlobalAlertRuleLister {
	return &globalAlertRuleLister{
		controller: c,
	}
}

func (c *globalAlertRuleController) AddHandler(ctx context.Context, name string, handler GlobalAlertRuleHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*GlobalAlertRule); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *globalAlertRuleController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler GlobalAlertRuleHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*GlobalAlertRule); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *globalAlertRuleController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler GlobalAlertRuleHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*GlobalAlertRule); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *globalAlertRuleController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler GlobalAlertRuleHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*GlobalAlertRule); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type globalAlertRuleFactory struct {
}

func (c globalAlertRuleFactory) Object() runtime.Object {
	return &GlobalAlertRule{}
}

func (c globalAlertRuleFactory) List() runtime.Object {
	return &GlobalAlertRuleList{}
}

func (s *globalAlertRuleClient) Controller() GlobalAlertRuleController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.globalAlertRuleControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(GlobalAlertRuleGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &globalAlertRuleController{
		GenericController: genericController,
	}

	s.client.globalAlertRuleControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type globalAlertRuleClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   GlobalAlertRuleController
}

func (s *globalAlertRuleClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *globalAlertRuleClient) Create(o *GlobalAlertRule) (*GlobalAlertRule, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*GlobalAlertRule), err
}

func (s *globalAlertRuleClient) Get(name string, opts metav1.GetOptions) (*GlobalAlertRule, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*GlobalAlertRule), err
}

func (s *globalAlertRuleClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*GlobalAlertRule, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*GlobalAlertRule), err
}

func (s *globalAlertRuleClient) Update(o *GlobalAlertRule) (*GlobalAlertRule, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*GlobalAlertRule), err
}

func (s *globalAlertRuleClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *globalAlertRuleClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *globalAlertRuleClient) List(opts metav1.ListOptions) (*GlobalAlertRuleList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*GlobalAlertRuleList), err
}

func (s *globalAlertRuleClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *globalAlertRuleClient) Patch(o *GlobalAlertRule, patchType types.PatchType, data []byte, subresources ...string) (*GlobalAlertRule, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*GlobalAlertRule), err
}

func (s *globalAlertRuleClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *globalAlertRuleClient) AddHandler(ctx context.Context, name string, sync GlobalAlertRuleHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *globalAlertRuleClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync GlobalAlertRuleHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *globalAlertRuleClient) AddLifecycle(ctx context.Context, name string, lifecycle GlobalAlertRuleLifecycle) {
	sync := NewGlobalAlertRuleLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *globalAlertRuleClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle GlobalAlertRuleLifecycle) {
	sync := NewGlobalAlertRuleLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *globalAlertRuleClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync GlobalAlertRuleHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *globalAlertRuleClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync GlobalAlertRuleHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *globalAlertRuleClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle GlobalAlertRuleLifecycle) {
	sync := NewGlobalAlertRuleLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *globalAlertRuleClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle GlobalAlertRuleLifecycle) {
	sync := NewGlobalAlertRuleLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

type GlobalAlertRuleIndexer func(obj *GlobalAlertRule) ([]string, error)

type GlobalAlertRuleClientCache interface {
	Get(namespace, name string) (*GlobalAlertRule, error)
	List(namespace string, selector labels.Selector) ([]*GlobalAlertRule, error)

	Index(name string, indexer GlobalAlertRuleIndexer)
	GetIndexed(name, key string) ([]*GlobalAlertRule, error)
}

type GlobalAlertRuleClient interface {
	Create(*GlobalAlertRule) (*GlobalAlertRule, error)
	Get(namespace, name string, opts metav1.GetOptions) (*GlobalAlertRule, error)
	Update(*GlobalAlertRule) (*GlobalAlertRule, error)
	Delete(namespace, name string, options *metav1.DeleteOptions) error
	List(namespace string, opts metav1.ListOptions) (*GlobalAlertRuleList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)

	Cache() GlobalAlertRuleClientCache

	OnCreate(ctx context.Context, name string, sync GlobalAlertRuleChangeHandlerFunc)
	OnChange(ctx context.Context, name string, sync GlobalAlertRuleChangeHandlerFunc)
	OnRemove(ctx context.Context, name string, sync GlobalAlertRuleChangeHandlerFunc)
	Enqueue(namespace, name string)

	Generic() controller.GenericController
	ObjectClient() *objectclient.ObjectClient
	Interface() GlobalAlertRuleInterface
}

type globalAlertRuleClientCache struct {
	client *globalAlertRuleClient2
}

type globalAlertRuleClient2 struct {
	iface      GlobalAlertRuleInterface
	controller GlobalAlertRuleController
}

func (n *globalAlertRuleClient2) Interface() GlobalAlertRuleInterface {
	return n.iface
}

func (n *globalAlertRuleClient2) Generic() controller.GenericController {
	return n.iface.Controller().Generic()
}

func (n *globalAlertRuleClient2) ObjectClient() *objectclient.ObjectClient {
	return n.Interface().ObjectClient()
}

func (n *globalAlertRuleClient2) Enqueue(namespace, name string) {
	n.iface.Controller().Enqueue(namespace, name)
}

func (n *globalAlertRuleClient2) Create(obj *GlobalAlertRule) (*GlobalAlertRule, error) {
	return n.iface.Create(obj)
}

func (n *globalAlertRuleClient2) Get(namespace, name string, opts metav1.GetOptions) (*GlobalAlertRule, error) {
	return n.iface.GetNamespaced(namespace, name, opts)
}

func (n *globalAlertRuleClient2) Update(obj *GlobalAlertRule) (*GlobalAlertRule, error) {
	return n.iface.Update(obj)
}

func (n *globalAlertRuleClient2) Delete(namespace, name string, options *metav1.DeleteOptions) error {
	return n.iface.DeleteNamespaced(namespace, name, options)
}

func (n *globalAlertRuleClient2) List(namespace string, opts metav1.ListOptions) (*GlobalAlertRuleList, error) {
	return n.iface.List(opts)
}

func (n *globalAlertRuleClient2) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return n.iface.Watch(opts)
}

func (n *globalAlertRuleClientCache) Get(namespace, name string) (*GlobalAlertRule, error) {
	return n.client.controller.Lister().Get(namespace, name)
}

func (n *globalAlertRuleClientCache) List(namespace string, selector labels.Selector) ([]*GlobalAlertRule, error) {
	return n.client.controller.Lister().List(namespace, selector)
}

func (n *globalAlertRuleClient2) Cache() GlobalAlertRuleClientCache {
	n.loadController()
	return &globalAlertRuleClientCache{
		client: n,
	}
}

func (n *globalAlertRuleClient2) OnCreate(ctx context.Context, name string, sync GlobalAlertRuleChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-create", &globalAlertRuleLifecycleDelegate{create: sync})
}

func (n *globalAlertRuleClient2) OnChange(ctx context.Context, name string, sync GlobalAlertRuleChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name+"-change", &globalAlertRuleLifecycleDelegate{update: sync})
}

func (n *globalAlertRuleClient2) OnRemove(ctx context.Context, name string, sync GlobalAlertRuleChangeHandlerFunc) {
	n.loadController()
	n.iface.AddLifecycle(ctx, name, &globalAlertRuleLifecycleDelegate{remove: sync})
}

func (n *globalAlertRuleClientCache) Index(name string, indexer GlobalAlertRuleIndexer) {
	err := n.client.controller.Informer().GetIndexer().AddIndexers(map[string]cache.IndexFunc{
		name: func(obj interface{}) ([]string, error) {
			if v, ok := obj.(*GlobalAlertRule); ok {
				return indexer(v)
			}
			return nil, nil
		},
	})

	if err != nil {
		panic(err)
	}
}

func (n *globalAlertRuleClientCache) GetIndexed(name, key string) ([]*GlobalAlertRule, error) {
	var result []*GlobalAlertRule
	objs, err := n.client.controller.Informer().GetIndexer().ByIndex(name, key)
	if err != nil {
		return nil, err
	}
	for _, obj := range objs {
		if v, ok := obj.(*GlobalAlertRule); ok {
			result = append(result, v)
		}
	}

	return result, nil
}

func (n *globalAlertRuleClient2) loadController() {
	if n.controller == nil {
		n.controller = n.iface.Controller()
	}
}

type globalAlertRuleLifecycleDelegate struct {
	create GlobalAlertRuleChangeHandlerFunc
	update GlobalAlertRuleChangeHandlerFunc
	remove GlobalAlertRuleChangeHandlerFunc
}

func (n *globalAlertRuleLifecycleDelegate) HasCreate() bool {
	return n.create != nil
}

func (n *globalAlertRuleLifecycleDelegate) Create(obj *GlobalAlertRule) (runtime.Object, error) {
	if n.create == nil {
		return obj, nil
	}
	return n.create(obj)
}

func (n *globalAlertRuleLifecycleDelegate) HasFinalize() bool {
	return n.remove != nil
}

func (n *globalAlertRuleLifecycleDelegate) Remove(obj *GlobalAlertRule) (runtime.Object, error) {
	if n.remove == nil {
		return obj, nil
	}
	return n.remove(obj)
}

func (n *globalAlertRuleLifecycleDelegate) Updated(obj *GlobalAlertRule) (runtime.Object, error) {
	if n.update == nil {
		return obj, nil
	}
	return n.update(obj)
}
