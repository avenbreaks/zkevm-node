package jsonrpc

import (
	"github.com/0xPolygonHermez/zkevm-node/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricPrefix        = "jsonrpc_"
	metricRequestPrefix = metricPrefix + "request_"
	requestsMetricName  = metricRequestPrefix + "counter"
	requestDurationName = metricRequestPrefix + "duration"

	requestMetricLabelName = "status"
)

type requestMetricLabel string

const (
	requestMetricLabelInvalid requestMetricLabel = "invalid"
	requestMetricLabelSingle  requestMetricLabel = "single"
	requestMetricLabelBatch   requestMetricLabel = "batch"
)

func registerMetrics() {

	var (
		counterVecs []metrics.CounterVecOpts
		histograms  []prometheus.HistogramOpts
	)

	counterVecs = []metrics.CounterVecOpts{
		{
			CounterOpts: prometheus.CounterOpts{
				Name: requestsMetricName,
				Help: "JSONRPC number of requests received",
			},
			Labels: []string{requestMetricLabelName},
		},
	}

	start := 0.1
	width := 0.1
	count := 10
	histograms = []prometheus.HistogramOpts{
		{
			Name:    requestDurationName,
			Help:    "JSONRPC Histogram for the runtime of requests",
			Buckets: prometheus.LinearBuckets(start, width, count),
		},
	}

	metrics.RegisterCounterVecs(counterVecs...)
	metrics.RegisterHistograms(histograms...)
}

func metricRequestInc(label requestMetricLabel) {
	metrics.CounterVecInc(requestsMetricName, string(label))
}
