package formula_test

import (
	. "gopkg.in/check.v1"
	"wallpaper/entities/formula"
)

type RosetteFormulaTest struct {
}

var _ = Suite(&RosetteFormulaTest{})

func (suite *RosetteFormulaTest) SetUpTest(checker *C) {
}

func (suite *RosetteFormulaTest) TestCalculateRosetteFormula(checker *C) {
	rosetteFormula := formula.RosetteFormula{
		Terms: []*formula.ZExponentialFormulaTerm{
			{
				Multiplier:             complex(3, 0),
				PowerN:                 1,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: 1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.PlusMPlusN,
					},
				},
			},
		},
	}
	result := rosetteFormula.Calculate(complex(2,1))
	total := result.Total
	checker.Assert(real(total) - 12 < 0.01, Equals, true)
	checker.Assert(imag(total) - 0 < 0.01, Equals, true)
}

func (suite *RosetteFormulaTest) TestMultifoldSymmetry1Term(checker *C) {
	rosetteFormula := formula.RosetteFormula{
		Terms: []*formula.ZExponentialFormulaTerm{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 6,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: 1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.PlusMPlusN,
					},
				},
			},
		},
	}
	symmetriesDetected := rosetteFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.Multifold, Equals, 6)
}

func (suite *RosetteFormulaTest) TestMultifoldSymmetryIsAlwaysPositive(checker *C) {
	rosetteFormula := formula.RosetteFormula{
		Terms: []*formula.ZExponentialFormulaTerm{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 -6,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: 1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.PlusMPlusN,
					},
				},
			},
		},
	}
	symmetriesDetected := rosetteFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.Multifold, Equals, 6)
}

func (suite *RosetteFormulaTest) TestSymmetryUsesGreatestCommonDenominator(checker *C) {
	rosetteFormula := formula.RosetteFormula{
		Terms: []*formula.ZExponentialFormulaTerm{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 -6,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: 1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.PlusMPlusN,
					},
				},
			},
			{
				Multiplier:             complex(1, 0),
				PowerN:                 -8,
				PowerM:                 4,
				IgnoreComplexConjugate: false,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: 1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.PlusMPlusN,
					},
				},
			},
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: 1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.PlusMPlusN,
					},
				},
			},
		},
	}
	symmetriesDetected := rosetteFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.Multifold, Equals, 2)
}

func (suite *RosetteFormulaTest) TestGetContributionOfRosetteTerm(checker *C) {
	rosetteFormula := formula.RosetteFormula{
		Terms: []*formula.ZExponentialFormulaTerm{
			{
				Multiplier:             complex(3, 0),
				PowerN:                 1,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: 1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.PlusMPlusN,
					},
				},
			},
		},
	}
	result := rosetteFormula.Calculate(complex(2,1))
	checker.Assert(result.ContributionByTerm, HasLen, 1)
	contributionByFirstTerm := result.ContributionByTerm[0]
	checker.Assert(real(contributionByFirstTerm) - 12 < 0.01, Equals, true)
	checker.Assert(imag(contributionByFirstTerm) - 0 < 0.01, Equals, true)
}

func (suite *RosetteFormulaTest) TestZExponentialFormula(checker *C) {
	form := formula.ZExponentialFormulaTerm{
		Multiplier:             complex(3, 0),
		PowerN:                 2,
		PowerM:                 0,
		IgnoreComplexConjugate: true,
	}
	total := form.Calculate(complex(3,2))
	checker.Assert(real(total) - 15 < 0.01, Equals, true)
	checker.Assert(imag(total) - 36 < 0.01, Equals, true)
}

