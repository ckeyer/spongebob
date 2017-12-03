package server

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
// nodeinfoCollector = prometheus.NewSummaryVec{}
)

func init() {
	prometheus.MustRegister()
}
