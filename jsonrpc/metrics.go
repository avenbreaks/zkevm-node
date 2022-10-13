package jsonrpc

import (
	"time"

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
	requestMetricLabelSingle                     = "single"
	requestMetricLabelBatch                      = "batch"
)

func (s *Server) registerMetrics() {
<<<<<<< Updated upstream
	if s.metrics == nil {
=======
	if !s.metricsEnabled {
>>>>>>> Stashed changes
		return
	}

	var (
		// gauges     map[string]prometheus.GaugeOpts
		counterVecs []metrics.CounterVecOpts
		histograms  []prometheus.HistogramOpts
		// summaries  map[string]prometheus.SummaryOpts
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

	s.metrics.RegisterCounterVecs(counterVecs...)
	s.metrics.RegisterHistograms(histograms...)
}

func (s *Server) requestMetricInc(label requestMetricLabel) {
	if !s.metricsEnabled {
		return
	}

	s.metrics.CounterVecInc(requestsMetricName, string(label))
}

func (s *Server) requestDurationMetric(start time.Time) {
	if !s.metricsEnabled {
		return
	}

	s.metrics.ObserveHistogram(requestDurationName, start)
}
