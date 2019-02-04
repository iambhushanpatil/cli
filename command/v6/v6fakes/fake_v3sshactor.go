// Code generated by counterfeiter. DO NOT EDIT.
package v6fakes

import (
	sync "sync"

	v3action "code.cloudfoundry.org/cli/actor/v3action"
	v6 "code.cloudfoundry.org/cli/command/v6"
)

type FakeV3SSHActor struct {
	GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexStub        func(string, string, string, uint) (v3action.SSHAuthentication, v3action.Warnings, error)
	getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex       sync.RWMutex
	getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 uint
	}
	getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturns struct {
		result1 v3action.SSHAuthentication
		result2 v3action.Warnings
		result3 error
	}
	getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall map[int]struct {
		result1 v3action.SSHAuthentication
		result2 v3action.Warnings
		result3 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeV3SSHActor) GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndex(arg1 string, arg2 string, arg3 string, arg4 uint) (v3action.SSHAuthentication, v3action.Warnings, error) {
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.Lock()
	ret, specificReturn := fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall[len(fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexArgsForCall)]
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexArgsForCall = append(fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 uint
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndex", []interface{}{arg1, arg2, arg3, arg4})
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.Unlock()
	if fake.GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexStub != nil {
		return fake.GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeV3SSHActor) GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexCallCount() int {
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.RLock()
	defer fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.RUnlock()
	return len(fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexArgsForCall)
}

func (fake *FakeV3SSHActor) GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexCalls(stub func(string, string, string, uint) (v3action.SSHAuthentication, v3action.Warnings, error)) {
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.Lock()
	defer fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.Unlock()
	fake.GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexStub = stub
}

func (fake *FakeV3SSHActor) GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexArgsForCall(i int) (string, string, string, uint) {
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.RLock()
	defer fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.RUnlock()
	argsForCall := fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeV3SSHActor) GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturns(result1 v3action.SSHAuthentication, result2 v3action.Warnings, result3 error) {
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.Lock()
	defer fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.Unlock()
	fake.GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexStub = nil
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturns = struct {
		result1 v3action.SSHAuthentication
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3SSHActor) GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall(i int, result1 v3action.SSHAuthentication, result2 v3action.Warnings, result3 error) {
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.Lock()
	defer fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.Unlock()
	fake.GetSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexStub = nil
	if fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall == nil {
		fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall = make(map[int]struct {
			result1 v3action.SSHAuthentication
			result2 v3action.Warnings
			result3 error
		})
	}
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexReturnsOnCall[i] = struct {
		result1 v3action.SSHAuthentication
		result2 v3action.Warnings
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeV3SSHActor) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.RLock()
	defer fake.getSecureShellConfigurationByApplicationNameSpaceProcessTypeAndIndexMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeV3SSHActor) recordInvocation(key string, args []interface{}) {
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

var _ v6.V3SSHActor = new(FakeV3SSHActor)