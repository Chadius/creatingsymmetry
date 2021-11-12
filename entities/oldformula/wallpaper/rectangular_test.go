package wallpaper_test

import (
	"github.com/Chadius/creating-symmetry/entities/oldformula"
	"github.com/Chadius/creating-symmetry/entities/oldformula/wallpaper"
	"github.com/Chadius/creating-symmetry/entities/utility"
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"
)

type RectangularWallpaper struct {
	newFormula *wallpaper.Formula
}

var _ = Suite(&RectangularWallpaper{})

func (suite *RectangularWallpaper) SetUpTest(checker *C) {
	suite.newFormula = &wallpaper.Formula{
		LatticeType: wallpaper.Rectangular,
		LatticeSize: &wallpaper.Dimensions{
			Width:  2,
			Height: 0.5,
		},
		Lattice:    nil,
		Multiplier: complex(1, 0),
		WavePackets: []*wallpaper.WavePacket{
			{
				Terms: []*oldformula.EisensteinFormulaTerm{
					{
						PowerN: 1,
						PowerM: -2,
					},
				},
				Multiplier: complex(1, 0),
			},
		},
	}

	suite.newFormula.Setup()
}

func (suite *RectangularWallpaper) TestSetupCreatesLatticeVectors(checker *C) {
	checker.Assert(real(suite.newFormula.Lattice.XLatticeVector), utility.NumericallyCloseEnough{}, 1, 1e-6)
	checker.Assert(imag(suite.newFormula.Lattice.XLatticeVector), utility.NumericallyCloseEnough{}, 0, 1e-6)

	checker.Assert(real(suite.newFormula.Lattice.YLatticeVector), utility.NumericallyCloseEnough{}, 0, 1e-6)
	checker.Assert(imag(suite.newFormula.Lattice.YLatticeVector), utility.NumericallyCloseEnough{}, 0.5, 1e-6)
}

func (suite *RectangularWallpaper) TestCalculationOfPoints(checker *C) {
	calculation := suite.newFormula.Calculate(complex(0.75, -0.25))
	total := calculation.Total

	expectedAnswer := cmplx.Exp(complex(0, math.Pi*7/2))

	checker.Assert(real(total), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(total), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}

type RectangularWallpaperHasSymmetryTest struct {
	baseWavePacketWithEvenPowerNAndOddPowerSum *wallpaper.WavePacket
	baseWavePacketWithOddPowerNAndEvenPowerSum *wallpaper.WavePacket
	wallpaperMultiplier                        complex128
}

var _ = Suite(&RectangularWallpaperHasSymmetryTest{})

func (suite *RectangularWallpaperHasSymmetryTest) SetUpTest(checker *C) {
	suite.baseWavePacketWithEvenPowerNAndOddPowerSum = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: 8,
				PowerM: -3,
			},
		},
		Multiplier: complex(1, 0),
	}
	suite.baseWavePacketWithOddPowerNAndEvenPowerSum = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: 7,
				PowerM: -3,
			},
		},
		Multiplier: complex(1, 0),
	}
	suite.wallpaperMultiplier = complex(1, 0)
}

func (suite *RectangularWallpaperHasSymmetryTest) TestRectangularHasNoSymmetry(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Rectangular,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0,
			Height: 1.5,
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacketWithEvenPowerNAndOddPowerSum,
		},
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, false)
}

func (suite *RectangularWallpaperHasSymmetryTest) TestRectangularMayHaveSymmetryForPm(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Rectangular,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0,
			Height: 1.5,
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacketWithEvenPowerNAndOddPowerSum,
			{
				Terms: []*oldformula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerN,
						PowerM: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerM * -1,
					},
				},
				Multiplier: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier,
			},
		},
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, false)
}

