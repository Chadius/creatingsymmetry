package formula_test

import (
	. "gopkg.in/check.v1"
	"math"
	"testing"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/coefficient"
	"wallpaper/entities/formula/exponential"
	"wallpaper/entities/utility"
)

func Test(t *testing.T) { TestingT(t) }

type FriezeFormulaSuite struct {}

var _ = Suite(&FriezeFormulaSuite{})

func (suite *FriezeFormulaSuite) SetUpTest(checker *C) {
}

func (suite *FriezeFormulaSuite) TestFriezeFormula(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*exponential.Term{
			{
				Multiplier:             complex(2, 0),
				PowerN:                 1,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{coefficient.PlusMPlusN},
			},
		},
	}
	result := friezeFormula.Calculate(complex(math.Pi/6, 1))
	total := result.Total

	expectedResult := complex(math.Exp(-1), 0) * complex(math.Sqrt(3) * 2, 0)
	checker.Assert(real(total), utility.NumericallyCloseEnough{}, real(expectedResult), 1e-6)
	checker.Assert(imag(total), utility.NumericallyCloseEnough{}, imag(expectedResult), 1e-6)
}

func (suite *FriezeFormulaSuite) TestP211Frieze(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*exponential.Term{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{coefficient.MinusNMinusM},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P211, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP1m1Frieze(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*exponential.Term{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{coefficient.PlusMPlusN},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P1m1, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP11mFrieze(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*exponential.Term{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{coefficient.MinusMMinusN},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P11m, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP11gFrieze(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*exponential.Term{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 1,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{coefficient.MinusMMinusNMaybeFlipScale},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P11g, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP11mFriezeIfP11gHasEvenSumPowers (checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*exponential.Term{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{coefficient.MinusMMinusNMaybeFlipScale},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P11m, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP2mmFrieze(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*exponential.Term{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{
					coefficient.MinusNMinusM,
					coefficient.PlusMPlusN,
					coefficient.MinusMMinusN,
				},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P2mm, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP2mgFrieze(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*exponential.Term{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 -1,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{
					coefficient.MinusNMinusM,
					coefficient.PlusMPlusNMaybeFlipScale,
					coefficient.MinusMMinusNMaybeFlipScale,
				},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P2mg, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP2mmFriezeEvenIfP2mgHasEvenSumPowers(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*exponential.Term{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{
					coefficient.MinusNMinusM,
					coefficient.PlusMPlusNMaybeFlipScale,
					coefficient.MinusMMinusNMaybeFlipScale,
				},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P2mm, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP111Frieze(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*exponential.Term{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P111, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP111FriezeComplexConjugateIgnored(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*exponential.Term{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: true,
				CoefficientRelationships: []coefficient.Relationship{coefficient.MinusNMinusM},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P111, Equals, true)
	checker.Assert(symmetriesDetected.P211, Equals, false)
}

func (suite *FriezeFormulaSuite) TestContributionOfFriezeFormula(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*exponential.Term{
			{
				Multiplier:             complex(2, 0),
				PowerN:                 1,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{
					coefficient.PlusMPlusN,
				},
			},
		},
	}
	result := friezeFormula.Calculate(complex(math.Pi/6, 1))

	checker.Assert(result.ContributionByTerm, HasLen, 1)
	contributionByFirstTerm := result.ContributionByTerm[0]

	expectedResult := complex(math.Exp(-1), 0) * complex(math.Sqrt(3) * 2, 0)
	checker.Assert(real(contributionByFirstTerm), utility.NumericallyCloseEnough{}, real(expectedResult), 1e-6)
	checker.Assert(imag(contributionByFirstTerm), utility.NumericallyCloseEnough{}, imag(expectedResult), 1e-6)
}

func (suite *FriezeFormulaSuite) TestCreateFriezeFormulaWithYAML(checker *C) {
	yamlByteStream := []byte(`terms:
  -
    multiplier:
      real: -1.0
      imaginary: 2e-2
    power_n: 3
    power_m: 0
    coefficient_relationships:
      - -M-N
      - "+M+NF"
  -
    multiplier:
      real: 1e-10
      imaginary: 0
    power_n: 1
    power_m: 1
    coefficient_relationships:
      - -M-NF
`)
	rosetteFormula, err := formula.NewFriezeFormulaFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(rosetteFormula.Terms, HasLen, 2)
	checker.Assert(rosetteFormula.Terms[0].PowerN, Equals, 3)
	checker.Assert(rosetteFormula.Terms[0].IgnoreComplexConjugate, Equals, false)
	checker.Assert(rosetteFormula.Terms[1].CoefficientRelationships[0], Equals, coefficient.Relationship(coefficient.MinusMMinusNMaybeFlipScale))
}

func (suite *FriezeFormulaSuite) TestCreateFriezeFormulaWithJSON(checker *C) {
	jsonByteStream := []byte(`{
				"terms": [
					{
						"multiplier": {
							"real": -1.0,
							"imaginary": 2e-2
						},
						"power_n": 3,
						"power_m": 0,
						"coefficient_relationships": ["-M-N", "+M+NF"]
					},
					{
						"multiplier": {
							"real": 1e-10,
							"imaginary": 0
						},
						"power_n": 1,
						"power_m": 1,
						"coefficient_relationships": ["-M-NF"]
					}
				]
			}`)
	rosetteFormula, err := formula.NewFriezeFormulaFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(rosetteFormula.Terms, HasLen, 2)
	checker.Assert(rosetteFormula.Terms[0].PowerN, Equals, 3)
	checker.Assert(rosetteFormula.Terms[0].IgnoreComplexConjugate, Equals, false)
	checker.Assert(rosetteFormula.Terms[1].CoefficientRelationships[0], Equals, coefficient.Relationship(coefficient.MinusMMinusNMaybeFlipScale))
}
