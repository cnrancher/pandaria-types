package status

import (
	"os"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rancher/norman/types/slice"
)

// PANDARIA: add cluster state metrics
var (
	ClusterState = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Subsystem: "pandaria_cluster",
			Name:      "unavailable",
			Help:      "Set a cluster is unavailable",
		},
		[]string{"cluster", "name"},
	)

	customCheckList = os.Getenv("PANDARIA_METRICS_CLUSTER_STATE_CHECK")
	checkList       = []string{
		"error",
		"unavailable",
		"failed",
	}

	PrometheusMetrics = false
)

func setClusterState(state, name, clusterID string) {
	if customCheckList != "" {
		checkList = strings.Split(customCheckList, ",")
	}

	value := 0
	if slice.ContainsString(checkList, state) {
		value = 1
	}

	ClusterState.With(
		prometheus.Labels{
			"cluster": clusterID,
			"name":    name,
		}).Set(float64(value))
}
