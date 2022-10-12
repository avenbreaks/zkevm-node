package jsonrpc

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricPrefix        = "jsonrpc_"
	metricRequestPrefix = metricPrefix + "request_"
	requestsMetricName  = metricRequestPrefix + "counter"
	requestDurationName = metricRequestPrefix + "duration"

	invalidRequestMetricLabel = metricRequestPrefix + "invalid"
	singleRequestMetricLabel  = metricRequestPrefix + "single"
	batchRequestMetricLabel   = metricRequestPrefix + "batch"
	totalMetricLabel          = metricRequestPrefix + "total"
)

func (s *Server) registerMetrics() {
	var (
		// gauges     map[string]prometheus.GaugeOpts
		counters   []prometheus.CounterOpts
		histograms []prometheus.HistogramOpts
		// summaries  map[string]prometheus.SummaryOpts
	)

	counters = []prometheus.CounterOpts{
		{
			Name: requestsMetricName,
			Help: "JSONRPC number of requests received",
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

	s.metrics.RegisterCounters(counters...)
	s.metrics.RegisterHistograms(histograms...)
}

func (s *Server) requestMetricInc(label string) {
	if s.metrics == nil {
		return
	}

	s.metrics.IncCounterVec(requestsMetricName, label)
	s.metrics.IncCounterVec(requestsMetricName, totalMetricLabel)
}

func (s *Server) requestDurationMetric(start time.Time) {
	if s.metrics == nil {
		return
	}

	s.metrics.ObserveHistogram(requestDurationName, start)
}
