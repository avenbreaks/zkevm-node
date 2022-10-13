package sequencer

import (
	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricPrefix                 = "sequencer_"
	sequenceTotalCountMetricName = metricPrefix + "sequence"
)

func (s *Sequencer) registerMetrics() {
<<<<<<< Updated upstream
	if s.metrics == nil {
		return
	}

=======
	if !s.metricsEnabled {
		return
	}
>>>>>>> Stashed changes
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
	if !s.metricsEnabled {
		return
	}

	s.metrics.CounterAdd(sequenceTotalCountMetricName, value)
}
