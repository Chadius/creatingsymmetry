package imageoutput_test

import (
	"github.com/Chadius/creating-symmetry/entities/imageoutput"
	. "gopkg.in/check.v1"
	"math"
)

type CoordinateCollectionBuilderTests struct {
}

var _ = Suite(&CoordinateCollectionBuilderTests{})

func (suite *CoordinateCollectionBuilderTests) TestSetupCreateDataRangeWithArray(checker *C) {
	coordinates := []*imageoutput.MappedCoordinate{
		imageoutput.NewMappedCoordinate(-10, 20),
		imageoutput.NewMappedCoordinate(20, 0),
		imageoutput.NewMappedCoordinate(0, -100),
		imageoutput.NewMappedCoordinate(0, 200),
	}
	collection := imageoutput.CoordinateCollectionBuilder().WithCoordinates(&coordinates).Build()
	checker.Assert(collection.Coordinates(), Equals, &coordinates)
}

func (suite *CoordinateCollectionBuilderTests) TestSetupCreateDataRangeWithComplexNumbers(checker *C) {
	coordinates := []complex128{
		complex(-10, 20),
		complex(20, 0),
		complex(0, -100),
		complex(0, 200),
	}
	collection := imageoutput.CoordinateCollectionBuilder().WithComplexNumbers(&coordinates).Build()
	checker.Assert(*collection.Coordinates(), HasLen, 4)
	checker.Assert((*collection.Coordinates())[0].X(), Equals, -10.0)
	checker.Assert((*collection.Coordinates())[0].Y(), Equals, 20.0)
	checker.Assert((*collection.Coordinates())[1].X(), Equals, 20.0)
	checker.Assert((*collection.Coordinates())[1].Y(), Equals, 0.0)
	checker.Assert((*collection.Coordinates())[2].X(), Equals, 0.0)
	checker.Assert((*collection.Coordinates())[2].Y(), Equals, -100.0)
	checker.Assert((*collection.Coordinates())[3].X(), Equals, 0.0)
	checker.Assert((*collection.Coordinates())[3].Y(), Equals, 200.0)
}

type CoordinateCollectionTests struct {
}

var _ = Suite(&CoordinateCollectionTests{})

func (suite *CoordinateCollectionTests) TestReturnsMinimumAndMaximums(checker *C) {
	coordinates := []*imageoutput.MappedCoordinate{
		imageoutput.NewMappedCoordinate(-10, 20),
		imageoutput.NewMappedCoordinate(20, 0),
		imageoutput.NewMappedCoordinate(0, -100),
		imageoutput.NewMappedCoordinate(0, 200),
	}
	collection := imageoutput.CoordinateCollectionBuilder().WithCoordinates(&coordinates).Build()
	(*collection.Coordinates())[0].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[1].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[2].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[3].MarkAsSatisfyingFilter()
	checker.Assert(collection.MinimumX(), Equals, float64(-10))
	checker.Assert(collection.MaximumX(), Equals, float64(20))
	checker.Assert(collection.MinimumY(), Equals, float64(-100))
	checker.Assert(collection.MaximumY(), Equals, float64(200))
}

func (suite *CoordinateCollectionTests) TestReturnsMinimumAndMaximumsRespectingSatisfiedFilter(checker *C) {
	coordinates := []*imageoutput.MappedCoordinate{
		imageoutput.NewMappedCoordinate(-1, 0),
		imageoutput.NewMappedCoordinate(1, 0),
		imageoutput.NewMappedCoordinate(0, -2),
		imageoutput.NewMappedCoordinate(0, 2),
		imageoutput.NewMappedCoordinate(-100, 0),
		imageoutput.NewMappedCoordinate(100, 0),
		imageoutput.NewMappedCoordinate(0, -100),
		imageoutput.NewMappedCoordinate(0, 100),
	}
	collection := imageoutput.CoordinateCollectionBuilder().WithCoordinates(&coordinates).Build()
	(*collection.Coordinates())[0].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[1].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[2].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[3].MarkAsSatisfyingFilter()
	checker.Assert(collection.MinimumX(), Equals, float64(-1))
	checker.Assert(collection.MaximumX(), Equals, float64(1))
	checker.Assert(collection.MinimumY(), Equals, float64(-2))
	checker.Assert(collection.MaximumY(), Equals, float64(2))
}

func (suite *CoordinateCollectionTests) TestReturnsMinimumAndMaximumsIgnoringInfinity(checker *C) {
	coordinates := []*imageoutput.MappedCoordinate{
		imageoutput.NewMappedCoordinate(-1, 0),
		imageoutput.NewMappedCoordinate(1, 0),
		imageoutput.NewMappedCoordinate(0, -2),
		imageoutput.NewMappedCoordinate(0, 2),
		imageoutput.NewMappedCoordinate(-100, math.Inf(-1)),
		imageoutput.NewMappedCoordinate(100, math.Inf(1)),
		imageoutput.NewMappedCoordinate(math.Inf(1), -100),
		imageoutput.NewMappedCoordinate(math.Inf(-1), 100),
	}
	collection := imageoutput.CoordinateCollectionBuilder().WithCoordinates(&coordinates).Build()
	(*collection.Coordinates())[0].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[1].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[2].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[3].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[4].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[5].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[6].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[7].MarkAsSatisfyingFilter()
	checker.Assert(collection.MinimumX(), Equals, float64(-1))
	checker.Assert(collection.MaximumX(), Equals, float64(1))
	checker.Assert(collection.MinimumY(), Equals, float64(-2))
	checker.Assert(collection.MaximumY(), Equals, float64(2))
}

func (suite *CoordinateCollectionTests) TestReturnsNaNMinimumIfAllCoordinatesAreInvalid(checker *C) {
	coordinates := []*imageoutput.MappedCoordinate{
		imageoutput.NewMappedCoordinate(math.Inf(-1), -10),
		imageoutput.NewMappedCoordinate(20, 0),
	}
	collection := imageoutput.CoordinateCollectionBuilder().WithCoordinates(&coordinates).Build()
	(*collection.Coordinates())[0].MarkAsSatisfyingFilter()
	checker.Assert(math.IsNaN(collection.MinimumX()), Equals, true)
	checker.Assert(math.IsNaN(collection.MinimumY()), Equals, true)
}
