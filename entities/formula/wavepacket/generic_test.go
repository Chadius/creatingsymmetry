package wavepacket_test

import (
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/wavepacket"
	"wallpaper/entities/utility"
)

type GenericWallpaper struct {
	GenericFormula *wavepacket.GenericWallpaperFormula
}

var _ = Suite(&GenericWallpaper{})

func (suite *GenericWallpaper) SetUpTest (checker *C) {
	suite.GenericFormula = &wavepacket.GenericWallpaperFormula{
		Formula: &wavepacket.WallpaperFormula{
			WavePackets: []*wavepacket.WavePacket{
				{
					Terms: []*formula.EisensteinFormulaTerm{
						{
							PowerN: 3,
							PowerM: -4,
						},
					},
					Multiplier: complex(1, 0),
				},
			},
			Multiplier: complex(1, 0),
		},
		VectorWidth: 2,
		VectorHeight: -0.5,
	}
	suite.GenericFormula.SetUp()
}

func (suite *GenericWallpaper) TestSetupCreatesLatticeVectors (checker *C) {
	checker.Assert(real(suite.GenericFormula.Formula.Lattice.XLatticeVector), utility.NumericallyCloseEnough{}, 1, 1e-6)
	checker.Assert(imag(suite.GenericFormula.Formula.Lattice.XLatticeVector), utility.NumericallyCloseEnough{}, 0, 1e-6)

	checker.Assert(real(suite.GenericFormula.Formula.Lattice.YLatticeVector), utility.NumericallyCloseEnough{}, 2, 1e-6)
	checker.Assert(imag(suite.GenericFormula.Formula.Lattice.YLatticeVector), utility.NumericallyCloseEnough{}, -0.5, 1e-6)
}

func (suite *GenericWallpaper) TestRaiseErrorIfHeightIsZero (checker *C) {
	GenericFormulaWithNoHeight := &wavepacket.GenericWallpaperFormula{
		Formula: &wavepacket.WallpaperFormula{
			WavePackets: []*wavepacket.WavePacket{
				{
					Terms: []*formula.EisensteinFormulaTerm{
						{
							PowerN: 1,
							PowerM: -2,
						},
					},
					Multiplier: complex(1, 0),
				},
			},
			Multiplier: complex(1, 0),
		},
		VectorWidth: -1,
		VectorHeight: 0,
	}
	err := GenericFormulaWithNoHeight.SetUp()

	checker.Assert(err, ErrorMatches, "vectors cannot be collinear: (.*,.*) and (.*,.*)")
}

func (suite *GenericWallpaper) TestSetupDoesNotAddLockedPairs (checker *C) {
	checker.Assert(suite.GenericFormula.Formula.WavePackets[0].Terms, HasLen, 1)
}