func (suite *RosetteFormulaTest) TestZExponentialFormulaWithLockedPairs(checker *C) {
	form := formula.ZExponentialFormulaTerm{
		Multiplier:             complex(3, 0),
		PowerN:                 2,
		PowerM:                 0,
		IgnoreComplexConjugate: true,
		CoefficientPairs: formula.LockedCoefficientPair{
			Multiplier: -1,
			OtherCoefficientRelationships: []formula.CoefficientRelationship{
				formula.PlusMPlusN,
			},
		},
	}
	total := form.Calculate(complex(3,2))
	checker.Assert(real(total) - 12 < 0.01, Equals, true)
	checker.Assert(imag(total) - 36 < 0.01, Equals, true)
}

func (suite *RosetteFormulaTest) TestZExponentialWithComplexConjugate(checker *C) {
	form := formula.ZExponentialFormulaTerm{
		Multiplier:             complex(3, 0),
		PowerN:                 2,
		PowerM:                 1,
		IgnoreComplexConjugate: false,
	}
	total := form.Calculate(complex(3,2))
	checker.Assert(real(total) - 117 < 0.01, Equals, true)
	checker.Assert(imag(total) - 78 < 0.01, Equals, true)
}

func (suite *RosetteFormulaTest) TestKeepCoefficientOrder(checker *C) {
	form := formula.ZExponentialFormulaTerm{
		Multiplier:             complex(1, 0),
		PowerN:                 1,
		PowerM:                 0,
		IgnoreComplexConjugate: true,
		CoefficientPairs: formula.LockedCoefficientPair{
			Multiplier: -1,
			OtherCoefficientRelationships: []formula.CoefficientRelationship{
				formula.PlusNPlusM,
			},
		},
	}
	total := form.Calculate(complex(3,2))
	checker.Assert(real(total) - 0 < 0.01, Equals, true)
	checker.Assert(imag(total) - 0 < 0.01, Equals, true)
}

func (suite *RosetteFormulaTest) TestSwapCoefficientOrder(checker *C) {
	form := formula.ZExponentialFormulaTerm{
		Multiplier:             complex(1, 0),
		PowerN:                 1,
		PowerM:                 0,
		IgnoreComplexConjugate: true,
		CoefficientPairs: formula.LockedCoefficientPair{
			Multiplier: -1,
			OtherCoefficientRelationships: []formula.CoefficientRelationship{
				formula.PlusMPlusN,
			},
		},
	}
	total := form.Calculate(complex(3,2))
	checker.Assert(real(total) - 2 < 0.01, Equals, true)
	checker.Assert(imag(total) - 2 < 0.01, Equals, true)
}

func (suite *RosetteFormulaTest) TestCreateZExponentialFromYAML(checker *C) {
	yamlByteStream := []byte(`
multiplier:
  real: -1.0
  imaginary: 2e-2
power_n: 12
power_m: -10
ignore_complex_conjugate: true
coefficient_pairs: 
  multiplier: 1
  relationships:
  - -M-N
  - +M+NF
`)

	zExponentialFormulaTerm, err := formula.NewZExponentialFormulaTermFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(real(zExponentialFormulaTerm.Multiplier) - -1.0 < 0.01, Equals, true)
	checker.Assert(imag(zExponentialFormulaTerm.Multiplier) - 2e-2 < 0.01, Equals, true)
	checker.Assert(zExponentialFormulaTerm.PowerN, Equals, 12)
	checker.Assert(zExponentialFormulaTerm.PowerM, Equals, -10)
	checker.Assert(zExponentialFormulaTerm.IgnoreComplexConjugate, Equals, true)
	checker.Assert(zExponentialFormulaTerm.CoefficientPairs.Multiplier - 1 < 0.01, Equals, true)
	checker.Assert(zExponentialFormulaTerm.CoefficientPairs.OtherCoefficientRelationships, HasLen, 2)
	checker.Assert(zExponentialFormulaTerm.CoefficientPairs.OtherCoefficientRelationships[0], Equals, formula.CoefficientRelationship(formula.MinusMMinusN))
	checker.Assert(zExponentialFormulaTerm.CoefficientPairs.OtherCoefficientRelationships[1], Equals, formula.CoefficientRelationship(formula.PlusMPlusNMaybeFlipScale))
}

