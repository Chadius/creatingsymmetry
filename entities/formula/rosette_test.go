package formula_test

import (
	. "gopkg.in/check.v1"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/coefficient"
	"wallpaper/entities/formula/exponential"
	"wallpaper/entities/utility"
)

type RosetteFormulaTest struct {}

var _ = Suite(&RosetteFormulaTest{})

func (suite *RosetteFormulaTest) SetUpTest(checker *C) {
}

func (suite *RosetteFormulaTest) TestCalculateRosetteFormula(checker *C) {
	rosetteFormula := formula.RosetteFormula{
		Terms: []*exponential.Term{
			{
				Multiplier:             complex(3, 0),
				PowerN:                 1,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{
					coefficient.PlusMPlusN,
				},
			},
		},
	}
	result := rosetteFormula.Calculate(complex(2,1))
	total := result.Total
	checker.Assert(real(total), utility.NumericallyCloseEnough{}, 12, 1e-6)
	checker.Assert(imag(total), utility.NumericallyCloseEnough{}, 0, 1e-6)
}

func (suite *RosetteFormulaTest) TestMultifoldSymmetry1Term(checker *C) {
	rosetteFormula := formula.RosetteFormula{
		Terms: []*exponential.Term{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 6,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{
					coefficient.PlusMPlusN,
				},
			},
		},
	}
	symmetriesDetected := rosetteFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.Multifold, Equals, 6)
}

func (suite *RosetteFormulaTest) TestMultifoldSymmetryIsAlwaysPositive(checker *C) {
	rosetteFormula := formula.RosetteFormula{
		Terms: []*exponential.Term{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 -6,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{
					coefficient.PlusMPlusN,
				},
			},
		},
	}
	symmetriesDetected := rosetteFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.Multifold, Equals, 6)
}

func (suite *RosetteFormulaTest) TestSymmetryUsesGreatestCommonDenominator(checker *C) {
	rosetteFormula := formula.RosetteFormula{
		Terms: []*exponential.Term{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 -6,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{
					coefficient.PlusMPlusN,
				},
			},
			{
				Multiplier:             complex(1, 0),
				PowerN:                 -8,
				PowerM:                 4,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{
					coefficient.PlusMPlusN,
				},
			},
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{
						coefficient.PlusMPlusN,
				},
			},
		},
	}
	symmetriesDetected := rosetteFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.Multifold, Equals, 2)
}

func (suite *RosetteFormulaTest) TestGetContributionOfRosetteTerm(checker *C) {
	rosetteFormula := formula.RosetteFormula{
		Terms: []*exponential.Term{
			{
				Multiplier:             complex(3, 0),
				PowerN:                 1,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{
					coefficient.PlusMPlusN,
				},
			},
		},
	}
	result := rosetteFormula.Calculate(complex(2,1))
	checker.Assert(result.ContributionByTerm, HasLen, 1)
	contributionByFirstTerm := result.ContributionByTerm[0]
	checker.Assert(real(contributionByFirstTerm), utility.NumericallyCloseEnough{}, 12, 1e-6)
	checker.Assert(imag(contributionByFirstTerm), utility.NumericallyCloseEnough{}, 0, 1e-6)
}

func (suite *RosetteFormulaTest) TestRosetteFormulaFromYAML(checker *C) {
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
	rosetteFormula, err := formula.NewRosetteFormulaFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(rosetteFormula.Terms, HasLen, 2)
	checker.Assert(rosetteFormula.Terms[0].PowerN, Equals, 3)
	checker.Assert(rosetteFormula.Terms[0].IgnoreComplexConjugate, Equals, false)
	checker.Assert(rosetteFormula.Terms[1].CoefficientRelationships[0], Equals, coefficient.Relationship(coefficient.MinusMMinusNMaybeFlipScale))
}

func (suite *RosetteFormulaTest) TestRosetteFormulaFromJSON(checker *C) {
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
	rosetteFormula, err := formula.NewRosetteFormulaFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(rosetteFormula.Terms, HasLen, 2)
	checker.Assert(rosetteFormula.Terms[0].PowerN, Equals, 3)
	checker.Assert(rosetteFormula.Terms[0].IgnoreComplexConjugate, Equals, false)
	checker.Assert(rosetteFormula.Terms[1].CoefficientRelationships[0], Equals, coefficient.Relationship(coefficient.MinusMMinusNMaybeFlipScale))
}
