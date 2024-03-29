package v1

import (
	v1 "github.com/F5Networks/k8s-bigip-ctlr/config/apis/cis/v1"
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type ExternalDNSLifecycle interface {
	Create(obj *v1.ExternalDNS) (runtime.Object, error)
	Remove(obj *v1.ExternalDNS) (runtime.Object, error)
	Updated(obj *v1.ExternalDNS) (runtime.Object, error)
}

type externalDNSLifecycleAdapter struct {
	lifecycle ExternalDNSLifecycle
}

func (w *externalDNSLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *externalDNSLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *externalDNSLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v1.ExternalDNS))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *externalDNSLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v1.ExternalDNS))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *externalDNSLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v1.ExternalDNS))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewExternalDNSLifecycleAdapter(name string, clusterScoped bool, client ExternalDNSInterface, l ExternalDNSLifecycle) ExternalDNSHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(ExternalDNSGroupVersionResource)
	}
	adapter := &externalDNSLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v1.ExternalDNS) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
