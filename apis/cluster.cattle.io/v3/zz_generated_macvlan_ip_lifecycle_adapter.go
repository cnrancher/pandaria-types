package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type MacvlanIPLifecycle interface {
	Create(obj *MacvlanIP) (runtime.Object, error)
	Remove(obj *MacvlanIP) (runtime.Object, error)
	Updated(obj *MacvlanIP) (runtime.Object, error)
}

type macvlanIPLifecycleAdapter struct {
	lifecycle MacvlanIPLifecycle
}

func (w *macvlanIPLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *macvlanIPLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *macvlanIPLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*MacvlanIP))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *macvlanIPLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*MacvlanIP))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *macvlanIPLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*MacvlanIP))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewMacvlanIPLifecycleAdapter(name string, clusterScoped bool, client MacvlanIPInterface, l MacvlanIPLifecycle) MacvlanIPHandlerFunc {
	adapter := &macvlanIPLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *MacvlanIP) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
