package status

import (
	"strings"

	"github.com/rancher/norman/types/convert"
	v3 "github.com/rancher/types/apis/management.cattle.io/v3"
	"github.com/rancher/wrangler/pkg/summary"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

func Set(data map[string]interface{}) {
	if data == nil {
		return
	}
	summary := summary.Summarize(&unstructured.Unstructured{Object: data})
	data["state"] = summary.State
	data["transitioning"] = "no"
	if summary.Error {
		data["transitioning"] = "error"
	} else if summary.Transitioning {
		data["transitioning"] = "yes"
	}
	data["transitioningMessage"] = strings.Join(summary.Message, "; ")
	if data["kind"] == "Cluster" && PrometheusMetrics {
		cluster := &v3.Cluster{}
		if err := convert.ToObj(data, cluster); err == nil {
			setClusterState(summary.State, cluster.Spec.DisplayName, cluster.Name)
		}
	}
}
