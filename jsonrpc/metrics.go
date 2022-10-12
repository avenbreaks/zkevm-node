package jsonrpc

import (
	"github.com/prometheus/client_golang/prometheus"
)

func (s *Server) registerMetrics() {
	var (
		// gauges     map[string]prometheus.GaugeOpts
		counters   map[string]prometheus.CounterOpts
		histograms map[string]prometheus.HistogramOpts
		// summaries  map[string]prometheus.SummaryOpts
	)

	counters = map[string]prometheus.CounterOpts{
		"requests": {
			Name: "jsonrpc_requests",
			Help: "JSONRPC number of requests received",
		},
	}

	start := 0.1
	width := 0.1
	count := 10
	histograms = map[string]prometheus.HistogramOpts{
		"requestDuration": {
			Name:    "jsonrpc_request_duration",
			Help:    "JSONRPC Histogram for the runtime of requests",
			Buckets: prometheus.LinearBuckets(start, width, count),
		},
	}

	for _, counter := range counters {
		s.metrics.RegisterCounters(counter)
	}

	for _, histogram := range histograms {
		s.metrics.RegisterHistograms(histogram)
	}
}

func (s *Server) requestHandled(label string) {
	if counterVec, ok := s.metrics.GetCounterVec("requests"); ok {
		counterVec.WithLabelValues(label).Inc()
	}
}

func (s *Server) requestDuration(value float64) {
	if histo, ok := s.metrics.GetHistogram("requestDuration"); ok {
		histo.Observe(value)
	}
}
