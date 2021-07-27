package wallpaper_test

import (
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/wallpaper"
	"wallpaper/entities/utility"
)

type SquareWallpaper struct {
	newFormula *wallpaper.Formula
}

var _ = Suite(&SquareWallpaper{})

func (suite *SquareWallpaper) SetUpTest (checker *C) {
	suite.newFormula = &wallpaper.Formula{
		LatticeType:     wallpaper.Square,
		LatticeSize:     nil,
		Lattice:         nil,
		Multiplier:      complex(1, 0),
		WavePackets:     []*wallpaper.WavePacket{
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

func (suite *SquareWallpaper) TestSetupCreatesLatticeVectors (checker *C) {
	checker.Assert(real(suite.newFormula.Lattice.XLatticeVector), utility.NumericallyCloseEnough{}, 1, 1e-6)
	checker.Assert(imag(suite.newFormula.Lattice.XLatticeVector), utility.NumericallyCloseEnough{}, 0, 1e-6)

	checker.Assert(real(suite.newFormula.Lattice.YLatticeVector), utility.NumericallyCloseEnough{}, 0, 1e-6)
	checker.Assert(imag(suite.newFormula.Lattice.YLatticeVector), utility.NumericallyCloseEnough{}, 1, 1e-6)
}

func (suite *SquareWallpaper) TestSetupAddsLockedPairs (checker *C) {
	checker.Assert(suite.newFormula.WavePackets[0].Terms, HasLen, 4)
	checker.Assert(suite.newFormula.WavePackets[0].Terms[1].PowerN, Equals, suite.newFormula.WavePackets[0].Terms[0].PowerM)
	checker.Assert(suite.newFormula.WavePackets[0].Terms[1].PowerM, Equals, suite.newFormula.WavePackets[0].Terms[0].PowerN * -1)

	checker.Assert(suite.newFormula.WavePackets[0].Terms[2].PowerN, Equals, suite.newFormula.WavePackets[0].Terms[0].PowerN * -1)
	checker.Assert(suite.newFormula.WavePackets[0].Terms[2].PowerM, Equals, suite.newFormula.WavePackets[0].Terms[0].PowerM * -1)

	checker.Assert(suite.newFormula.WavePackets[0].Terms[3].PowerN, Equals, suite.newFormula.WavePackets[0].Terms[0].PowerM * -1)
	checker.Assert(suite.newFormula.WavePackets[0].Terms[3].PowerM, Equals, suite.newFormula.WavePackets[0].Terms[0].PowerN)
}

func (suite *SquareWallpaper) TestCalculationOfPoints (checker *C) {
	calculation := suite.newFormula.Calculate(complex(2, 0.5))
	total := calculation.Total

	expectedAnswer :=
		(
			cmplx.Exp(complex(0, 2 * math.Pi)) +
				cmplx.Exp(complex(0, 2 * math.Pi * -3.5)) +
				cmplx.Exp(complex(0, 2 * math.Pi * -1)) +
				cmplx.Exp(complex(0, 2 * math.Pi * 4.5)))/4

	checker.Assert(real(total), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(total), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)

}

type SquareWallpaperHasSymmetryTest struct {
	baseWavePacket *wallpaper.WavePacket // TODO delete this
	wallpaperMultiplier complex128
}

var _ = Suite(&SquareWallpaperHasSymmetryTest{})

func (suite *SquareWallpaperHasSymmetryTest) SetUpTest(checker *C) {
	suite.wallpaperMultiplier = complex(1, 0)
}

func (suite *SquareWallpaperHasSymmetryTest) TestP4mSymmetryDetectedAcrossSinglePair(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType:     wallpaper.Square,
		LatticeSize:     nil,
		Lattice:         nil,
		Multiplier:      complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			{
				Terms: []*formula.EisensteinFormulaTerm{
					{
						PowerN: -2,
						PowerM: 1,
					},
				},
				Multiplier: complex(1, 0),
			},
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
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.HasSymmetry(wallpaper.P4), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.P4m), Equals, true)
}

func (suite *SquareWallpaperHasSymmetryTest) TestP4mSymmetryDetectedAcrossMultiplePairs(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType:     wallpaper.Square,
		LatticeSize:     nil,
		Lattice:         nil,
		Multiplier:      complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			{
				Terms: []*formula.EisensteinFormulaTerm{
					{
						PowerN: -5,
						PowerM: 8,
					},
				},
				Multiplier: complex(1, 0),
			},
			{
				Terms: []*formula.EisensteinFormulaTerm{
					{
						PowerN: 2,
						PowerM: -1,
					},
				},
				Multiplier: complex(1, 0),
			},
			{
				Terms: []*formula.EisensteinFormulaTerm{
					{
						PowerN: 8,
						PowerM: -5,
					},
				},
				Multiplier: complex(1, 0),
			},
			{
				Terms: []*formula.EisensteinFormulaTerm{
					{
						PowerN: -1,
						PowerM: 2,
					},
				},
				Multiplier: complex(1, 0),
			},
		},
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.HasSymmetry(wallpaper.P4), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.P4m), Equals, true)
}

func (suite *SquareWallpaperHasSymmetryTest) TestP4SymmetryIsAlwaysTrueForSquarePatterns(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType:     wallpaper.Square,
		LatticeSize:     nil,
		Lattice:         nil,
		Multiplier:      complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			{
				Terms: []*formula.EisensteinFormulaTerm{
					{
						PowerN: -2,
						PowerM: 1,
					},
				},
				Multiplier: complex(1, 0),
			},
		},
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.HasSymmetry(wallpaper.P4), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.P4m), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.P4g), Equals, false)
}

