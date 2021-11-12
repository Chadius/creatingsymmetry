package formula_test

import (
	"github.com/Chadius/creating-symmetry/entities/formula"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type IdentityTest struct {
	identityFormula formula.Arbitrary
}

var _ = Suite(&IdentityTest{})

func (suite *IdentityTest) SetUpTest(checker *C) {
	suite.identityFormula = &formula.Identity{}
}

func (suite *IdentityTest) TestIdentityFormulaTransforms(checker *C) {
	basePoint := complex(-2e3, 5e-7)
	transformedPoint := suite.identityFormula.Calculate(basePoint)
	checker.Assert(transformedPoint, Equals, basePoint)
}
