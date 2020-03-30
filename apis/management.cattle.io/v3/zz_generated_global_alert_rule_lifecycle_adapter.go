package v3

import (
	"github.com/rancher/norman/lifecycle"
	"github.com/rancher/norman/resource"
	"k8s.io/apimachinery/pkg/runtime"
)

type GlobalAlertRuleLifecycle interface {
	Create(obj *GlobalAlertRule) (runtime.Object, error)
	Remove(obj *GlobalAlertRule) (runtime.Object, error)
	Updated(obj *GlobalAlertRule) (runtime.Object, error)
}

type globalAlertRuleLifecycleAdapter struct {
	lifecycle GlobalAlertRuleLifecycle
}

func (w *globalAlertRuleLifecycleAdapter) HasCreate() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasCreate()
}

func (w *globalAlertRuleLifecycleAdapter) HasFinalize() bool {
	o, ok := w.lifecycle.(lifecycle.ObjectLifecycleCondition)
	return !ok || o.HasFinalize()
}

func (w *globalAlertRuleLifecycleAdapter) Create(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Create(obj.(*GlobalAlertRule))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *globalAlertRuleLifecycleAdapter) Finalize(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Remove(obj.(*GlobalAlertRule))
	if o == nil {
		return nil, err
	}
	return o, err
}

func (w *globalAlertRuleLifecycleAdapter) Updated(obj runtime.Object) (runtime.Object, error) {
	o, err := w.lifecycle.Updated(obj.(*GlobalAlertRule))
	if o == nil {
		return nil, err
	}
	return o, err
}

func NewGlobalAlertRuleLifecycleAdapter(name string, clusterScoped bool, client GlobalAlertRuleInterface, l GlobalAlertRuleLifecycle) GlobalAlertRuleHandlerFunc {
	if clusterScoped {
		resource.PutClusterScoped(GlobalAlertRuleGroupVersionResource)
	}
	adapter := &globalAlertRuleLifecycleAdapter{lifecycle: l}
	syncFn := lifecycle.NewObjectLifecycleAdapter(name, clusterScoped, adapter, client.ObjectClient())
	return func(key string, obj *GlobalAlertRule) (runtime.Object, error) {
		newObj, err := syncFn(key, obj)
		if o, ok := newObj.(runtime.Object); ok {
			return o, err
		}
		return nil, err
	}
}
