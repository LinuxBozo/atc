// This file was generated by counterfeiter
package teamserverfakes

import (
	"sync"

	"github.com/concourse/atc/api/teamserver"
	"github.com/concourse/atc/db"
)

type FakeTeamsDB struct {
	GetTeamsStub        func() ([]db.SavedTeam, error)
	getTeamsMutex       sync.RWMutex
	getTeamsArgsForCall []struct{}
	getTeamsReturns     struct {
		result1 []db.SavedTeam
		result2 error
	}
	CreateTeamStub        func(data db.Team) (db.SavedTeam, error)
	createTeamMutex       sync.RWMutex
	createTeamArgsForCall []struct {
		data db.Team
	}
	createTeamReturns struct {
		result1 db.SavedTeam
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTeamsDB) GetTeams() ([]db.SavedTeam, error) {
	fake.getTeamsMutex.Lock()
	fake.getTeamsArgsForCall = append(fake.getTeamsArgsForCall, struct{}{})
	fake.recordInvocation("GetTeams", []interface{}{})
	fake.getTeamsMutex.Unlock()
	if fake.GetTeamsStub != nil {
		return fake.GetTeamsStub()
	} else {
		return fake.getTeamsReturns.result1, fake.getTeamsReturns.result2
	}
}

func (fake *FakeTeamsDB) GetTeamsCallCount() int {
	fake.getTeamsMutex.RLock()
	defer fake.getTeamsMutex.RUnlock()
	return len(fake.getTeamsArgsForCall)
}

func (fake *FakeTeamsDB) GetTeamsReturns(result1 []db.SavedTeam, result2 error) {
	fake.GetTeamsStub = nil
	fake.getTeamsReturns = struct {
		result1 []db.SavedTeam
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamsDB) CreateTeam(data db.Team) (db.SavedTeam, error) {
	fake.createTeamMutex.Lock()
	fake.createTeamArgsForCall = append(fake.createTeamArgsForCall, struct {
		data db.Team
	}{data})
	fake.recordInvocation("CreateTeam", []interface{}{data})
	fake.createTeamMutex.Unlock()
	if fake.CreateTeamStub != nil {
		return fake.CreateTeamStub(data)
	} else {
		return fake.createTeamReturns.result1, fake.createTeamReturns.result2
	}
}

func (fake *FakeTeamsDB) CreateTeamCallCount() int {
	fake.createTeamMutex.RLock()
	defer fake.createTeamMutex.RUnlock()
	return len(fake.createTeamArgsForCall)
}

func (fake *FakeTeamsDB) CreateTeamArgsForCall(i int) db.Team {
	fake.createTeamMutex.RLock()
	defer fake.createTeamMutex.RUnlock()
	return fake.createTeamArgsForCall[i].data
}

func (fake *FakeTeamsDB) CreateTeamReturns(result1 db.SavedTeam, result2 error) {
	fake.CreateTeamStub = nil
	fake.createTeamReturns = struct {
		result1 db.SavedTeam
		result2 error
	}{result1, result2}
}

func (fake *FakeTeamsDB) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getTeamsMutex.RLock()
	defer fake.getTeamsMutex.RUnlock()
	fake.createTeamMutex.RLock()
	defer fake.createTeamMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTeamsDB) recordInvocation(key string, args []interface{}) {
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

var _ teamserver.TeamsDB = new(FakeTeamsDB)
