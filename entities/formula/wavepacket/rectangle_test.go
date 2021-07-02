package wavepacket_test

import (
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/wavepacket"
	"wallpaper/entities/utility"
)

type RectangularWallpaper struct {
	RectangularFormula *wavepacket.RectangularWallpaperFormula
}

var _ = Suite(&RectangularWallpaper{})

func (suite *RectangularWallpaper) SetUpTest (checker *C) {
	suite.RectangularFormula = &wavepacket.RectangularWallpaperFormula{
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
		LatticeHeight: 0.5,
	}
	suite.RectangularFormula.SetUp()
}

func (suite *RectangularWallpaper) TestSetupCreatesLatticeVectors (checker *C) {
	checker.Assert(real(suite.RectangularFormula.Formula.Lattice.XLatticeVector), utility.NumericallyCloseEnough{}, 1, 1e-6)
	checker.Assert(imag(suite.RectangularFormula.Formula.Lattice.XLatticeVector), utility.NumericallyCloseEnough{}, 0, 1e-6)

	checker.Assert(real(suite.RectangularFormula.Formula.Lattice.YLatticeVector), utility.NumericallyCloseEnough{}, 0, 1e-6)
	checker.Assert(imag(suite.RectangularFormula.Formula.Lattice.YLatticeVector), utility.NumericallyCloseEnough{}, 0.5, 1e-6)
}

func (suite *RectangularWallpaper) TestRaiseErrorIfHeightIsZero (checker *C) {
	RectangularFormulaWithNoHeight := &wavepacket.RectangularWallpaperFormula{
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
		LatticeHeight: 0,
	}
	err := RectangularFormulaWithNoHeight.SetUp()

	checker.Assert(err, ErrorMatches, "lattice vectors cannot be \\(0,0\\)")
}

func (suite *RectangularWallpaper) TestSetupLocksPairs (checker *C) {
	checker.Assert(suite.RectangularFormula.Formula.WavePackets[0].Terms, HasLen, 1)
}

