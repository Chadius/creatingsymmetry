package utility_test

import (
	. "gopkg.in/check.v1"
	"testing"
	"github.com/Chadius/creating-symmetry/entities/utility"
)

func Test(t *testing.T) { TestingT(t) }

type CloseEnoughTests struct {
}

var _ = Suite(&CloseEnoughTests{})

func (suite *CloseEnoughTests) SetUpTest(checker *C) {}

func (suite *CloseEnoughTests) TestExactNumbersAreCloseEnough(checker *C) {
	checker.Assert(0.0, utility.NumericallyCloseEnough{}, 0.0, 0.1)
}

func (suite *CloseEnoughTests) TestCloseNumbersAreCloseEnough(checker *C) {
	checker.Assert(0.0, utility.NumericallyCloseEnough{}, 0.1, 0.2)
}

func (suite *CloseEnoughTests) TestVeryDifferentNumbersAreNotCloseEnough(checker *C) {
	numericallyCloseEnoughCheck := utility.NumericallyCloseEnough{}

	matcherArguments := []float64{0.0, 10.0, 0.1}
	interfaceOfFloats := make([]interface{}, len(matcherArguments))
	for index, flo := range matcherArguments {
		interfaceOfFloats[index] = flo
	}

	result, err := numericallyCloseEnoughCheck.Check( interfaceOfFloats, []string{"obtained", "expected", "tolerance"})

	checker.Assert(result, Equals, false)
	checker.Assert(err, Equals, "0.000000 and 10.000000 are outside of tolerance 0.100000")
}

func (suite *CloseEnoughTests) TestDifferentNumbersCanBeWithinTolerance(checker *C) {
	checker.Assert(0.0, utility.NumericallyCloseEnough{}, 10.0, 20.0)
}
