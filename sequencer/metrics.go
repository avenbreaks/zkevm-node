package sequencer

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricPrefix                 = "sequencer_"
	sequenceTotalCountMetricName = metricPrefix + "sequence"
)

func (s *Sequencer) registerMetrics() {
	var counters []prometheus.CounterOpts

	counters = []prometheus.CounterOpts{
		{
			Name: sequenceTotalCountMetricName,
			Help: "SEQUENCER total sequence processed",
		},
	}
	s.metrics.RegisterCounters(counters...)
}

func (s *Sequencer) batchesMetricAdd(value float64) {
	if s.metrics == nil {
		return
	}

	s.metrics.CounterAdd(sequenceTotalCountMetricName, value)
}
