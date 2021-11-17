// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"code.cloudfoundry.org/cf-k8s-controllers/api/apis"
	"code.cloudfoundry.org/cf-k8s-controllers/api/repositories"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type CFBuildRepository struct {
	CreateBuildStub        func(context.Context, client.Client, repositories.BuildCreateMessage) (repositories.BuildRecord, error)
	createBuildMutex       sync.RWMutex
	createBuildArgsForCall []struct {
		arg1 context.Context
		arg2 client.Client
		arg3 repositories.BuildCreateMessage
	}
	createBuildReturns struct {
		result1 repositories.BuildRecord
		result2 error
	}
	createBuildReturnsOnCall map[int]struct {
		result1 repositories.BuildRecord
		result2 error
	}
	FetchBuildStub        func(context.Context, client.Client, string) (repositories.BuildRecord, error)
	fetchBuildMutex       sync.RWMutex
	fetchBuildArgsForCall []struct {
		arg1 context.Context
		arg2 client.Client
		arg3 string
	}
	fetchBuildReturns struct {
		result1 repositories.BuildRecord
		result2 error
	}
	fetchBuildReturnsOnCall map[int]struct {
		result1 repositories.BuildRecord
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CFBuildRepository) CreateBuild(arg1 context.Context, arg2 client.Client, arg3 repositories.BuildCreateMessage) (repositories.BuildRecord, error) {
	fake.createBuildMutex.Lock()
	ret, specificReturn := fake.createBuildReturnsOnCall[len(fake.createBuildArgsForCall)]
	fake.createBuildArgsForCall = append(fake.createBuildArgsForCall, struct {
		arg1 context.Context
		arg2 client.Client
		arg3 repositories.BuildCreateMessage
	}{arg1, arg2, arg3})
	stub := fake.CreateBuildStub
	fakeReturns := fake.createBuildReturns
	fake.recordInvocation("CreateBuild", []interface{}{arg1, arg2, arg3})
	fake.createBuildMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFBuildRepository) CreateBuildCallCount() int {
	fake.createBuildMutex.RLock()
	defer fake.createBuildMutex.RUnlock()
	return len(fake.createBuildArgsForCall)
}

func (fake *CFBuildRepository) CreateBuildCalls(stub func(context.Context, client.Client, repositories.BuildCreateMessage) (repositories.BuildRecord, error)) {
	fake.createBuildMutex.Lock()
	defer fake.createBuildMutex.Unlock()
	fake.CreateBuildStub = stub
}

func (fake *CFBuildRepository) CreateBuildArgsForCall(i int) (context.Context, client.Client, repositories.BuildCreateMessage) {
	fake.createBuildMutex.RLock()
	defer fake.createBuildMutex.RUnlock()
	argsForCall := fake.createBuildArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFBuildRepository) CreateBuildReturns(result1 repositories.BuildRecord, result2 error) {
	fake.createBuildMutex.Lock()
	defer fake.createBuildMutex.Unlock()
	fake.CreateBuildStub = nil
	fake.createBuildReturns = struct {
		result1 repositories.BuildRecord
		result2 error
	}{result1, result2}
}

func (fake *CFBuildRepository) CreateBuildReturnsOnCall(i int, result1 repositories.BuildRecord, result2 error) {
	fake.createBuildMutex.Lock()
	defer fake.createBuildMutex.Unlock()
	fake.CreateBuildStub = nil
	if fake.createBuildReturnsOnCall == nil {
		fake.createBuildReturnsOnCall = make(map[int]struct {
			result1 repositories.BuildRecord
			result2 error
		})
	}
	fake.createBuildReturnsOnCall[i] = struct {
		result1 repositories.BuildRecord
		result2 error
	}{result1, result2}
}

func (fake *CFBuildRepository) FetchBuild(arg1 context.Context, arg2 client.Client, arg3 string) (repositories.BuildRecord, error) {
	fake.fetchBuildMutex.Lock()
	ret, specificReturn := fake.fetchBuildReturnsOnCall[len(fake.fetchBuildArgsForCall)]
	fake.fetchBuildArgsForCall = append(fake.fetchBuildArgsForCall, struct {
		arg1 context.Context
		arg2 client.Client
		arg3 string
	}{arg1, arg2, arg3})
	stub := fake.FetchBuildStub
	fakeReturns := fake.fetchBuildReturns
	fake.recordInvocation("FetchBuild", []interface{}{arg1, arg2, arg3})
	fake.fetchBuildMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *CFBuildRepository) FetchBuildCallCount() int {
	fake.fetchBuildMutex.RLock()
	defer fake.fetchBuildMutex.RUnlock()
	return len(fake.fetchBuildArgsForCall)
}

func (fake *CFBuildRepository) FetchBuildCalls(stub func(context.Context, client.Client, string) (repositories.BuildRecord, error)) {
	fake.fetchBuildMutex.Lock()
	defer fake.fetchBuildMutex.Unlock()
	fake.FetchBuildStub = stub
}

func (fake *CFBuildRepository) FetchBuildArgsForCall(i int) (context.Context, client.Client, string) {
	fake.fetchBuildMutex.RLock()
	defer fake.fetchBuildMutex.RUnlock()
	argsForCall := fake.fetchBuildArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *CFBuildRepository) FetchBuildReturns(result1 repositories.BuildRecord, result2 error) {
	fake.fetchBuildMutex.Lock()
	defer fake.fetchBuildMutex.Unlock()
	fake.FetchBuildStub = nil
	fake.fetchBuildReturns = struct {
		result1 repositories.BuildRecord
		result2 error
	}{result1, result2}
}

func (fake *CFBuildRepository) FetchBuildReturnsOnCall(i int, result1 repositories.BuildRecord, result2 error) {
	fake.fetchBuildMutex.Lock()
	defer fake.fetchBuildMutex.Unlock()
	fake.FetchBuildStub = nil
	if fake.fetchBuildReturnsOnCall == nil {
		fake.fetchBuildReturnsOnCall = make(map[int]struct {
			result1 repositories.BuildRecord
			result2 error
		})
	}
	fake.fetchBuildReturnsOnCall[i] = struct {
		result1 repositories.BuildRecord
		result2 error
	}{result1, result2}
}

func (fake *CFBuildRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createBuildMutex.RLock()
	defer fake.createBuildMutex.RUnlock()
	fake.fetchBuildMutex.RLock()
	defer fake.fetchBuildMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CFBuildRepository) recordInvocation(key string, args []interface{}) {
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

var _ apis.CFBuildRepository = new(CFBuildRepository)