func (suite *RosetteFormulaTest) TestCreateZExponentialFromJSON(checker *C) {
	jsonByteStream := []byte(`{
				"multiplier": {
					"real": -1.0,
					"imaginary": 2e-2
				},
				"power_n": 12,
				"power_m": -10,
				"ignore_complex_conjugate": true,
				"coefficient_pairs": {
				  "multiplier": 1,
				  "relationships": ["-M-N", "+M+NF"]
				}
			}`)
	zExponentialFormulaTerm, err := formula.NewZExponentialFormulaTermFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(real(zExponentialFormulaTerm.Multiplier) - -1.0 < 0.01, Equals, true)
	checker.Assert(imag(zExponentialFormulaTerm.Multiplier) - 2e-2 < 0.01, Equals, true)
	checker.Assert(zExponentialFormulaTerm.PowerN, Equals, 12)
	checker.Assert(zExponentialFormulaTerm.PowerM, Equals, -10)
	checker.Assert(zExponentialFormulaTerm.IgnoreComplexConjugate, Equals, true)
	checker.Assert(zExponentialFormulaTerm.CoefficientPairs.Multiplier - 1 < 0.01, Equals, true)
	checker.Assert(zExponentialFormulaTerm.CoefficientPairs.OtherCoefficientRelationships, HasLen, 2)
	checker.Assert(zExponentialFormulaTerm.CoefficientPairs.OtherCoefficientRelationships[0], Equals, formula.CoefficientRelationship(formula.MinusMMinusN))
	checker.Assert(zExponentialFormulaTerm.CoefficientPairs.OtherCoefficientRelationships[1], Equals, formula.CoefficientRelationship(formula.PlusMPlusNMaybeFlipScale))
}

func (suite *RosetteFormulaTest) TestRosetteFormulaFromYAML(checker *C) {
	yamlByteStream := []byte(`terms:
  -
    multiplier:
      real: -1.0
      imaginary: 2e-2
    power_n: 3
    power_m: 0
    coefficient_pairs: 
      multiplier: 1
      relationships:
      - -M-N
      - "+M+NF"
  -
    multiplier:
      real: 1e-10
      imaginary: 0
    power_n: 1
    power_m: 1
    coefficient_pairs:
      multiplier: 1
      relationships:
      - -M-NF
`)
	rosetteFormula, err := formula.NewRosetteFormulaFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(rosetteFormula.Terms, HasLen, 2)
	checker.Assert(rosetteFormula.Terms[0].PowerN, Equals, 3)
	checker.Assert(rosetteFormula.Terms[0].IgnoreComplexConjugate, Equals, false)
	checker.Assert(rosetteFormula.Terms[1].CoefficientPairs.OtherCoefficientRelationships[0], Equals, formula.CoefficientRelationship(formula.MinusMMinusNMaybeFlipScale))
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
						"coefficient_pairs": {
						  "multiplier": 1,
						  "relationships": ["-M-N", "+M+NF"]
						}
					},
					{
						"multiplier": {
							"real": 1e-10,
							"imaginary": 0
						},
						"power_n": 1,
						"power_m": 1,
						"coefficient_pairs": {
						  "multiplier": 1,
						  "relationships": ["-M-NF"]
						}
					}
				]
			}`)
	rosetteFormula, err := formula.NewRosetteFormulaFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(rosetteFormula.Terms, HasLen, 2)
	checker.Assert(rosetteFormula.Terms[0].PowerN, Equals, 3)
	checker.Assert(rosetteFormula.Terms[0].IgnoreComplexConjugate, Equals, false)
	checker.Assert(rosetteFormula.Terms[1].CoefficientPairs.OtherCoefficientRelationships[0], Equals, formula.CoefficientRelationship(formula.MinusMMinusNMaybeFlipScale))
}
