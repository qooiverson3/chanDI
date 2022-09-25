package service

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	opsProcessed = prometheus.NewCounterVec(prometheus.CounterOpts{
		Namespace: "myapp",
		Name:      "processed_ops_total",
		Help:      "The total number of processed events",
	}, []string{"name"})
)

func init() {
	prometheus.MustRegister(opsProcessed)
}

func (s service) Exporter(n string) {
	go func() {
		for {
			opsProcessed.WithLabelValues(n).Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}
