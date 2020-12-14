package v1

import (
	v1 "github.com/F5Networks/k8s-bigip-ctlr/config/apis/cis/v1"
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type VirtualServerLifecycle interface {
	Create(obj *v1.VirtualServer) (runtime.Object, error)
	Remove(obj *v1.VirtualServer) (runtime.Object, error)
	Updated(obj *v1.VirtualServer) (runtime.Object, error)
}

type virtualServerLifecycleAdapter struct {
	lifecycle VirtualServerLifecycle
}

func (w *virtualServerLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *virtualServerLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *virtualServerLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*v1.VirtualServer))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *virtualServerLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*v1.VirtualServer))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *virtualServerLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*v1.VirtualServer))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewVirtualServerLifecycleAdapter(name string, clusterScoped bool, client VirtualServerInterface, l VirtualServerLifecycle) VirtualServerHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(VirtualServerGroupVersionResource)
	}
	adapter := &virtualServerLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *v1.VirtualServer) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
