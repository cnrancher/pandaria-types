package v3

import (
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type GlobalAlertGroupLifecycle interface {
	Create(obj *GlobalAlertGroup) (runtime.Object, error)
	Remove(obj *GlobalAlertGroup) (runtime.Object, error)
	Updated(obj *GlobalAlertGroup) (runtime.Object, error)
}

type globalAlertGroupLifecycleAdapter struct {
	lifecycle GlobalAlertGroupLifecycle
}

func (w *globalAlertGroupLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *globalAlertGroupLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *globalAlertGroupLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*GlobalAlertGroup))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *globalAlertGroupLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*GlobalAlertGroup))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *globalAlertGroupLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*GlobalAlertGroup))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewGlobalAlertGroupLifecycleAdapter(name string, clusterScoped bool, client GlobalAlertGroupInterface, l GlobalAlertGroupLifecycle) GlobalAlertGroupHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(GlobalAlertGroupGroupVersionResource)
	}
	adapter := &globalAlertGroupLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *GlobalAlertGroup) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
