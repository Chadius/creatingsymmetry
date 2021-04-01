package formula_test

import (
	. "gopkg.in/check.v1"
	"math"
	"wallpaper/entities/formula"
)

type FriezeFormulaSuite struct {
}

var _ = Suite(&FriezeFormulaSuite{})

func (suite *FriezeFormulaSuite) SetUpTest(checker *C) {
}

func (suite *FriezeFormulaSuite) TestEulerFormulaCalculation(checker *C) {
	form := formula.EulerFormulaTerm{
		Multiplier:             complex(3, 0),
		PowerN:                 2,
		PowerM:                 0,
		IgnoreComplexConjugate: true,
	}
	result := form.Calculate(complex(math.Pi / 6.0,1))
	checker.Assert(math.Abs(real(result) - 3 * math.Exp(-2) * 1.0 / 2.0) < 0.01, Equals, true)
	checker.Assert(math.Abs(imag(result) - 3 * math.Exp(-2) * math.Sqrt(3.0) / 2.0) < 0.01, Equals, true)
}

func (suite *FriezeFormulaSuite) TestLockedCoefficientPairs(checker *C) {
	form := formula.EulerFormulaTerm{
		Multiplier:             complex(3, 0),
		PowerN:                 2,
		PowerM:                 0,
		IgnoreComplexConjugate: true,
		CoefficientPairs: formula.LockedCoefficientPair{
			Multiplier: 1,
			OtherCoefficientRelationships: []formula.CoefficientRelationship{
				formula.PlusMPlusN,
			},
		},
	}
	result := form.Calculate(complex(math.Pi / 6.0,1))
	checker.Assert(math.Abs(real(result) - 3 * ((math.Exp(-2) * 1.0 / 2.0) + 1.0)) < 0.01, Equals, true)
	checker.Assert(math.Abs(imag(result) - 3 * math.Exp(-2) * math.Sqrt(3.0) / 2.0) < 0.01, Equals, true)
}

func (suite *FriezeFormulaSuite) TestUseComplexConjugate(checker *C) {
	form := formula.EulerFormulaTerm{
		Multiplier:             complex(3, 0),
		PowerN:                 2,
		PowerM:                 1,
		IgnoreComplexConjugate: false,
	}
	result := form.Calculate(complex(math.Pi / 6.0,2))
	checker.Assert(real(result), Equals, 3 * math.Exp(-6) * math.Sqrt(3.0) / 2.0)
	checker.Assert(imag(result), Equals, 3 * math.Exp(-6) * 1.0 / 2.0)
}

