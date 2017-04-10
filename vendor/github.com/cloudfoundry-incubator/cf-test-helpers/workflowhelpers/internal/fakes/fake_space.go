// This file was generated by counterfeiter
package fakes

import "sync"

type FakeSpace struct {
	CreateStub              func()
	createMutex             sync.RWMutex
	createArgsForCall       []struct{}
	DestroyStub             func()
	destroyMutex            sync.RWMutex
	destroyArgsForCall      []struct{}
	ShouldRemainStub        func() bool
	shouldRemainMutex       sync.RWMutex
	shouldRemainArgsForCall []struct{}
	shouldRemainReturns     struct {
		result1 bool
	}
	OrganizationNameStub        func() string
	organizationNameMutex       sync.RWMutex
	organizationNameArgsForCall []struct{}
	organizationNameReturns     struct {
		result1 string
	}
}

func (fake *FakeSpace) Create() {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct{}{})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		fake.CreateStub()
	}
}

func (fake *FakeSpace) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeSpace) Destroy() {
	fake.destroyMutex.Lock()
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct{}{})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		fake.DestroyStub()
	}
}

func (fake *FakeSpace) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeSpace) ShouldRemain() bool {
	fake.shouldRemainMutex.Lock()
	fake.shouldRemainArgsForCall = append(fake.shouldRemainArgsForCall, struct{}{})
	fake.shouldRemainMutex.Unlock()
	if fake.ShouldRemainStub != nil {
		return fake.ShouldRemainStub()
	} else {
		return fake.shouldRemainReturns.result1
	}
}

func (fake *FakeSpace) ShouldRemainCallCount() int {
	fake.shouldRemainMutex.RLock()
	defer fake.shouldRemainMutex.RUnlock()
	return len(fake.shouldRemainArgsForCall)
}

func (fake *FakeSpace) ShouldRemainReturns(result1 bool) {
	fake.ShouldRemainStub = nil
	fake.shouldRemainReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeSpace) OrganizationName() string {
	fake.organizationNameMutex.Lock()
	fake.organizationNameArgsForCall = append(fake.organizationNameArgsForCall, struct{}{})
	fake.organizationNameMutex.Unlock()
	if fake.OrganizationNameStub != nil {
		return fake.OrganizationNameStub()
	} else {
		return fake.organizationNameReturns.result1
	}
}

func (fake *FakeSpace) OrganizationNameCallCount() int {
	fake.organizationNameMutex.RLock()
	defer fake.organizationNameMutex.RUnlock()
	return len(fake.organizationNameArgsForCall)
}

func (fake *FakeSpace) OrganizationNameReturns(result1 string) {
	fake.OrganizationNameStub = nil
	fake.organizationNameReturns = struct {
		result1 string
	}{result1}
}