// Code generated by mockery v2.14.0. DO NOT EDIT.

package jsonrpc

import (
	time "time"

	prometheus "github.com/prometheus/client_golang/prometheus"
	mock "github.com/stretchr/testify/mock"
)

// metricsMock is an autogenerated mock type for the metricsInterface type
type metricsMock struct {
	mock.Mock
}

// IncCounterVec provides a mock function with given fields: name, label
func (_m *metricsMock) IncCounterVec(name string, label string) {
	_m.Called(name, label)
}

// ObserveHistogram provides a mock function with given fields: name, start
func (_m *metricsMock) ObserveHistogram(name string, start time.Time) {
	_m.Called(name, start)
}

// RegisterCounters provides a mock function with given fields: opts
func (_m *metricsMock) RegisterCounters(opts ...prometheus.CounterOpts) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

// RegisterHistograms provides a mock function with given fields: opts
func (_m *metricsMock) RegisterHistograms(opts ...prometheus.HistogramOpts) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	_m.Called(_ca...)
}

type mockConstructorTestingTnewMetricsMock interface {
	mock.TestingT
	Cleanup(func())
}

// newMetricsMock creates a new instance of metricsMock. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMetricsMock(t mockConstructorTestingTnewMetricsMock) *metricsMock {
	mock := &metricsMock{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