func (suite *FriezeFormulaSuite) TestFriezeFormula(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*formula.EulerFormulaTerm{
			{
				Multiplier:             complex(2, 0),
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
	result := friezeFormula.Calculate(complex(math.Pi/6, 1))
	total := result.Total

	expectedResult := complex(math.Exp(-1), 0) * complex(math.Sqrt(3) * 2, 0)
	checker.Assert(real(total), Equals, real(expectedResult))
	checker.Assert(imag(total), Equals, imag(expectedResult))
}

func (suite *FriezeFormulaSuite) TestP211Frieze(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*formula.EulerFormulaTerm{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: 1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.MinusNMinusM,
					},
				},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P211, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP1m1Frieze(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*formula.EulerFormulaTerm{
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
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P1m1, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP11mFrieze(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*formula.EulerFormulaTerm{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: 1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.MinusMMinusN,
					},
				},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P11m, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP11gFrieze(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*formula.EulerFormulaTerm{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 1,
				IgnoreComplexConjugate: false,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: 1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.MinusMMinusNMaybeFlipScale,
					},
				},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P11g, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP11mFriezeIfP11gHasEvenSumPowers (checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*formula.EulerFormulaTerm{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: 1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.MinusMMinusNMaybeFlipScale,
					},
				},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P11m, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP2mmFrieze(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*formula.EulerFormulaTerm{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: 1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.MinusNMinusM,
						formula.PlusMPlusN,
						formula.MinusMMinusN,
					},
				},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P2mm, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP2mgFrieze(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*formula.EulerFormulaTerm{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 -1,
				IgnoreComplexConjugate: false,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: 1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.MinusNMinusM,
						formula.PlusMPlusNMaybeFlipScale,
						formula.MinusMMinusNMaybeFlipScale,
					},
				},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P2mg, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP2mmFriezeEvenIfP2mgHasEvenSumPowers(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*formula.EulerFormulaTerm{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: 1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.MinusNMinusM,
						formula.PlusMPlusNMaybeFlipScale,
						formula.MinusMMinusNMaybeFlipScale,
					},
				},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P2mm, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP111Frieze(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*formula.EulerFormulaTerm{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientPairs:       formula.LockedCoefficientPair{},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P111, Equals, true)
}

func (suite *FriezeFormulaSuite) TestP111FriezeComplexConjugateIgnored(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*formula.EulerFormulaTerm{
			{
				Multiplier:             complex(1, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: true,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: 1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.MinusNMinusM,
					},
				},
			},
		},
	}
	symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
	checker.Assert(symmetriesDetected.P111, Equals, true)
	checker.Assert(symmetriesDetected.P211, Equals, false)
}

func (suite *FriezeFormulaSuite) TestContributionOfFriezeFormula(checker *C) {
	friezeFormula := formula.FriezeFormula{
		Terms: []*formula.EulerFormulaTerm{
			{
				Multiplier:             complex(2, 0),
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
	result := friezeFormula.Calculate(complex(math.Pi/6, 1))

	checker.Assert(result.ContributionByTerm, HasLen, 1)
	contributionByFirstTerm := result.ContributionByTerm[0]

	expectedResult := complex(math.Exp(-1), 0) * complex(math.Sqrt(3) * 2, 0)
	checker.Assert(real(contributionByFirstTerm), Equals, real(expectedResult))
	checker.Assert(imag(contributionByFirstTerm), Equals, imag(expectedResult))
}

func (suite *FriezeFormulaSuite) TestCreateEulerFormulaWithYAML(checker *C) {
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
	eulerExponentialFormulaTerm, err := formula.NewEulerFormulaTermFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(real(eulerExponentialFormulaTerm.Multiplier), Equals, -1.0)
	checker.Assert(imag(eulerExponentialFormulaTerm.Multiplier), Equals, 2e-2)
	checker.Assert(eulerExponentialFormulaTerm.PowerN, Equals, 12)
	checker.Assert(eulerExponentialFormulaTerm.PowerM, Equals, -10)
	checker.Assert(eulerExponentialFormulaTerm.IgnoreComplexConjugate, Equals, true)
	checker.Assert(eulerExponentialFormulaTerm.CoefficientPairs.Multiplier, Equals, 1.0)
	checker.Assert(eulerExponentialFormulaTerm.CoefficientPairs.OtherCoefficientRelationships, HasLen, 2)
	checker.Assert(eulerExponentialFormulaTerm.CoefficientPairs.OtherCoefficientRelationships[0], Equals, formula.CoefficientRelationship(formula.MinusMMinusN))
	checker.Assert(eulerExponentialFormulaTerm.CoefficientPairs.OtherCoefficientRelationships[1], Equals, formula.CoefficientRelationship(formula.PlusMPlusNMaybeFlipScale))
}

func (suite *FriezeFormulaSuite) TestCreateEulerFormulaWithJSON(checker *C) {
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
	eulerExponentialFormulaTerm, err := formula.NewEulerFormulaTermFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(real(eulerExponentialFormulaTerm.Multiplier), Equals, -1.0)
	checker.Assert(imag(eulerExponentialFormulaTerm.Multiplier), Equals, 2e-2)
	checker.Assert(eulerExponentialFormulaTerm.PowerN, Equals, 12)
	checker.Assert(eulerExponentialFormulaTerm.PowerM, Equals, -10)
	checker.Assert(eulerExponentialFormulaTerm.IgnoreComplexConjugate, Equals, true)
	checker.Assert(eulerExponentialFormulaTerm.CoefficientPairs.Multiplier, Equals, 1.0)
	checker.Assert(eulerExponentialFormulaTerm.CoefficientPairs.OtherCoefficientRelationships, HasLen, 2)
	checker.Assert(eulerExponentialFormulaTerm.CoefficientPairs.OtherCoefficientRelationships[0], Equals, formula.CoefficientRelationship(formula.MinusMMinusN))
	checker.Assert(eulerExponentialFormulaTerm.CoefficientPairs.OtherCoefficientRelationships[1], Equals, formula.CoefficientRelationship(formula.PlusMPlusNMaybeFlipScale))
}

func (suite *FriezeFormulaSuite) TestCreateFriezeFormulaWithYAML(checker *C) {
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
	rosetteFormula, err := formula.NewFriezeFormulaFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(rosetteFormula.Terms, HasLen, 2)
	checker.Assert(rosetteFormula.Terms[0].PowerN, Equals, 3)
	checker.Assert(rosetteFormula.Terms[0].IgnoreComplexConjugate, Equals, false)
	checker.Assert(rosetteFormula.Terms[1].CoefficientPairs.OtherCoefficientRelationships[0], Equals, formula.CoefficientRelationship(formula.MinusMMinusNMaybeFlipScale))
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
	rosetteFormula, err := formula.NewFriezeFormulaFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(rosetteFormula.Terms, HasLen, 2)
	checker.Assert(rosetteFormula.Terms[0].PowerN, Equals, 3)
	checker.Assert(rosetteFormula.Terms[0].IgnoreComplexConjugate, Equals, false)
	checker.Assert(rosetteFormula.Terms[1].CoefficientPairs.OtherCoefficientRelationships[0], Equals, formula.CoefficientRelationship(formula.MinusMMinusNMaybeFlipScale))
}
