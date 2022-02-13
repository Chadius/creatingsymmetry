package formula_test

import (
	"github.com/chadius/creatingsymmetry/entities/formula"
	"github.com/chadius/creatingsymmetry/entities/utility"
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"
	"reflect"
)

type SquareWallpaper struct {
	newFormula formula.Arbitrary
}

var _ = Suite(&SquareWallpaper{})

func (suite *SquareWallpaper) SetUpTest(checker *C) {
	suite.newFormula, _ = formula.NewBuilder().
		Square().
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
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

type SquareWallpaperHasSymmetryTest struct {
	wallpaperMultiplier complex128
}

var _ = Suite(&SquareWallpaperHasSymmetryTest{})

func (suite *SquareWallpaperHasSymmetryTest) SetUpTest(checker *C) {
	suite.wallpaperMultiplier = complex(1, 0)
}

func (suite *SquareWallpaperHasSymmetryTest) TestP4mSymmetryDetectedAcrossSinglePair(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Square().
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						-2,
						1,
					),
				).
				Build(),
		).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						1,
						-2,
					),
				).
				Build(),
		).
		Build()

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 3)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.P4)
	checker.Assert(newFormula.SymmetriesFound()[2], Equals, formula.P4m)
}

func (suite *SquareWallpaperHasSymmetryTest) TestP4mSymmetryDetectedAcrossMultiplePairs(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Square().
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						-5,
						8,
					),
				).
				Build(),
		).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						2,
						-1,
					),
				).
				Build(),
		).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						8,
						-5,
					),
				).
				Build(),
		).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						-1,
						2,
					),
				).
				Build(),
		).
		Build()

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 3)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.P4)
	checker.Assert(newFormula.SymmetriesFound()[2], Equals, formula.P4m)
}

func (suite *SquareWallpaperHasSymmetryTest) TestP4SymmetryIsAlwaysTrueForSquarePatterns(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Square().
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						-2,
						1,
					),
				).
				Build(),
		).
		Build()

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 2)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.P4)
}

func (suite *SquareWallpaperHasSymmetryTest) TestP4g(checker *C) {
	p4gOddSum, _ := formula.NewBuilder().
		Square().
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						-2,
						1,
					),
				).
				Build(),
		).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(-1, 0)).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						1,
						-2,
					),
				).
				Build(),
		).
		Build()

	checker.Assert(p4gOddSum.SymmetriesFound(), HasLen, 3)
	checker.Assert(p4gOddSum.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(p4gOddSum.SymmetriesFound()[1], Equals, formula.P4)
	checker.Assert(p4gOddSum.SymmetriesFound()[2], Equals, formula.P4g)

	p4gEvenSum, _ := formula.NewBuilder().
		Square().
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						-3,
						1,
					),
				).
				Build(),
		).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						1,
						-3,
					),
				).
				Build(),
		).
		Build()

	checker.Assert(p4gEvenSum.SymmetriesFound(), HasLen, 4)
	checker.Assert(p4gEvenSum.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(p4gEvenSum.SymmetriesFound()[1], Equals, formula.P4)
	checker.Assert(p4gEvenSum.SymmetriesFound()[2], Equals, formula.P4m)
	checker.Assert(p4gEvenSum.SymmetriesFound()[3], Equals, formula.P4g)
}

type SquareCreatedWithDesiredSymmetry struct {
	baseWavePacketWithOddSumFormula  *formula.WavePacket
	baseWavePacketWithEvenSumFormula *formula.WavePacket
	wallpaperMultiplier              complex128
}

var _ = Suite(&SquareCreatedWithDesiredSymmetry{})

