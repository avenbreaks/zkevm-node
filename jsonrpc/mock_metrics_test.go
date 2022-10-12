// Code generated by mockery v2.14.0. DO NOT EDIT.

package jsonrpc

import (
	prometheus "github.com/prometheus/client_golang/prometheus"
	mock "github.com/stretchr/testify/mock"
)

// metricsMock is an autogenerated mock type for the metricsInterface type
type metricsMock struct {
	mock.Mock
}

// GetCounterVec provides a mock function with given fields: name
func (_m *metricsMock) GetCounterVec(name string) (*prometheus.CounterVec, bool) {
	ret := _m.Called(name)

	var r0 *prometheus.CounterVec
	if rf, ok := ret.Get(0).(func(string) *prometheus.CounterVec); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*prometheus.CounterVec)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// GetHistogram provides a mock function with given fields: name
func (_m *metricsMock) GetHistogram(name string) (prometheus.Histogram, bool) {
	ret := _m.Called(name)

	var r0 prometheus.Histogram
	if rf, ok := ret.Get(0).(func(string) prometheus.Histogram); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(prometheus.Histogram)
		}
	}

	var r1 bool
	if rf, ok := ret.Get(1).(func(string) bool); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
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
