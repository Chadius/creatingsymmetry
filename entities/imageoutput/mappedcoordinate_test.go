package imageoutput_test

import (
	"github.com/chadius/creatingsymmetry/entities/imageoutput"
	. "gopkg.in/check.v1"
	"math"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type MappedCoordinateTest struct {
}

var _ = Suite(&MappedCoordinateTest{})

func (suite *MappedCoordinateTest) TestCreateMappedCoordinateUsingOutputImageCoordinates(checker *C) {
	coordinate := imageoutput.NewMappedCoordinateUsingOutputImageCoordinates(200, 30)
	checker.Assert(coordinate.OutputImageX(), Equals, 200)
	checker.Assert(coordinate.OutputImageY(), Equals, 30)
}

func (suite *MappedCoordinateTest) TestCreateMappedCoordinateUsingTransformedCoordinates(checker *C) {
	coordinate := imageoutput.NewMappedCoordinateUsingTransformedCoordinates(100e0, -20e-3)
	checker.Assert(coordinate.TransformedX(), Equals, 100e0)
	checker.Assert(coordinate.TransformedY(), Equals, -20e-3)
}

func (suite *MappedCoordinateTest) TestUpdateTransformedCoordinates(checker *C) {
	coordinate := imageoutput.NewMappedCoordinateUsingTransformedCoordinates(100e0, -20e-3)

	coordinate.UpdateTransformedCoordinates(30e0, -5e3)

	checker.Assert(coordinate.TransformedX(), Equals, 30e0)
	checker.Assert(coordinate.TransformedY(), Equals, -5e3)
}

func (suite *MappedCoordinateTest) TestUpdatePatternViewportCoordinates(checker *C) {
	coordinate := imageoutput.NewMappedCoordinateUsingTransformedCoordinates(100e0, -20e-3)

	coordinate.UpdatePatternViewportCoordinates(30e0, -5e3)

	checker.Assert(coordinate.PatternViewportX(), Equals, 30e0)
	checker.Assert(coordinate.PatternViewportY(), Equals, -5e3)
}

func (suite *MappedCoordinateTest) TestUpdateInputImageCoordinates(checker *C) {
	coordinate := imageoutput.NewMappedCoordinateUsingInputImageCoordinates(30, -20)
	checker.Assert(coordinate.InputImageX(), Equals, 30)
	checker.Assert(coordinate.InputImageY(), Equals, -20)
}

func (suite *MappedCoordinateTest) TestChecksIfOnePartCannotBeCompared(checker *C) {
	coordinate := imageoutput.NewMappedCoordinateUsingTransformedCoordinates(100e0, -20e-3)
	checker.Assert(coordinate.CanBeCompared(), Equals, true)
	coordinateWithXInfinity := imageoutput.NewMappedCoordinateUsingTransformedCoordinates(math.Inf(1), -20e-3)
	checker.Assert(coordinateWithXInfinity.CanBeCompared(), Equals, false)
	coordinateWithYInfinity := imageoutput.NewMappedCoordinateUsingTransformedCoordinates(100e0, math.Inf(-1))
	checker.Assert(coordinateWithYInfinity.CanBeCompared(), Equals, false)

	coordinateWithXNan := imageoutput.NewMappedCoordinateUsingTransformedCoordinates(math.NaN(), 0)
	checker.Assert(coordinateWithXNan.CanBeCompared(), Equals, false)

	coordinateWithYNan := imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, math.NaN())
	checker.Assert(coordinateWithYNan.CanBeCompared(), Equals, false)
}

func (suite *MappedCoordinateTest) TestMarkCoordinateAsFiltered(checker *C) {
	coordinate := imageoutput.NewMappedCoordinateUsingTransformedCoordinates(100e0, -20e-3)
	checker.Assert(coordinate.SatisfiesFilter(), Equals, false)
	coordinate.MarkAsSatisfyingFilter()
	checker.Assert(coordinate.SatisfiesFilter(), Equals, true)
}

func (suite *MappedCoordinateTest) TestStoreMappedCoordinate(checker *C) {
	coordinate := imageoutput.NewMappedCoordinateUsingTransformedCoordinates(100e0, -20e-3)
	checker.Assert(coordinate.HasMappedCoordinate(), Equals, false)
	coordinate.StoreMappedCoordinate(2e0, -3e-3)
	checker.Assert(coordinate.HasMappedCoordinate(), Equals, true)

	x, y := coordinate.MappedCoordinate()
	checker.Assert(x, Equals, 2e0)
	checker.Assert(y, Equals, -3e-3)
}