func (suite *RectangularWallpaper) TestCalculationOfPoints (checker *C) {
	calculation := suite.RectangularFormula.Calculate(complex(0.75, -0.25))
	total := calculation.Total

	expectedAnswer := cmplx.Exp(complex(0, math.Pi * 7 / 2))

	checker.Assert(real(total), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(total), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}

func (suite *RectangularWallpaper) TestUnmarshalFromJSON(checker *C) {
	jsonByteStream := []byte(`{
				"lattice_height": 0.3,
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
	RectangularFormula, err := wavepacket.NewRectangularWallpaperFormulaFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(RectangularFormula.LatticeHeight, utility.NumericallyCloseEnough{}, 0.3, 1e-6)
	checker.Assert(real(RectangularFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(RectangularFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
	checker.Assert(RectangularFormula.Formula.WavePackets, HasLen, 1)
	checker.Assert(RectangularFormula.Formula.WavePackets[0].Terms[0].PowerN, Equals, 12)
	checker.Assert(RectangularFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -10)
}

func (suite *RectangularWallpaper) TestUnmarshalFromYAML(checker *C) {
	yamlByteStream := []byte(`
lattice_height: 0.3
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
	RectangularFormula, err := wavepacket.NewRectangularWallpaperFormulaFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(RectangularFormula.LatticeHeight, utility.NumericallyCloseEnough{}, 0.3, 1e-6)
	checker.Assert(real(RectangularFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(RectangularFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
	checker.Assert(RectangularFormula.Formula.WavePackets, HasLen, 1)
	checker.Assert(RectangularFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -10)
}

//type RectangularWaveSymmetry struct {
//	baseWavePacket *wavepacket.WavePacket
//}
//
//var _ = Suite(&RectangularWaveSymmetry{})
//
//func (suite *RectangularWaveSymmetry) SetUpTest(checker *C) {
//	suite.baseWavePacket = &wavepacket.WavePacket{
//		Terms:[]*formula.EisensteinFormulaTerm{
//			{
//				PowerN:         8,
//				PowerM:         -3,
//			},
//		},
//		Multiplier: complex(1, 0),
//	}
//}
//
//func (suite *RectangularWaveSymmetry) TestNoSymmetryFound(checker *C) {
//	RectangularFormula := wavepacket.RectangularWallpaperFormula{
//		Formula: &wavepacket.WallpaperFormula{
//			WavePackets: []*wavepacket.WavePacket{
//				suite.baseWavePacket,
//			},
//			Multiplier: complex(1, 0),
//		},
//	}
//
//	RectangularFormula.SetUp()
//	checker.Assert(RectangularFormula.HasSymmetry(wavepacket.Cm), Equals, false)
//	checker.Assert(RectangularFormula.HasSymmetry(wavepacket.Cmm), Equals, false)
//}
//
//func (suite *RectangularWaveSymmetry) TestCm(checker *C) {
//	RectangularFormula := wavepacket.RectangularWallpaperFormula{
//		Formula: &wavepacket.WallpaperFormula{
//			WavePackets: []*wavepacket.WavePacket{
//				suite.baseWavePacket,
//				{
//					Terms: []*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacket.Terms[0].PowerM,
//							PowerM:         suite.baseWavePacket.Terms[0].PowerN,
//						},
//					},
//				},
//			},
//			Multiplier: complex(1, 0),
//		},
//	}
//
//	RectangularFormula.SetUp()
//	checker.Assert(RectangularFormula.HasSymmetry(wavepacket.Cm), Equals, true)
//	checker.Assert(RectangularFormula.HasSymmetry(wavepacket.Cmm), Equals, false)
//}
//
//func (suite *RectangularWaveSymmetry) TestCmm(checker *C) {
//	RectangularFormula := wavepacket.RectangularWallpaperFormula{
//		Formula: &wavepacket.WallpaperFormula{
//			WavePackets: []*wavepacket.WavePacket{
//				suite.baseWavePacket,
//				{
//					Terms: []*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacket.Terms[0].PowerM,
//							PowerM:         suite.baseWavePacket.Terms[0].PowerN,
//						},
//					},
//				},
//				{
//					Terms: []*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacket.Terms[0].PowerM * -1,
//							PowerM:         suite.baseWavePacket.Terms[0].PowerN * -1,
//						},
//					},
//				},
//				{
//					Terms: []*formula.EisensteinFormulaTerm{
//						{
//							PowerN:         suite.baseWavePacket.Terms[0].PowerN * -1,
//							PowerM:         suite.baseWavePacket.Terms[0].PowerM * -1,
//						},
//					},
//				},
//			},
//			Multiplier: complex(1, 0),
//		},
//	}
//
//	RectangularFormula.SetUp()
//	checker.Assert(RectangularFormula.HasSymmetry(wavepacket.Cm), Equals, true)
//	checker.Assert(RectangularFormula.HasSymmetry(wavepacket.Cmm), Equals, true)
//}
//
//type RectangularCreatedWithDesiredSymmetry struct {
//	singleEisensteinFormulaTerm []*formula.EisensteinFormulaTerm
//	wallpaperMultiplier complex128
//	LatticeHeight float64
//}
//
//var _ = Suite(&RectangularCreatedWithDesiredSymmetry{})
//
//func (suite *RectangularCreatedWithDesiredSymmetry) SetUpTest (checker *C) {
//	suite.singleEisensteinFormulaTerm = []*formula.EisensteinFormulaTerm{
//		{
//			PowerN: 1,
//			PowerM: -2,
//			Multiplier: complex(1, 0),
//		},
//	}
//	suite.wallpaperMultiplier = complex(1, 0)
//	suite.LatticeHeight = 1.0
//}
//
//func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateWallpaperWithCm(checker *C) {
//	RectangularFormula, err := wavepacket.NewRectangularWallpaperFormulaWithSymmetry(
//		suite.singleEisensteinFormulaTerm,
//		suite.wallpaperMultiplier,
//		suite.LatticeHeight,
//		wavepacket.Cm,
//	)
//
//	checker.Assert(err, IsNil)
//	checker.Assert(RectangularFormula.Formula.WavePackets, HasLen, 2)
//	checker.Assert(RectangularFormula.Formula.WavePackets[0].Terms, HasLen, 2)
//
//	checker.Assert(RectangularFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, suite.singleEisensteinFormulaTerm[0].PowerM)
//	checker.Assert(RectangularFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, suite.singleEisensteinFormulaTerm[0].PowerN)
//
//	checker.Assert(RectangularFormula.HasSymmetry(wavepacket.Cm), Equals, true)
//	checker.Assert(RectangularFormula.HasSymmetry(wavepacket.Cmm), Equals, false)
//}
//
//func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateWallpaperWithCmm(checker *C) {
//	RectangularFormula, err := wavepacket.NewRectangularWallpaperFormulaWithSymmetry(
//		suite.singleEisensteinFormulaTerm,
//		suite.wallpaperMultiplier,
//		suite.LatticeHeight,
//		wavepacket.Cmm,
//	)
//
//	checker.Assert(err, IsNil)
//	checker.Assert(RectangularFormula.Formula.WavePackets, HasLen, 4)
//	checker.Assert(RectangularFormula.Formula.WavePackets[0].Terms, HasLen, 2)
//
//	checker.Assert(RectangularFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, suite.singleEisensteinFormulaTerm[0].PowerN * -1)
//	checker.Assert(RectangularFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, suite.singleEisensteinFormulaTerm[0].PowerM * -1)
//	checker.Assert(RectangularFormula.Formula.WavePackets[2].Terms[0].PowerN, Equals, suite.singleEisensteinFormulaTerm[0].PowerM)
//	checker.Assert(RectangularFormula.Formula.WavePackets[2].Terms[0].PowerM, Equals, suite.singleEisensteinFormulaTerm[0].PowerN)
//	checker.Assert(RectangularFormula.Formula.WavePackets[3].Terms[0].PowerN, Equals, suite.singleEisensteinFormulaTerm[0].PowerM * -1)
//	checker.Assert(RectangularFormula.Formula.WavePackets[3].Terms[0].PowerM, Equals, suite.singleEisensteinFormulaTerm[0].PowerN * -1)
//
//	checker.Assert(RectangularFormula.HasSymmetry(wavepacket.Cm), Equals, true)
//	checker.Assert(RectangularFormula.HasSymmetry(wavepacket.Cmm), Equals, true)
//}
//
//func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateDesiredSymmetryWithYAML(checker *C) {
//	yamlByteStream := []byte(`
//lattice_height: 0.3
//formula:
//  desired_symmetry: cm
//  multiplier:
//    real: 1.0
//    imaginary: 0
//  wave_packets:
//    -
//      multiplier:
//        real: 1.0
//        imaginary: 0
//      terms:
//        -
//          power_n: 1
//          power_m: -2
//`)
//	RectangularFormula, err := wavepacket.NewRectangularWallpaperFormulaFromYAML(yamlByteStream)
//	checker.Assert(err, IsNil)
//
//	checker.Assert(RectangularFormula.Formula.WavePackets, HasLen, 2)
//	checker.Assert(RectangularFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -2)
//
//	checker.Assert(RectangularFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, -2)
//	checker.Assert(RectangularFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, 1)
//
//	checker.Assert(RectangularFormula.HasSymmetry(wavepacket.Cm), Equals, true)
//	checker.Assert(RectangularFormula.HasSymmetry(wavepacket.Cmm), Equals, false)
//}
//
//func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateDesiredSymmetryWithJSON(checker *C) {
//	jsonByteStream := []byte(`{
//				"lattice_height": 0.3,
//				"formula": {
//					"desired_symmetry": "cmm",
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
//	RectangularFormula, err := wavepacket.NewRectangularWallpaperFormulaFromJSON(jsonByteStream)
//	checker.Assert(err, IsNil)
//
//	checker.Assert(RectangularFormula.Formula.WavePackets, HasLen, 4)
//	checker.Assert(RectangularFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -2)
//
//	checker.Assert(RectangularFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, -1)
//	checker.Assert(RectangularFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, 2)
//	checker.Assert(RectangularFormula.Formula.WavePackets[2].Terms[0].PowerN, Equals, -2)
//	checker.Assert(RectangularFormula.Formula.WavePackets[2].Terms[0].PowerM, Equals, 1)
//	checker.Assert(RectangularFormula.Formula.WavePackets[3].Terms[0].PowerN, Equals, 2)
//	checker.Assert(RectangularFormula.Formula.WavePackets[3].Terms[0].PowerM, Equals, -1)
//
//	checker.Assert(RectangularFormula.HasSymmetry(wavepacket.Cm), Equals, true)
//	checker.Assert(RectangularFormula.HasSymmetry(wavepacket.Cmm), Equals, true)
//}