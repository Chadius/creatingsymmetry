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

func (suite *CoordinateFilterTests) TestFilterMarksMappedCoordinates(checker *C) {
	filter := imageoutput.CoordinateFilterBuilder().WithMinimumX(-1e-5).WithMaximumX(2e1).WithMinimumY(-6e6).WithMaximumY(5e2).Build()

	coordinateThatSatisfiesFilter := imageoutput.NewMappedCoordinate(1e1, 2e2)
	coordinateThatDoesNotSatisfyFilterBecauseItIsOutsideInXDirection := imageoutput.NewMappedCoordinate(-1e1, 2e2)
	coordinateThatDoesNotSatisfyFilterBecauseItIsOutsideInYDirection := imageoutput.NewMappedCoordinate(1e1, 2e5)
	coordinateThatDoesNotSatisfyFilterBecauseItIsAtInfinity := imageoutput.NewMappedCoordinate(0, math.Inf(1))

	filter.FilterAndMarkMappedCoordinate(coordinateThatSatisfiesFilter)
	filter.FilterAndMarkMappedCoordinate(coordinateThatDoesNotSatisfyFilterBecauseItIsOutsideInXDirection)
	filter.FilterAndMarkMappedCoordinate(coordinateThatDoesNotSatisfyFilterBecauseItIsOutsideInYDirection)
	filter.FilterAndMarkMappedCoordinate(coordinateThatDoesNotSatisfyFilterBecauseItIsAtInfinity)

	checker.Assert(coordinateThatSatisfiesFilter.SatisfiesFilter(), Equals, true)
	checker.Assert(coordinateThatDoesNotSatisfyFilterBecauseItIsOutsideInXDirection.SatisfiesFilter(), Equals, false)
	checker.Assert(coordinateThatDoesNotSatisfyFilterBecauseItIsOutsideInYDirection.SatisfiesFilter(), Equals, false)
	checker.Assert(coordinateThatDoesNotSatisfyFilterBecauseItIsAtInfinity.SatisfiesFilter(), Equals, false)
}

func (suite *CoordinateFilterTests) TestFilterMarksCoordinateCollection(checker *C) {
	coordinates := []*imageoutput.MappedCoordinate{
		imageoutput.NewMappedCoordinate(-10, 20),
		imageoutput.NewMappedCoordinate(20, 0),
		imageoutput.NewMappedCoordinate(0, -100),
		imageoutput.NewMappedCoordinate(0, 200),
	}
	collection := imageoutput.CoordinateCollectionBuilder().WithCoordinates(&coordinates).Build()

	filter := imageoutput.CoordinateFilterBuilder().WithMinimumX(0).WithMaximumX(25).WithMinimumY(-6e6).WithMaximumY(50).Build()
	filter.FilterAndMarkMappedCoordinateCollection(collection)

	checker.Assert((*collection.Coordinates())[0].SatisfiesFilter(), Equals, false)
	checker.Assert((*collection.Coordinates())[1].SatisfiesFilter(), Equals, true)
	checker.Assert((*collection.Coordinates())[2].SatisfiesFilter(), Equals, true)
	checker.Assert((*collection.Coordinates())[3].SatisfiesFilter(), Equals, false)
}
