package oldformula_test

import (
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	"github.com/Chadius/creating-symmetry/entities/oldformula"
	"github.com/Chadius/creating-symmetry/entities/oldformula/latticevector"
	"github.com/Chadius/creating-symmetry/entities/utility"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type EisensteinFormulaSuite struct{}

var _ = Suite(&EisensteinFormulaSuite{})

func (suite *EisensteinFormulaSuite) SetUpTest(checker *C) {
}

func (suite *EisensteinFormulaSuite) TestCalculateEisensteinTermForGivenPoint(checker *C) {
	squareLatticePair := latticevector.Pair{
		XLatticeVector: complex(1, 0),
		YLatticeVector: complex(0, 1),
	}

	latticeCoordinate := squareLatticePair.ConvertToLatticeCoordinates(complex(1.0, 1.5))

	eisenstein := oldformula.EisensteinFormulaTerm{
		PowerN: 1,
		PowerM: 1,
	}
	calculatedCoordinate := eisenstein.Calculate(latticeCoordinate)

	checker.Assert(real(calculatedCoordinate), utility.NumericallyCloseEnough{}, -1, 1e-6)
	checker.Assert(imag(calculatedCoordinate), utility.NumericallyCloseEnough{}, 0.0, 1e-6)
}

func (suite *EisensteinFormulaSuite) TestCreateFormulaWithJSON(checker *C) {
	jsonByteStream := []byte(`{
				"power_n": 12,
				"power_m": -10
			}`)
	term, err := oldformula.NewEisensteinFormulaTermFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(term.PowerN, Equals, 12)
	checker.Assert(term.PowerM, Equals, -10)
}

func (suite *EisensteinFormulaSuite) TestCreateFormulaWithYAML(checker *C) {
	yamlByteStream := []byte(`
power_n: 12
power_m: -10
`)
	term, err := oldformula.NewEisensteinFormulaTermFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(term.PowerN, Equals, 12)
	checker.Assert(term.PowerM, Equals, -10)
}

type EisensteinRelationshipTest struct {
	aPlusNPlusMOddTerm         *oldformula.EisensteinFormulaTerm
	aPlusMMinusNOddTerm        *oldformula.EisensteinFormulaTerm
	aPlusMPlusNOddTerm         *oldformula.EisensteinFormulaTerm
	aMinusNMinusMOddTerm       *oldformula.EisensteinFormulaTerm
	aMinusMMinusNOddTerm       *oldformula.EisensteinFormulaTerm
	aMinusMPlusNOddTerm        *oldformula.EisensteinFormulaTerm
	aPlusNPlusMEvenTerm        *oldformula.EisensteinFormulaTerm
	aPlusMPlusNEvenTerm        *oldformula.EisensteinFormulaTerm
	aMinusMMinusNEvenTerm      *oldformula.EisensteinFormulaTerm
	aPlusMMinusSumNAndMOddTerm *oldformula.EisensteinFormulaTerm
	aMinusSumNAndMPlusN        *oldformula.EisensteinFormulaTerm
}

var _ = Suite(&EisensteinRelationshipTest{})

func (suite *EisensteinRelationshipTest) SetUpTest(checker *C) {
	suite.aPlusNPlusMOddTerm = &oldformula.EisensteinFormulaTerm{
		PowerN: -1,
		PowerM: 2,
	}
	suite.aPlusMPlusNOddTerm = &oldformula.EisensteinFormulaTerm{
		PowerN: 2,
		PowerM: -1,
	}
	suite.aMinusNMinusMOddTerm = &oldformula.EisensteinFormulaTerm{
		PowerN: 1,
		PowerM: -2,
	}
	suite.aMinusMMinusNOddTerm = &oldformula.EisensteinFormulaTerm{
		PowerN: -2,
		PowerM: 1,
	}
	suite.aPlusMMinusNOddTerm = &oldformula.EisensteinFormulaTerm{
		PowerN: 2,
		PowerM: 1,
	}
	suite.aMinusMPlusNOddTerm = &oldformula.EisensteinFormulaTerm{
		PowerN: -2,
		PowerM: -1,
	}
	suite.aMinusSumNAndMPlusN = &oldformula.EisensteinFormulaTerm{
		PowerN: -1,
		PowerM: -1,
	}
	suite.aPlusMMinusSumNAndMOddTerm = &oldformula.EisensteinFormulaTerm{
		PowerN: 2,
		PowerM: -1,
	}
	suite.aPlusNPlusMEvenTerm = &oldformula.EisensteinFormulaTerm{
		PowerN: -6,
		PowerM: 2,
	}
	suite.aPlusMPlusNEvenTerm = &oldformula.EisensteinFormulaTerm{
		PowerN: 2,
		PowerM: -6,
	}
	suite.aMinusMMinusNEvenTerm = &oldformula.EisensteinFormulaTerm{
		PowerN: -2,
		PowerM: 6,
	}
}

func (suite *EisensteinRelationshipTest) TestPlusMPlusN(checker *C) {
	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMPlusN,
	), Equals, true)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aMinusNMinusMOddTerm,
		suite.aMinusMMinusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMPlusN,
	), Equals, true)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aMinusMMinusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMPlusN,
	), Equals, false)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(-2, -1),
		coefficient.PlusMPlusN,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestMultipliersMustBeTheSameOrNegated(checker *C) {
	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		&oldformula.EisensteinFormulaTerm{
			PowerN: suite.aPlusMPlusNOddTerm.PowerN,
			PowerM: suite.aPlusMPlusNOddTerm.PowerM,
		},
		complex(2, 1),
		complex(-4, -2),
		coefficient.PlusMPlusN,
	), Equals, false)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		&oldformula.EisensteinFormulaTerm{
			PowerN: suite.aPlusMPlusNOddTerm.PowerN,
			PowerM: suite.aPlusMPlusNOddTerm.PowerM,
		},
		complex(2, 1),
		complex(4, 2),
		coefficient.PlusMPlusN,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestMinusNMinusM(checker *C) {
	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aMinusNMinusMOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusNMinusM,
	), Equals, true)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusNMinusM,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestPlusNPlusM(checker *C) {
	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusNPlusMOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusNPlusM,
	), Equals, true)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aMinusMMinusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusNPlusM,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestMinusMMinusN(checker *C) {
	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aMinusMMinusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusMMinusN,
	), Equals, true)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusMMinusN,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestPlusMPlusNMaybeFlipScale(checker *C) {
	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(-2, -1),
		coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum,
	), Equals, true)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum,
	), Equals, false)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMEvenTerm,
		suite.aPlusMPlusNEvenTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum,
	), Equals, true)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMEvenTerm,
		suite.aPlusMPlusNEvenTerm,
		complex(2, 1),
		complex(-2, -1),
		coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum,
	), Equals, false)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestMinusMMinusNMaybeFlipScale(checker *C) {
	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aMinusMMinusNOddTerm,
		complex(2, 1),
		complex(-2, -1),
		coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum,
	), Equals, true)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aMinusMMinusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum,
	), Equals, false)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMEvenTerm,
		suite.aMinusMMinusNEvenTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum,
	), Equals, true)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMEvenTerm,
		suite.aMinusMMinusNEvenTerm,
		complex(2, 1),
		complex(-2, -1),
		coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum,
	), Equals, false)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestPlusMMinusSumNAndM(checker *C) {
	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMMinusSumNAndMOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMMinusSumNAndM,
	), Equals, true)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusNPlusMOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMMinusSumNAndM,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestMinusSumNAndMPlusN(checker *C) {
	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aMinusSumNAndMPlusN,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusSumNAndMPlusN,
	), Equals, true)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusNPlusMOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusSumNAndMPlusN,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestPlusMMinusN(checker *C) {
	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMMinusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMMinusN,
	), Equals, true)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMMinusN,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestMinusMPlusN(checker *C) {
	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aMinusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusMPlusN,
	), Equals, true)

	checker.Assert(oldformula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusMPlusN,
	), Equals, false)
}
