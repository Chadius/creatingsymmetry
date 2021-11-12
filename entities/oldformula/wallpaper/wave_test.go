package wallpaper_test

import (
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	"github.com/Chadius/creating-symmetry/entities/oldformula"
	"github.com/Chadius/creating-symmetry/entities/oldformula/latticevector"
	"github.com/Chadius/creating-symmetry/entities/oldformula/wallpaper"
	"github.com/Chadius/creating-symmetry/entities/utility"
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"
)

type WaveFormulaTests struct {
	hexLatticeVectors   *latticevector.Pair
	hexagonalWavePacket *wallpaper.WavePacket
}

var _ = Suite(&WaveFormulaTests{})

func (suite *WaveFormulaTests) SetUpTest(checker *C) {
	suite.hexLatticeVectors = &latticevector.Pair{
		XLatticeVector: complex(1, 0),
		YLatticeVector: complex(-0.5, math.Sqrt(3.0)/2.0),
	}
	suite.hexagonalWavePacket = &wallpaper.WavePacket{

		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: 1,
				PowerM: -2,
			},
			{
				PowerN: -2,
				PowerM: 1,
			},
			{
				PowerN: 1,
				PowerM: 1,
			},
		},
		Multiplier: complex(1, 0),
	}
}

func (suite *WaveFormulaTests) TestWaveFormulaCombinesEisensteinTerms(checker *C) {
	zInLatticeCoordinates := suite.hexLatticeVectors.ConvertToLatticeCoordinates(complex(math.Sqrt(3), -1*math.Sqrt(3)))
	calculation := suite.hexagonalWavePacket.Calculate(zInLatticeCoordinates)
	total := calculation.Total

	expectedAnswer := cmplx.Exp(complex(0, 2*math.Pi*(3+math.Sqrt(3)))) +
		cmplx.Exp(complex(0, 2*math.Pi*(-2*math.Sqrt(3)))) +
		cmplx.Exp(complex(0, 2*math.Pi*(-3+math.Sqrt(3))))

	checker.Assert(real(total), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(total), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}

func (suite *WaveFormulaTests) TestWaveFormulaShowsContributionsPerTerm(checker *C) {
	zInLatticeCoordinates := suite.hexLatticeVectors.ConvertToLatticeCoordinates(complex(math.Sqrt(3), -1*math.Sqrt(3)))
	calculation := suite.hexagonalWavePacket.Calculate(zInLatticeCoordinates)

	checker.Assert(calculation.ContributionByTerm, HasLen, 3)

	contributionOfTerm1 := cmplx.Exp(complex(0, 2*math.Pi*(3+math.Sqrt(3))))
	checker.Assert(real(calculation.ContributionByTerm[0]), utility.NumericallyCloseEnough{}, real(contributionOfTerm1), 1e-6)
	checker.Assert(imag(calculation.ContributionByTerm[0]), utility.NumericallyCloseEnough{}, imag(contributionOfTerm1), 1e-6)

	contributionOfTerm2 := cmplx.Exp(complex(0, 2*math.Pi*(-2*math.Sqrt(3))))
	checker.Assert(real(calculation.ContributionByTerm[1]), utility.NumericallyCloseEnough{}, real(contributionOfTerm2), 1e-6)
	checker.Assert(imag(calculation.ContributionByTerm[1]), utility.NumericallyCloseEnough{}, imag(contributionOfTerm2), 1e-6)

	contributionOfTerm3 := cmplx.Exp(complex(0, 2*math.Pi*(-3+math.Sqrt(3))))
	checker.Assert(real(calculation.ContributionByTerm[2]), utility.NumericallyCloseEnough{}, real(contributionOfTerm3), 1e-6)
	checker.Assert(imag(calculation.ContributionByTerm[2]), utility.NumericallyCloseEnough{}, imag(contributionOfTerm3), 1e-6)
}

func (suite *WaveFormulaTests) TestWaveFormulaUsesMultiplier(checker *C) {
	suite.hexagonalWavePacket.Multiplier = complex(1/3.0, 0)
	zInLatticeCoordinates := suite.hexLatticeVectors.ConvertToLatticeCoordinates(complex(math.Sqrt(3), -1*math.Sqrt(3)))
	calculation := suite.hexagonalWavePacket.Calculate(zInLatticeCoordinates)
	total := calculation.Total

	expectedAnswer := (cmplx.Exp(complex(0, 2*math.Pi*(3+math.Sqrt(3)))) +
		cmplx.Exp(complex(0, 2*math.Pi*(-2*math.Sqrt(3)))) +
		cmplx.Exp(complex(0, 2*math.Pi*(-3+math.Sqrt(3))))) / 3

	checker.Assert(real(total), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(total), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}

func (suite *WaveFormulaTests) TestWaveFormulaMarshalFromJson(checker *C) {
	jsonByteStream := []byte(`{
				"multiplier": {
					"real": -1.0,
					"imaginary": 2e-2
				},
				"terms": [
					{
						"power_n": 12,
						"power_m": -10,
						"multiplier": {
							"real": -1.0,
							"imaginary": 2e-2
						}
					}
				]
	}`)
	wave, err := wallpaper.NewWaveFormulaFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(real(wave.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(wave.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
	checker.Assert(wave.Terms, HasLen, 1)
	checker.Assert(wave.Terms[0].PowerN, Equals, 12)
}

func (suite *WaveFormulaTests) TestWaveFormulaMarshalFromYAML(checker *C) {
	yamlByteStream := []byte(`
multiplier:
  real: -1.0
  imaginary: 2e-2
terms:
  -
    power_n: 12
    power_m: -10
    multiplier:
      real: -1.0
      imaginary: 2e-2
`)
	wave, err := wallpaper.NewWaveFormulaFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(real(wave.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(wave.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
	checker.Assert(wave.Terms, HasLen, 1)
	checker.Assert(wave.Terms[0].PowerN, Equals, 12)
}

type WavePacketRelationshipTest struct {
	aPlusNPlusMOddWavePacket         *wallpaper.WavePacket
	aPlusMMinusNOddWavePacket        *wallpaper.WavePacket
	aPlusMPlusNOddWavePacket         *wallpaper.WavePacket
	aMinusNMinusMOddWavePacket       *wallpaper.WavePacket
	aMinusMMinusNOddWavePacket       *wallpaper.WavePacket
	aMinusMPlusNOddWavePacket        *wallpaper.WavePacket
	aPlusMPlusNOddNegatedWavePacket  *wallpaper.WavePacket
	aMinusSumNAndMPlusNOddWavePacket *wallpaper.WavePacket

	aPlusNPlusMEvenWavePacket          *wallpaper.WavePacket
	aPlusMPlusNEvenWavePacket          *wallpaper.WavePacket
	aPlusMPlusNEvenNegatedWavePacket   *wallpaper.WavePacket
	aMinusMMinusNOddNegatedWavePacket  *wallpaper.WavePacket
	aMinusMMinusNEvenWavePacket        *wallpaper.WavePacket
	aMinusMMinusNEvenNegatedWavePacket *wallpaper.WavePacket
	aPlusMMinusSumNAndMOddWavePacket   *wallpaper.WavePacket
	aMinusSumNAndMPlusNWavePacket      *wallpaper.WavePacket
}

var _ = Suite(&WavePacketRelationshipTest{})

func (suite *WavePacketRelationshipTest) SetUpTest(checker *C) {
	suite.aPlusNPlusMOddWavePacket = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: -1,
				PowerM: 4,
			},
		},
		Multiplier: complex(2, 1),
	}
	suite.aPlusMPlusNOddWavePacket = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: 4,
				PowerM: -1,
			},
		},
		Multiplier: complex(2, 1),
	}
	suite.aMinusNMinusMOddWavePacket = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: 1,
				PowerM: -4,
			},
		},
		Multiplier: complex(2, 1),
	}
	suite.aMinusMMinusNOddWavePacket = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: -4,
				PowerM: 1,
			},
		},
		Multiplier: complex(2, 1),
	}
	suite.aPlusMMinusNOddWavePacket = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: 4,
				PowerM: 1,
			},
		},
		Multiplier: complex(2, 1),
	}
	suite.aMinusMPlusNOddWavePacket = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: -4,
				PowerM: -1,
			},
		},
		Multiplier: complex(2, 1),
	}
	suite.aPlusMPlusNOddNegatedWavePacket = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: 4,
				PowerM: -1,
			},
		},
		Multiplier: complex(-2, -1),
	}
	suite.aMinusMMinusNOddNegatedWavePacket = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: -4,
				PowerM: 1,
			},
		},
		Multiplier: complex(-2, -1),
	}
	suite.aPlusMMinusSumNAndMOddWavePacket = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: 4,
				PowerM: -3,
			},
		},
		Multiplier: complex(2, 1),
	}
	suite.aMinusSumNAndMPlusNWavePacket = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: -3,
				PowerM: -1,
			},
		},
		Multiplier: complex(2, 1),
	}
	suite.aMinusSumNAndMPlusNOddWavePacket = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: -3,
				PowerM: -1,
			},
		},
		Multiplier: complex(2, 1),
	}

	suite.aPlusNPlusMEvenWavePacket = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: -6,
				PowerM: 2,
			},
		},
		Multiplier: complex(2, 1),
	}
	suite.aPlusMPlusNEvenWavePacket = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: 2,
				PowerM: -6,
			},
		},
		Multiplier: complex(2, 1),
	}
	suite.aPlusMPlusNEvenNegatedWavePacket = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: 2,
				PowerM: -6,
			},
		},
		Multiplier: complex(-2, -1),
	}
	suite.aMinusMMinusNEvenWavePacket = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: -2,
				PowerM: 6,
			},
		},
		Multiplier: complex(2, 1),
	}
	suite.aMinusMMinusNEvenNegatedWavePacket = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: -2,
				PowerM: 6,
			},
		},
		Multiplier: complex(-2, -1),
	}
}

