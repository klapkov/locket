// Code generated by counterfeiter. DO NOT EDIT.
package metricsfakes

import (
	"sync"

	"code.cloudfoundry.org/lager"
	"code.cloudfoundry.org/locket/metrics"
)

type FakeLockDBMetrics struct {
	OpenConnectionsStub        func() int
	openConnectionsMutex       sync.RWMutex
	openConnectionsArgsForCall []struct{}
	openConnectionsReturns     struct {
		result1 int
	}
	openConnectionsReturnsOnCall map[int]struct {
		result1 int
	}
	CountStub        func(logger lager.Logger, lockType string) (int, error)
	countMutex       sync.RWMutex
	countArgsForCall []struct {
		logger   lager.Logger
		lockType string
	}
	countReturns struct {
		result1 int
		result2 error
	}
	countReturnsOnCall map[int]struct {
		result1 int
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeLockDBMetrics) OpenConnections() int {
	fake.openConnectionsMutex.Lock()
	ret, specificReturn := fake.openConnectionsReturnsOnCall[len(fake.openConnectionsArgsForCall)]
	fake.openConnectionsArgsForCall = append(fake.openConnectionsArgsForCall, struct{}{})
	fake.recordInvocation("OpenConnections", []interface{}{})
	fake.openConnectionsMutex.Unlock()
	if fake.OpenConnectionsStub != nil {
		return fake.OpenConnectionsStub()
	}
	if specificReturn {
		return ret.result1
	}
	return fake.openConnectionsReturns.result1
}

func (fake *FakeLockDBMetrics) OpenConnectionsCallCount() int {
	fake.openConnectionsMutex.RLock()
	defer fake.openConnectionsMutex.RUnlock()
	return len(fake.openConnectionsArgsForCall)
}

func (fake *FakeLockDBMetrics) OpenConnectionsReturns(result1 int) {
	fake.OpenConnectionsStub = nil
	fake.openConnectionsReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeLockDBMetrics) OpenConnectionsReturnsOnCall(i int, result1 int) {
	fake.OpenConnectionsStub = nil
	if fake.openConnectionsReturnsOnCall == nil {
		fake.openConnectionsReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.openConnectionsReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeLockDBMetrics) Count(logger lager.Logger, lockType string) (int, error) {
	fake.countMutex.Lock()
	ret, specificReturn := fake.countReturnsOnCall[len(fake.countArgsForCall)]
	fake.countArgsForCall = append(fake.countArgsForCall, struct {
		logger   lager.Logger
		lockType string
	}{logger, lockType})
	fake.recordInvocation("Count", []interface{}{logger, lockType})
	fake.countMutex.Unlock()
	if fake.CountStub != nil {
		return fake.CountStub(logger, lockType)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.countReturns.result1, fake.countReturns.result2
}

func (fake *FakeLockDBMetrics) CountCallCount() int {
	fake.countMutex.RLock()
	defer fake.countMutex.RUnlock()
	return len(fake.countArgsForCall)
}

func (fake *FakeLockDBMetrics) CountArgsForCall(i int) (lager.Logger, string) {
	fake.countMutex.RLock()
	defer fake.countMutex.RUnlock()
	return fake.countArgsForCall[i].logger, fake.countArgsForCall[i].lockType
}

func (fake *FakeLockDBMetrics) CountReturns(result1 int, result2 error) {
	fake.CountStub = nil
	fake.countReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeLockDBMetrics) CountReturnsOnCall(i int, result1 int, result2 error) {
	fake.CountStub = nil
	if fake.countReturnsOnCall == nil {
		fake.countReturnsOnCall = make(map[int]struct {
			result1 int
			result2 error
		})
	}
	fake.countReturnsOnCall[i] = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeLockDBMetrics) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.openConnectionsMutex.RLock()
	defer fake.openConnectionsMutex.RUnlock()
	fake.countMutex.RLock()
	defer fake.countMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeLockDBMetrics) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ metrics.LockDBMetrics = new(FakeLockDBMetrics)
