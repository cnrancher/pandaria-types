package v3

import (
	"github.com/rancher/norman/lifecycle"
	"k8s.io/apimachinery/pkg/runtime"
)

type MacvlanSubnetLifecycle interface {
	Create(obj *MacvlanSubnet) (runtime.Object, error)
	Remove(obj *MacvlanSubnet) (runtime.Object, error)
	Updated(obj *MacvlanSubnet) (runtime.Object, error)
}

type macvlanSubnetLifecycleAdapter struct {
	lifecycle MacvlanSubnetLifecycle
}

func (w *macvlanSubnetLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *macvlanSubnetLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *macvlanSubnetLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*MacvlanSubnet))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *macvlanSubnetLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*MacvlanSubnet))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *macvlanSubnetLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*MacvlanSubnet))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewMacvlanSubnetLifecycleAdapter(name string, clusterScoped bool, client MacvlanSubnetInterface, l MacvlanSubnetLifecycle) MacvlanSubnetHandlerFunc {
	adapter := &macvlanSubnetLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *MacvlanSubnet) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
