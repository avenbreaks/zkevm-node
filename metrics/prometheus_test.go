package metrics

import (
	"runtime"
	"sync"
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/testutil"
	dto "github.com/prometheus/client_model/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var (
	gaugeName           = "gaugeName"
	gaugeOpts           = prometheus.GaugeOpts{Name: gaugeName}
	gauge               prometheus.Gauge
	counterName         = "counterName"
	counterOpts         = prometheus.CounterOpts{Name: counterName}
	counter             prometheus.Counter
	counterVecName      = "counterVecName"
	counterVecLabelName = "counterVecLabelName"
	counterVecLabelVal  = "counterVecLabelVal"
	counterVecOpts      = CounterVecOpts{prometheus.CounterOpts{Name: counterVecName}, []string{counterVecLabelName}}
	counterVec          *prometheus.CounterVec
	histogramName       = "histogramName"
	histogramOpts       = prometheus.HistogramOpts{Name: histogramName, Buckets: []float64{0.5, 10, 20}}
	histogram           prometheus.Histogram
	summaryName         = "summaryName"
	summaryOpts         = prometheus.SummaryOpts{Name: summaryName}
	summary             = prometheus.NewSummary(summaryOpts)
)

func setup() {
	Initialize()
	gauge = prometheus.NewGauge(gaugeOpts)
	counter = prometheus.NewCounter(counterOpts)
	counterVec = prometheus.NewCounterVec(counterVecOpts.CounterOpts, counterVecOpts.Labels)
	histogram = prometheus.NewHistogram(histogramOpts)
	summary = prometheus.NewSummary(summaryOpts)

	// Overriding registerer to be able to do the unit tests independently
	registerer = prometheus.NewRegistry()
}

func cleanup() {
	initialized = false
	initOnce = sync.Once{}
}

func TestHandler(t *testing.T) {
	setup()
	defer cleanup()

	actual := Handler()

	assert.NotNil(t, actual)
}

func TestRegisterGauges(t *testing.T) {
	setup()
	defer cleanup()
	gaugesOpts := []prometheus.GaugeOpts{gaugeOpts}

	RegisterGauges(gaugesOpts...)

	assert.Len(t, gauges, 1)
}

func TestGauge(t *testing.T) {
	setup()
	defer cleanup()
	gauges[gaugeName] = gauge

	actual, exist := Gauge(gaugeName)

	assert.True(t, exist)
	assert.Equal(t, gauge, actual)
}

func TestGaugeSet(t *testing.T) {
	setup()
	defer cleanup()
	gauges[gaugeName] = gauge
	expected := float64(2)

	GaugeSet(gaugeName, expected)
	actual := testutil.ToFloat64(gauge)

	assert.Equal(t, expected, actual)
}

func TestUnregisterGauges(t *testing.T) {
	setup()
	defer cleanup()
	RegisterGauges(gaugeOpts)

	UnregisterGauges(gaugeName)

	assert.Len(t, gauges, 0)
}

func TestRegisterCounters(t *testing.T) {
	setup()
	defer cleanup()
	countersOpts := []prometheus.CounterOpts{counterOpts}

	RegisterCounters(countersOpts...)

	assert.Len(t, counters, 1)
}

func TestCounter(t *testing.T) {
	setup()
	defer cleanup()
	counters[counterName] = counter

	actual, exist := Counter(counterName)

	assert.True(t, exist)
	assert.Equal(t, counter, actual)
}

func TestCounterInc(t *testing.T) {
	setup()
	defer cleanup()
	counters[counterName] = counter
	expected := float64(1)

	CounterInc(counterName)
	actual := testutil.ToFloat64(counter)

	assert.Equal(t, expected, actual)
}

func TestCounterAdd(t *testing.T) {
	setup()
	defer cleanup()
	counters[counterName] = counter
	expected := float64(2)

	CounterAdd(counterName, expected)
	actual := testutil.ToFloat64(counter)

	assert.Equal(t, expected, actual)
}

func TestUnregisterCounters(t *testing.T) {
	setup()
	defer cleanup()
	RegisterCounters(counterOpts)

	UnregisterCounters(counterName)

	assert.Len(t, counters, 0)
}

func TestRegisterCounterVecs(t *testing.T) {
	setup()
	defer cleanup()
	counterVecsOpts := []CounterVecOpts{counterVecOpts}

	RegisterCounterVecs(counterVecsOpts...)

	assert.Len(t, counterVecs, 1)
}

