// Code generated by counterfeiter. DO NOT EDIT.
package fixturesfakes

import (
	"sync"
)

type FakeUnexportedInterface struct {
	MethodStub        func(string, map[string]interface{}) string
	methodMutex       sync.RWMutex
	methodArgsForCall []struct {
		arg1 string
		arg2 map[string]interface{}
	}
	methodReturns struct {
		result1 string
	}
	methodReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeUnexportedInterface) Method(arg1 string, arg2 map[string]interface{}) string {
	fake.methodMutex.Lock()
	defer fake.methodMutex.Unlock()
	ret, specificReturn := fake.methodReturnsOnCall[len(fake.methodArgsForCall)]
	fake.methodArgsForCall = append(fake.methodArgsForCall, struct {
		arg1 string
		arg2 map[string]interface{}
	}{arg1, arg2})
	fake.recordInvocation("Method", []interface{}{arg1, arg2})
	if fake.MethodStub != nil {
		return fake.MethodStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.methodReturns
	return fakeReturns.result1
}

func (fake *FakeUnexportedInterface) MethodCallCount() int {
	fake.methodMutex.RLock()
	defer fake.methodMutex.RUnlock()
	return len(fake.methodArgsForCall)
}

func (fake *FakeUnexportedInterface) MethodCalls(stub func(string, map[string]interface{}) string) {
	fake.methodMutex.Lock()
	defer fake.methodMutex.Unlock()
	fake.MethodStub = stub
}

func (fake *FakeUnexportedInterface) MethodArgsForCall(i int) (string, map[string]interface{}) {
	fake.methodMutex.RLock()
	defer fake.methodMutex.RUnlock()
	argsForCall := fake.methodArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeUnexportedInterface) MethodReturns(result1 string) {
	fake.methodMutex.Lock()
	defer fake.methodMutex.Unlock()
	fake.MethodStub = nil
	fake.methodReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeUnexportedInterface) MethodReturnsOnCall(i int, result1 string) {
	fake.methodMutex.Lock()
	defer fake.methodMutex.Unlock()
	fake.MethodStub = nil
	if fake.methodReturnsOnCall == nil {
		fake.methodReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.methodReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeUnexportedInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.methodMutex.RLock()
	defer fake.methodMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeUnexportedInterface) recordInvocation(key string, args []interface{}) {
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
