// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"code.cloudfoundry.org/cf-k8s-controllers/api/apis"
	"code.cloudfoundry.org/cf-k8s-controllers/api/repositories"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CFDropletRepository struct {
	FetchDropletStub        func(context.Context, client.Client, string) (repositories.DropletRecord, error)
	fetchDropletMutex       sync.RWMutex
	fetchDropletArgsForCall []struct {
		arg1 context.Context
		arg2 client.Client
		arg3 string
	}
	fetchDropletReturns struct {
		result1 repositories.DropletRecord
		result2 error
	}
	fetchDropletReturnsOnCall map[int]struct {
		result1 repositories.DropletRecord
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CFDropletRepository) FetchDroplet(arg1 context.Context, arg2 client.Client, arg3 string) (repositories.DropletRecord, error) {
	fake.fetchDropletMutex.Lock()
	ret, specificReturn := fake.fetchDropletReturnsOnCall[len(fake.fetchDropletArgsForCall)]
	fake.fetchDropletArgsForCall = append(fake.fetchDropletArgsForCall, struct {
		arg1 context.Context
		arg2 client.Client
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.FetchDropletStub
	fakeReturns := fake.fetchDropletReturns
	fake.recordInvocation("FetchDroplet", []interface{}{arg1, arg2, arg3})
	fake.fetchDropletMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFDropletRepository) FetchDropletCallCount() int {
	fake.fetchDropletMutex.RLock()
	defer fake.fetchDropletMutex.RUnlock()
	return len(fake.fetchDropletArgsForCall)
}

func (fake *CFDropletRepository) FetchDropletCalls(stub func(context.Context, client.Client, string) (repositories.DropletRecord, error)) {
	fake.fetchDropletMutex.Lock()
	defer fake.fetchDropletMutex.Unlock()
	fake.FetchDropletStub = stub
}

func (fake *CFDropletRepository) FetchDropletArgsForCall(i int) (context.Context, client.Client, string) {
	fake.fetchDropletMutex.RLock()
	defer fake.fetchDropletMutex.RUnlock()
	argsForCall := fake.fetchDropletArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFDropletRepository) FetchDropletReturns(result1 repositories.DropletRecord, result2 error) {
	fake.fetchDropletMutex.Lock()
	defer fake.fetchDropletMutex.Unlock()
	fake.FetchDropletStub = nil
	fake.fetchDropletReturns = struct {
		result1 repositories.DropletRecord
		result2 error
	}{result1, result2}
}

func (fake *CFDropletRepository) FetchDropletReturnsOnCall(i int, result1 repositories.DropletRecord, result2 error) {
	fake.fetchDropletMutex.Lock()
	defer fake.fetchDropletMutex.Unlock()
	fake.FetchDropletStub = nil
	if fake.fetchDropletReturnsOnCall == nil {
		fake.fetchDropletReturnsOnCall = make(map[int]struct {
			result1 repositories.DropletRecord
			result2 error
		})
	}
	fake.fetchDropletReturnsOnCall[i] = struct {
		result1 repositories.DropletRecord
		result2 error
	}{result1, result2}
}

func (fake *CFDropletRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchDropletMutex.RLock()
	defer fake.fetchDropletMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CFDropletRepository) recordInvocation(key string, args []interface{}) {
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

var _ apis.CFDropletRepository = new(CFDropletRepository)