func (suite *RectangularWallpaperHasSymmetryTest) TestRectangularMayHaveSymmetryForPg(checker *C) {
	newFormulaWithEvenPowerN := wallpaper.Formula{
		LatticeType: wallpaper.Rectangular,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0,
			Height: 1.5,
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacketWithEvenPowerNAndOddPowerSum,
			{
				Terms: []*oldformula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerN,
						PowerM: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerM * -1,
					},
				},
				Multiplier: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier,
			},
		},
	}
	err := newFormulaWithEvenPowerN.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormulaWithEvenPowerN.HasSymmetry(wallpaper.Pm), Equals, true)
	checker.Assert(newFormulaWithEvenPowerN.HasSymmetry(wallpaper.Pg), Equals, true)
	checker.Assert(newFormulaWithEvenPowerN.HasSymmetry(wallpaper.Pmm), Equals, false)
	checker.Assert(newFormulaWithEvenPowerN.HasSymmetry(wallpaper.Pmg), Equals, false)
	checker.Assert(newFormulaWithEvenPowerN.HasSymmetry(wallpaper.Pgg), Equals, false)

	newFormulaWithOddPowerN := wallpaper.Formula{
		LatticeType: wallpaper.Rectangular,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0,
			Height: 1.5,
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
			{
				Terms: []*oldformula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN,
						PowerM: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM * -1,
					},
				},
				Multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier * -1,
			},
		},
	}
	oddErr := newFormulaWithOddPowerN.Setup()
	checker.Assert(oddErr, IsNil)

	checker.Assert(newFormulaWithOddPowerN.HasSymmetry(wallpaper.Pm), Equals, false)
	checker.Assert(newFormulaWithOddPowerN.HasSymmetry(wallpaper.Pg), Equals, true)
	checker.Assert(newFormulaWithOddPowerN.HasSymmetry(wallpaper.Pmm), Equals, false)
	checker.Assert(newFormulaWithOddPowerN.HasSymmetry(wallpaper.Pmg), Equals, false)
	checker.Assert(newFormulaWithOddPowerN.HasSymmetry(wallpaper.Pgg), Equals, false)
}

func (suite *RectangularWallpaperHasSymmetryTest) TestPmmAndPmgWithEvenPowerN(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Rectangular,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0,
			Height: 1.5,
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacketWithEvenPowerNAndOddPowerSum,
			{
				Terms: []*oldformula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerN * -1,
						PowerM: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerM * -1,
					},
				},
				Multiplier: complex(1, 0),
			},
			{
				Terms: []*oldformula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerN,
						PowerM: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerM * -1,
					},
				},
				Multiplier: complex(1, 0),
			},
			{
				Terms: []*oldformula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerN * -1,
						PowerM: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerM,
					},
				},
				Multiplier: complex(1, 0),
			},
		},
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, false)
}

func (suite *RectangularWallpaperHasSymmetryTest) TestPmgWithOddPowerN(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Rectangular,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0,
			Height: 1.5,
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
			{
				Terms: []*oldformula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN * -1,
						PowerM: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM * -1,
					},
				},
				Multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier,
			},
			{
				Terms: []*oldformula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN,
						PowerM: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM * -1,
					},
				},
				Multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier * -1,
			},
			{
				Terms: []*oldformula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN * -1,
						PowerM: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM,
					},
				},
				Multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier * -1,
			},
		},
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, false)
}

func (suite *RectangularWallpaperHasSymmetryTest) TestPgg(checker *C) {
	newFormulaWithOddPowerSum := wallpaper.Formula{
		LatticeType: wallpaper.Rectangular,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0,
			Height: 1.5,
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacketWithEvenPowerNAndOddPowerSum,
			{
				Terms: []*oldformula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerN * -1,
						PowerM: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerM * -1,
					},
				},
				Multiplier: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier,
			},
			{
				Terms: []*oldformula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerN,
						PowerM: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerM * -1,
					},
				},
				Multiplier: -1 * suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier,
			},
			{
				Terms: []*oldformula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerN * -1,
						PowerM: suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms[0].PowerM,
					},
				},
				Multiplier: -1 * suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier,
			},
		},
	}
	oddErr := newFormulaWithOddPowerSum.Setup()
	checker.Assert(oddErr, IsNil)

	checker.Assert(newFormulaWithOddPowerSum.HasSymmetry(wallpaper.Pm), Equals, false)
	checker.Assert(newFormulaWithOddPowerSum.HasSymmetry(wallpaper.Pg), Equals, false)
	checker.Assert(newFormulaWithOddPowerSum.HasSymmetry(wallpaper.Pmm), Equals, false)
	checker.Assert(newFormulaWithOddPowerSum.HasSymmetry(wallpaper.Pmg), Equals, false)
	checker.Assert(newFormulaWithOddPowerSum.HasSymmetry(wallpaper.Pgg), Equals, true)

	newFormulaWithEvenPowerSum := wallpaper.Formula{
		LatticeType: wallpaper.Rectangular,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0,
			Height: 1.5,
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
			{
				Terms: []*oldformula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN * -1,
						PowerM: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM * -1,
					},
				},
				Multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier,
			},
			{
				Terms: []*oldformula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN,
						PowerM: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM * -1,
					},
				},
				Multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier,
			},
			{
				Terms: []*oldformula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN * -1,
						PowerM: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM,
					},
				},
				Multiplier: suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier,
			},
		},
	}
	evenErr := newFormulaWithEvenPowerSum.Setup()
	checker.Assert(evenErr, IsNil)

	checker.Assert(newFormulaWithEvenPowerSum.HasSymmetry(wallpaper.Pm), Equals, true)
	checker.Assert(newFormulaWithEvenPowerSum.HasSymmetry(wallpaper.Pg), Equals, false)
	checker.Assert(newFormulaWithEvenPowerSum.HasSymmetry(wallpaper.Pmm), Equals, true)
	checker.Assert(newFormulaWithEvenPowerSum.HasSymmetry(wallpaper.Pmg), Equals, false)
	checker.Assert(newFormulaWithEvenPowerSum.HasSymmetry(wallpaper.Pgg), Equals, true)
}

