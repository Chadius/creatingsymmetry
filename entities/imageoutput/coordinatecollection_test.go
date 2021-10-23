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
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(-10, 20),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(20, 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, -100),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, 200),
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
	checker.Assert((*collection.Coordinates())[0].TransformedX(), Equals, -10.0)
	checker.Assert((*collection.Coordinates())[0].TransformedY(), Equals, 20.0)
	checker.Assert((*collection.Coordinates())[1].TransformedX(), Equals, 20.0)
	checker.Assert((*collection.Coordinates())[1].TransformedY(), Equals, 0.0)
	checker.Assert((*collection.Coordinates())[2].TransformedX(), Equals, 0.0)
	checker.Assert((*collection.Coordinates())[2].TransformedY(), Equals, -100.0)
	checker.Assert((*collection.Coordinates())[3].TransformedX(), Equals, 0.0)
	checker.Assert((*collection.Coordinates())[3].TransformedY(), Equals, 200.0)
}

type CoordinateCollectionTests struct {
}

var _ = Suite(&CoordinateCollectionTests{})

func (suite *CoordinateCollectionTests) TestReturnsMinimumAndMaximums(checker *C) {
	coordinates := []*imageoutput.MappedCoordinate{
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(-10, 20),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(20, 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, -100),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, 200),
	}
	collection := imageoutput.CoordinateCollectionBuilder().WithCoordinates(&coordinates).Build()
	(*collection.Coordinates())[0].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[1].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[2].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[3].MarkAsSatisfyingFilter()
	checker.Assert(collection.MinimumTransformedX(), Equals, float64(-10))
	checker.Assert(collection.MaximumTransformedX(), Equals, float64(20))
	checker.Assert(collection.MinimumTransformedY(), Equals, float64(-100))
	checker.Assert(collection.MaximumTransformedY(), Equals, float64(200))
}

func (suite *CoordinateCollectionTests) TestReturnsMinimumAndMaximumsRespectingSatisfiedFilter(checker *C) {
	coordinates := []*imageoutput.MappedCoordinate{
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(-1, 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(1, 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, -2),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, 2),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(-100, 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(100, 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, -100),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, 100),
	}
	collection := imageoutput.CoordinateCollectionBuilder().WithCoordinates(&coordinates).Build()
	(*collection.Coordinates())[0].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[1].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[2].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[3].MarkAsSatisfyingFilter()
	checker.Assert(collection.MinimumTransformedX(), Equals, float64(-1))
	checker.Assert(collection.MaximumTransformedX(), Equals, float64(1))
	checker.Assert(collection.MinimumTransformedY(), Equals, float64(-2))
	checker.Assert(collection.MaximumTransformedY(), Equals, float64(2))
}

func (suite *CoordinateCollectionTests) TestReturnsMinimumAndMaximumsIgnoringInfinity(checker *C) {
	coordinates := []*imageoutput.MappedCoordinate{
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(-1, 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(1, 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, -2),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, 2),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(-100, math.Inf(-1)),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(100, math.Inf(1)),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(math.Inf(1), -100),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(math.Inf(-1), 100),
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
	checker.Assert(collection.MinimumTransformedX(), Equals, float64(-1))
	checker.Assert(collection.MaximumTransformedX(), Equals, float64(1))
	checker.Assert(collection.MinimumTransformedY(), Equals, float64(-2))
	checker.Assert(collection.MaximumTransformedY(), Equals, float64(2))
}

func (suite *CoordinateCollectionTests) TestReturnsNaNMinimumIfAllCoordinatesAreInvalid(checker *C) {
	coordinates := []*imageoutput.MappedCoordinate{
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(math.Inf(-1), -10),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(20, 0),
	}
	collection := imageoutput.CoordinateCollectionBuilder().WithCoordinates(&coordinates).Build()
	(*collection.Coordinates())[0].MarkAsSatisfyingFilter()
	checker.Assert(math.IsNaN(collection.MinimumTransformedX()), Equals, true)
	checker.Assert(math.IsNaN(collection.MinimumTransformedY()), Equals, true)
}
