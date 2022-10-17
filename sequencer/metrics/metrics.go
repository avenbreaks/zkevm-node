package metrics

import (
	"github.com/0xPolygonHermez/zkevm-node/metrics"
	"github.com/prometheus/client_golang/prometheus"
)

const (
	metricPrefix                       = "sequencer_"
	metricBatchesLastNumberName        = metricPrefix + "batches_last_number"
	metricSequencesSentToL1CountName   = metricPrefix + "sequences_sent_to_L1_count"
	metricGasPriceEstimatedAverageName = metricPrefix + "gas_price_estimated_average"
)

// Register the metrics for the sequencer package.
func Register() {
	var (
		counters []prometheus.CounterOpts
		gauges   []prometheus.GaugeOpts
	)

	counters = []prometheus.CounterOpts{
		{
			Name: metricSequencesSentToL1CountName,
			Help: "[SEQUENCER] total count of sequences sent to L1",
		},
	}

	gauges = []prometheus.GaugeOpts{
		{
			Name: metricBatchesLastNumberName,
			Help: "[SEQUENCER] last batch number processed",
		},
		{
			Name: metricGasPriceEstimatedAverageName,
			Help: "[SEQUENCER] average gas price estimated",
		},
	}

	metrics.RegisterCounters(counters...)
	metrics.RegisterGauges(gauges...)
}

// LastSyncedBatchNumber sets the gauge to the provided batch number.
func LastSyncedBatchNumber(batchNum float64) {
	metrics.GaugeSet(metricBatchesLastNumberName, float64(batchNum))
}

// AverageGasPrice sets the gauge to the given average gas price.
func AverageGasPrice(price float64) {
	metrics.GaugeSet(metricGasPriceEstimatedAverageName, float64(price))
}

// SequencesSentToL1 increases the counter by the provided number of sequences
// sent to L1.
func SequencesSentToL1(numSequences float64) {
	metrics.CounterAdd(metricSequencesSentToL1CountName, float64(numSequences))
}
