package formula_test

import (
	"github.com/Chadius/creating-symmetry/entities/formula"
	"github.com/Chadius/creating-symmetry/entities/utility"
	. "gopkg.in/check.v1"
)

type RectangularWallpaper struct {
	newFormula formula.Arbitrary
}

var _ = Suite(&RectangularWallpaper{})

func (suite *RectangularWallpaper) SetUpTest(checker *C) {
	suite.newFormula, _ = formula.NewBuilder().
		Rectangular().
		LatticeHeight(0.5).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1,0)).
				AddTerm(
					formula.NewTermBuilder().Multiplier(complex(1, 0)).PowerN(1).PowerM(-2).Build(),
				).
			Build(),
		).
	Build()
}

func (suite *RectangularWallpaper) TestSetupCreatesLatticeVectors(checker *C) {
	checker.Assert(real(suite.newFormula.LatticeVectors()[0]), utility.NumericallyCloseEnough{}, 1, 1e-6)
	checker.Assert(imag(suite.newFormula.LatticeVectors()[0]), utility.NumericallyCloseEnough{}, 0, 1e-6)

	checker.Assert(real(suite.newFormula.LatticeVectors()[1]), utility.NumericallyCloseEnough{}, 0, 1e-6)
	checker.Assert(imag(suite.newFormula.LatticeVectors()[1]), utility.NumericallyCloseEnough{}, 0.5, 1e-6)
}

//func (suite *RectangularWallpaper) TestCalculationOfPoints(checker *C) {
//	calculation := suite.newFormula.Calculate(complex(0.75, -0.25))
//
//	expectedAnswer := cmplx.Exp(complex(0, math.Pi*7/2))
//	checker.Assert(real(calculation), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
//	checker.Assert(imag(calculation), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
//}