func TestCounterVec(t *testing.T) {
	setup()
	defer cleanup()
	counterVecs[counterVecName] = counterVec

	actual, exist := CounterVec(counterVecName)

	assert.True(t, exist)
	assert.Equal(t, counterVec, actual)
}

func TestCounterVecInc(t *testing.T) {
	setup()
	defer cleanup()
	counterVecs[counterVecName] = counterVec
	expected := float64(1)

	CounterVecInc(counterVecName, counterVecLabelVal)
	currCounterVec, err := counterVec.GetMetricWithLabelValues(counterVecLabelVal)
	require.NoError(t, err)
	actual := testutil.ToFloat64(currCounterVec)

	assert.Equal(t, expected, actual)
}

func TestCounterVecAdd(t *testing.T) {
	setup()
	defer cleanup()
	counterVecs[counterVecName] = counterVec
	expected := float64(2)

	CounterVecAdd(counterVecName, counterVecLabelVal, expected)
	currCounterVec, err := counterVec.GetMetricWithLabelValues(counterVecLabelVal)
	require.NoError(t, err)
	actual := testutil.ToFloat64(currCounterVec)

	assert.Equal(t, expected, actual)
}

func TestUnregisterCounterVecs(t *testing.T) {
	setup()
	defer cleanup()
	RegisterCounterVecs(counterVecOpts)

	UnregisterCounterVecs(counterVecName)

	assert.Len(t, counterVecs, 0)
}

func TestRegisterHistograms(t *testing.T) {
	setup()
	defer cleanup()
	histogramsOpts := []prometheus.HistogramOpts{histogramOpts}

	RegisterHistograms(histogramsOpts...)

	assert.Len(t, histograms, 1)
}

func TestHistogram(t *testing.T) {
	setup()
	defer cleanup()
	histograms[histogramName] = histogram

	actual, exist := Histogram(histogramName)

	assert.True(t, exist)
	assert.Equal(t, histogram, actual)
}

func TestHistogramObserve(t *testing.T) {
	setup()
	defer cleanup()
	histograms[histogramName] = histogram

	// NOTE: this is taken from prometheus tests for Histogram. In order to validate that
	//       our implementation doesn't break the main prometheus functionality.
	var (
		quit = make(chan struct{})
	)

	defer func() { close(quit) }()

	observe := func() {
		for {
			select {
			case <-quit:
				return
			default:
				HistogramObserve(histogramName, 1)
			}
		}
	}

	go observe()
	go observe()
	go observe()

	for i := 0; i < 100; i++ {
		m := &dto.Metric{}
		if err := histogram.Write(m); err != nil {
			t.Fatal("unexpected error writing histogram:", err)
		}
		h := m.GetHistogram()
		if h.GetSampleCount() != uint64(h.GetSampleSum()) ||
			h.GetSampleCount() != h.GetBucket()[1].GetCumulativeCount() ||
			h.GetSampleCount() != h.GetBucket()[2].GetCumulativeCount() {
			t.Fatalf(
				"inconsistent counts in histogram: count=%d sum=%f buckets=[%d, %d]",
				h.GetSampleCount(), h.GetSampleSum(),
				h.GetBucket()[1].GetCumulativeCount(), h.GetBucket()[2].GetCumulativeCount(),
			)
		}
		runtime.Gosched()
	}
}

func TestUnregisterHistograms(t *testing.T) {
	setup()
	defer cleanup()
	RegisterHistograms(histogramOpts)

	UnregisterHistogram(histogramName)

	assert.Len(t, histograms, 0)
}

func TestRegisterSummaries(t *testing.T) {
	setup()
	defer cleanup()
	summariesOpts := []prometheus.SummaryOpts{summaryOpts}

	RegisterSummaries(summariesOpts...)

	assert.Len(t, summaries, 1)
}

func TestSummary(t *testing.T) {
	setup()
	defer cleanup()
	summaries[summaryName] = summary

	actual, exist := Summary(summaryName)

	assert.True(t, exist)
	assert.Equal(t, summary, actual)
}

func TestUnregisterSummaries(t *testing.T) {
	setup()
	defer cleanup()
	RegisterSummaries(summaryOpts)

	UnregisterSummaries(summaryName)

	assert.Len(t, summaries, 0)
}
