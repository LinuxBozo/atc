// This file was generated by counterfeiter
package enginefakes

import (
	"sync"

	"github.com/concourse/atc"
	"github.com/concourse/atc/engine"
	"github.com/pivotal-golang/lager"
)

type FakeBuild struct {
	MetadataStub        func() string
	metadataMutex       sync.RWMutex
	metadataArgsForCall []struct{}
	metadataReturns     struct {
		result1 string
	}
	PublicPlanStub        func(lager.Logger) (atc.PublicBuildPlan, bool, error)
	publicPlanMutex       sync.RWMutex
	publicPlanArgsForCall []struct {
		arg1 lager.Logger
	}
	publicPlanReturns struct {
		result1 atc.PublicBuildPlan
		result2 bool
		result3 error
	}
	AbortStub        func(lager.Logger) error
	abortMutex       sync.RWMutex
	abortArgsForCall []struct {
		arg1 lager.Logger
	}
	abortReturns struct {
		result1 error
	}
	ResumeStub        func(lager.Logger)
	resumeMutex       sync.RWMutex
	resumeArgsForCall []struct {
		arg1 lager.Logger
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBuild) Metadata() string {
	fake.metadataMutex.Lock()
	fake.metadataArgsForCall = append(fake.metadataArgsForCall, struct{}{})
	fake.recordInvocation("Metadata", []interface{}{})
	fake.metadataMutex.Unlock()
	if fake.MetadataStub != nil {
		return fake.MetadataStub()
	} else {
		return fake.metadataReturns.result1
	}
}

func (fake *FakeBuild) MetadataCallCount() int {
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	return len(fake.metadataArgsForCall)
}

func (fake *FakeBuild) MetadataReturns(result1 string) {
	fake.MetadataStub = nil
	fake.metadataReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBuild) PublicPlan(arg1 lager.Logger) (atc.PublicBuildPlan, bool, error) {
	fake.publicPlanMutex.Lock()
	fake.publicPlanArgsForCall = append(fake.publicPlanArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("PublicPlan", []interface{}{arg1})
	fake.publicPlanMutex.Unlock()
	if fake.PublicPlanStub != nil {
		return fake.PublicPlanStub(arg1)
	} else {
		return fake.publicPlanReturns.result1, fake.publicPlanReturns.result2, fake.publicPlanReturns.result3
	}
}

func (fake *FakeBuild) PublicPlanCallCount() int {
	fake.publicPlanMutex.RLock()
	defer fake.publicPlanMutex.RUnlock()
	return len(fake.publicPlanArgsForCall)
}

func (fake *FakeBuild) PublicPlanArgsForCall(i int) lager.Logger {
	fake.publicPlanMutex.RLock()
	defer fake.publicPlanMutex.RUnlock()
	return fake.publicPlanArgsForCall[i].arg1
}

func (fake *FakeBuild) PublicPlanReturns(result1 atc.PublicBuildPlan, result2 bool, result3 error) {
	fake.PublicPlanStub = nil
	fake.publicPlanReturns = struct {
		result1 atc.PublicBuildPlan
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeBuild) Abort(arg1 lager.Logger) error {
	fake.abortMutex.Lock()
	fake.abortArgsForCall = append(fake.abortArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Abort", []interface{}{arg1})
	fake.abortMutex.Unlock()
	if fake.AbortStub != nil {
		return fake.AbortStub(arg1)
	} else {
		return fake.abortReturns.result1
	}
}

func (fake *FakeBuild) AbortCallCount() int {
	fake.abortMutex.RLock()
	defer fake.abortMutex.RUnlock()
	return len(fake.abortArgsForCall)
}

func (fake *FakeBuild) AbortArgsForCall(i int) lager.Logger {
	fake.abortMutex.RLock()
	defer fake.abortMutex.RUnlock()
	return fake.abortArgsForCall[i].arg1
}

func (fake *FakeBuild) AbortReturns(result1 error) {
	fake.AbortStub = nil
	fake.abortReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBuild) Resume(arg1 lager.Logger) {
	fake.resumeMutex.Lock()
	fake.resumeArgsForCall = append(fake.resumeArgsForCall, struct {
		arg1 lager.Logger
	}{arg1})
	fake.recordInvocation("Resume", []interface{}{arg1})
	fake.resumeMutex.Unlock()
	if fake.ResumeStub != nil {
		fake.ResumeStub(arg1)
	}
}

func (fake *FakeBuild) ResumeCallCount() int {
	fake.resumeMutex.RLock()
	defer fake.resumeMutex.RUnlock()
	return len(fake.resumeArgsForCall)
}

func (fake *FakeBuild) ResumeArgsForCall(i int) lager.Logger {
	fake.resumeMutex.RLock()
	defer fake.resumeMutex.RUnlock()
	return fake.resumeArgsForCall[i].arg1
}

func (fake *FakeBuild) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.metadataMutex.RLock()
	defer fake.metadataMutex.RUnlock()
	fake.publicPlanMutex.RLock()
	defer fake.publicPlanMutex.RUnlock()
	fake.abortMutex.RLock()
	defer fake.abortMutex.RUnlock()
	fake.resumeMutex.RLock()
	defer fake.resumeMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeBuild) recordInvocation(key string, args []interface{}) {
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

var _ engine.Build = new(FakeBuild)