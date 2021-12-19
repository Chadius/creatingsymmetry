package imageoutput_test

import (
	"github.com/Chadius/creating-symmetry/entities/imageoutput"
	. "gopkg.in/check.v1"
	"math"
)

type CoordinateFilterTests struct {
}

var _ = Suite(&CoordinateFilterTests{})

func (suite *CoordinateFilterTests) TestCreateFilterWithBoundaries(checker *C) {
	filter := imageoutput.CoordinateFilterBuilder().WithMinimumX(-1e-5).WithMaximumX(2e1).WithMinimumY(-6e6).WithMaximumY(5e2).Build()
	checker.Assert(filter.MinimumX(), Equals, -1e-5)
	checker.Assert(filter.MaximumX(), Equals, 2e1)
	checker.Assert(filter.MinimumY(), Equals, -6e6)
	checker.Assert(filter.MaximumY(), Equals, 5e2)
}

func (suite *CoordinateFilterTests) TestFilterMarksCoordinateCollection(checker *C) {
	coordinates := []*imageoutput.MappedCoordinate{
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(-10, 20),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(20, 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, -100),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, 200),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, math.Inf(1)),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(math.NaN(), 0),
	}
	collection := imageoutput.CoordinateCollectionBuilder().WithCoordinates(&coordinates).Build()

	filter := imageoutput.CoordinateFilterBuilder().WithMinimumX(0).WithMaximumX(25).WithMinimumY(-6e6).WithMaximumY(50).Build()
	filter.FilterAndMarkMappedCoordinateCollection(collection)

	checker.Assert((*collection.Coordinates())[0].SatisfiesFilter(), Equals, false)
	checker.Assert((*collection.Coordinates())[1].SatisfiesFilter(), Equals, true)
	checker.Assert((*collection.Coordinates())[2].SatisfiesFilter(), Equals, true)
	checker.Assert((*collection.Coordinates())[3].SatisfiesFilter(), Equals, false)
	checker.Assert((*collection.Coordinates())[4].SatisfiesFilter(), Equals, false)
	checker.Assert((*collection.Coordinates())[5].SatisfiesFilter(), Equals, false)
}

func (suite *CoordinateFilterTests) TestNoFilterMeansAllCountableCoordinatesAreAccepted(checker *C) {
	coordinates := []*imageoutput.MappedCoordinate{
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(-10, 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(10, 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, -10),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, 10),

		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(math.NaN(), 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, math.Inf(-1)),
	}
	collection := imageoutput.CoordinateCollectionBuilder().WithCoordinates(&coordinates).Build()

	filter := imageoutput.CoordinateFilterBuilder().Build()
	filter.FilterAndMarkMappedCoordinateCollection(collection)

	checker.Assert((*collection.Coordinates())[0].SatisfiesFilter(), Equals, true)
	checker.Assert((*collection.Coordinates())[1].SatisfiesFilter(), Equals, true)
	checker.Assert((*collection.Coordinates())[2].SatisfiesFilter(), Equals, true)
	checker.Assert((*collection.Coordinates())[3].SatisfiesFilter(), Equals, true)
	checker.Assert((*collection.Coordinates())[4].SatisfiesFilter(), Equals, false)
	checker.Assert((*collection.Coordinates())[5].SatisfiesFilter(), Equals, false)
}