func (suite *GenericWallpaper) TestCalculationOfPoints (checker *C) {
	calculation := suite.GenericFormula.Calculate(complex(1.5, 10))
	total := calculation.Total

	expectedAnswer := cmplx.Exp(complex(0, math.Pi ))

	checker.Assert(real(total), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(total), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}

func (suite *GenericWallpaper) TestUnmarshalFromJSON(checker *C) {
	jsonByteStream := []byte(`{
				"vector_width": 0.8,
				"vector_height": 0.3,
				"formula": {
					"multiplier": {
						"real": -1.0,
						"imaginary": 2e-2
					},
					"wave_packets": [
						{
							"multiplier": {
								"real": -1.0,
								"imaginary": 2e-2
							},
							"terms": [
								{
									"power_n": 12,
									"power_m": -10
								}
							]
						}
					]
				}
			}`)
	GenericFormula, err := wavepacket.NewGenericWallpaperFormulaFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(GenericFormula.VectorWidth, utility.NumericallyCloseEnough{}, 0.8, 1e-6)
	checker.Assert(GenericFormula.VectorHeight, utility.NumericallyCloseEnough{}, 0.3, 1e-6)
	checker.Assert(real(GenericFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(GenericFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
	checker.Assert(GenericFormula.Formula.WavePackets, HasLen, 1)
	checker.Assert(GenericFormula.Formula.WavePackets[0].Terms[0].PowerN, Equals, 12)
	checker.Assert(GenericFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -10)
}

func (suite *GenericWallpaper) TestUnmarshalFromYAML(checker *C) {
	yamlByteStream := []byte(`
vector_width: 0.8
vector_height: 0.3
formula:
  multiplier:
    real: -1.0
    imaginary: 2e-2
  wave_packets:
    -
      multiplier:
        real: -1.0
        imaginary: 2e-2
      terms:
        -
          power_n: 12
          power_m: -10
`)
	GenericFormula, err := wavepacket.NewGenericWallpaperFormulaFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(GenericFormula.VectorWidth, utility.NumericallyCloseEnough{}, 0.8, 1e-6)
	checker.Assert(GenericFormula.VectorHeight, utility.NumericallyCloseEnough{}, 0.3, 1e-6)
	checker.Assert(real(GenericFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(GenericFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
	checker.Assert(GenericFormula.Formula.WavePackets, HasLen, 1)
	checker.Assert(GenericFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -10)
}

//type GenericWaveSymmetry struct {
//	baseWavePacketWithEvenPowerNAndOddPowerSum *wavepacket.WavePacket
//	baseWavePacketWithOddPowerNAndEvenPowerSum *wavepacket.WavePacket
//}
//
//var _ = Suite(&GenericWaveSymmetry{})
//
//func (suite *GenericWaveSymmetry) SetUpTest(checker *C) {
//	suite.baseWavePacketWithEvenPowerNAndOddPowerSum = &wavepacket.WavePacket{
//		Terms:[]*formula.EisensteinFormulaTerm{
//			{
//				PowerN:         8,
//				PowerM:         -3,
//			},
//		},
//		Multiplier: complex(1, 0),
//	}
//
//	suite.baseWavePacketWithOddPowerNAndEvenPowerSum = &wavepacket.WavePacket{
//		Terms:[]*formula.EisensteinFormulaTerm{
//			{
//				PowerN:         7,
//				PowerM:         -3,
//			},
//		},
//		Multiplier: complex(1, 0),
//	}
//}
//
//func (suite *GenericWaveSymmetry) TestNoSymmetryFound(checker *C) {
//	GenericFormula := wavepacket.GenericWallpaperFormula{
//		Formula: &wavepacket.WallpaperFormula{
//			WavePackets: []*wavepacket.WavePacket{
//				suite.baseWavePacketWithEvenPowerNAndOddPowerSum,
//			},
//			Multiplier: complex(1, 0),
//		},
//	}
//
//	GenericFormula.SetUp()
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pm), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pg), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmm), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmg), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pgg), Equals, false)
//}
//
//func (suite *GenericWaveSymmetry) TestPm(checker *C) {
//	GenericFormula := wavepacket.GenericWallpaperFormula{
//		Formula: &wavepacket.WallpaperFormula{
//			WavePackets: []*wavepacket.WavePacket{
//				suite.baseWavePacketWithEvenPowerNAndOddPowerSum,
//				{
//					Terms: []*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerN,
//							PowerM:         suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerM * -1,
//						},
//					},
//					Multiplier: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier,
//				},
//			},
//			Multiplier: complex(1, 0),
//		},
//	}
//
//	GenericFormula.SetUp()
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pm), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pg), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmm), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmg), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pgg), Equals, false)
//}
//
//func (suite *GenericWaveSymmetry) TestPg(checker *C) {
//	GenericFormulaWithEvenPowerN := wavepacket.GenericWallpaperFormula{
//		Formula: &wavepacket.WallpaperFormula{
//			WavePackets: []*wavepacket.WavePacket{
//				suite.baseWavePacketWithEvenPowerNAndOddPowerSum,
//				{
//					Terms: []*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerN,
//							PowerM:         suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerM * -1,
//						},
//					},
//					Multiplier: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier,
//				},
//			},
//			Multiplier: complex(1,0),
//		},
//	}
//
//	GenericFormulaWithEvenPowerN.SetUp()
//	checker.Assert(GenericFormulaWithEvenPowerN.HasSymmetry(wavepacket.Pm), Equals, true)
//	checker.Assert(GenericFormulaWithEvenPowerN.HasSymmetry(wavepacket.Pg), Equals, true)
//	checker.Assert(GenericFormulaWithEvenPowerN.HasSymmetry(wavepacket.Pmm), Equals, false)
//	checker.Assert(GenericFormulaWithEvenPowerN.HasSymmetry(wavepacket.Pmg), Equals, false)
//	checker.Assert(GenericFormulaWithEvenPowerN.HasSymmetry(wavepacket.Pgg), Equals, false)
//
//	GenericFormulaWithOddPowerN := wavepacket.GenericWallpaperFormula{
//		Formula: &wavepacket.WallpaperFormula{
//			WavePackets: []*wavepacket.WavePacket{
//				suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
//				{
//					Terms: []*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN,
//							PowerM:         suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM * -1,
//						},
//					},
//					Multiplier: 	suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier * -1,
//				},
//			},
//			Multiplier: 	complex(1, 0),
//		},
//	}
//
//	GenericFormulaWithOddPowerN.SetUp()
//	checker.Assert(GenericFormulaWithOddPowerN.HasSymmetry(wavepacket.Pm), Equals, false)
//	checker.Assert(GenericFormulaWithOddPowerN.HasSymmetry(wavepacket.Pg), Equals, true)
//	checker.Assert(GenericFormulaWithOddPowerN.HasSymmetry(wavepacket.Pmm), Equals, false)
//	checker.Assert(GenericFormulaWithOddPowerN.HasSymmetry(wavepacket.Pmg), Equals, false)
//	checker.Assert(GenericFormulaWithOddPowerN.HasSymmetry(wavepacket.Pgg), Equals, false)
//}
//
//func (suite *GenericWaveSymmetry) TestPmmAndPmgWithEvenPowerN(checker *C) {
//	GenericFormula := wavepacket.GenericWallpaperFormula{
//		Formula: &wavepacket.WallpaperFormula{
//			WavePackets: []*wavepacket.WavePacket{
//				suite.baseWavePacketWithEvenPowerNAndOddPowerSum,
//				{
//					Terms:[]*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerN * -1,
//							PowerM:         suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerM * -1,
//						},
//					},
//					Multiplier: complex(1, 0),
//				},
//				{
//					Terms:[]*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerN,
//							PowerM:         suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerM * -1,
//						},
//					},
//					Multiplier: complex(1, 0),
//				},
//				{
//					Terms:[]*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerN * -1,
//							PowerM:         suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerM,
//						},
//					},
//					Multiplier: complex(1, 0),
//				},
//			},
//			Multiplier: complex(1, 0),
//		},
//	}
//
//	GenericFormula.SetUp()
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pm), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pg), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmm), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmg), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pgg), Equals, false)
//}
//
//func (suite *GenericWaveSymmetry) TestPmgWithOddPowerN(checker *C) {
//	GenericFormulaWithOddPowerN := wavepacket.GenericWallpaperFormula{
//		Formula: &wavepacket.WallpaperFormula{
//			WavePackets: []*wavepacket.WavePacket{
//				suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
//				{
//					Terms:[]*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN * -1,
//							PowerM:         suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM * -1,
//						},
//					},
//					Multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier,
//				},
//				{
//					Terms:[]*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN,
//							PowerM:         suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM * -1,
//						},
//					},
//					Multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier * -1,
//				},
//				{
//					Terms:[]*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN * -1,
//							PowerM:         suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM,
//						},
//					},
//					Multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier * -1,
//				},
//			},
//			Multiplier: complex(1, 0),
//		},
//	}
//
//	GenericFormulaWithOddPowerN.SetUp()
//	checker.Assert(GenericFormulaWithOddPowerN.HasSymmetry(wavepacket.Pm), Equals, false)
//	checker.Assert(GenericFormulaWithOddPowerN.HasSymmetry(wavepacket.Pg), Equals, true)
//	checker.Assert(GenericFormulaWithOddPowerN.HasSymmetry(wavepacket.Pmm), Equals, false)
//	checker.Assert(GenericFormulaWithOddPowerN.HasSymmetry(wavepacket.Pmg), Equals, true)
//	checker.Assert(GenericFormulaWithOddPowerN.HasSymmetry(wavepacket.Pgg), Equals, false)
//}
//
//func (suite *GenericWaveSymmetry) TestPgg(checker *C) {
//	GenericFormulaWithOddPowerSum := wavepacket.GenericWallpaperFormula{
//		Formula: &wavepacket.WallpaperFormula{
//			WavePackets: []*wavepacket.WavePacket{
//				suite.baseWavePacketWithEvenPowerNAndOddPowerSum,
//				{
//					Terms:[]*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerN * -1,
//							PowerM:         suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerM * -1,
//						},
//					},
//					Multiplier: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier,
//				},
//				{
//					Terms:[]*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerN,
//							PowerM:         suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerM * -1,
//						},
//					},
//					Multiplier: -1 * suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier,
//				},
//				{
//					Terms:[]*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerN * -1,
//							PowerM:         suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerM,
//						},
//					},
//					Multiplier: -1 * suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier,
//				},
//			},
//			Multiplier: complex(1, 0),
//		},
//	}
//
//	GenericFormulaWithOddPowerSum.SetUp()
//	checker.Assert(GenericFormulaWithOddPowerSum.HasSymmetry(wavepacket.Pm), Equals, false)
//	checker.Assert(GenericFormulaWithOddPowerSum.HasSymmetry(wavepacket.Pg), Equals, false)
//	checker.Assert(GenericFormulaWithOddPowerSum.HasSymmetry(wavepacket.Pmm), Equals, false)
//	checker.Assert(GenericFormulaWithOddPowerSum.HasSymmetry(wavepacket.Pmg), Equals, false)
//	checker.Assert(GenericFormulaWithOddPowerSum.HasSymmetry(wavepacket.Pgg), Equals, true)
//
//	GenericFormulaWithEvenPowerSum := wavepacket.GenericWallpaperFormula{
//		Formula: &wavepacket.WallpaperFormula{
//			WavePackets: []*wavepacket.WavePacket{
//				suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
//				{
//					Terms:[]*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN * -1,
//							PowerM:         suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM * -1,
//						},
//					},
//					Multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier,
//				},
//				{
//					Terms:[]*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN,
//							PowerM:         suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM * -1,
//						},
//					},
//					Multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier,
//				},
//				{
//					Terms:[]*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN * -1,
//							PowerM:         suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM,
//						},
//					},
//					Multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier,
//				},
//			},
//			Multiplier: complex(1, 0),
//		},
//	}
//
//	GenericFormulaWithEvenPowerSum.SetUp()
//	checker.Assert(GenericFormulaWithEvenPowerSum.HasSymmetry(wavepacket.Pm), Equals, true)
//	checker.Assert(GenericFormulaWithEvenPowerSum.HasSymmetry(wavepacket.Pg), Equals, false)
//	checker.Assert(GenericFormulaWithEvenPowerSum.HasSymmetry(wavepacket.Pmm), Equals, true)
//	checker.Assert(GenericFormulaWithEvenPowerSum.HasSymmetry(wavepacket.Pmg), Equals, false)
//	checker.Assert(GenericFormulaWithEvenPowerSum.HasSymmetry(wavepacket.Pgg), Equals, true)
//}

