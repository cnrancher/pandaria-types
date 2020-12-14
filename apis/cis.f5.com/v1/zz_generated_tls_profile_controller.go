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
	TLSProfileGroupVersionKind = schema.GroupVersionKind{
		Version: Version,
		Group:   GroupName,
		Kind:    "TLSProfile",
	}
	TLSProfileResource = metav1.APIResource{
		Name:         "tlsprofiles",
		SingularName: "tlsprofile",
		Namespaced:   true,

		Kind: TLSProfileGroupVersionKind.Kind,
	}

	TLSProfileGroupVersionResource = schema.GroupVersionResource{
		Group:    GroupName,
		Version:  Version,
		Resource: "tlsprofiles",
	}
)

func init() {
	resource.Put(TLSProfileGroupVersionResource)
}

func NewTLSProfile(namespace, name string, obj v1.TLSProfile) *v1.TLSProfile {
	obj.APIVersion, obj.Kind = TLSProfileGroupVersionKind.ToAPIVersionAndKind()
	obj.Name = name
	obj.Namespace = namespace
	return &obj
}

type TLSProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []v1.TLSProfile `json:"items"`
}

type TLSProfileHandlerFunc func(key string, obj *v1.TLSProfile) (runtime.Object, error)

type TLSProfileChangeHandlerFunc func(obj *v1.TLSProfile) (runtime.Object, error)

type TLSProfileLister interface {
	List(namespace string, selector labels.Selector) (ret []*v1.TLSProfile, err error)
	Get(namespace, name string) (*v1.TLSProfile, error)
}

type TLSProfileController interface {
	Generic() controller.GenericController
	Informer() cache.SharedIndexInformer
	Lister() TLSProfileLister
	AddHandler(ctx context.Context, name string, handler TLSProfileHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync TLSProfileHandlerFunc)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, handler TLSProfileHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, handler TLSProfileHandlerFunc)
	Enqueue(namespace, name string)
	EnqueueAfter(namespace, name string, after time.Duration)
	Sync(ctx context.Context) error
	Start(ctx context.Context, threadiness int) error
}

type TLSProfileInterface interface {
	ObjectClient() *objectclient.ObjectClient
	Create(*v1.TLSProfile) (*v1.TLSProfile, error)
	GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.TLSProfile, error)
	Get(name string, opts metav1.GetOptions) (*v1.TLSProfile, error)
	Update(*v1.TLSProfile) (*v1.TLSProfile, error)
	Delete(name string, options *metav1.DeleteOptions) error
	DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error
	List(opts metav1.ListOptions) (*TLSProfileList, error)
	ListNamespaced(namespace string, opts metav1.ListOptions) (*TLSProfileList, error)
	Watch(opts metav1.ListOptions) (watch.Interface, error)
	DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Controller() TLSProfileController
	AddHandler(ctx context.Context, name string, sync TLSProfileHandlerFunc)
	AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync TLSProfileHandlerFunc)
	AddLifecycle(ctx context.Context, name string, lifecycle TLSProfileLifecycle)
	AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle TLSProfileLifecycle)
	AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync TLSProfileHandlerFunc)
	AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync TLSProfileHandlerFunc)
	AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle TLSProfileLifecycle)
	AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle TLSProfileLifecycle)
}

type tlsProfileLister struct {
	controller *tlsProfileController
}

func (l *tlsProfileLister) List(namespace string, selector labels.Selector) (ret []*v1.TLSProfile, err error) {
	err = cache.ListAllByNamespace(l.controller.Informer().GetIndexer(), namespace, selector, func(obj interface{}) {
		ret = append(ret, obj.(*v1.TLSProfile))
	})
	return
}

func (l *tlsProfileLister) Get(namespace, name string) (*v1.TLSProfile, error) {
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
			Group:    TLSProfileGroupVersionKind.Group,
			Resource: "tlsProfile",
		}, key)
	}
	return obj.(*v1.TLSProfile), nil
}

type tlsProfileController struct {
	controller.GenericController
}

func (c *tlsProfileController) Generic() controller.GenericController {
	return c.GenericController
}

func (c *tlsProfileController) Lister() TLSProfileLister {
	return &tlsProfileLister{
		controller: c,
	}
}

func (c *tlsProfileController) AddHandler(ctx context.Context, name string, handler TLSProfileHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.TLSProfile); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *tlsProfileController) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, handler TLSProfileHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.TLSProfile); ok {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *tlsProfileController) AddClusterScopedHandler(ctx context.Context, name, cluster string, handler TLSProfileHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.TLSProfile); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

func (c *tlsProfileController) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, cluster string, handler TLSProfileHandlerFunc) {
	c.GenericController.AddHandler(ctx, name, func(key string, obj interface{}) (interface{}, error) {
		if !enabled() {
			return nil, nil
		} else if obj == nil {
			return handler(key, nil)
		} else if v, ok := obj.(*v1.TLSProfile); ok && controller.ObjectInCluster(cluster, obj) {
			return handler(key, v)
		} else {
			return nil, nil
		}
	})
}

type tlsProfileFactory struct {
}

