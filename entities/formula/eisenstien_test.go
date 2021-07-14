package formula_test

import (
	. "gopkg.in/check.v1"
	"testing"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/coefficient"
	"wallpaper/entities/utility"
)

func Test(t *testing.T) { TestingT(t) }

type EisensteinFormulaSuite struct {}

var _ = Suite(&EisensteinFormulaSuite{})

func (suite *EisensteinFormulaSuite) SetUpTest(checker *C) {
}

func (suite *EisensteinFormulaSuite) TestCalculateEisensteinTermForGivenPoint(checker *C) {
	squareLatticePair := formula.LatticeVectorPair{
		XLatticeVector: complex(1, 0),
		YLatticeVector: complex(0, 1),
	}

	latticeCoordinate := squareLatticePair.ConvertToLatticeCoordinates(complex(1.0,1.5))

	eisenstein := formula.EisensteinFormulaTerm{
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
	term, err := formula.NewEisensteinFormulaTermFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(term.PowerN, Equals, 12)
	checker.Assert(term.PowerM, Equals, -10)
}

func (suite *EisensteinFormulaSuite) TestCreateFormulaWithYAML(checker *C) {
	yamlByteStream := []byte(`
power_n: 12
power_m: -10
`)
	term, err := formula.NewEisensteinFormulaTermFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(term.PowerN, Equals, 12)
	checker.Assert(term.PowerM, Equals, -10)
}

type EisensteinRelationshipTest struct {
	aPlusNPlusMOddTerm *formula.EisensteinFormulaTerm
	aPlusMMinusNOddTerm *formula.EisensteinFormulaTerm
	aPlusMPlusNOddTerm *formula.EisensteinFormulaTerm
	aMinusNMinusMOddTerm *formula.EisensteinFormulaTerm
	aMinusMMinusNOddTerm *formula.EisensteinFormulaTerm
	aMinusMPlusNOddTerm *formula.EisensteinFormulaTerm
	aPlusNPlusMEvenTerm *formula.EisensteinFormulaTerm
	aPlusMPlusNEvenTerm *formula.EisensteinFormulaTerm
	aMinusMMinusNEvenTerm *formula.EisensteinFormulaTerm
	aPlusMMinusSumNAndMOddTerm *formula.EisensteinFormulaTerm
	aMinusSumNAndMPlusN *formula.EisensteinFormulaTerm
}

var _ = Suite(&EisensteinRelationshipTest{})

func (suite *EisensteinRelationshipTest) SetUpTest(checker *C) {
	suite.aPlusNPlusMOddTerm = &formula.EisensteinFormulaTerm{
		PowerN:     -1,
		PowerM:     2,
	}
	suite.aPlusMPlusNOddTerm = &formula.EisensteinFormulaTerm{
		PowerN:     2,
		PowerM:     -1,
	}
	suite.aMinusNMinusMOddTerm = &formula.EisensteinFormulaTerm{
		PowerN:     1,
		PowerM:     -2,
	}
	suite.aMinusMMinusNOddTerm = &formula.EisensteinFormulaTerm{
		PowerN:     -2,
		PowerM:     1,
	}
	suite.aPlusMMinusNOddTerm = &formula.EisensteinFormulaTerm{
		PowerN:     2,
		PowerM:     1,
	}
	suite.aMinusMPlusNOddTerm = &formula.EisensteinFormulaTerm{
		PowerN:     -2,
		PowerM:     -1,
	}
	suite.aMinusSumNAndMPlusN = &formula.EisensteinFormulaTerm{
		PowerN:     -1,
		PowerM:     -1,
	}
	suite.aPlusMMinusSumNAndMOddTerm = &formula.EisensteinFormulaTerm{
		PowerN:     2,
		PowerM:     -1,
	}
	suite.aPlusNPlusMEvenTerm = &formula.EisensteinFormulaTerm{
		PowerN:     -6,
		PowerM:     2,
	}
	suite.aPlusMPlusNEvenTerm = &formula.EisensteinFormulaTerm{
		PowerN:     2,
		PowerM:     -6,
	}
	suite.aMinusMMinusNEvenTerm = &formula.EisensteinFormulaTerm{
		PowerN:     -2,
		PowerM:     6,
	}
}

func (suite *EisensteinRelationshipTest) TestPlusMPlusN(checker *C) {
	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMPlusN,
	), Equals, true)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aMinusNMinusMOddTerm,
		suite.aMinusMMinusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMPlusN,
	), Equals, true)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aMinusMMinusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMPlusN,
	), Equals, false)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(-2, -1),
		coefficient.PlusMPlusN,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestMultipliersMustBeTheSameOrNegated(checker *C) {
	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		&formula.EisensteinFormulaTerm{
			PowerN: suite.aPlusMPlusNOddTerm.PowerN,
			PowerM: suite.aPlusMPlusNOddTerm.PowerM,
		},
		complex(2, 1),
		complex(-4, -2),
		coefficient.PlusMPlusN,
	), Equals, false)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		&formula.EisensteinFormulaTerm{
			PowerN: suite.aPlusMPlusNOddTerm.PowerN,
			PowerM: suite.aPlusMPlusNOddTerm.PowerM,
		},
		complex(2, 1),
		complex(4, 2),
		coefficient.PlusMPlusN,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestMinusNMinusM(checker *C) {
	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aMinusNMinusMOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusNMinusM,
	), Equals, true)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusNMinusM,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestPlusNPlusM(checker *C) {
	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusNPlusMOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusNPlusM,
	), Equals, true)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aMinusMMinusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusNPlusM,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestMinusMMinusN(checker *C) {
	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aMinusMMinusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusMMinusN,
	), Equals, true)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusMMinusN,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestPlusMPlusNMaybeFlipScale(checker *C) {
	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(-2, -1),
		coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum,
	), Equals, true)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum,
	), Equals, false)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMEvenTerm,
		suite.aPlusMPlusNEvenTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum,
	), Equals, true)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMEvenTerm,
		suite.aPlusMPlusNEvenTerm,
		complex(2, 1),
		complex(-2, -1),
		coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum,
	), Equals, false)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestMinusMMinusNMaybeFlipScale(checker *C) {
	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aMinusMMinusNOddTerm,
		complex(2, 1),
		complex(-2, -1),
		coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum,
	), Equals, true)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aMinusMMinusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum,
	), Equals, false)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMEvenTerm,
		suite.aMinusMMinusNEvenTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum,
	), Equals, true)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMEvenTerm,
		suite.aMinusMMinusNEvenTerm,
		complex(2, 1),
		complex(-2, -1),
		coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum,
	), Equals, false)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestPlusMMinusSumNAndM(checker *C) {
	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMMinusSumNAndMOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMMinusSumNAndM,
	), Equals, true)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusNPlusMOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMMinusSumNAndM,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestMinusSumNAndMPlusN(checker *C) {
	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aMinusSumNAndMPlusN,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusSumNAndMPlusN,
	), Equals, true)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusNPlusMOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusSumNAndMPlusN,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestPlusMMinusN(checker *C) {
	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMMinusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMMinusN,
	), Equals, true)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.PlusMMinusN,
	), Equals, false)
}

func (suite *EisensteinRelationshipTest) TestMinusMPlusN(checker *C) {
	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aMinusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusMPlusN,
	), Equals, true)

	checker.Assert(formula.SatisfiesRelationship(
		suite.aPlusNPlusMOddTerm,
		suite.aPlusMPlusNOddTerm,
		complex(2, 1),
		complex(2, 1),
		coefficient.MinusMPlusN,
	), Equals, false)
}
