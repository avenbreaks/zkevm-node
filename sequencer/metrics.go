package sequencer

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricPrefix                = "sequencer_"
	batchesTotalCountMetricName = metricPrefix + "batch"
)

func (s *Sequencer) registerMetrics() {
	var (
		counters []prometheus.CounterOpts
	)

	counters = []prometheus.CounterOpts{
		{
			Name: batchesTotalCountMetricName,
			Help: "Total batches processed",
		},
	}
	s.metrics.RegisterCounters(counters...)
}

func (s *Sequencer) batchesMetricAdd(value float64) {
	if s.metrics == nil {
		return
	}

	s.metrics.CounterAdd(batchesTotalCountMetricName, value)
}