func (c tlsProfileFactory) Object() runtime.Object {
	return &v1.TLSProfile{}
}

func (c tlsProfileFactory) List() runtime.Object {
	return &TLSProfileList{}
}

func (s *tlsProfileClient) Controller() TLSProfileController {
	s.client.Lock()
	defer s.client.Unlock()

	c, ok := s.client.tlsProfileControllers[s.ns]
	if ok {
		return c
	}

	genericController := controller.NewGenericController(TLSProfileGroupVersionKind.Kind+"Controller",
		s.objectClient)

	c = &tlsProfileController{
		GenericController: genericController,
	}

	s.client.tlsProfileControllers[s.ns] = c
	s.client.starters = append(s.client.starters, c)

	return c
}

type tlsProfileClient struct {
	client       *Client
	ns           string
	objectClient *objectclient.ObjectClient
	controller   TLSProfileController
}

func (s *tlsProfileClient) ObjectClient() *objectclient.ObjectClient {
	return s.objectClient
}

func (s *tlsProfileClient) Create(o *v1.TLSProfile) (*v1.TLSProfile, error) {
	obj, err := s.objectClient.Create(o)
	return obj.(*v1.TLSProfile), err
}

func (s *tlsProfileClient) Get(name string, opts metav1.GetOptions) (*v1.TLSProfile, error) {
	obj, err := s.objectClient.Get(name, opts)
	return obj.(*v1.TLSProfile), err
}

func (s *tlsProfileClient) GetNamespaced(namespace, name string, opts metav1.GetOptions) (*v1.TLSProfile, error) {
	obj, err := s.objectClient.GetNamespaced(namespace, name, opts)
	return obj.(*v1.TLSProfile), err
}

func (s *tlsProfileClient) Update(o *v1.TLSProfile) (*v1.TLSProfile, error) {
	obj, err := s.objectClient.Update(o.Name, o)
	return obj.(*v1.TLSProfile), err
}

func (s *tlsProfileClient) Delete(name string, options *metav1.DeleteOptions) error {
	return s.objectClient.Delete(name, options)
}

func (s *tlsProfileClient) DeleteNamespaced(namespace, name string, options *metav1.DeleteOptions) error {
	return s.objectClient.DeleteNamespaced(namespace, name, options)
}

func (s *tlsProfileClient) List(opts metav1.ListOptions) (*TLSProfileList, error) {
	obj, err := s.objectClient.List(opts)
	return obj.(*TLSProfileList), err
}

func (s *tlsProfileClient) ListNamespaced(namespace string, opts metav1.ListOptions) (*TLSProfileList, error) {
	obj, err := s.objectClient.ListNamespaced(namespace, opts)
	return obj.(*TLSProfileList), err
}

func (s *tlsProfileClient) Watch(opts metav1.ListOptions) (watch.Interface, error) {
	return s.objectClient.Watch(opts)
}

// Patch applies the patch and returns the patched deployment.
func (s *tlsProfileClient) Patch(o *v1.TLSProfile, patchType types.PatchType, data []byte, subresources ...string) (*v1.TLSProfile, error) {
	obj, err := s.objectClient.Patch(o.Name, o, patchType, data, subresources...)
	return obj.(*v1.TLSProfile), err
}

func (s *tlsProfileClient) DeleteCollection(deleteOpts *metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	return s.objectClient.DeleteCollection(deleteOpts, listOpts)
}

func (s *tlsProfileClient) AddHandler(ctx context.Context, name string, sync TLSProfileHandlerFunc) {
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *tlsProfileClient) AddFeatureHandler(ctx context.Context, enabled func() bool, name string, sync TLSProfileHandlerFunc) {
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *tlsProfileClient) AddLifecycle(ctx context.Context, name string, lifecycle TLSProfileLifecycle) {
	sync := NewTLSProfileLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddHandler(ctx, name, sync)
}

func (s *tlsProfileClient) AddFeatureLifecycle(ctx context.Context, enabled func() bool, name string, lifecycle TLSProfileLifecycle) {
	sync := NewTLSProfileLifecycleAdapter(name, false, s, lifecycle)
	s.Controller().AddFeatureHandler(ctx, enabled, name, sync)
}

func (s *tlsProfileClient) AddClusterScopedHandler(ctx context.Context, name, clusterName string, sync TLSProfileHandlerFunc) {
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *tlsProfileClient) AddClusterScopedFeatureHandler(ctx context.Context, enabled func() bool, name, clusterName string, sync TLSProfileHandlerFunc) {
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}

func (s *tlsProfileClient) AddClusterScopedLifecycle(ctx context.Context, name, clusterName string, lifecycle TLSProfileLifecycle) {
	sync := NewTLSProfileLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedHandler(ctx, name, clusterName, sync)
}

func (s *tlsProfileClient) AddClusterScopedFeatureLifecycle(ctx context.Context, enabled func() bool, name, clusterName string, lifecycle TLSProfileLifecycle) {
	sync := NewTLSProfileLifecycleAdapter(name+"_"+clusterName, true, s, lifecycle)
	s.Controller().AddClusterScopedFeatureHandler(ctx, enabled, name, clusterName, sync)
}