type RectangularCreatedWithDesiredSymmetry struct {
	baseWavePacketWithOddPowerNAndEvenPowerSum *wallpaper.WavePacket
}

var _ = Suite(&RectangularCreatedWithDesiredSymmetry{})

func (suite *RectangularCreatedWithDesiredSymmetry) SetUpTest(checker *C) {
	suite.baseWavePacketWithOddPowerNAndEvenPowerSum = &wallpaper.WavePacket{
		Terms: []*oldformula.EisensteinFormulaTerm{
			{
				PowerN: 7,
				PowerM: -3,
			},
		},
		Multiplier: complex(1, 0),
	}
}

func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateWallpaperWithPm(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Rectangular,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0,
			Height: 2.0,
		},
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
		},
		DesiredSymmetry: wallpaper.Pm,
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.WavePackets, HasLen, 2)
	checker.Assert(newFormula.WavePackets[0].Terms, HasLen, 1)

	checker.Assert(newFormula.WavePackets[1].Multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM*-1)

	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, false)
}

func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateWallpaperWithPg(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Rectangular,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0,
			Height: 2.0,
		},
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
		},
		DesiredSymmetry: wallpaper.Pg,
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.WavePackets, HasLen, 2)
	checker.Assert(newFormula.WavePackets[0].Terms, HasLen, 1)

	checker.Assert(newFormula.WavePackets[1].Multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier*-1)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM*-1)

	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, false)
}

func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateWallpaperWithPmm(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Rectangular,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0,
			Height: 2.0,
		},
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
		},
		DesiredSymmetry: wallpaper.Pmm,
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.WavePackets, HasLen, 4)
	checker.Assert(newFormula.WavePackets[0].Terms, HasLen, 1)

	checker.Assert(newFormula.WavePackets[1].Multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN*-1)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM*-1)

	checker.Assert(newFormula.WavePackets[1].Multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN*-1)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM*-1)

	checker.Assert(newFormula.WavePackets[2].Multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier)
	checker.Assert(newFormula.WavePackets[2].Terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN*-1)
	checker.Assert(newFormula.WavePackets[2].Terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM)

	checker.Assert(newFormula.WavePackets[3].Multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier)
	checker.Assert(newFormula.WavePackets[3].Terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN)
	checker.Assert(newFormula.WavePackets[3].Terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM*-1)

	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, true)
}

func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateWallpaperWithPmg(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Rectangular,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0,
			Height: 2.0,
		},
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
		},
		DesiredSymmetry: wallpaper.Pmg,
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.WavePackets, HasLen, 4)
	checker.Assert(newFormula.WavePackets[0].Terms, HasLen, 1)

	checker.Assert(newFormula.WavePackets[1].Multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN*-1)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM*-1)

	checker.Assert(newFormula.WavePackets[2].Multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier*-1)
	checker.Assert(newFormula.WavePackets[2].Terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN*-1)
	checker.Assert(newFormula.WavePackets[2].Terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM)

	checker.Assert(newFormula.WavePackets[3].Multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier*-1)
	checker.Assert(newFormula.WavePackets[3].Terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN)
	checker.Assert(newFormula.WavePackets[3].Terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM*-1)

	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, false)
}

func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateWallpaperWithPgg(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Rectangular,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0,
			Height: 2.0,
		},
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacketWithOddPowerNAndEvenPowerSum,
		},
		DesiredSymmetry: wallpaper.Pgg,
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.WavePackets, HasLen, 4)
	checker.Assert(newFormula.WavePackets[0].Terms, HasLen, 1)

	checker.Assert(newFormula.WavePackets[1].Multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN*-1)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM*-1)

	checker.Assert(newFormula.WavePackets[2].Multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier)
	checker.Assert(newFormula.WavePackets[2].Terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN*-1)
	checker.Assert(newFormula.WavePackets[2].Terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM)

	checker.Assert(newFormula.WavePackets[3].Multiplier, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier)
	checker.Assert(newFormula.WavePackets[3].Terms[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerN)
	checker.Assert(newFormula.WavePackets[3].Terms[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms[0].PowerM*-1)

	checker.Assert(newFormula.HasSymmetry(wallpaper.Pm), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pg), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmm), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pmg), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Pgg), Equals, true)
}