//type RectangularWallpaperHasSymmetryTest struct {
//	baseWavePacketWithEvenPowerNAndOddPowerSum *wallpaper.WavePacket
//	baseWavePacketWithOddPowerNAndEvenPowerSum *wallpaper.WavePacket
//	wallpaperMultiplier                        complex128
//}
//
//var _ = Suite(&RectangularWallpaperHasSymmetryTest{})
//
//func (suite *RectangularWallpaperHasSymmetryTest) SetUpTest(checker *C) {
//	suite.baseWavePacketWithEvenPowerNAndOddPowerSum = &wallpaper.WavePacket{
//		terms: []*eisenstien.EisensteinFormulaTerm{
//			{
//				PowerN: 8,
//				PowerM: -3,
//			},
//		},
//		multiplier: complex(1, 0),
//	}
//	suite.baseWavePacketWithOddPowerNAndEvenPowerSum = &wallpaper.WavePacket{
//		terms: []*eisenstien.EisensteinFormulaTerm{
//			{
//				PowerN: 7,
//				PowerM: -3,
//			},
//		},
//		multiplier: complex(1, 0),
//	}
//	suite.wallpaperMultiplier = complex(1, 0)
//}
//
//func (suite *RectangularWallpaperHasSymmetryTest) TestRectangularHasNoSymmetry(checker *C) {
//	newFormula := wallpaper.Formula{
//		LatticeType: wallpaper.Rectangular,
//		LatticeSize: &wallpaper.Dimensions{
//			Width:  0,
//			Height: 1.5,
//		},
//		Lattice:    nil,
//		multiplier: complex(2, 0),
//		WavePackets: []*wallpaper.WavePacket{
//			suite.baseWavePacketWithEvenPowerNAndOddPowerSum,
//		},
//	}
//	err := newFormula.Setup()
//	checker.Assert(err, IsNil)
//
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, false)
//}
//
//func (suite *RectangularWallpaperHasSymmetryTest) TestRectangularMayHaveSymmetryForPm(checker *C) {
//	newFormula := wallpaper.Formula{
//		LatticeType: wallpaper.Rectangular,
//		LatticeSize: &wallpaper.Dimensions{
//			Width:  0,
//			Height: 1.5,
//		},
//		Lattice:    nil,
//		multiplier: complex(2, 0),
//		WavePackets: []*wallpaper.WavePacket{
//			suite.baseWavePacketWithEvenPowerNAndOddPowerSum,
//			{
//				terms: []*eisenstien.EisensteinFormulaTerm{
//					{
//						PowerN: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.terms[0].PowerN,
//						PowerM: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.terms[0].PowerM * -1,
//					},
//				},
//				multiplier: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.multiplier,
//			},
//		},
//	}
//	err := newFormula.Setup()
//	checker.Assert(err, IsNil)
//
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, true)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, true)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, false)
//}
//
//func (suite *RectangularWallpaperHasSymmetryTest) TestRectangularMayHaveSymmetryForPg(checker *C) {
//	newFormulaWithEvenPowerN := wallpaper.Formula{
//		LatticeType: wallpaper.Rectangular,
//		LatticeSize: &wallpaper.Dimensions{
//			Width:  0,
//			Height: 1.5,
//		},
//		Lattice:    nil,
//		multiplier: complex(2, 0),
//		WavePackets: []*wallpaper.WavePacket{
//			suite.baseWavePacketWithEvenPowerNAndOddPowerSum,
//			{
//				terms: []*eisenstien.EisensteinFormulaTerm{
//					{
//						PowerN: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.terms[0].PowerN,
//						PowerM: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.terms[0].PowerM * -1,
//					},
//				},
//				multiplier: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.multiplier,
//			},
//		},
//	}
//	err := newFormulaWithEvenPowerN.Setup()
//	checker.Assert(err, IsNil)
//
//	checker.Assert(newFormulaWithEvenPowerN.HasSymmetry(wallpaper.Pm), Equals, true)
//	checker.Assert(newFormulaWithEvenPowerN.HasSymmetry(wallpaper.Pg), Equals, true)
//	checker.Assert(newFormulaWithEvenPowerN.HasSymmetry(wallpaper.Pmm), Equals, false)
//	checker.Assert(newFormulaWithEvenPowerN.HasSymmetry(wallpaper.Pmg), Equals, false)
//	checker.Assert(newFormulaWithEvenPowerN.HasSymmetry(wallpaper.Pgg), Equals, false)
//
//	newFormulaWithOddPowerN := wallpaper.Formula{
//		LatticeType: wallpaper.Rectangular,
//		LatticeSize: &wallpaper.Dimensions{
//			Width:  0,
//			Height: 1.5,
//		},
//		Lattice:    nil,
//		multiplier: complex(2, 0),
//		WavePackets: []*wallpaper.WavePacket{
//			suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
//			{
//				terms: []*eisenstien.EisensteinFormulaTerm{
//					{
//						PowerN: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN,
//						PowerM: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM * -1,
//					},
//				},
//				multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier * -1,
//			},
//		},
//	}
//	oddErr := newFormulaWithOddPowerN.Setup()
//	checker.Assert(oddErr, IsNil)
//
//	checker.Assert(newFormulaWithOddPowerN.HasSymmetry(wallpaper.Pm), Equals, false)
//	checker.Assert(newFormulaWithOddPowerN.HasSymmetry(wallpaper.Pg), Equals, true)
//	checker.Assert(newFormulaWithOddPowerN.HasSymmetry(wallpaper.Pmm), Equals, false)
//	checker.Assert(newFormulaWithOddPowerN.HasSymmetry(wallpaper.Pmg), Equals, false)
//	checker.Assert(newFormulaWithOddPowerN.HasSymmetry(wallpaper.Pgg), Equals, false)
//}
//
//func (suite *RectangularWallpaperHasSymmetryTest) TestPmmAndPmgWithEvenPowerN(checker *C) {
//	newFormula := wallpaper.Formula{
//		LatticeType: wallpaper.Rectangular,
//		LatticeSize: &wallpaper.Dimensions{
//			Width:  0,
//			Height: 1.5,
//		},
//		Lattice:    nil,
//		multiplier: complex(2, 0),
//		WavePackets: []*wallpaper.WavePacket{
//			suite.baseWavePacketWithEvenPowerNAndOddPowerSum,
//			{
//				terms: []*eisenstien.EisensteinFormulaTerm{
//					{
//						PowerN: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.terms[0].PowerN * -1,
//						PowerM: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.terms[0].PowerM * -1,
//					},
//				},
//				multiplier: complex(1, 0),
//			},
//			{
//				terms: []*eisenstien.EisensteinFormulaTerm{
//					{
//						PowerN: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.terms[0].PowerN,
//						PowerM: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.terms[0].PowerM * -1,
//					},
//				},
//				multiplier: complex(1, 0),
//			},
//			{
//				terms: []*eisenstien.EisensteinFormulaTerm{
//					{
//						PowerN: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.terms[0].PowerN * -1,
//						PowerM: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.terms[0].PowerM,
//					},
//				},
//				multiplier: complex(1, 0),
//			},
//		},
//	}
//	err := newFormula.Setup()
//	checker.Assert(err, IsNil)
//
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, true)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, true)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, true)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, true)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, false)
//}
//
//func (suite *RectangularWallpaperHasSymmetryTest) TestPmgWithOddPowerN(checker *C) {
//	newFormula := wallpaper.Formula{
//		LatticeType: wallpaper.Rectangular,
//		LatticeSize: &wallpaper.Dimensions{
//			Width:  0,
//			Height: 1.5,
//		},
//		Lattice:    nil,
//		multiplier: complex(2, 0),
//		WavePackets: []*wallpaper.WavePacket{
//			suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
//			{
//				terms: []*eisenstien.EisensteinFormulaTerm{
//					{
//						PowerN: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN * -1,
//						PowerM: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM * -1,
//					},
//				},
//				multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier,
//			},
//			{
//				terms: []*eisenstien.EisensteinFormulaTerm{
//					{
//						PowerN: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN,
//						PowerM: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM * -1,
//					},
//				},
//				multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier * -1,
//			},
//			{
//				terms: []*eisenstien.EisensteinFormulaTerm{
//					{
//						PowerN: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN * -1,
//						PowerM: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM,
//					},
//				},
//				multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier * -1,
//			},
//		},
//	}
//	err := newFormula.Setup()
//	checker.Assert(err, IsNil)
//
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, true)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, true)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, false)
//}
//
//func (suite *RectangularWallpaperHasSymmetryTest) TestPgg(checker *C) {
//	newFormulaWithOddPowerSum := wallpaper.Formula{
//		LatticeType: wallpaper.Rectangular,
//		LatticeSize: &wallpaper.Dimensions{
//			Width:  0,
//			Height: 1.5,
//		},
//		Lattice:    nil,
//		multiplier: complex(2, 0),
//		WavePackets: []*wallpaper.WavePacket{
//			suite.baseWavePacketWithEvenPowerNAndOddPowerSum,
//			{
//				terms: []*eisenstien.EisensteinFormulaTerm{
//					{
//						PowerN: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.terms[0].PowerN * -1,
//						PowerM: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.terms[0].PowerM * -1,
//					},
//				},
//				multiplier: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.multiplier,
//			},
//			{
//				terms: []*eisenstien.EisensteinFormulaTerm{
//					{
//						PowerN: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.terms[0].PowerN,
//						PowerM: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.terms[0].PowerM * -1,
//					},
//				},
//				multiplier: -1 * suite.baseWavePacketWithEvenPowerNAndOddPowerSum.multiplier,
//			},
//			{
//				terms: []*eisenstien.EisensteinFormulaTerm{
//					{
//						PowerN: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.terms[0].PowerN * -1,
//						PowerM: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.terms[0].PowerM,
//					},
//				},
//				multiplier: -1 * suite.baseWavePacketWithEvenPowerNAndOddPowerSum.multiplier,
//			},
//		},
//	}
//	oddErr := newFormulaWithOddPowerSum.Setup()
//	checker.Assert(oddErr, IsNil)
//
//	checker.Assert(newFormulaWithOddPowerSum.HasSymmetry(wallpaper.Pm), Equals, false)
//	checker.Assert(newFormulaWithOddPowerSum.HasSymmetry(wallpaper.Pg), Equals, false)
//	checker.Assert(newFormulaWithOddPowerSum.HasSymmetry(wallpaper.Pmm), Equals, false)
//	checker.Assert(newFormulaWithOddPowerSum.HasSymmetry(wallpaper.Pmg), Equals, false)
//	checker.Assert(newFormulaWithOddPowerSum.HasSymmetry(wallpaper.Pgg), Equals, true)
//
//	newFormulaWithEvenPowerSum := wallpaper.Formula{
//		LatticeType: wallpaper.Rectangular,
//		LatticeSize: &wallpaper.Dimensions{
//			Width:  0,
//			Height: 1.5,
//		},
//		Lattice:    nil,
//		multiplier: complex(2, 0),
//		WavePackets: []*wallpaper.WavePacket{
//			suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
//			{
//				terms: []*eisenstien.EisensteinFormulaTerm{
//					{
//						PowerN: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN * -1,
//						PowerM: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM * -1,
//					},
//				},
//				multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier,
//			},
//			{
//				terms: []*eisenstien.EisensteinFormulaTerm{
//					{
//						PowerN: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN,
//						PowerM: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM * -1,
//					},
//				},
//				multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier,
//			},
//			{
//				terms: []*eisenstien.EisensteinFormulaTerm{
//					{
//						PowerN: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN * -1,
//						PowerM: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM,
//					},
//				},
//				multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier,
//			},
//		},
//	}
//	evenErr := newFormulaWithEvenPowerSum.Setup()
//	checker.Assert(evenErr, IsNil)
//
//	checker.Assert(newFormulaWithEvenPowerSum.HasSymmetry(wallpaper.Pm), Equals, true)
//	checker.Assert(newFormulaWithEvenPowerSum.HasSymmetry(wallpaper.Pg), Equals, false)
//	checker.Assert(newFormulaWithEvenPowerSum.HasSymmetry(wallpaper.Pmm), Equals, true)
//	checker.Assert(newFormulaWithEvenPowerSum.HasSymmetry(wallpaper.Pmg), Equals, false)
//	checker.Assert(newFormulaWithEvenPowerSum.HasSymmetry(wallpaper.Pgg), Equals, true)
//}
//
//type RectangularCreatedWithDesiredSymmetry struct {
//	baseWavePacketWithOddPowerNAndEvenPowerSum *wallpaper.WavePacket
//}
//
//var _ = Suite(&RectangularCreatedWithDesiredSymmetry{})
//
//func (suite *RectangularCreatedWithDesiredSymmetry) SetUpTest(checker *C) {
//	suite.baseWavePacketWithOddPowerNAndEvenPowerSum = &wallpaper.WavePacket{
//		terms: []*eisenstien.EisensteinFormulaTerm{
//			{
//				PowerN: 7,
//				PowerM: -3,
//			},
//		},
//		multiplier: complex(1, 0),
//	}
//}
//
//func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateWallpaperWithPm(checker *C) {
//	newFormula := wallpaper.Formula{
//		LatticeType: wallpaper.Rectangular,
//		LatticeSize: &wallpaper.Dimensions{
//			Width:  0,
//			Height: 2.0,
//		},
//		multiplier: complex(2, 0),
//		WavePackets: []*wallpaper.WavePacket{
//			suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
//		},
//		DesiredSymmetry: wallpaper.Pm,
//	}
//	err := newFormula.Setup()
//	checker.Assert(err, IsNil)
//
//	checker.Assert(newFormula.WavePackets, HasLen, 2)
//	checker.Assert(newFormula.WavePackets[0].terms, HasLen, 1)
//
//	checker.Assert(newFormula.WavePackets[1].multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier)
//	checker.Assert(newFormula.WavePackets[1].terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN)
//	checker.Assert(newFormula.WavePackets[1].terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM*-1)
//
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, true)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, false)
//}
//
//func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateWallpaperWithPg(checker *C) {
//	newFormula := wallpaper.Formula{
//		LatticeType: wallpaper.Rectangular,
//		LatticeSize: &wallpaper.Dimensions{
//			Width:  0,
//			Height: 2.0,
//		},
//		multiplier: complex(2, 0),
//		WavePackets: []*wallpaper.WavePacket{
//			suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
//		},
//		DesiredSymmetry: wallpaper.Pg,
//	}
//	err := newFormula.Setup()
//	checker.Assert(err, IsNil)
//
//	checker.Assert(newFormula.WavePackets, HasLen, 2)
//	checker.Assert(newFormula.WavePackets[0].terms, HasLen, 1)
//
//	checker.Assert(newFormula.WavePackets[1].multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier*-1)
//	checker.Assert(newFormula.WavePackets[1].terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN)
//	checker.Assert(newFormula.WavePackets[1].terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM*-1)
//
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, true)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, false)
//}
//
//func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateWallpaperWithPmm(checker *C) {
//	newFormula := wallpaper.Formula{
//		LatticeType: wallpaper.Rectangular,
//		LatticeSize: &wallpaper.Dimensions{
//			Width:  0,
//			Height: 2.0,
//		},
//		multiplier: complex(2, 0),
//		WavePackets: []*wallpaper.WavePacket{
//			suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
//		},
//		DesiredSymmetry: wallpaper.Pmm,
//	}
//	err := newFormula.Setup()
//	checker.Assert(err, IsNil)
//
//	checker.Assert(newFormula.WavePackets, HasLen, 4)
//	checker.Assert(newFormula.WavePackets[0].terms, HasLen, 1)
//
//	checker.Assert(newFormula.WavePackets[1].multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier)
//	checker.Assert(newFormula.WavePackets[1].terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN*-1)
//	checker.Assert(newFormula.WavePackets[1].terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM*-1)
//
//	checker.Assert(newFormula.WavePackets[1].multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier)
//	checker.Assert(newFormula.WavePackets[1].terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN*-1)
//	checker.Assert(newFormula.WavePackets[1].terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM*-1)
//
//	checker.Assert(newFormula.WavePackets[2].multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier)
//	checker.Assert(newFormula.WavePackets[2].terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN*-1)
//	checker.Assert(newFormula.WavePackets[2].terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM)
//
//	checker.Assert(newFormula.WavePackets[3].multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier)
//	checker.Assert(newFormula.WavePackets[3].terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN)
//	checker.Assert(newFormula.WavePackets[3].terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM*-1)
//
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, true)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, true)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, true)
//}
//
//func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateWallpaperWithPmg(checker *C) {
//	newFormula := wallpaper.Formula{
//		LatticeType: wallpaper.Rectangular,
//		LatticeSize: &wallpaper.Dimensions{
//			Width:  0,
//			Height: 2.0,
//		},
//		multiplier: complex(2, 0),
//		WavePackets: []*wallpaper.WavePacket{
//			suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
//		},
//		DesiredSymmetry: wallpaper.Pmg,
//	}
//	err := newFormula.Setup()
//	checker.Assert(err, IsNil)
//
//	checker.Assert(newFormula.WavePackets, HasLen, 4)
//	checker.Assert(newFormula.WavePackets[0].terms, HasLen, 1)
//
//	checker.Assert(newFormula.WavePackets[1].multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier)
//	checker.Assert(newFormula.WavePackets[1].terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN*-1)
//	checker.Assert(newFormula.WavePackets[1].terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM*-1)
//
//	checker.Assert(newFormula.WavePackets[2].multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier*-1)
//	checker.Assert(newFormula.WavePackets[2].terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN*-1)
//	checker.Assert(newFormula.WavePackets[2].terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM)
//
//	checker.Assert(newFormula.WavePackets[3].multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier*-1)
//	checker.Assert(newFormula.WavePackets[3].terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN)
//	checker.Assert(newFormula.WavePackets[3].terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM*-1)
//
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, true)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, true)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, false)
//}
//
//func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateWallpaperWithPgg(checker *C) {
//	newFormula := wallpaper.Formula{
//		LatticeType: wallpaper.Rectangular,
//		LatticeSize: &wallpaper.Dimensions{
//			Width:  0,
//			Height: 2.0,
//		},
//		multiplier: complex(2, 0),
//		WavePackets: []*wallpaper.WavePacket{
//			suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
//		},
//		DesiredSymmetry: wallpaper.Pgg,
//	}
//	err := newFormula.Setup()
//	checker.Assert(err, IsNil)
//
//	checker.Assert(newFormula.WavePackets, HasLen, 4)
//	checker.Assert(newFormula.WavePackets[0].terms, HasLen, 1)
//
//	checker.Assert(newFormula.WavePackets[1].multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier)
//	checker.Assert(newFormula.WavePackets[1].terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN*-1)
//	checker.Assert(newFormula.WavePackets[1].terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM*-1)
//
//	checker.Assert(newFormula.WavePackets[2].multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier)
//	checker.Assert(newFormula.WavePackets[2].terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN*-1)
//	checker.Assert(newFormula.WavePackets[2].terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM)
//
//	checker.Assert(newFormula.WavePackets[3].multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.multiplier)
//	checker.Assert(newFormula.WavePackets[3].terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerN)
//	checker.Assert(newFormula.WavePackets[3].terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.terms[0].PowerM*-1)
//
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, true)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, true)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, false)
//	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, true)
//}
//
