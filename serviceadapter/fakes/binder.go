// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/pivotal-cf/on-demand-services-sdk/serviceadapter"
)

type FakeBinder struct {
	CreateBindingStub        func(serviceadapter.CreateBindingParams) (serviceadapter.Binding, error)
	createBindingMutex       sync.RWMutex
	createBindingArgsForCall []struct {
		arg1 serviceadapter.CreateBindingParams
	}
	createBindingReturns struct {
		result1 serviceadapter.Binding
		result2 error
	}
	createBindingReturnsOnCall map[int]struct {
		result1 serviceadapter.Binding
		result2 error
	}
	DeleteBindingStub        func(serviceadapter.DeleteBindingParams) error
	deleteBindingMutex       sync.RWMutex
	deleteBindingArgsForCall []struct {
		arg1 serviceadapter.DeleteBindingParams
	}
	deleteBindingReturns struct {
		result1 error
	}
	deleteBindingReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBinder) CreateBinding(arg1 serviceadapter.CreateBindingParams) (serviceadapter.Binding, error) {
	fake.createBindingMutex.Lock()
	ret, specificReturn := fake.createBindingReturnsOnCall[len(fake.createBindingArgsForCall)]
	fake.createBindingArgsForCall = append(fake.createBindingArgsForCall, struct {
		arg1 serviceadapter.CreateBindingParams
	}{arg1})
	fake.recordInvocation("CreateBinding", []interface{}{arg1})
	fake.createBindingMutex.Unlock()
	if fake.CreateBindingStub != nil {
		return fake.CreateBindingStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createBindingReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBinder) CreateBindingCallCount() int {
	fake.createBindingMutex.RLock()
	defer fake.createBindingMutex.RUnlock()
	return len(fake.createBindingArgsForCall)
}

func (fake *FakeBinder) CreateBindingCalls(stub func(serviceadapter.CreateBindingParams) (serviceadapter.Binding, error)) {
	fake.createBindingMutex.Lock()
	defer fake.createBindingMutex.Unlock()
	fake.CreateBindingStub = stub
}

func (fake *FakeBinder) CreateBindingArgsForCall(i int) serviceadapter.CreateBindingParams {
	fake.createBindingMutex.RLock()
	defer fake.createBindingMutex.RUnlock()
	argsForCall := fake.createBindingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBinder) CreateBindingReturns(result1 serviceadapter.Binding, result2 error) {
	fake.createBindingMutex.Lock()
	defer fake.createBindingMutex.Unlock()
	fake.CreateBindingStub = nil
	fake.createBindingReturns = struct {
		result1 serviceadapter.Binding
		result2 error
	}{result1, result2}
}

func (fake *FakeBinder) CreateBindingReturnsOnCall(i int, result1 serviceadapter.Binding, result2 error) {
	fake.createBindingMutex.Lock()
	defer fake.createBindingMutex.Unlock()
	fake.CreateBindingStub = nil
	if fake.createBindingReturnsOnCall == nil {
		fake.createBindingReturnsOnCall = make(map[int]struct {
			result1 serviceadapter.Binding
			result2 error
		})
	}
	fake.createBindingReturnsOnCall[i] = struct {
		result1 serviceadapter.Binding
		result2 error
	}{result1, result2}
}

func (fake *FakeBinder) DeleteBinding(arg1 serviceadapter.DeleteBindingParams) error {
	fake.deleteBindingMutex.Lock()
	ret, specificReturn := fake.deleteBindingReturnsOnCall[len(fake.deleteBindingArgsForCall)]
	fake.deleteBindingArgsForCall = append(fake.deleteBindingArgsForCall, struct {
		arg1 serviceadapter.DeleteBindingParams
	}{arg1})
	fake.recordInvocation("DeleteBinding", []interface{}{arg1})
	fake.deleteBindingMutex.Unlock()
	if fake.DeleteBindingStub != nil {
		return fake.DeleteBindingStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.deleteBindingReturns
	return fakeReturns.result1
}

func (fake *FakeBinder) DeleteBindingCallCount() int {
	fake.deleteBindingMutex.RLock()
	defer fake.deleteBindingMutex.RUnlock()
	return len(fake.deleteBindingArgsForCall)
}

func (fake *FakeBinder) DeleteBindingCalls(stub func(serviceadapter.DeleteBindingParams) error) {
	fake.deleteBindingMutex.Lock()
	defer fake.deleteBindingMutex.Unlock()
	fake.DeleteBindingStub = stub
}

func (fake *FakeBinder) DeleteBindingArgsForCall(i int) serviceadapter.DeleteBindingParams {
	fake.deleteBindingMutex.RLock()
	defer fake.deleteBindingMutex.RUnlock()
	argsForCall := fake.deleteBindingArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBinder) DeleteBindingReturns(result1 error) {
	fake.deleteBindingMutex.Lock()
	defer fake.deleteBindingMutex.Unlock()
	fake.DeleteBindingStub = nil
	fake.deleteBindingReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBinder) DeleteBindingReturnsOnCall(i int, result1 error) {
	fake.deleteBindingMutex.Lock()
	defer fake.deleteBindingMutex.Unlock()
	fake.DeleteBindingStub = nil
	if fake.deleteBindingReturnsOnCall == nil {
		fake.deleteBindingReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.deleteBindingReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBinder) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createBindingMutex.RLock()
	defer fake.createBindingMutex.RUnlock()
	fake.deleteBindingMutex.RLock()
	defer fake.deleteBindingMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBinder) recordInvocation(key string, args []interface{}) {
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

var _ serviceadapter.Binder = new(FakeBinder)