func (suite *WavePacketRelationshipTest) TestLessThanTwoWavePacketsHasNoRelationship(checker *C) {
	checker.Assert(
		wallpaper.GetWavePacketRelationship(nil, nil),
		HasLen, 0)

	checker.Assert(
		wallpaper.GetWavePacketRelationship(
			suite.aPlusNPlusMOddWavePacket,
			nil,
		),
		HasLen, 0)
}

func (suite *WavePacketRelationshipTest) TestPlusNPlusM(checker *C) {
	relationshipsFound := wallpaper.GetWavePacketRelationship(
		suite.aPlusNPlusMOddWavePacket,
		suite.aPlusNPlusMOddWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(relationshipsFound[0], Equals, coefficient.PlusNPlusM)
}

func (suite *WavePacketRelationshipTest) TestMinusNMinusM(checker *C) {
	relationshipsFound := wallpaper.GetWavePacketRelationship(
		suite.aPlusNPlusMOddWavePacket,
		suite.aMinusNMinusMOddWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(relationshipsFound[0], Equals, coefficient.MinusNMinusM)
}

func (suite *WavePacketRelationshipTest) TestMinusMMinusN(checker *C) {
	relationshipsFound := wallpaper.GetWavePacketRelationship(
		suite.aPlusNPlusMOddWavePacket,
		suite.aMinusMMinusNOddWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(relationshipsFound[0], Equals, coefficient.MinusMMinusN)
}

func (suite *WavePacketRelationshipTest) TestPlusMPlusNMaybeFlipScale(checker *C) {
	relationshipsFound := wallpaper.GetWavePacketRelationship(
		suite.aPlusNPlusMOddWavePacket,
		suite.aPlusMPlusNOddNegatedWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(wallpaper.ContainsRelationship(relationshipsFound, coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum), Equals, true)

	relationshipsFound = wallpaper.GetWavePacketRelationship(
		suite.aPlusNPlusMEvenWavePacket,
		suite.aPlusMPlusNEvenNegatedWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 0)

	relationshipsFound = wallpaper.GetWavePacketRelationship(
		suite.aPlusNPlusMEvenWavePacket,
		suite.aPlusMPlusNEvenWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 2)
	checker.Assert(wallpaper.ContainsRelationship(relationshipsFound, coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum), Equals, true)
	checker.Assert(wallpaper.ContainsRelationship(relationshipsFound, coefficient.PlusMPlusN), Equals, true)
}

func (suite *WavePacketRelationshipTest) TestMinusMMinusNMaybeFlipScale(checker *C) {
	relationshipsFound := wallpaper.GetWavePacketRelationship(
		suite.aPlusNPlusMOddWavePacket,
		suite.aMinusMMinusNOddNegatedWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(relationshipsFound[0], Equals, coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum)

	relationshipsFound = wallpaper.GetWavePacketRelationship(
		suite.aPlusNPlusMEvenWavePacket,
		suite.aMinusMMinusNEvenWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 2)
	checker.Assert(wallpaper.ContainsRelationship(relationshipsFound, coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum), Equals, true)
	checker.Assert(wallpaper.ContainsRelationship(relationshipsFound, coefficient.MinusMMinusN), Equals, true)

	relationshipsFound = wallpaper.GetWavePacketRelationship(
		suite.aPlusNPlusMEvenWavePacket,
		suite.aMinusMMinusNEvenNegatedWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 0)
}

func (suite *WavePacketRelationshipTest) TestPlusMMinusSumNAndM(checker *C) {
	relationshipsFound := wallpaper.GetWavePacketRelationship(
		suite.aPlusNPlusMOddWavePacket,
		suite.aPlusMMinusSumNAndMOddWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(relationshipsFound[0], Equals, coefficient.PlusMMinusSumNAndM)
}

func (suite *WavePacketRelationshipTest) TestMinusSumNAndMPlusN(checker *C) {
	relationshipsFound := wallpaper.GetWavePacketRelationship(
		suite.aPlusNPlusMOddWavePacket,
		suite.aMinusSumNAndMPlusNOddWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(relationshipsFound[0], Equals, coefficient.MinusSumNAndMPlusN)
}

func (suite *WavePacketRelationshipTest) TestPlusMMinusN(checker *C) {
	relationshipsFound := wallpaper.GetWavePacketRelationship(
		suite.aPlusNPlusMOddWavePacket,
		suite.aPlusMMinusNOddWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(relationshipsFound[0], Equals, coefficient.PlusMMinusN)
}

func (suite *WavePacketRelationshipTest) TestMinusMPlusN(checker *C) {
	relationshipsFound := wallpaper.GetWavePacketRelationship(
		suite.aPlusNPlusMOddWavePacket,
		suite.aMinusMPlusNOddWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(relationshipsFound[0], Equals, coefficient.MinusMPlusN)
}
