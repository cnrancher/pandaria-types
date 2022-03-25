package v1

import (
	"github.com/F5Networks/k8s-bigip-ctlr/config/apis/cis/v1"
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type TLSProfileLifecycle interface {
	Create(obj *v1.TLSProfile) (runtime.Object, error)
	Remove(obj *v1.TLSProfile) (runtime.Object, error)
	Updated(obj *v1.TLSProfile) (runtime.Object, error)
}

type tlsProfileLifecycleAdapter struct {
	lifecycle TLSProfileLifecycle
}

func (w *tlsProfileLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *tlsProfileLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *tlsProfileLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v1.TLSProfile))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *tlsProfileLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v1.TLSProfile))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *tlsProfileLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v1.TLSProfile))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewTLSProfileLifecycleAdapter(name string, clusterScoped bool, client TLSProfileInterface, l TLSProfileLifecycle) TLSProfileHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(TLSProfileGroupVersionResource)
	}
	adapter := &tlsProfileLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v1.TLSProfile) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
