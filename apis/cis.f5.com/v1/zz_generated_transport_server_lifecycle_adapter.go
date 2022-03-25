package v1

import (
	"github.com/F5Networks/k8s-bigip-ctlr/config/apis/cis/v1"
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type TransportServerLifecycle interface {
	Create(obj *v1.TransportServer) (runtime.Object, error)
	Remove(obj *v1.TransportServer) (runtime.Object, error)
	Updated(obj *v1.TransportServer) (runtime.Object, error)
}

type transportServerLifecycleAdapter struct {
	lifecycle TransportServerLifecycle
}

func (w *transportServerLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *transportServerLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *transportServerLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v1.TransportServer))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *transportServerLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v1.TransportServer))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *transportServerLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v1.TransportServer))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewTransportServerLifecycleAdapter(name string, clusterScoped bool, client TransportServerInterface, l TransportServerLifecycle) TransportServerHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(TransportServerGroupVersionResource)
	}
	adapter := &transportServerLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v1.TransportServer) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
