// Code generated by counterfeiter. DO NOT EDIT.
package imageoutputfakes

import (
	"sync"

	"github.com/chadius/creatingsymmetry/entities/imageoutput"
)

type FakeCoordinateThreshold struct {
	FilterAndMarkMappedCoordinateCollectionStub        func(*imageoutput.CoordinateCollection)
	filterAndMarkMappedCoordinateCollectionMutex       sync.RWMutex
	filterAndMarkMappedCoordinateCollectionArgsForCall []struct {
		arg1 *imageoutput.CoordinateCollection
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCoordinateThreshold) FilterAndMarkMappedCoordinateCollection(arg1 *imageoutput.CoordinateCollection) {
	fake.filterAndMarkMappedCoordinateCollectionMutex.Lock()
	fake.filterAndMarkMappedCoordinateCollectionArgsForCall = append(fake.filterAndMarkMappedCoordinateCollectionArgsForCall, struct {
		arg1 *imageoutput.CoordinateCollection
	}{arg1})
	stub := fake.FilterAndMarkMappedCoordinateCollectionStub
	fake.recordInvocation("FilterAndMarkMappedCoordinateCollection", []interface{}{arg1})
	fake.filterAndMarkMappedCoordinateCollectionMutex.Unlock()
	if stub != nil {
		fake.FilterAndMarkMappedCoordinateCollectionStub(arg1)
	}
}

func (fake *FakeCoordinateThreshold) FilterAndMarkMappedCoordinateCollectionCallCount() int {
	fake.filterAndMarkMappedCoordinateCollectionMutex.RLock()
	defer fake.filterAndMarkMappedCoordinateCollectionMutex.RUnlock()
	return len(fake.filterAndMarkMappedCoordinateCollectionArgsForCall)
}

func (fake *FakeCoordinateThreshold) FilterAndMarkMappedCoordinateCollectionCalls(stub func(*imageoutput.CoordinateCollection)) {
	fake.filterAndMarkMappedCoordinateCollectionMutex.Lock()
	defer fake.filterAndMarkMappedCoordinateCollectionMutex.Unlock()
	fake.FilterAndMarkMappedCoordinateCollectionStub = stub
}

func (fake *FakeCoordinateThreshold) FilterAndMarkMappedCoordinateCollectionArgsForCall(i int) *imageoutput.CoordinateCollection {
	fake.filterAndMarkMappedCoordinateCollectionMutex.RLock()
	defer fake.filterAndMarkMappedCoordinateCollectionMutex.RUnlock()
	argsForCall := fake.filterAndMarkMappedCoordinateCollectionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeCoordinateThreshold) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.filterAndMarkMappedCoordinateCollectionMutex.RLock()
	defer fake.filterAndMarkMappedCoordinateCollectionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeCoordinateThreshold) recordInvocation(key string, args []interface{}) {
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

var _ imageoutput.CoordinateThreshold = new(FakeCoordinateThreshold)
