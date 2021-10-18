package imageoutput_test

import (
	"github.com/Chadius/creating-symmetry/entities/imageoutput"
	. "gopkg.in/check.v1"
	"math"
)

type CoordinateCollectionFactoryTests struct {
}

var _ = Suite(&CoordinateCollectionFactoryTests{})

func (suite *CoordinateCollectionFactoryTests) TestSetupCreateDataRangeWithArray(checker *C) {
	coordinates := []complex128{
		complex(-10, 20),
		complex(20, 0),
		complex(0, -100),
		complex(0, 200),
	}
	collection := imageoutput.CoordinateCollectionFactory().WithCoordinates(&coordinates).Build()
	checker.Assert(collection.Coordinates(), Equals, &coordinates)
}

type CoordinateCollectionTests struct {
}

var _ = Suite(&CoordinateCollectionTests{})

func (suite *CoordinateCollectionTests) TestReturnsMinimumAndMaximums(checker *C) {
	coordinates := []complex128{
		complex(-10, 20),
		complex(20, 0),
		complex(0, -100),
		complex(0, 200),
	}
	collection := imageoutput.CoordinateCollectionFactory().WithCoordinates(&coordinates).Build()
	checker.Assert(collection.MinimumX(), Equals, float64(-10))
	checker.Assert(collection.MaximumX(), Equals, float64(20))
	checker.Assert(collection.MinimumY(), Equals, float64(-100))
	checker.Assert(collection.MaximumY(), Equals, float64(200))
}

func (suite *CoordinateCollectionTests) TestIfCoordinateHasInfinity(checker *C) {

}

func (suite *CoordinateCollectionTests) TestReturnsMinimumAndMaximumsIgnoringInfinity(checker *C) {
	coordinates := []complex128{
		complex(-1, 0),
		complex(1, 0),
		complex(0, -2),
		complex(0, 2),
		complex(-100, math.Inf(-1)),
		complex(100, math.Inf(1)),
		complex(math.Inf(1), -100),
		complex(math.Inf(-1), 100),
	}
	collection := imageoutput.CoordinateCollectionFactory().WithCoordinates(&coordinates).Build()
	checker.Assert(collection.MinimumX(), Equals, float64(-1))
	checker.Assert(collection.MaximumX(), Equals, float64(1))
	checker.Assert(collection.MinimumY(), Equals, float64(-2))
	checker.Assert(collection.MaximumY(), Equals, float64(2))
}
