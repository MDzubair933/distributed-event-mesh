package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	MessagesPublished = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "event_mesh_published_total",
			Help: "Total number of events published to topic",
		},
		[]string{"topic", "partition"},
	)
	PublishLatency = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "event_mesh_publish_duration_seconds",
			Help:    "Latency of message publishing into WAL",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"topic"},
	)
)

func init() {
	prometheus.MustRegister(MessagesPublished, PublishLatency)
}
