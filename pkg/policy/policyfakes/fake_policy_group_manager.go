// Code generated by counterfeiter. DO NOT EDIT.
package policyfakes

import (
	"context"
	"sync"

	"github.com/rode/rode/pkg/policy"
	"github.com/rode/rode/proto/v1alpha1"
)

type FakePolicyGroupManager struct {
	CreatePolicyGroupStub        func(context.Context, *v1alpha1.PolicyGroup) (*v1alpha1.PolicyGroup, error)
	createPolicyGroupMutex       sync.RWMutex
	createPolicyGroupArgsForCall []struct {
		arg1 context.Context
		arg2 *v1alpha1.PolicyGroup
	}
	createPolicyGroupReturns struct {
		result1 *v1alpha1.PolicyGroup
		result2 error
	}
	createPolicyGroupReturnsOnCall map[int]struct {
		result1 *v1alpha1.PolicyGroup
		result2 error
	}
	GetPolicyGroupStub        func(context.Context, *v1alpha1.GetPolicyGroupRequest) (*v1alpha1.PolicyGroup, error)
	getPolicyGroupMutex       sync.RWMutex
	getPolicyGroupArgsForCall []struct {
		arg1 context.Context
		arg2 *v1alpha1.GetPolicyGroupRequest
	}
	getPolicyGroupReturns struct {
		result1 *v1alpha1.PolicyGroup
		result2 error
	}
	getPolicyGroupReturnsOnCall map[int]struct {
		result1 *v1alpha1.PolicyGroup
		result2 error
	}
	ListPolicyGroupsStub        func(context.Context, *v1alpha1.ListPolicyGroupsRequest) (*v1alpha1.ListPolicyGroupsResponse, error)
	listPolicyGroupsMutex       sync.RWMutex
	listPolicyGroupsArgsForCall []struct {
		arg1 context.Context
		arg2 *v1alpha1.ListPolicyGroupsRequest
	}
	listPolicyGroupsReturns struct {
		result1 *v1alpha1.ListPolicyGroupsResponse
		result2 error
	}
	listPolicyGroupsReturnsOnCall map[int]struct {
		result1 *v1alpha1.ListPolicyGroupsResponse
		result2 error
	}
	UpdatePolicyGroupStub        func(context.Context, *v1alpha1.PolicyGroup) (*v1alpha1.PolicyGroup, error)
	updatePolicyGroupMutex       sync.RWMutex
	updatePolicyGroupArgsForCall []struct {
		arg1 context.Context
		arg2 *v1alpha1.PolicyGroup
	}
	updatePolicyGroupReturns struct {
		result1 *v1alpha1.PolicyGroup
		result2 error
	}
	updatePolicyGroupReturnsOnCall map[int]struct {
		result1 *v1alpha1.PolicyGroup
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePolicyGroupManager) CreatePolicyGroup(arg1 context.Context, arg2 *v1alpha1.PolicyGroup) (*v1alpha1.PolicyGroup, error) {
	fake.createPolicyGroupMutex.Lock()
	ret, specificReturn := fake.createPolicyGroupReturnsOnCall[len(fake.createPolicyGroupArgsForCall)]
	fake.createPolicyGroupArgsForCall = append(fake.createPolicyGroupArgsForCall, struct {
		arg1 context.Context
		arg2 *v1alpha1.PolicyGroup
	}{arg1, arg2})
	stub := fake.CreatePolicyGroupStub
	fakeReturns := fake.createPolicyGroupReturns
	fake.recordInvocation("CreatePolicyGroup", []interface{}{arg1, arg2})
	fake.createPolicyGroupMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePolicyGroupManager) CreatePolicyGroupCallCount() int {
	fake.createPolicyGroupMutex.RLock()
	defer fake.createPolicyGroupMutex.RUnlock()
	return len(fake.createPolicyGroupArgsForCall)
}

func (fake *FakePolicyGroupManager) CreatePolicyGroupCalls(stub func(context.Context, *v1alpha1.PolicyGroup) (*v1alpha1.PolicyGroup, error)) {
	fake.createPolicyGroupMutex.Lock()
	defer fake.createPolicyGroupMutex.Unlock()
	fake.CreatePolicyGroupStub = stub
}

func (fake *FakePolicyGroupManager) CreatePolicyGroupArgsForCall(i int) (context.Context, *v1alpha1.PolicyGroup) {
	fake.createPolicyGroupMutex.RLock()
	defer fake.createPolicyGroupMutex.RUnlock()
	argsForCall := fake.createPolicyGroupArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePolicyGroupManager) CreatePolicyGroupReturns(result1 *v1alpha1.PolicyGroup, result2 error) {
	fake.createPolicyGroupMutex.Lock()
	defer fake.createPolicyGroupMutex.Unlock()
	fake.CreatePolicyGroupStub = nil
	fake.createPolicyGroupReturns = struct {
		result1 *v1alpha1.PolicyGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyGroupManager) CreatePolicyGroupReturnsOnCall(i int, result1 *v1alpha1.PolicyGroup, result2 error) {
	fake.createPolicyGroupMutex.Lock()
	defer fake.createPolicyGroupMutex.Unlock()
	fake.CreatePolicyGroupStub = nil
	if fake.createPolicyGroupReturnsOnCall == nil {
		fake.createPolicyGroupReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.PolicyGroup
			result2 error
		})
	}
	fake.createPolicyGroupReturnsOnCall[i] = struct {
		result1 *v1alpha1.PolicyGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyGroupManager) GetPolicyGroup(arg1 context.Context, arg2 *v1alpha1.GetPolicyGroupRequest) (*v1alpha1.PolicyGroup, error) {
	fake.getPolicyGroupMutex.Lock()
	ret, specificReturn := fake.getPolicyGroupReturnsOnCall[len(fake.getPolicyGroupArgsForCall)]
	fake.getPolicyGroupArgsForCall = append(fake.getPolicyGroupArgsForCall, struct {
		arg1 context.Context
		arg2 *v1alpha1.GetPolicyGroupRequest
	}{arg1, arg2})
	stub := fake.GetPolicyGroupStub
	fakeReturns := fake.getPolicyGroupReturns
	fake.recordInvocation("GetPolicyGroup", []interface{}{arg1, arg2})
	fake.getPolicyGroupMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePolicyGroupManager) GetPolicyGroupCallCount() int {
	fake.getPolicyGroupMutex.RLock()
	defer fake.getPolicyGroupMutex.RUnlock()
	return len(fake.getPolicyGroupArgsForCall)
}

func (fake *FakePolicyGroupManager) GetPolicyGroupCalls(stub func(context.Context, *v1alpha1.GetPolicyGroupRequest) (*v1alpha1.PolicyGroup, error)) {
	fake.getPolicyGroupMutex.Lock()
	defer fake.getPolicyGroupMutex.Unlock()
	fake.GetPolicyGroupStub = stub
}

func (fake *FakePolicyGroupManager) GetPolicyGroupArgsForCall(i int) (context.Context, *v1alpha1.GetPolicyGroupRequest) {
	fake.getPolicyGroupMutex.RLock()
	defer fake.getPolicyGroupMutex.RUnlock()
	argsForCall := fake.getPolicyGroupArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePolicyGroupManager) GetPolicyGroupReturns(result1 *v1alpha1.PolicyGroup, result2 error) {
	fake.getPolicyGroupMutex.Lock()
	defer fake.getPolicyGroupMutex.Unlock()
	fake.GetPolicyGroupStub = nil
	fake.getPolicyGroupReturns = struct {
		result1 *v1alpha1.PolicyGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyGroupManager) GetPolicyGroupReturnsOnCall(i int, result1 *v1alpha1.PolicyGroup, result2 error) {
	fake.getPolicyGroupMutex.Lock()
	defer fake.getPolicyGroupMutex.Unlock()
	fake.GetPolicyGroupStub = nil
	if fake.getPolicyGroupReturnsOnCall == nil {
		fake.getPolicyGroupReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.PolicyGroup
			result2 error
		})
	}
	fake.getPolicyGroupReturnsOnCall[i] = struct {
		result1 *v1alpha1.PolicyGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyGroupManager) ListPolicyGroups(arg1 context.Context, arg2 *v1alpha1.ListPolicyGroupsRequest) (*v1alpha1.ListPolicyGroupsResponse, error) {
	fake.listPolicyGroupsMutex.Lock()
	ret, specificReturn := fake.listPolicyGroupsReturnsOnCall[len(fake.listPolicyGroupsArgsForCall)]
	fake.listPolicyGroupsArgsForCall = append(fake.listPolicyGroupsArgsForCall, struct {
		arg1 context.Context
		arg2 *v1alpha1.ListPolicyGroupsRequest
	}{arg1, arg2})
	stub := fake.ListPolicyGroupsStub
	fakeReturns := fake.listPolicyGroupsReturns
	fake.recordInvocation("ListPolicyGroups", []interface{}{arg1, arg2})
	fake.listPolicyGroupsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePolicyGroupManager) ListPolicyGroupsCallCount() int {
	fake.listPolicyGroupsMutex.RLock()
	defer fake.listPolicyGroupsMutex.RUnlock()
	return len(fake.listPolicyGroupsArgsForCall)
}

func (fake *FakePolicyGroupManager) ListPolicyGroupsCalls(stub func(context.Context, *v1alpha1.ListPolicyGroupsRequest) (*v1alpha1.ListPolicyGroupsResponse, error)) {
	fake.listPolicyGroupsMutex.Lock()
	defer fake.listPolicyGroupsMutex.Unlock()
	fake.ListPolicyGroupsStub = stub
}

func (fake *FakePolicyGroupManager) ListPolicyGroupsArgsForCall(i int) (context.Context, *v1alpha1.ListPolicyGroupsRequest) {
	fake.listPolicyGroupsMutex.RLock()
	defer fake.listPolicyGroupsMutex.RUnlock()
	argsForCall := fake.listPolicyGroupsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePolicyGroupManager) ListPolicyGroupsReturns(result1 *v1alpha1.ListPolicyGroupsResponse, result2 error) {
	fake.listPolicyGroupsMutex.Lock()
	defer fake.listPolicyGroupsMutex.Unlock()
	fake.ListPolicyGroupsStub = nil
	fake.listPolicyGroupsReturns = struct {
		result1 *v1alpha1.ListPolicyGroupsResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyGroupManager) ListPolicyGroupsReturnsOnCall(i int, result1 *v1alpha1.ListPolicyGroupsResponse, result2 error) {
	fake.listPolicyGroupsMutex.Lock()
	defer fake.listPolicyGroupsMutex.Unlock()
	fake.ListPolicyGroupsStub = nil
	if fake.listPolicyGroupsReturnsOnCall == nil {
		fake.listPolicyGroupsReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.ListPolicyGroupsResponse
			result2 error
		})
	}
	fake.listPolicyGroupsReturnsOnCall[i] = struct {
		result1 *v1alpha1.ListPolicyGroupsResponse
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyGroupManager) UpdatePolicyGroup(arg1 context.Context, arg2 *v1alpha1.PolicyGroup) (*v1alpha1.PolicyGroup, error) {
	fake.updatePolicyGroupMutex.Lock()
	ret, specificReturn := fake.updatePolicyGroupReturnsOnCall[len(fake.updatePolicyGroupArgsForCall)]
	fake.updatePolicyGroupArgsForCall = append(fake.updatePolicyGroupArgsForCall, struct {
		arg1 context.Context
		arg2 *v1alpha1.PolicyGroup
	}{arg1, arg2})
	stub := fake.UpdatePolicyGroupStub
	fakeReturns := fake.updatePolicyGroupReturns
	fake.recordInvocation("UpdatePolicyGroup", []interface{}{arg1, arg2})
	fake.updatePolicyGroupMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakePolicyGroupManager) UpdatePolicyGroupCallCount() int {
	fake.updatePolicyGroupMutex.RLock()
	defer fake.updatePolicyGroupMutex.RUnlock()
	return len(fake.updatePolicyGroupArgsForCall)
}

func (fake *FakePolicyGroupManager) UpdatePolicyGroupCalls(stub func(context.Context, *v1alpha1.PolicyGroup) (*v1alpha1.PolicyGroup, error)) {
	fake.updatePolicyGroupMutex.Lock()
	defer fake.updatePolicyGroupMutex.Unlock()
	fake.UpdatePolicyGroupStub = stub
}

func (fake *FakePolicyGroupManager) UpdatePolicyGroupArgsForCall(i int) (context.Context, *v1alpha1.PolicyGroup) {
	fake.updatePolicyGroupMutex.RLock()
	defer fake.updatePolicyGroupMutex.RUnlock()
	argsForCall := fake.updatePolicyGroupArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakePolicyGroupManager) UpdatePolicyGroupReturns(result1 *v1alpha1.PolicyGroup, result2 error) {
	fake.updatePolicyGroupMutex.Lock()
	defer fake.updatePolicyGroupMutex.Unlock()
	fake.UpdatePolicyGroupStub = nil
	fake.updatePolicyGroupReturns = struct {
		result1 *v1alpha1.PolicyGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyGroupManager) UpdatePolicyGroupReturnsOnCall(i int, result1 *v1alpha1.PolicyGroup, result2 error) {
	fake.updatePolicyGroupMutex.Lock()
	defer fake.updatePolicyGroupMutex.Unlock()
	fake.UpdatePolicyGroupStub = nil
	if fake.updatePolicyGroupReturnsOnCall == nil {
		fake.updatePolicyGroupReturnsOnCall = make(map[int]struct {
			result1 *v1alpha1.PolicyGroup
			result2 error
		})
	}
	fake.updatePolicyGroupReturnsOnCall[i] = struct {
		result1 *v1alpha1.PolicyGroup
		result2 error
	}{result1, result2}
}

func (fake *FakePolicyGroupManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createPolicyGroupMutex.RLock()
	defer fake.createPolicyGroupMutex.RUnlock()
	fake.getPolicyGroupMutex.RLock()
	defer fake.getPolicyGroupMutex.RUnlock()
	fake.listPolicyGroupsMutex.RLock()
	defer fake.listPolicyGroupsMutex.RUnlock()
	fake.updatePolicyGroupMutex.RLock()
	defer fake.updatePolicyGroupMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakePolicyGroupManager) recordInvocation(key string, args []interface{}) {
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

var _ policy.PolicyGroupManager = new(FakePolicyGroupManager)
