// Code generated by counterfeiter. DO NOT EDIT.
package imageoutputfakes

import (
	"image/color"
	"sync"

	"github.com/chadius/creatingsymmetry/entities/imageoutput"
)

type FakeEyedropper struct {
	ConvertCoordinatesToColorsStub        func(*imageoutput.CoordinateCollection) *[]color.Color
	convertCoordinatesToColorsMutex       sync.RWMutex
	convertCoordinatesToColorsArgsForCall []struct {
		arg1 *imageoutput.CoordinateCollection
	}
	convertCoordinatesToColorsReturns struct {
		result1 *[]color.Color
	}
	convertCoordinatesToColorsReturnsOnCall map[int]struct {
		result1 *[]color.Color
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeEyedropper) ConvertCoordinatesToColors(arg1 *imageoutput.CoordinateCollection) *[]color.Color {
	fake.convertCoordinatesToColorsMutex.Lock()
	ret, specificReturn := fake.convertCoordinatesToColorsReturnsOnCall[len(fake.convertCoordinatesToColorsArgsForCall)]
	fake.convertCoordinatesToColorsArgsForCall = append(fake.convertCoordinatesToColorsArgsForCall, struct {
		arg1 *imageoutput.CoordinateCollection
	}{arg1})
	stub := fake.ConvertCoordinatesToColorsStub
	fakeReturns := fake.convertCoordinatesToColorsReturns
	fake.recordInvocation("ConvertCoordinatesToColors", []interface{}{arg1})
	fake.convertCoordinatesToColorsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeEyedropper) ConvertCoordinatesToColorsCallCount() int {
	fake.convertCoordinatesToColorsMutex.RLock()
	defer fake.convertCoordinatesToColorsMutex.RUnlock()
	return len(fake.convertCoordinatesToColorsArgsForCall)
}

func (fake *FakeEyedropper) ConvertCoordinatesToColorsCalls(stub func(*imageoutput.CoordinateCollection) *[]color.Color) {
	fake.convertCoordinatesToColorsMutex.Lock()
	defer fake.convertCoordinatesToColorsMutex.Unlock()
	fake.ConvertCoordinatesToColorsStub = stub
}

func (fake *FakeEyedropper) ConvertCoordinatesToColorsArgsForCall(i int) *imageoutput.CoordinateCollection {
	fake.convertCoordinatesToColorsMutex.RLock()
	defer fake.convertCoordinatesToColorsMutex.RUnlock()
	argsForCall := fake.convertCoordinatesToColorsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeEyedropper) ConvertCoordinatesToColorsReturns(result1 *[]color.Color) {
	fake.convertCoordinatesToColorsMutex.Lock()
	defer fake.convertCoordinatesToColorsMutex.Unlock()
	fake.ConvertCoordinatesToColorsStub = nil
	fake.convertCoordinatesToColorsReturns = struct {
		result1 *[]color.Color
	}{result1}
}

func (fake *FakeEyedropper) ConvertCoordinatesToColorsReturnsOnCall(i int, result1 *[]color.Color) {
	fake.convertCoordinatesToColorsMutex.Lock()
	defer fake.convertCoordinatesToColorsMutex.Unlock()
	fake.ConvertCoordinatesToColorsStub = nil
	if fake.convertCoordinatesToColorsReturnsOnCall == nil {
		fake.convertCoordinatesToColorsReturnsOnCall = make(map[int]struct {
			result1 *[]color.Color
		})
	}
	fake.convertCoordinatesToColorsReturnsOnCall[i] = struct {
		result1 *[]color.Color
	}{result1}
}

func (fake *FakeEyedropper) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.convertCoordinatesToColorsMutex.RLock()
	defer fake.convertCoordinatesToColorsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeEyedropper) recordInvocation(key string, args []interface{}) {
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

var _ imageoutput.Eyedropper = new(FakeEyedropper)