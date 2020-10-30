package v3

import (
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type SensitiveFilterLifecycle interface {
	Create(obj *SensitiveFilter) (runtime.Object, error)
	Remove(obj *SensitiveFilter) (runtime.Object, error)
	Updated(obj *SensitiveFilter) (runtime.Object, error)
}

type sensitiveFilterLifecycleAdapter struct {
	lifecycle SensitiveFilterLifecycle
}

func (w *sensitiveFilterLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *sensitiveFilterLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *sensitiveFilterLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*SensitiveFilter))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *sensitiveFilterLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*SensitiveFilter))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *sensitiveFilterLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*SensitiveFilter))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewSensitiveFilterLifecycleAdapter(name string, clusterScoped bool, client SensitiveFilterInterface, l SensitiveFilterLifecycle) SensitiveFilterHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(SensitiveFilterGroupVersionResource)
	}
	adapter := &sensitiveFilterLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *SensitiveFilter) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