func (suite *SquareWallpaperHasSymmetryTest) TestP4g (checker *C) {
	p4gOddSum := wallpaper.Formula{
		LatticeType:     wallpaper.Square,
		LatticeSize:     nil,
		Lattice:         nil,
		Multiplier:      complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			{
				Terms: []*formula.EisensteinFormulaTerm{
					{
						PowerN: -2,
						PowerM: 1,
					},
				},
				Multiplier: complex(1, 0),
			},
			{
				Terms: []*formula.EisensteinFormulaTerm{
					{
						PowerN: 1,
						PowerM: -2,
					},
				},
				Multiplier: complex(-1, 0),
			},
		},
	}
	p4gOddSum.Setup()

	checker.Assert(p4gOddSum.HasSymmetry(wallpaper.P4g), Equals, true)

	p4gEvenSum := wallpaper.Formula{
		LatticeType:     wallpaper.Square,
		LatticeSize:     nil,
		Lattice:         nil,
		Multiplier:      complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			{
				Terms: []*formula.EisensteinFormulaTerm{
					{
						PowerN: -3,
						PowerM: 1,
					},
				},
				Multiplier: complex(1, 0),
			},
			{
				Terms: []*formula.EisensteinFormulaTerm{
					{
						PowerN: 1,
						PowerM: -3,
					},
				},
				Multiplier: complex(1, 0),
			},
		},
	}
	p4gEvenSum.Setup()

	checker.Assert(p4gEvenSum.HasSymmetry(wallpaper.P4g), Equals, true)
}

type SquareCreatedWithDesiredSymmetry struct {
	baseWavePacketWithOddSumFormula  *wallpaper.WavePacket
	baseWavePacketWithEvenSumFormula *wallpaper.WavePacket
	wallpaperMultiplier              complex128
}

var _ = Suite(&SquareCreatedWithDesiredSymmetry{})

func (suite *SquareCreatedWithDesiredSymmetry) SetUpTest (checker *C) {
	suite.baseWavePacketWithOddSumFormula = &wallpaper.WavePacket{
		Terms: []*formula.EisensteinFormulaTerm{
			{
				PowerN: 1,
				PowerM: -2,
			},
		},
		Multiplier: complex(1, 0),
	}
	suite.baseWavePacketWithEvenSumFormula = &wallpaper.WavePacket{
		Terms: []*formula.EisensteinFormulaTerm{
			{
				PowerM: 3,
				PowerN: -1,
			},
		},
		Multiplier: complex(1, 0),
	}
	suite.wallpaperMultiplier = complex(1, 0)
}

func (suite *SquareCreatedWithDesiredSymmetry) TestCreateWallpaperWithP4m(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType:     wallpaper.Square,
		LatticeSize:     &wallpaper.Dimensions{
			Width:  0.5,
			Height: 1,
		},
		Multiplier:      complex(2, 0),
		WavePackets:     []*wallpaper.WavePacket{
			suite.baseWavePacketWithOddSumFormula,
		},
		DesiredSymmetry: wallpaper.P4m,
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)
	checker.Assert(newFormula.WavePackets, HasLen, 2)
	checker.Assert(newFormula.WavePackets[0].Terms, HasLen, 4)

	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerN, Equals, -2)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerM, Equals, 1)

	checker.Assert(newFormula.HasSymmetry(wallpaper.P4), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.P4m), Equals, true)
}

func (suite *SquareCreatedWithDesiredSymmetry) TestCreateWallpaperWithP4gAndOddSumPowers(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType:     wallpaper.Square,
		LatticeSize:     &wallpaper.Dimensions{
			Width:  0.5,
			Height: 1,
		},
		Multiplier:      complex(2, 0),
		WavePackets:     []*wallpaper.WavePacket{
			suite.baseWavePacketWithOddSumFormula,
		},
		DesiredSymmetry: wallpaper.P4g,
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)
	checker.Assert(newFormula.WavePackets, HasLen, 2)
	checker.Assert(newFormula.WavePackets[0].Terms, HasLen, 4)

	checker.Assert(real(newFormula.WavePackets[1].Multiplier), utility.NumericallyCloseEnough{}, real(suite.wallpaperMultiplier) * -1, 1e-6)
	checker.Assert(imag(newFormula.WavePackets[1].Multiplier), utility.NumericallyCloseEnough{}, imag(suite.wallpaperMultiplier) * -1, 1e-6)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerM, Equals, 1)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerN, Equals, -2)

	checker.Assert(newFormula.HasSymmetry(wallpaper.P4), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.P4g), Equals, true)
}

func (suite *SquareCreatedWithDesiredSymmetry) TestCreateWallpaperWithP4gAndEvenSumPowers(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType:     wallpaper.Square,
		LatticeSize:     &wallpaper.Dimensions{
			Width:  0.5,
			Height: 1,
		},
		Multiplier:      complex(2, 0),
		WavePackets:     []*wallpaper.WavePacket{
			suite.baseWavePacketWithEvenSumFormula,
		},
		DesiredSymmetry: wallpaper.P4g,
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.WavePackets, HasLen, 2)
	checker.Assert(newFormula.WavePackets[0].Terms, HasLen, 4)

	checker.Assert(real(newFormula.WavePackets[1].Multiplier), utility.NumericallyCloseEnough{}, real(suite.wallpaperMultiplier), 1e-6)
	checker.Assert(imag(newFormula.WavePackets[1].Multiplier), utility.NumericallyCloseEnough{}, imag(suite.wallpaperMultiplier), 1e-6)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerM, Equals, -1)
	checker.Assert(newFormula.WavePackets[1].Terms[0].PowerN, Equals, 3)

	checker.Assert(newFormula.HasSymmetry(wallpaper.P4), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.P4m), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.P4g), Equals, true)
}