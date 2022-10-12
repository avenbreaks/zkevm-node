package metrics

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/stretchr/testify/assert"
)

var (
	gaugeName      = "gaugeName"
	gaugeOpts      = prometheus.GaugeOpts{Name: gaugeName}
	gauge          = prometheus.NewGauge(gaugeOpts)
	counterName    = "counterName"
	counterOpts    = prometheus.CounterOpts{Name: counterName}
	counter        = prometheus.NewCounter(counterOpts)
	counterVecName = "counterVecName"
	counterVecOpts = CounterVecOpts{prometheus.CounterOpts{Name: counterVecName}, []string{}}
	counterVec     = prometheus.NewCounterVec(counterVecOpts.CounterOpts, counterVecOpts.labels)
	histogramName  = "histogramName"
	histogramOpts  = prometheus.HistogramOpts{Name: histogramName}
	histogram      = prometheus.NewHistogram(histogramOpts)
	summaryName    = "summaryName"
	summaryOpts    = prometheus.SummaryOpts{Name: summaryName}
	summary        = prometheus.NewSummary(summaryOpts)
)

func newTestEnv() *Prometheus {
	p := NewPrometheus()
	// Overriding registerer to be able to do the unit tests independently
	p.registerer = prometheus.NewRegistry()
	return p
}

func TestHandler(t *testing.T) {
	instance := newTestEnv()

	actual := instance.Handler()

	assert.NotNil(t, actual)
}

func TestRegisterGauges(t *testing.T) {
	instance := newTestEnv()
	gaugesOpts := []prometheus.GaugeOpts{gaugeOpts}

	instance.RegisterGauges(gaugesOpts...)

	assert.Len(t, instance.gauges, 1)
}

func TestGetGauge(t *testing.T) {
	instance := newTestEnv()
	instance.gauges[gaugeName] = gauge

	actual, exist := instance.GetGauge(gaugeName)

	assert.True(t, exist)
	assert.Equal(t, gauge, actual)
}

func TestUnregisterGauges(t *testing.T) {
	instance := newTestEnv()
	instance.RegisterGauges(gaugeOpts)

	instance.UnregisterGauges(gaugeName)

	assert.Len(t, instance.gauges, 0)
}

func TestRegisterCounters(t *testing.T) {
	instance := newTestEnv()
	countersOpts := []prometheus.CounterOpts{counterOpts}

	instance.RegisterCounters(countersOpts...)

	assert.Len(t, instance.counters, 1)
}

func TestGetCounter(t *testing.T) {
	instance := newTestEnv()
	instance.counters[counterName] = counter

	actual, exist := instance.GetCounter(counterName)

	assert.True(t, exist)
	assert.Equal(t, counter, actual)
}

func TestUnregisterCounters(t *testing.T) {
	instance := newTestEnv()
	instance.RegisterCounters(counterOpts)

	instance.UnregisterCounters(counterName)

	assert.Len(t, instance.counters, 0)
}

func TestRegisterCounterVecs(t *testing.T) {
	instance := newTestEnv()
	counterVecsOpts := []CounterVecOpts{counterVecOpts}

	instance.RegisterCounterVecs(counterVecsOpts...)

	assert.Len(t, instance.counterVecs, 1)
}

func TestGetCounterVec(t *testing.T) {
	instance := newTestEnv()
	instance.counterVecs[counterVecName] = counterVec

	actual, exist := instance.GetCounterVec(counterVecName)

	assert.True(t, exist)
	assert.Equal(t, counterVec, actual)
}

func TestUnregisterCounterVecs(t *testing.T) {
	instance := newTestEnv()
	instance.RegisterCounterVecs(counterVecOpts)

	instance.UnregisterCounterVecs(counterVecName)

	assert.Len(t, instance.counterVecs, 0)
}

func TestRegisterHistograms(t *testing.T) {
	instance := newTestEnv()
	histogramsOpts := []prometheus.HistogramOpts{histogramOpts}

	instance.RegisterHistograms(histogramsOpts...)

	assert.Len(t, instance.histograms, 1)
}

func TestGetHistogram(t *testing.T) {
	instance := newTestEnv()
	instance.histograms[histogramName] = histogram

	actual, exist := instance.GetHistogram(histogramName)

	assert.True(t, exist)
	assert.Equal(t, histogram, actual)
}

func TestUnregisterHistograms(t *testing.T) {
	instance := newTestEnv()
	instance.RegisterHistograms(histogramOpts)

	instance.UnregisterHistogram(histogramName)

	assert.Len(t, instance.histograms, 0)
}

func TestRegisterSummaries(t *testing.T) {
	instance := newTestEnv()
	summariesOpts := []prometheus.SummaryOpts{summaryOpts}

	instance.RegisterSummaries(summariesOpts...)

	assert.Len(t, instance.summaries, 1)
}

func TestGetSummary(t *testing.T) {
	instance := newTestEnv()
	instance.summaries[summaryName] = summary

	actual, exist := instance.GetSummary(summaryName)

	assert.True(t, exist)
	assert.Equal(t, summary, actual)
}

func TestUnregisterSummaries(t *testing.T) {
	instance := newTestEnv()
	instance.RegisterSummaries(summaryOpts)

	instance.UnregisterSummaries(summaryName)

	assert.Len(t, instance.summaries, 0)
}
