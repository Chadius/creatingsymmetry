package wallpaper_test

import (
	"github.com/Chadius/creating-symmetry/entities/formula"
	"github.com/Chadius/creating-symmetry/entities/formula/wallpaper"
	"github.com/Chadius/creating-symmetry/entities/utility"
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"
)

type GenericWallpaper struct {
	newFormula *wallpaper.Formula
}

var _ = Suite(&GenericWallpaper{})

func (suite *GenericWallpaper) SetUpTest(checker *C) {
	suite.newFormula = &wallpaper.Formula{
		LatticeType: wallpaper.Generic,
		LatticeSize: &wallpaper.Dimensions{
			Width:  2,
			Height: -0.5,
		},
		Lattice:    nil,
		Multiplier: complex(1, 0),
		WavePackets: []*wallpaper.WavePacket{
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
		DesiredSymmetry: wallpaper.P1,
	}

	suite.newFormula.Setup()
}

func (suite *GenericWallpaper) TestSetupCreatesLatticeVectors(checker *C) {
	checker.Assert(real(suite.newFormula.Lattice.XLatticeVector), utility.NumericallyCloseEnough{}, 1, 1e-6)
	checker.Assert(imag(suite.newFormula.Lattice.XLatticeVector), utility.NumericallyCloseEnough{}, 0, 1e-6)

	checker.Assert(real(suite.newFormula.Lattice.YLatticeVector), utility.NumericallyCloseEnough{}, 2, 1e-6)
	checker.Assert(imag(suite.newFormula.Lattice.YLatticeVector), utility.NumericallyCloseEnough{}, -0.5, 1e-6)
}

func (suite *GenericWallpaper) TestSetupDoesNotAddLockedPairs(checker *C) {
	checker.Assert(suite.newFormula.WavePackets[0].Terms, HasLen, 1)
}

func (suite *GenericWallpaper) TestCalculationOfPoints(checker *C) {
	calculation := suite.newFormula.Calculate(complex(1.5, 10))
	total := calculation.Total

	expectedAnswer := cmplx.Exp(complex(0, math.Pi))

	checker.Assert(real(total), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(total), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}

type GenericWallpaperDesiredSymmetryTest struct {
	eisensteinTerm      []*formula.EisensteinFormulaTerm
	wallpaperMultiplier complex128
	latticeSize         complex128
}

var _ = Suite(&GenericWallpaperDesiredSymmetryTest{})

func (suite *GenericWallpaperDesiredSymmetryTest) SetUpTest(checker *C) {
	suite.eisensteinTerm = []*formula.EisensteinFormulaTerm{
		{
			PowerN: 8,
			PowerM: -3,
		},
	}

	suite.wallpaperMultiplier = complex(1, 0)
	suite.latticeSize = complex(2.0, 1.5)
}

func (suite *GenericWallpaperDesiredSymmetryTest) TestCreateGenericWithP1(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Generic,
		LatticeSize: &wallpaper.Dimensions{
			Width:  real(suite.latticeSize),
			Height: imag(suite.latticeSize),
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			{
				Multiplier: complex(1, 0),
				Terms: []*formula.EisensteinFormulaTerm{
					suite.eisensteinTerm[0],
				},
			},
		},
		DesiredSymmetry: wallpaper.P1,
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.Lattice.YLatticeVector, Equals, suite.latticeSize)

	checker.Assert(newFormula.WavePackets, HasLen, 1)

	checker.Assert(newFormula.WavePackets[0].Terms, HasLen, 1)
	checker.Assert(newFormula.WavePackets[0].Terms[0].PowerN, Equals, suite.eisensteinTerm[0].PowerN)
	checker.Assert(newFormula.WavePackets[0].Terms[0].PowerM, Equals, suite.eisensteinTerm[0].PowerM)
}

func (suite *GenericWallpaperDesiredSymmetryTest) TestCreateGenericWithP2(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Generic,
		LatticeSize: &wallpaper.Dimensions{
			Width:  real(suite.latticeSize),
			Height: imag(suite.latticeSize),
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			{
				Multiplier: complex(1, 0),
				Terms: []*formula.EisensteinFormulaTerm{
					suite.eisensteinTerm[0],
				},
			},
		},
		DesiredSymmetry: wallpaper.P2,
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.Lattice.YLatticeVector, Equals, suite.latticeSize)

	checker.Assert(newFormula.WavePackets, HasLen, 2)

	checker.Assert(newFormula.WavePackets[1].Terms, HasLen, 1)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerN, Equals, suite.eisensteinTerm[0].PowerN*-1)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerM, Equals, suite.eisensteinTerm[0].PowerM*-1)
}

type GenericWaveSymmetry struct {
	baseWavePacket *wallpaper.WavePacket
}

var _ = Suite(&GenericWaveSymmetry{})

func (suite *GenericWaveSymmetry) SetUpTest(checker *C) {
	suite.baseWavePacket = &wallpaper.WavePacket{
		Terms: []*formula.EisensteinFormulaTerm{
			{
				PowerN: 8,
				PowerM: -3,
			},
		},
		Multiplier: complex(1, 0),
	}
}

func (suite *GenericWaveSymmetry) TestOnlyP1SymmetryFound(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Generic,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0.5,
			Height: 2.4,
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacket,
		},
		DesiredSymmetry: wallpaper.P1,
	}
	newFormula.Setup()
	checker.Assert(newFormula.HasSymmetry(wallpaper.P1), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.P2), Equals, false)
}

func (suite *GenericWaveSymmetry) TestP2(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Generic,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0.5,
			Height: 2.4,
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacket,
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
		DesiredSymmetry: wallpaper.P1,
	}
	newFormula.Setup()
	checker.Assert(newFormula.HasSymmetry(wallpaper.P1), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.P2), Equals, true)

}
