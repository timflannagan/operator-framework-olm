// Code generated by counterfeiter. DO NOT EDIT.
package appregistryfakes

import (
	"sync"

	"github.com/openshift/operator-framework-olm/staging/operator-registry/pkg/appregistry"
)

type FakeImageAppender struct {
	AppendStub        func(string, string, string) error
	appendMutex       sync.RWMutex
	appendArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
	}
	appendReturns struct {
		result1 error
	}
	appendReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImageAppender) Append(arg1 string, arg2 string, arg3 string) error {
	fake.appendMutex.Lock()
	ret, specificReturn := fake.appendReturnsOnCall[len(fake.appendArgsForCall)]
	fake.appendArgsForCall = append(fake.appendArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
	}{arg1, arg2, arg3})
	fake.recordInvocation("Append", []interface{}{arg1, arg2, arg3})
	fake.appendMutex.Unlock()
	if fake.AppendStub != nil {
		return fake.AppendStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.appendReturns
	return fakeReturns.result1
}

func (fake *FakeImageAppender) AppendCallCount() int {
	fake.appendMutex.RLock()
	defer fake.appendMutex.RUnlock()
	return len(fake.appendArgsForCall)
}

func (fake *FakeImageAppender) AppendCalls(stub func(string, string, string) error) {
	fake.appendMutex.Lock()
	defer fake.appendMutex.Unlock()
	fake.AppendStub = stub
}

func (fake *FakeImageAppender) AppendArgsForCall(i int) (string, string, string) {
	fake.appendMutex.RLock()
	defer fake.appendMutex.RUnlock()
	argsForCall := fake.appendArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImageAppender) AppendReturns(result1 error) {
	fake.appendMutex.Lock()
	defer fake.appendMutex.Unlock()
	fake.AppendStub = nil
	fake.appendReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageAppender) AppendReturnsOnCall(i int, result1 error) {
	fake.appendMutex.Lock()
	defer fake.appendMutex.Unlock()
	fake.AppendStub = nil
	if fake.appendReturnsOnCall == nil {
		fake.appendReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.appendReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImageAppender) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.appendMutex.RLock()
	defer fake.appendMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImageAppender) recordInvocation(key string, args []interface{}) {
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

var _ appregistry.ImageAppender = new(FakeImageAppender)
