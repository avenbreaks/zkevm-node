package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var (
	gaugeName           = "gaugeName"
	gaugeOpts           = prometheus.GaugeOpts{Name: gaugeName}
	gauge               = prometheus.NewGauge(gaugeOpts)
	counterName         = "counterName"
	counterOpts         = prometheus.CounterOpts{Name: counterName}
	counter             = prometheus.NewCounter(counterOpts)
	counterVecName      = "counterVecName"
	counterVecLabelName = "counterVecLabelName"
	counterVecLabelVal  = "counterVecLabelVal"
	counterVecOpts      = CounterVecOpts{prometheus.CounterOpts{Name: counterVecName}, []string{counterVecLabelName}}
	counterVec          = prometheus.NewCounterVec(counterVecOpts.CounterOpts, counterVecOpts.Labels)
	histogramName       = "histogramName"
	histogramOpts       = prometheus.HistogramOpts{Name: histogramName}
	histogram           = prometheus.NewHistogram(histogramOpts)
	summaryName         = "summaryName"
	summaryOpts         = prometheus.SummaryOpts{Name: summaryName}
	summary             = prometheus.NewSummary(summaryOpts)
)

func TestMain(m *testing.M) {
	Initialize()
	// Overriding registerer to be able to do the unit tests independently
	registerer = prometheus.NewRegistry()

	code := m.Run()

	initialized = false
	os.Exit(code)
}

func TestHandler(t *testing.T) {
	actual := Handler()

	assert.NotNil(t, actual)
}

func TestRegisterGauges(t *testing.T) {
	gaugesOpts := []prometheus.GaugeOpts{gaugeOpts}

	RegisterGauges(gaugesOpts...)

	assert.Len(t, gauges, 1)
}

func TestGauge(t *testing.T) {
	gauges[gaugeName] = gauge

	actual, exist := Gauge(gaugeName)

	assert.True(t, exist)
	assert.Equal(t, gauge, actual)
}

func TestGaugeSet(t *testing.T) {
	gauges[gaugeName] = gauge
	expected := float64(2)

	GaugeSet(gaugeName, expected)
	actual := testutil.ToFloat64(gauge)

	assert.Equal(t, expected, actual)
}

func TestUnregisterGauges(t *testing.T) {
	RegisterGauges(gaugeOpts)

	UnregisterGauges(gaugeName)

	assert.Len(t, gauges, 0)
}

func TestRegisterCounters(t *testing.T) {
	countersOpts := []prometheus.CounterOpts{counterOpts}

	RegisterCounters(countersOpts...)

	assert.Len(t, counters, 1)
}

func TestCounter(t *testing.T) {
	counters[counterName] = counter

	actual, exist := Counter(counterName)

	assert.True(t, exist)
	assert.Equal(t, counter, actual)
}

func TestCounterInc(t *testing.T) {
	counters[counterName] = counter
	expected := float64(1)

	CounterInc(counterName)
	actual := testutil.ToFloat64(counter)

	assert.Equal(t, expected, actual)
}

func TestCounterAdd(t *testing.T) {
	counters[counterName] = counter
	expected := float64(2)

	CounterAdd(counterName, expected)
	actual := testutil.ToFloat64(counter)

	assert.Equal(t, expected, actual)
}

func TestUnregisterCounters(t *testing.T) {
	RegisterCounters(counterOpts)

	UnregisterCounters(counterName)

	assert.Len(t, counters, 0)
}

func TestRegisterCounterVecs(t *testing.T) {
	counterVecsOpts := []CounterVecOpts{counterVecOpts}

	RegisterCounterVecs(counterVecsOpts...)

	assert.Len(t, counterVecs, 1)
}

func TestCounterVec(t *testing.T) {
	counterVecs[counterVecName] = counterVec

	actual, exist := CounterVec(counterVecName)

	assert.True(t, exist)
	assert.Equal(t, counterVec, actual)
}

func TestCounterVecInc(t *testing.T) {
	counterVecs[counterVecName] = counterVec
	expected := float64(1)

	CounterVecInc(counterVecName, counterVecLabelVal)
	currCounterVec, err := counterVec.GetMetricWithLabelValues(counterVecLabelVal)
	require.NoError(t, err)
	actual := testutil.ToFloat64(currCounterVec)

	assert.Equal(t, expected, actual)
}

func TestCounterVecAdd(t *testing.T) {
	counterVecs[counterVecName] = counterVec
	expected := float64(2)

	CounterVecAdd(counterVecName, counterVecLabelVal, expected)
	currCounterVec, err := counterVec.GetMetricWithLabelValues(counterVecLabelVal)
	require.NoError(t, err)
	actual := testutil.ToFloat64(currCounterVec)

	assert.Equal(t, expected, actual)
}

func TestUnregisterCounterVecs(t *testing.T) {
	RegisterCounterVecs(counterVecOpts)

	UnregisterCounterVecs(counterVecName)

	assert.Len(t, counterVecs, 0)
}

func TestRegisterHistograms(t *testing.T) {
	histogramsOpts := []prometheus.HistogramOpts{histogramOpts}

	RegisterHistograms(histogramsOpts...)

	assert.Len(t, histograms, 1)
}

func TestHistogram(t *testing.T) {
	histograms[histogramName] = histogram

	actual, exist := Histogram(histogramName)

	assert.True(t, exist)
	assert.Equal(t, histogram, actual)
}

func TestHistogramObserve(t *testing.T) {
	histograms[histogramName] = histogram
	expected := float64(2)

	HistogramObserve(histogramName, expected)
	// TODO: Finish the test
}

func TestUnregisterHistograms(t *testing.T) {
	RegisterHistograms(histogramOpts)

	UnregisterHistogram(histogramName)

	assert.Len(t, histograms, 0)
}

func TestRegisterSummaries(t *testing.T) {
	summariesOpts := []prometheus.SummaryOpts{summaryOpts}

	RegisterSummaries(summariesOpts...)

	assert.Len(t, summaries, 1)
}

func TestSummary(t *testing.T) {
	summaries[summaryName] = summary

	actual, exist := Summary(summaryName)

	assert.True(t, exist)
	assert.Equal(t, summary, actual)
}

func TestUnregisterSummaries(t *testing.T) {
	RegisterSummaries(summaryOpts)

	UnregisterSummaries(summaryName)

	assert.Len(t, summaries, 0)
}
