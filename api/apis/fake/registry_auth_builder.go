// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"code.cloudfoundry.org/cf-k8s-controllers/api/apis"
	"github.com/google/go-containerregistry/pkg/v1/remote"
)

type RegistryAuthBuilder struct {
	Stub        func(context.Context) (remote.Option, error)
	mutex       sync.RWMutex
	argsForCall []struct {
		arg1 context.Context
	}
	returns struct {
		result1 remote.Option
		result2 error
	}
	returnsOnCall map[int]struct {
		result1 remote.Option
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *RegistryAuthBuilder) Spy(arg1 context.Context) (remote.Option, error) {
	fake.mutex.Lock()
	ret, specificReturn := fake.returnsOnCall[len(fake.argsForCall)]
	fake.argsForCall = append(fake.argsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.Stub
	returns := fake.returns
	fake.recordInvocation("RegistryAuthBuilder", []interface{}{arg1})
	fake.mutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return returns.result1, returns.result2
}

func (fake *RegistryAuthBuilder) CallCount() int {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return len(fake.argsForCall)
}

func (fake *RegistryAuthBuilder) Calls(stub func(context.Context) (remote.Option, error)) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = stub
}

func (fake *RegistryAuthBuilder) ArgsForCall(i int) context.Context {
	fake.mutex.RLock()
	defer fake.mutex.RUnlock()
	return fake.argsForCall[i].arg1
}

func (fake *RegistryAuthBuilder) Returns(result1 remote.Option, result2 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	fake.returns = struct {
		result1 remote.Option
		result2 error
	}{result1, result2}
}

func (fake *RegistryAuthBuilder) ReturnsOnCall(i int, result1 remote.Option, result2 error) {
	fake.mutex.Lock()
	defer fake.mutex.Unlock()
	fake.Stub = nil
	if fake.returnsOnCall == nil {
		fake.returnsOnCall = make(map[int]struct {
			result1 remote.Option
			result2 error
		})
	}
	fake.returnsOnCall[i] = struct {
		result1 remote.Option
		result2 error
	}{result1, result2}
}

func (fake *RegistryAuthBuilder) Invocations() map[string][][]interface{} {
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

func (fake *RegistryAuthBuilder) recordInvocation(key string, args []interface{}) {
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

var _ apis.RegistryAuthBuilder = new(RegistryAuthBuilder).Spy