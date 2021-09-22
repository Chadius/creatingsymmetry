package mathutility_test

import (
	. "gopkg.in/check.v1"
	"math/cmplx"
	"testing"
	"github.com/Chadius/creating-symmetry/entities/mathutility"
)

func Test(t *testing.T) { TestingT(t) }

type BoundsTestSuite struct {
}

var _ = Suite(&BoundsTestSuite{})

func (suite *BoundsTestSuite) SetUpTest(checker *C) {
}

func (suite *BoundsTestSuite) TestBoundAgainstMinValueWhenLess(checker *C) {
	checker.Assert(mathutility.ScaleValueBetweenTwoRanges(
		-200,
		-100.0,
		100.0,
		0,
		200.0) - 0 < 0.01, Equals, true)
}

func (suite *BoundsTestSuite) TestBoundAgainstMinValueWhenEqual(checker *C) {
	checker.Assert(mathutility.ScaleValueBetweenTwoRanges(
		-100.0,
		-100.0,
		100.0,
		0,
		200.0) - 0 < 0.01, Equals, true)
}

func (suite *BoundsTestSuite) TestBoundAgainstMaxValueWhenGreater(checker *C) {
	checker.Assert(mathutility.ScaleValueBetweenTwoRanges(
		9001,
		-100.0,
		100.0,
		0,
		200.0) - 200 < 0.01, Equals, true)
}

func (suite *BoundsTestSuite) TestBoundAgainstMaxValueWhenEqual(checker *C) {
	checker.Assert(mathutility.ScaleValueBetweenTwoRanges(
		100,
		-100.0,
		100.0,
		0,
		200.0) - 200 < 0.01, Equals, true)
}

func (suite *BoundsTestSuite) TestScaleNewRange(checker *C) {
	checker.Assert(mathutility.ScaleValueBetweenTwoRanges(
		0,
		-100.0,
		100.0,
		0,
		200.0) - 100 < 0.01, Equals, true)
}

func (suite *BoundsTestSuite) TestBoundingBoxCalculation(checker *C) {
	lotsOfComplexNumbers := []complex128{
		complex(0, 0),
		complex(10, 0),
		complex(0, -100),
		complex(0, -100.2),
		complex(0, 0),
		complex(-10, 0),
		complex(0, 25),
		complex(-100, 0),
		complex(0, 25.5),
		complex(9000.1, 0),
	}

	min, max := mathutility.GetBoundingBox(lotsOfComplexNumbers)

	checker.Assert(real(min)- -100 < 0.01, Equals, true)
	checker.Assert(imag(min)- -100.2 < 0.01, Equals, true)
	checker.Assert(real(max)- 9000.1 < 0.01, Equals, true)
	checker.Assert(imag(max)- 25.5 < 0.01, Equals, true)
}

func (suite *BoundsTestSuite) TestBoundingBoxDefault(checker *C) {
	min, max := mathutility.GetBoundingBox([]complex128{})

	checker.Assert(real(min)- 0 < 0.01, Equals, true)
	checker.Assert(imag(min)- 0 < 0.01, Equals, true)
	checker.Assert(real(max)- 0 < 0.01, Equals, true)
	checker.Assert(imag(max)- 0 < 0.01, Equals, true)
}

func (suite *BoundsTestSuite) TestBoundingBoxIgnoresInfinity(checker *C) {
	lotsOfComplexNumbers := []complex128{
		complex(-10, -10),
		complex(10, 10),
		cmplx.Inf(),
		-1 * cmplx.Inf(),
	}
	min, max := mathutility.GetBoundingBox(lotsOfComplexNumbers)

	checker.Assert(real(min)- -10 < 0.01, Equals, true)
	checker.Assert(imag(min)- -10 < 0.01, Equals, true)
	checker.Assert(real(max)- 10 < 0.01, Equals, true)
	checker.Assert(imag(max)- 10 < 0.01, Equals, true)
}