func (suite *SquareCreatedWithDesiredSymmetry) SetUpTest(checker *C) {
	suite.baseWavePacketWithOddSumFormula = formula.NewWavePacketBuilder().
		Multiplier(complex(1, 0)).
		AddTerm(
			formula.NewTermWithMultiplierAndPowers(
				complex(1, 0),
				1,
				-2,
			),
		).
		Build()

	suite.baseWavePacketWithEvenSumFormula = formula.NewWavePacketBuilder().
		Multiplier(complex(1, 0)).
		AddTerm(
			formula.NewTermWithMultiplierAndPowers(
				complex(1, 0),
				-1,
				3,
			),
		).
		Build()

	suite.wallpaperMultiplier = complex(1, 0)
}

func (suite *SquareCreatedWithDesiredSymmetry) TestCreateWallpaperWithP4m(checker *C) {
	newFormula, err := formula.NewBuilder().
		Square().
		AddWavePacket(suite.baseWavePacketWithOddSumFormula).
		DesiredSymmetry(formula.P4m).
		Build()

	checker.Assert(err, IsNil)

	checker.Assert(newFormula.WavePackets(), HasLen, 2)
	checker.Assert(newFormula.WavePackets()[0].Terms(), HasLen, 4)

	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerN, Equals, -2)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerM, Equals, 1)

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 3)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.P4)
	checker.Assert(newFormula.SymmetriesFound()[2], Equals, formula.P4m)
}

func (suite *SquareCreatedWithDesiredSymmetry) TestCreateWallpaperWithP4gAndOddSumPowers(checker *C) {
	newFormula, err := formula.NewBuilder().
		Square().
		AddWavePacket(suite.baseWavePacketWithOddSumFormula).
		DesiredSymmetry(formula.P4g).
		Build()

	checker.Assert(err, IsNil)

	checker.Assert(newFormula.WavePackets(), HasLen, 2)
	checker.Assert(newFormula.WavePackets()[0].Terms(), HasLen, 4)

	checker.Assert(real(newFormula.WavePackets()[1].Multiplier()), utility.NumericallyCloseEnough{}, real(suite.wallpaperMultiplier)*-1, 1e-6)
	checker.Assert(imag(newFormula.WavePackets()[1].Multiplier()), utility.NumericallyCloseEnough{}, imag(suite.wallpaperMultiplier)*-1, 1e-6)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerN, Equals, -2)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerM, Equals, 1)

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 3)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.P4)
	checker.Assert(newFormula.SymmetriesFound()[2], Equals, formula.P4g)
}

func (suite *SquareCreatedWithDesiredSymmetry) TestCreateWallpaperWithP4gAndEvenSumPowers(checker *C) {
	newFormula, err := formula.NewBuilder().
		Square().
		AddWavePacket(suite.baseWavePacketWithEvenSumFormula).
		DesiredSymmetry(formula.P4g).
		Build()

	checker.Assert(err, IsNil)

	checker.Assert(newFormula.WavePackets(), HasLen, 2)
	checker.Assert(newFormula.WavePackets()[0].Terms(), HasLen, 4)

	checker.Assert(real(newFormula.WavePackets()[1].Multiplier()), utility.NumericallyCloseEnough{}, real(suite.wallpaperMultiplier), 1e-6)
	checker.Assert(imag(newFormula.WavePackets()[1].Multiplier()), utility.NumericallyCloseEnough{}, imag(suite.wallpaperMultiplier), 1e-6)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerN, Equals, 3)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerM, Equals, -1)

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 4)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.P4)
	checker.Assert(newFormula.SymmetriesFound()[2], Equals, formula.P4m)
	checker.Assert(newFormula.SymmetriesFound()[3], Equals, formula.P4g)
}

func (suite *SquareCreatedWithDesiredSymmetry) TestWhenOtherDesiredSymmetry_BuilderReturnsError(checker *C) {
	newFormula, err := formula.NewBuilder().
		Square().
		AddWavePacket(suite.baseWavePacketWithEvenSumFormula).
		DesiredSymmetry(formula.P2).
		Build()

	checker.Assert(err, ErrorMatches, "square lattice can apply these desired symmetries: P1, P4, P4m, P4g")
	checker.Assert(reflect.TypeOf(newFormula).String(), Equals, "*formula.Identity")
}
