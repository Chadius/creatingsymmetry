package imageoutput_test

import (
	"github.com/Chadius/creating-symmetry/entities/imageoutput"
	. "gopkg.in/check.v1"
	"math"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MappedCoordinateTest struct {
}

var _ = Suite(&MappedCoordinateTest{})

func (suite *MappedCoordinateTest) TestCreateMappedCoordinate(checker *C) {
	coordinate := imageoutput.NewMappedCoordinate(100e0, -20e-3)
	checker.Assert(coordinate.X(), Equals, 100e0)
	checker.Assert(coordinate.Y(), Equals, -20e-3)
}

func (suite *MappedCoordinateTest) TestChecksIfOnePartIsInfinity(checker *C) {
	coordinate := imageoutput.NewMappedCoordinate(100e0, -20e-3)
	checker.Assert(coordinate.IsAtInfinity(), Equals, false)
	coordinateWithXInfinity := imageoutput.NewMappedCoordinate(math.Inf(1), -20e-3)
	checker.Assert(coordinateWithXInfinity.IsAtInfinity(), Equals, true)
	coordinateWithYInfinity := imageoutput.NewMappedCoordinate(100e0, math.Inf(-1))
	checker.Assert(coordinateWithYInfinity.IsAtInfinity(), Equals, true)
}

func (suite *MappedCoordinateTest) TestMarkCoordinateAsFiltered(checker *C) {
	coordinate := imageoutput.NewMappedCoordinate(100e0, -20e-3)
	checker.Assert(coordinate.SatisfiesFilter(), Equals, false)
	coordinate.MarkAsSatisfyingFilter()
	checker.Assert(coordinate.SatisfiesFilter(), Equals, true)
}

func (suite *MappedCoordinateTest) TestStoreMappedCoordinate(checker *C) {
	coordinate := imageoutput.NewMappedCoordinate(100e0, -20e-3)
	checker.Assert(coordinate.HasMappedCoordinate(), Equals, false)
	coordinate.StoreMappedCoordinate(2e0, -3e-3)
	checker.Assert(coordinate.HasMappedCoordinate(), Equals, true)

	x, y := coordinate.MappedCoordinate()
	checker.Assert(x, Equals, 2e0)
	checker.Assert(y, Equals, -3e-3)
}