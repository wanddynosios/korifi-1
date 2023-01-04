// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"net/http"
	"sync"

	"code.cloudfoundry.org/korifi/api/routing"
)

type Handler struct {
	Stub        func(*http.Request) (*routing.Response, error)
	mutex       sync.RWMutex
	argsForCall []struct {
		arg1 *http.Request
	}
	returns struct {
		result1 *routing.Response
		result2 error
	}
	returnsOnCall map[int]struct {
		result1 *routing.Response
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *Handler) Spy(arg1 *http.Request) (*routing.Response, error) {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		arg1 *http.Request
	}{arg1})
	stub := fake.Stub
	returns := fake.returns
	fake.recordInvocation("Handler", []interface{}{arg1})
	fake.mutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return returns.result1, returns.result2
}

func (fake *Handler) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *Handler) Calls(stub func(*http.Request) (*routing.Response, error)) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = stub
}

func (fake *Handler) ArgsForCall(i int) *http.Request {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].arg1
}

func (fake *Handler) Returns(result1 *routing.Response, result2 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	fake.returns = struct {
		result1 *routing.Response
		result2 error
	}{result1, result2}
}

func (fake *Handler) ReturnsOnCall(i int, result1 *routing.Response, result2 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 *routing.Response
			result2 error
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 *routing.Response
		result2 error
	}{result1, result2}
}

func (fake *Handler) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *Handler) recordInvocation(key string, args []interface{}) {
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

var _ routing.Handler = new(Handler).Spy
