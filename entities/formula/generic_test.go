package formula_test

import (
	"github.com/Chadius/creating-symmetry/entities/formula"
	"github.com/Chadius/creating-symmetry/entities/utility"
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"
	"reflect"
)

type GenericWallpaper struct {
	newFormula formula.Arbitrary
}

var _ = Suite(&GenericWallpaper{})

func (suite *GenericWallpaper) SetUpTest(checker *C) {
	suite.newFormula, _ = formula.NewBuilder().
		Generic().
		LatticeWidth(2).
		LatticeHeight(-0.5).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1,0)).
				AddTerm(
					formula.NewTermBuilder().PowerN(3).PowerM(-4).Build(),
				).
				Build(),
		).
		Build()
}

func (suite *GenericWallpaper) TestSetupCreatesLatticeVectors(checker *C) {
	checker.Assert(real(suite.newFormula.LatticeVectors()[0]), utility.NumericallyCloseEnough{}, 1, 1e-6)
	checker.Assert(imag(suite.newFormula.LatticeVectors()[0]), utility.NumericallyCloseEnough{}, 0, 1e-6)

	checker.Assert(real(suite.newFormula.LatticeVectors()[1]), utility.NumericallyCloseEnough{}, 2, 1e-6)
	checker.Assert(imag(suite.newFormula.LatticeVectors()[1]), utility.NumericallyCloseEnough{}, -0.5, 1e-6)
}

func (suite *GenericWallpaper) TestSetupDoesNotAddLockedPairs(checker *C) {
	checker.Assert(suite.newFormula.WavePackets()[0].Terms(), HasLen, 1)
}

func (suite *GenericWallpaper) TestCalculationOfPoints(checker *C) {
	calculation := suite.newFormula.Calculate(complex(1.5, 10))

	expectedAnswer := cmplx.Exp(complex(0, math.Pi))
	checker.Assert(real(calculation), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(calculation), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}

type GenericWallpaperDesiredSymmetryTest struct {
	eisensteinTerm      []formula.Term
	wallpaperMultiplier complex128
	latticeSize         complex128
}

var _ = Suite(&GenericWallpaperDesiredSymmetryTest{})

func (suite *GenericWallpaperDesiredSymmetryTest) SetUpTest(checker *C) {
	suite.eisensteinTerm = []formula.Term{
		*formula.NewTermBuilder().PowerN(8).PowerM(-3).Multiplier(complex(1, 0)).Build(),
	}
	suite.wallpaperMultiplier = complex(1, 0)
	suite.latticeSize = complex(2.0, 1.5)
}

func (suite *GenericWallpaperDesiredSymmetryTest) TestCreateGenericWithP1(checker *C) {
	newFormula, err := formula.NewBuilder().
		Generic().
		LatticeWidth(real(suite.latticeSize)).
		LatticeHeight(imag(suite.latticeSize)).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1,0)).
				AddTerm(&suite.eisensteinTerm[0]).
				Build(),
		).
		Build()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.LatticeVectors()[1], Equals, suite.latticeSize)

	checker.Assert(newFormula.WavePackets(), HasLen, 1)
	checker.Assert(newFormula.WavePackets()[0].Terms(), HasLen, 1)
	checker.Assert(newFormula.WavePackets()[0].Terms()[0].PowerN, Equals, suite.eisensteinTerm[0].PowerN)
	checker.Assert(newFormula.WavePackets()[0].Terms()[0].PowerM, Equals, suite.eisensteinTerm[0].PowerM)
}

func (suite *GenericWallpaperDesiredSymmetryTest) TestCreateGenericWithP2(checker *C) {
	newFormula, err := formula.NewBuilder().
		Generic().
		LatticeWidth(real(suite.latticeSize)).
		LatticeHeight(imag(suite.latticeSize)).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1,0)).
				AddTerm(&suite.eisensteinTerm[0]).
				Build(),
		).
		DesiredSymmetry(formula.P2).
		Build()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.LatticeVectors()[1], Equals, suite.latticeSize)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.P2)

	checker.Assert(newFormula.WavePackets(), HasLen, 2)
	checker.Assert(newFormula.WavePackets()[0].Terms(), HasLen, 1)
	checker.Assert(newFormula.WavePackets()[0].Terms()[0].PowerN, Equals, suite.eisensteinTerm[0].PowerN)
	checker.Assert(newFormula.WavePackets()[0].Terms()[0].PowerM, Equals, suite.eisensteinTerm[0].PowerM)
	checker.Assert(newFormula.WavePackets()[1].Terms(), HasLen, 1)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerN, Equals, suite.eisensteinTerm[0].PowerN * -1)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerM, Equals, suite.eisensteinTerm[0].PowerM * -1)
}

func (suite *GenericWallpaperDesiredSymmetryTest)TestWhenOtherDesiredSymmetry_BuilderReturnsError(checker *C) {
	newFormula, err := formula.NewBuilder().
		Generic().
		LatticeWidth(real(suite.latticeSize)).
		LatticeHeight(imag(suite.latticeSize)).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1,0)).
				AddTerm(&suite.eisensteinTerm[0]).
				Build(),
		).
		DesiredSymmetry(formula.P3).
		Build()

	checker.Assert(err, ErrorMatches, "generic lattice can apply these desired symmetries: P1, P2")
	checker.Assert(reflect.TypeOf(newFormula).String(), Equals, "*formula.Identity")
}

type GenericWaveSymmetry struct {
	baseWavePacket *formula.WavePacket
}

var _ = Suite(&GenericWaveSymmetry{})

func (suite *GenericWaveSymmetry) SetUpTest(checker *C) {
	suite.baseWavePacket = formula.NewWavePacketBuilder().
		AddTerm(formula.NewTermWithMultiplierAndPowers(complex(1, 0), 8, -3)).
		Build()
}

func (suite *GenericWaveSymmetry) TestOnlyP1SymmetryFound(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Generic().
		LatticeWidth(0.5).
		LatticeHeight(2.4).
		AddWavePacket(suite.baseWavePacket).
		DesiredSymmetry(formula.P1).
		Build()

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 1)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
}

func (suite *GenericWaveSymmetry) TestP2(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Generic().
		LatticeWidth(0.5).
		LatticeHeight(2.4).
		AddWavePacket(suite.baseWavePacket).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacket.Multiplier()).
				AddTerm(formula.NewTermWithMultiplierAndPowers(complex(1,0), suite.baseWavePacket.Terms()[0].PowerN * -1, suite.baseWavePacket.Terms()[0].PowerM * -1)).
				Build(),
		).
		DesiredSymmetry(formula.P1).
		Build()

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 2)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.P2)
}
