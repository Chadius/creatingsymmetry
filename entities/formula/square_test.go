package formula_test

import (
	"github.com/Chadius/creating-symmetry/entities/formula"
	"github.com/Chadius/creating-symmetry/entities/utility"
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"

	/*	"math"
	"math/cmplx"
*/)

type SquareWallpaper struct {
	newFormula formula.Arbitrary
}

var _ = Suite(&SquareWallpaper{})

func (suite *SquareWallpaper) SetUpTest(checker *C) {
	suite.newFormula, _ = formula.NewBuilder().
		Square().
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

func (suite *SquareWallpaper) TestSetupCreatesLatticeVectors(checker *C) {
	checker.Assert(real(suite.newFormula.LatticeVectors()[0]), utility.NumericallyCloseEnough{}, 1, 1e-6)
	checker.Assert(imag(suite.newFormula.LatticeVectors()[0]), utility.NumericallyCloseEnough{}, 0, 1e-6)

	checker.Assert(real(suite.newFormula.LatticeVectors()[1]), utility.NumericallyCloseEnough{}, 0, 1e-6)
	checker.Assert(imag(suite.newFormula.LatticeVectors()[1]), utility.NumericallyCloseEnough{}, 1, 1e-6)
}

func (suite *SquareWallpaper) TestSetupAddsLockedPairs(checker *C) {
	checker.Assert(suite.newFormula.WavePackets()[0].Terms(), HasLen, 4)
	checker.Assert(suite.newFormula.WavePackets()[0].Terms()[1].PowerN, Equals, suite.newFormula.WavePackets()[0].Terms()[0].PowerM)
	checker.Assert(suite.newFormula.WavePackets()[0].Terms()[1].PowerM, Equals, suite.newFormula.WavePackets()[0].Terms()[0].PowerN*-1)
	checker.Assert(suite.newFormula.WavePackets()[0].Terms()[2].PowerN, Equals, suite.newFormula.WavePackets()[0].Terms()[0].PowerN*-1)
	checker.Assert(suite.newFormula.WavePackets()[0].Terms()[2].PowerM, Equals, suite.newFormula.WavePackets()[0].Terms()[0].PowerM*-1)
	checker.Assert(suite.newFormula.WavePackets()[0].Terms()[3].PowerN, Equals, suite.newFormula.WavePackets()[0].Terms()[0].PowerM*-1)
	checker.Assert(suite.newFormula.WavePackets()[0].Terms()[3].PowerM, Equals, suite.newFormula.WavePackets()[0].Terms()[0].PowerN)
}

func (suite *SquareWallpaper) TestCalculationOfPoints(checker *C) {
	calculation := suite.newFormula.Calculate(complex(2, 0.5))

	expectedAnswer :=
		(cmplx.Exp(complex(0, 2*math.Pi)) +
			cmplx.Exp(complex(0, 2*math.Pi*-3.5)) +
			cmplx.Exp(complex(0, 2*math.Pi*-1)) +
			cmplx.Exp(complex(0, 2*math.Pi*4.5))) / 4

	checker.Assert(real(calculation), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(calculation), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}
/*
type SquareWallpaperHasSymmetryTest struct {
	wallpaperMultiplier complex128
}

var _ = Suite(&SquareWallpaperHasSymmetryTest{})

func (suite *SquareWallpaperHasSymmetryTest) SetUpTest(checker *C) {
	suite.wallpaperMultiplier = complex(1, 0)
}

func (suite *SquareWallpaperHasSymmetryTest) TestP4mSymmetryDetectedAcrossSinglePair(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Square,
		LatticeSize: nil,
		Lattice:     nil,
		multiplier:  complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			{
				terms: []*eisenstien.EisensteinFormulaTerm{
					{
						PowerN: -2,
						PowerM: 1,
					},
				},
				multiplier: complex(1, 0),
			},
			{
				terms: []*eisenstien.EisensteinFormulaTerm{
					{
						PowerN: 1,
						PowerM: -2,
					},
				},
				multiplier: complex(1, 0),
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
		LatticeType: wallpaper.Square,
		LatticeSize: nil,
		Lattice:     nil,
		multiplier:  complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			{
				terms: []*eisenstien.EisensteinFormulaTerm{
					{
						PowerN: -5,
						PowerM: 8,
					},
				},
				multiplier: complex(1, 0),
			},
			{
				terms: []*eisenstien.EisensteinFormulaTerm{
					{
						PowerN: 2,
						PowerM: -1,
					},
				},
				multiplier: complex(1, 0),
			},
			{
				terms: []*eisenstien.EisensteinFormulaTerm{
					{
						PowerN: 8,
						PowerM: -5,
					},
				},
				multiplier: complex(1, 0),
			},
			{
				terms: []*eisenstien.EisensteinFormulaTerm{
					{
						PowerN: -1,
						PowerM: 2,
					},
				},
				multiplier: complex(1, 0),
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
		LatticeType: wallpaper.Square,
		LatticeSize: nil,
		Lattice:     nil,
		multiplier:  complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			{
				terms: []*eisenstien.EisensteinFormulaTerm{
					{
						PowerN: -2,
						PowerM: 1,
					},
				},
				multiplier: complex(1, 0),
			},
		},
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.HasSymmetry(wallpaper.P4), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.P4m), Equals, false)
	checker.Assert(newFormula.HasSymmetry(wallpaper.P4g), Equals, false)
}

func (suite *SquareWallpaperHasSymmetryTest) TestP4g(checker *C) {
	p4gOddSum := wallpaper.Formula{
		LatticeType: wallpaper.Square,
		LatticeSize: nil,
		Lattice:     nil,
		multiplier:  complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			{
				terms: []*eisenstien.EisensteinFormulaTerm{
					{
						PowerN: -2,
						PowerM: 1,
					},
				},
				multiplier: complex(1, 0),
			},
			{
				terms: []*eisenstien.EisensteinFormulaTerm{
					{
						PowerN: 1,
						PowerM: -2,
					},
				},
				multiplier: complex(-1, 0),
			},
		},
	}
	p4gOddSum.Setup()

	checker.Assert(p4gOddSum.HasSymmetry(wallpaper.P4g), Equals, true)

	p4gEvenSum := wallpaper.Formula{
		LatticeType: wallpaper.Square,
		LatticeSize: nil,
		Lattice:     nil,
		multiplier:  complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			{
				terms: []*eisenstien.EisensteinFormulaTerm{
					{
						PowerN: -3,
						PowerM: 1,
					},
				},
				multiplier: complex(1, 0),
			},
			{
				terms: []*eisenstien.EisensteinFormulaTerm{
					{
						PowerN: 1,
						PowerM: -3,
					},
				},
				multiplier: complex(1, 0),
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

func (suite *SquareCreatedWithDesiredSymmetry) SetUpTest(checker *C) {
	suite.baseWavePacketWithOddSumFormula = &wallpaper.WavePacket{
		terms: []*eisenstien.EisensteinFormulaTerm{
			{
				PowerN: 1,
				PowerM: -2,
			},
		},
		multiplier: complex(1, 0),
	}
	suite.baseWavePacketWithEvenSumFormula = &wallpaper.WavePacket{
		terms: []*eisenstien.EisensteinFormulaTerm{
			{
				PowerM: 3,
				PowerN: -1,
			},
		},
		multiplier: complex(1, 0),
	}
	suite.wallpaperMultiplier = complex(1, 0)
}

func (suite *SquareCreatedWithDesiredSymmetry) TestCreateWallpaperWithP4m(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Square,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0.5,
			Height: 1,
		},
		multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacketWithOddSumFormula,
		},
		DesiredSymmetry: wallpaper.P4m,
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)
	checker.Assert(newFormula.WavePackets, HasLen, 2)
	checker.Assert(newFormula.WavePackets[0].terms, HasLen, 4)

	checker.Assert(newFormula.WavePackets[1].terms[0].PowerN, Equals, -2)
	checker.Assert(newFormula.WavePackets[1].terms[0].PowerM, Equals, 1)

	checker.Assert(newFormula.HasSymmetry(wallpaper.P4), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.P4m), Equals, true)
}

func (suite *SquareCreatedWithDesiredSymmetry) TestCreateWallpaperWithP4gAndOddSumPowers(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Square,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0.5,
			Height: 1,
		},
		multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacketWithOddSumFormula,
		},
		DesiredSymmetry: wallpaper.P4g,
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)
	checker.Assert(newFormula.WavePackets, HasLen, 2)
	checker.Assert(newFormula.WavePackets[0].terms, HasLen, 4)

	checker.Assert(real(newFormula.WavePackets[1].multiplier), utility.NumericallyCloseEnough{}, real(suite.wallpaperMultiplier)*-1, 1e-6)
	checker.Assert(imag(newFormula.WavePackets[1].multiplier), utility.NumericallyCloseEnough{}, imag(suite.wallpaperMultiplier)*-1, 1e-6)
	checker.Assert(newFormula.WavePackets[1].terms[0].PowerM, Equals, 1)
	checker.Assert(newFormula.WavePackets[1].terms[0].PowerN, Equals, -2)

	checker.Assert(newFormula.HasSymmetry(wallpaper.P4), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.P4g), Equals, true)
}

func (suite *SquareCreatedWithDesiredSymmetry) TestCreateWallpaperWithP4gAndEvenSumPowers(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Square,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0.5,
			Height: 1,
		},
		multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			suite.baseWavePacketWithEvenSumFormula,
		},
		DesiredSymmetry: wallpaper.P4g,
	}
	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.WavePackets, HasLen, 2)
	checker.Assert(newFormula.WavePackets[0].terms, HasLen, 4)

	checker.Assert(real(newFormula.WavePackets[1].multiplier), utility.NumericallyCloseEnough{}, real(suite.wallpaperMultiplier), 1e-6)
	checker.Assert(imag(newFormula.WavePackets[1].multiplier), utility.NumericallyCloseEnough{}, imag(suite.wallpaperMultiplier), 1e-6)
	checker.Assert(newFormula.WavePackets[1].terms[0].PowerM, Equals, -1)
	checker.Assert(newFormula.WavePackets[1].terms[0].PowerN, Equals, 3)

	checker.Assert(newFormula.HasSymmetry(wallpaper.P4), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.P4m), Equals, true)
	checker.Assert(newFormula.HasSymmetry(wallpaper.P4g), Equals, true)
}
*/