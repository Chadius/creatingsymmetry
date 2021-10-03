package wallpaper_test

import (
	"github.com/Chadius/creating-symmetry/entities/formula"
	"github.com/Chadius/creating-symmetry/entities/formula/wallpaper"
	"github.com/Chadius/creating-symmetry/entities/utility"
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"
)

type RhombicWallpaper struct {
	newFormula *wallpaper.Formula
}

var _ = Suite(&RhombicWallpaper{})

func (suite *RhombicWallpaper) SetUpTest(checker *C) {
	suite.newFormula = &wallpaper.Formula{
		LatticeType: wallpaper.Rhombic,
		LatticeSize: &wallpaper.Dimensions{
			Width:  2,
			Height: 1,
		},
		Lattice:    nil,
		Multiplier: complex(1, 0),
		WavePackets: []*wallpaper.WavePacket{
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
	}

	suite.newFormula.Setup()
}

func (suite *RhombicWallpaper) TestSetupCreatesLatticeVectors(checker *C) {
	checker.Assert(real(suite.newFormula.Lattice.XLatticeVector), utility.NumericallyCloseEnough{}, 0.5, 1e-6)
	checker.Assert(imag(suite.newFormula.Lattice.XLatticeVector), utility.NumericallyCloseEnough{}, 1.0, 1e-6)

	checker.Assert(real(suite.newFormula.Lattice.YLatticeVector), utility.NumericallyCloseEnough{}, 0.5, 1e-6)
	checker.Assert(imag(suite.newFormula.Lattice.YLatticeVector), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
}

func (suite *RhombicWallpaper) TestSetupAddsLockedPairs(checker *C) {
	checker.Assert(suite.newFormula.WavePackets[0].Terms, HasLen, 2)
	checker.Assert(suite.newFormula.WavePackets[0].Terms[1].PowerN, Equals, suite.newFormula.WavePackets[0].Terms[0].PowerM)
	checker.Assert(suite.newFormula.WavePackets[0].Terms[1].PowerM, Equals, suite.newFormula.WavePackets[0].Terms[0].PowerN)
}

func (suite *RhombicWallpaper) TestCalculationOfPoints(checker *C) {
	calculation := suite.newFormula.Calculate(complex(0.75, -0.25))
	total := calculation.Total

	expectedAnswer := (cmplx.Exp(complex(0, math.Pi*-9/4)) +
		cmplx.Exp(complex(0, math.Pi*-3/4))) / 2

	checker.Assert(real(total), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(total), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}

type RhombicWallpaperHasSymmetryTest struct {
	baseWavePacket      *wallpaper.WavePacket
	wallpaperMultiplier complex128
}

var _ = Suite(&RhombicWallpaperHasSymmetryTest{})

func (suite *RhombicWallpaperHasSymmetryTest) SetUpTest(checker *C) {
	suite.baseWavePacket = &wallpaper.WavePacket{
		Terms: []*formula.EisensteinFormulaTerm{
			{
				PowerN: 8,
				PowerM: -3,
			},
		},
		Multiplier: complex(1, 0),
	}

	suite.wallpaperMultiplier = complex(1, 0)
}

func (suite *RhombicWallpaperHasSymmetryTest) TestRhombicHasNoSymmetry(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Rhombic,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0.5,
			Height: 1,
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacket,
		},
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.WavePackets, HasLen, 1)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Cm), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Cmm), Equals, false)
}

func (suite *RhombicWallpaperHasSymmetryTest) TestRhombicMayHaveSymmetryForCm(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Rhombic,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0.5,
			Height: 1,
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacket,
			{
				Terms: []*formula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacket.Terms[0].PowerM,
						PowerM: suite.baseWavePacket.Terms[0].PowerN,
					},
				},
				Multiplier: suite.baseWavePacket.Multiplier,
			},
		},
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.HasSymmetry(wallpaper.Cm), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Cmm), Equals, false)
}

func (suite *RhombicWallpaperHasSymmetryTest) TestRhombicMayHaveSymmetryForCmm(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Rhombic,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0.5,
			Height: 1,
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacket,
			{
				Terms: []*formula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacket.Terms[0].PowerM,
						PowerM: suite.baseWavePacket.Terms[0].PowerN,
					},
				},
				Multiplier: suite.baseWavePacket.Multiplier,
			},
			{
				Terms: []*formula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacket.Terms[0].PowerM * -1,
						PowerM: suite.baseWavePacket.Terms[0].PowerN * -1,
					},
				},
				Multiplier: suite.baseWavePacket.Multiplier,
			},
			{
				Terms: []*formula.EisensteinFormulaTerm{
					{
						PowerN: suite.baseWavePacket.Terms[0].PowerN * -1,
						PowerM: suite.baseWavePacket.Terms[0].PowerM * -1,
					},
				},
				Multiplier: suite.baseWavePacket.Multiplier,
			},
		},
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.HasSymmetry(wallpaper.Cm), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Cmm), Equals, true)
}

type RhombicCreatedWithDesiredSymmetry struct {
	baseWavePacket      *wallpaper.WavePacket
	wallpaperMultiplier complex128
}

var _ = Suite(&RhombicCreatedWithDesiredSymmetry{})

func (suite *RhombicCreatedWithDesiredSymmetry) SetUpTest(checker *C) {
	suite.baseWavePacket = &wallpaper.WavePacket{
		Terms: []*formula.EisensteinFormulaTerm{
			{
				PowerN: 1,
				PowerM: -2,
			},
		},
		Multiplier: complex(1, 0),
	}
	suite.wallpaperMultiplier = complex(1, 0)
}

func (suite *RhombicCreatedWithDesiredSymmetry) TestCreateWallpaperWithCm(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Rhombic,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0.5,
			Height: 1,
		},
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacket,
		},
		DesiredSymmetry: wallpaper.Cm,
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.WavePackets, HasLen, 2)
	checker.Assert(newFormula.WavePackets[0].Terms, HasLen, 2)

	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerN, Equals, suite.baseWavePacket.Terms[0].PowerM)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerM, Equals, suite.baseWavePacket.Terms[0].PowerN)

	checker.Assert(newFormula.HasSymmetry(wallpaper.Cm), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Cmm), Equals, false)
}

func (suite *RhombicCreatedWithDesiredSymmetry) TestCreateWallpaperWithCmm(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Rhombic,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0.5,
			Height: 1,
		},
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacket,
		},
		DesiredSymmetry: wallpaper.Cmm,
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.WavePackets, HasLen, 4)
	checker.Assert(newFormula.WavePackets[0].Terms, HasLen, 2)

	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerN, Equals, suite.baseWavePacket.Terms[0].PowerN*-1)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerM, Equals, suite.baseWavePacket.Terms[0].PowerM*-1)
	checker.Assert(newFormula.WavePackets[2].Terms[0].PowerN, Equals, suite.baseWavePacket.Terms[0].PowerM)
	checker.Assert(newFormula.WavePackets[2].Terms[0].PowerM, Equals, suite.baseWavePacket.Terms[0].PowerN)
	checker.Assert(newFormula.WavePackets[3].Terms[0].PowerN, Equals, suite.baseWavePacket.Terms[0].PowerM*-1)
	checker.Assert(newFormula.WavePackets[3].Terms[0].PowerM, Equals, suite.baseWavePacket.Terms[0].PowerN*-1)

	checker.Assert(newFormula.HasSymmetry(wallpaper.Cm), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.Cmm), Equals, true)
}
