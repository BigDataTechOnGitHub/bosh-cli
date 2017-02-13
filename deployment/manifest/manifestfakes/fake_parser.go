// This file was generated by counterfeiter
package manifestfakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-cli/deployment/manifest"
	"github.com/cloudfoundry/bosh-cli/deployment/template"
)

type FakeParser struct {
	ParseStub        func(interpolatedTemplate template.InterpolatedTemplate, path string) (manifest.Manifest, error)
	parseMutex       sync.RWMutex
	parseArgsForCall []struct {
		interpolatedTemplate template.InterpolatedTemplate
		path                 string
	}
	parseReturns struct {
		result1 manifest.Manifest
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeParser) Parse(interpolatedTemplate template.InterpolatedTemplate, path string) (manifest.Manifest, error) {
	fake.parseMutex.Lock()
	fake.parseArgsForCall = append(fake.parseArgsForCall, struct {
		interpolatedTemplate template.InterpolatedTemplate
		path                 string
	}{interpolatedTemplate, path})
	fake.recordInvocation("Parse", []interface{}{interpolatedTemplate, path})
	fake.parseMutex.Unlock()
	if fake.ParseStub != nil {
		return fake.ParseStub(interpolatedTemplate, path)
	} else {
		return fake.parseReturns.result1, fake.parseReturns.result2
	}
}

func (fake *FakeParser) ParseCallCount() int {
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	return len(fake.parseArgsForCall)
}

func (fake *FakeParser) ParseArgsForCall(i int) (template.InterpolatedTemplate, string) {
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	return fake.parseArgsForCall[i].interpolatedTemplate, fake.parseArgsForCall[i].path
}

func (fake *FakeParser) ParseReturns(result1 manifest.Manifest, result2 error) {
	fake.ParseStub = nil
	fake.parseReturns = struct {
		result1 manifest.Manifest
		result2 error
	}{result1, result2}
}

func (fake *FakeParser) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.parseMutex.RLock()
	defer fake.parseMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeParser) recordInvocation(key string, args []interface{}) {
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

var _ manifest.Parser = new(FakeParser)