//type GenericCreatedWithDesiredSymmetry struct {
//	eisensteinTermWithEvenPowerNAndOddPowerSum []*formula.EisensteinFormulaTerm
//	eisensteinTermWithOddPowerNAndEvenPowerSum []*formula.EisensteinFormulaTerm
//	wallpaperMultiplier complex128
//	LatticeHeight float64
//}
//
//var _ = Suite(&GenericCreatedWithDesiredSymmetry{})
//
//func (suite *GenericCreatedWithDesiredSymmetry) SetUpTest (checker *C) {
//	suite.eisensteinTermWithEvenPowerNAndOddPowerSum = []*formula.EisensteinFormulaTerm{
//		{
//			PowerN:         8,
//			PowerM:         -3,
//		},
//	}
//	suite.eisensteinTermWithOddPowerNAndEvenPowerSum = []*formula.EisensteinFormulaTerm{
//			{
//				PowerN:         7,
//				PowerM:         -3,
//			},
//		}
//
//	suite.wallpaperMultiplier = complex(1, 0)
//	suite.LatticeHeight = 2.0
//}
//
//func (suite *GenericCreatedWithDesiredSymmetry) TestCreateWallpaperWithPm(checker *C) {
//	GenericFormula, err := wavepacket.NewGenericWallpaperFormulaWithSymmetry(
//		suite.eisensteinTermWithOddPowerNAndEvenPowerSum,
//		suite.wallpaperMultiplier,
//		suite.LatticeHeight,
//		wavepacket.Pm,
//	)
//
//	checker.Assert(err, IsNil)
//	checker.Assert(GenericFormula.Formula.WavePackets, HasLen, 2)
//	checker.Assert(GenericFormula.Formula.WavePackets[0].Terms, HasLen, 1)
//
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerN)
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerM * -1)
//
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pm), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pg), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmm), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmg), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pgg), Equals, false)
//}
//
//func (suite *GenericCreatedWithDesiredSymmetry) TestCreateWallpaperWithPg(checker *C) {
//	GenericFormula, err := wavepacket.NewGenericWallpaperFormulaWithSymmetry(
//		suite.eisensteinTermWithOddPowerNAndEvenPowerSum,
//		suite.wallpaperMultiplier,
//		suite.LatticeHeight,
//		wavepacket.Pg,
//	)
//
//	checker.Assert(err, IsNil)
//	checker.Assert(GenericFormula.Formula.WavePackets, HasLen, 2)
//	checker.Assert(GenericFormula.Formula.WavePackets[0].Terms, HasLen, 1)
//
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Multiplier, Equals, suite.wallpaperMultiplier * -1)
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerN)
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerM * -1)
//
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pm), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pg), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmm), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmg), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pgg), Equals, false)
//}
//
//func (suite *GenericCreatedWithDesiredSymmetry) TestCreateWallpaperWithPmm(checker *C) {
//	GenericFormula, err := wavepacket.NewGenericWallpaperFormulaWithSymmetry(
//		suite.eisensteinTermWithOddPowerNAndEvenPowerSum,
//		suite.wallpaperMultiplier,
//		suite.LatticeHeight,
//		wavepacket.Pmm,
//	)
//
//	checker.Assert(err, IsNil)
//	checker.Assert(GenericFormula.Formula.WavePackets, HasLen, 4)
//	checker.Assert(GenericFormula.Formula.WavePackets[0].Terms, HasLen, 1)
//
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Multiplier, Equals, suite.wallpaperMultiplier)
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerN * -1)
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerM * -1)
//
//	checker.Assert(GenericFormula.Formula.WavePackets[2].Multiplier, Equals, suite.wallpaperMultiplier)
//	checker.Assert(GenericFormula.Formula.WavePackets[2].Terms[0].PowerN, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerN * -1)
//	checker.Assert(GenericFormula.Formula.WavePackets[2].Terms[0].PowerM, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerM)
//
//	checker.Assert(GenericFormula.Formula.WavePackets[3].Multiplier, Equals, suite.wallpaperMultiplier)
//	checker.Assert(GenericFormula.Formula.WavePackets[3].Terms[0].PowerN, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerN)
//	checker.Assert(GenericFormula.Formula.WavePackets[3].Terms[0].PowerM, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerM * -1)
//
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pm), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pg), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmm), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmg), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pgg), Equals, true)
//}
//
//func (suite *GenericCreatedWithDesiredSymmetry) TestCreateWallpaperWithPmg(checker *C) {
//	GenericFormula, err := wavepacket.NewGenericWallpaperFormulaWithSymmetry(
//		suite.eisensteinTermWithOddPowerNAndEvenPowerSum,
//		suite.wallpaperMultiplier,
//		suite.LatticeHeight,
//		wavepacket.Pmg,
//	)
//
//	checker.Assert(err, IsNil)
//	checker.Assert(GenericFormula.Formula.WavePackets, HasLen, 4)
//	checker.Assert(GenericFormula.Formula.WavePackets[0].Terms, HasLen, 1)
//
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Multiplier, Equals, suite.wallpaperMultiplier)
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerN * -1)
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerM * -1)
//
//	checker.Assert(GenericFormula.Formula.WavePackets[2].Multiplier, Equals, suite.wallpaperMultiplier * -1)
//	checker.Assert(GenericFormula.Formula.WavePackets[2].Terms[0].PowerN, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerN * -1)
//	checker.Assert(GenericFormula.Formula.WavePackets[2].Terms[0].PowerM, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerM)
//
//	checker.Assert(GenericFormula.Formula.WavePackets[3].Multiplier, Equals, suite.wallpaperMultiplier * -1)
//	checker.Assert(GenericFormula.Formula.WavePackets[3].Terms[0].PowerN, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerN)
//	checker.Assert(GenericFormula.Formula.WavePackets[3].Terms[0].PowerM, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerM * -1)
//
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pm), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pg), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmm), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmg), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pgg), Equals, false)
//}
//
//func (suite *GenericCreatedWithDesiredSymmetry) TestCreateWallpaperWithPgg(checker *C) {
//	GenericFormula, err := wavepacket.NewGenericWallpaperFormulaWithSymmetry(
//		suite.eisensteinTermWithOddPowerNAndEvenPowerSum,
//		suite.wallpaperMultiplier,
//		suite.LatticeHeight,
//		wavepacket.Pgg,
//	)
//
//	checker.Assert(err, IsNil)
//	checker.Assert(GenericFormula.Formula.WavePackets, HasLen, 4)
//	checker.Assert(GenericFormula.Formula.WavePackets[0].Terms, HasLen, 1)
//
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Multiplier, Equals, suite.wallpaperMultiplier)
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerN * -1)
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerM * -1)
//
//	checker.Assert(GenericFormula.Formula.WavePackets[2].Multiplier, Equals, suite.wallpaperMultiplier)
//	checker.Assert(GenericFormula.Formula.WavePackets[2].Terms[0].PowerN, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerN * -1)
//	checker.Assert(GenericFormula.Formula.WavePackets[2].Terms[0].PowerM, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerM)
//
//	checker.Assert(GenericFormula.Formula.WavePackets[3].Multiplier, Equals, suite.wallpaperMultiplier)
//	checker.Assert(GenericFormula.Formula.WavePackets[3].Terms[0].PowerN, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerN)
//	checker.Assert(GenericFormula.Formula.WavePackets[3].Terms[0].PowerM, Equals, suite.eisensteinTermWithOddPowerNAndEvenPowerSum[0].PowerM * -1)
//
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pm), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pg), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmm), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmg), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pgg), Equals, true)
//}
//
//func (suite *GenericCreatedWithDesiredSymmetry) TestCreateDesiredSymmetryWithYAML(checker *C) {
//	yamlByteStream := []byte(`
//lattice_height: 0.3
//formula:
// desired_symmetry: pmm
// multiplier:
//   real: -1.0
//   imaginary: 2e-2
// wave_packets:
//   -
//     multiplier:
//       real: -1.0
//       imaginary: 2e-2
//     terms:
//       -
//         power_n: 12
//         power_m: -9
//`)
//
//	GenericFormula, err := wavepacket.NewGenericWallpaperFormulaFromYAML(yamlByteStream)
//	checker.Assert(err, IsNil)
//
//	checker.Assert(GenericFormula.Formula.WavePackets, HasLen, 4)
//	checker.Assert(GenericFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -9)
//
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, -12)
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, 9)
//
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pm), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pg), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmm), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmg), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pgg), Equals, false)
//}
//
//func (suite *GenericCreatedWithDesiredSymmetry) TestCreateDesiredSymmetryWithJSON(checker *C) {
//	jsonByteStream := []byte(`{
//				"lattice_height": 0.3,
//				"formula": {
//					"desired_symmetry": "pmg",
//					"multiplier": {
//						"real": 1.0,
//						"imaginary": 0
//					},
//					"wave_packets": [
//						{
//							"multiplier": {
//								"real": 1.0,
//								"imaginary": 0
//							},
//							"terms": [
//								{
//									"power_n": 1,
//									"power_m": -2
//								}
//							]
//						}
//					]
//				}
//			}`)
//	GenericFormula, err := wavepacket.NewGenericWallpaperFormulaFromJSON(jsonByteStream)
//	checker.Assert(err, IsNil)
//
//	checker.Assert(GenericFormula.Formula.WavePackets, HasLen, 4)
//	checker.Assert(GenericFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -2)
//
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Multiplier, Equals, complex(1, 0))
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, -1)
//	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, 2)
//
//	checker.Assert(GenericFormula.Formula.WavePackets[2].Multiplier, Equals, complex(-1, 0))
//	checker.Assert(GenericFormula.Formula.WavePackets[2].Terms[0].PowerN, Equals, -1)
//	checker.Assert(GenericFormula.Formula.WavePackets[2].Terms[0].PowerM, Equals, -2)
//
//	checker.Assert(GenericFormula.Formula.WavePackets[3].Multiplier, Equals, complex(-1, 0))
//	checker.Assert(GenericFormula.Formula.WavePackets[3].Terms[0].PowerN, Equals, 1)
//	checker.Assert(GenericFormula.Formula.WavePackets[3].Terms[0].PowerM, Equals, 2)
//
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pm), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pg), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmm), Equals, false)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pmg), Equals, true)
//	checker.Assert(GenericFormula.HasSymmetry(wavepacket.Pgg), Equals, true)
//}