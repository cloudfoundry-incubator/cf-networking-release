// This file was generated by counterfeiter
package fakes

import (
	"net"
	"sync"
)

type NetAdapter struct {
	InterfaceByIndexStub        func(int) (*net.Interface, error)
	interfaceByIndexMutex       sync.RWMutex
	interfaceByIndexArgsForCall []struct {
		arg1 int
	}
	interfaceByIndexReturns struct {
		result1 *net.Interface
		result2 error
	}
	interfaceByIndexReturnsOnCall map[int]struct {
		result1 *net.Interface
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *NetAdapter) InterfaceByIndex(arg1 int) (*net.Interface, error) {
	fake.interfaceByIndexMutex.Lock()
	ret, specificReturn := fake.interfaceByIndexReturnsOnCall[len(fake.interfaceByIndexArgsForCall)]
	fake.interfaceByIndexArgsForCall = append(fake.interfaceByIndexArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("InterfaceByIndex", []interface{}{arg1})
	fake.interfaceByIndexMutex.Unlock()
	if fake.InterfaceByIndexStub != nil {
		return fake.InterfaceByIndexStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.interfaceByIndexReturns.result1, fake.interfaceByIndexReturns.result2
}

func (fake *NetAdapter) InterfaceByIndexCallCount() int {
	fake.interfaceByIndexMutex.RLock()
	defer fake.interfaceByIndexMutex.RUnlock()
	return len(fake.interfaceByIndexArgsForCall)
}

func (fake *NetAdapter) InterfaceByIndexArgsForCall(i int) int {
	fake.interfaceByIndexMutex.RLock()
	defer fake.interfaceByIndexMutex.RUnlock()
	return fake.interfaceByIndexArgsForCall[i].arg1
}

func (fake *NetAdapter) InterfaceByIndexReturns(result1 *net.Interface, result2 error) {
	fake.InterfaceByIndexStub = nil
	fake.interfaceByIndexReturns = struct {
		result1 *net.Interface
		result2 error
	}{result1, result2}
}

func (fake *NetAdapter) InterfaceByIndexReturnsOnCall(i int, result1 *net.Interface, result2 error) {
	fake.InterfaceByIndexStub = nil
	if fake.interfaceByIndexReturnsOnCall == nil {
		fake.interfaceByIndexReturnsOnCall = make(map[int]struct {
			result1 *net.Interface
			result2 error
		})
	}
	fake.interfaceByIndexReturnsOnCall[i] = struct {
		result1 *net.Interface
		result2 error
	}{result1, result2}
}

func (fake *NetAdapter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.interfaceByIndexMutex.RLock()
	defer fake.interfaceByIndexMutex.RUnlock()
	return fake.invocations
}

func (fake *NetAdapter) recordInvocation(key string, args []interface{}) {
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