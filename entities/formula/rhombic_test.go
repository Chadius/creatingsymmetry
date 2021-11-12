package formula_test

import (
	"github.com/Chadius/creating-symmetry/entities/formula"
	"github.com/Chadius/creating-symmetry/entities/utility"
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"
	"reflect"
)

type RhombicWallpaper struct {
	newFormula formula.Arbitrary
}

var _ = Suite(&RhombicWallpaper{})

func (suite *RhombicWallpaper) SetUpTest(checker *C) {
	suite.newFormula, _ = formula.NewBuilder().
		Rhombic().
		LatticeHeight(1).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
				AddTerm(
					formula.NewTermBuilder().PowerN(1).PowerM(-2).Build(),
				).
				Build(),
		).
		Build()
}

func (suite *RhombicWallpaper) TestSetupCreatesLatticeVectors(checker *C) {
	checker.Assert(real(suite.newFormula.LatticeVectors()[0]), utility.NumericallyCloseEnough{}, 0.5, 1e-6)
	checker.Assert(imag(suite.newFormula.LatticeVectors()[0]), utility.NumericallyCloseEnough{}, 1.0, 1e-6)

	checker.Assert(real(suite.newFormula.LatticeVectors()[1]), utility.NumericallyCloseEnough{}, 0.5, 1e-6)
	checker.Assert(imag(suite.newFormula.LatticeVectors()[1]), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
}

func (suite *RhombicWallpaper) TestSetupAddsLockedPairs(checker *C) {
	checker.Assert(suite.newFormula.WavePackets()[0].Terms(), HasLen, 2)
	checker.Assert(suite.newFormula.WavePackets()[0].Terms()[1].PowerN, Equals, suite.newFormula.WavePackets()[0].Terms()[0].PowerM)
	checker.Assert(suite.newFormula.WavePackets()[0].Terms()[1].PowerM, Equals, suite.newFormula.WavePackets()[0].Terms()[0].PowerN)
}

func (suite *RhombicWallpaper) TestCalculationOfPoints(checker *C) {
	calculation := suite.newFormula.Calculate(complex(0.75, -0.25))

	expectedAnswer := (cmplx.Exp(complex(0, math.Pi*-9/4)) +
		cmplx.Exp(complex(0, math.Pi*-3/4))) / 2

	checker.Assert(real(calculation), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(calculation), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}

type RhombicWallpaperHasSymmetryTest struct {
	baseWavePacket      *formula.WavePacket
	wallpaperMultiplier complex128
}

var _ = Suite(&RhombicWallpaperHasSymmetryTest{})

func (suite *RhombicWallpaperHasSymmetryTest) SetUpTest(checker *C) {
	suite.baseWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(1, 0)).
		AddTerm(
			formula.NewTermWithMultiplierAndPowers(
				complex(1, 0),
				8,
				-3,
			),
		).
		Build()

	suite.wallpaperMultiplier = complex(1, 0)
}

func (suite *RhombicWallpaperHasSymmetryTest) TestRhombicHasNoSymmetry(checker *C) {
	newFormula, err := formula.NewBuilder().
		Rhombic().
		LatticeHeight(1.0).
		AddWavePacket(suite.baseWavePacket).
		Build()

	checker.Assert(err, IsNil)

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 1)
	checker.Assert(newFormula.SymmetriesFound(), Not(Equals), formula.Cm)
	checker.Assert(newFormula.SymmetriesFound(), Not(Equals), formula.Cmm)
}

func (suite *RhombicWallpaperHasSymmetryTest) TestRhombicMayHaveSymmetryForCm(checker *C) {
	newFormula, err := formula.NewBuilder().
		Rhombic().
		LatticeHeight(1.0).
		AddWavePacket(suite.baseWavePacket).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacket.Multiplier()).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						suite.baseWavePacket.Terms()[0].PowerM,
						suite.baseWavePacket.Terms()[0].PowerN,
					),
				).
				Build(),
		).
		Build()

	checker.Assert(err, IsNil)

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 2)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.Cm)
}

func (suite *RhombicWallpaperHasSymmetryTest) TestRhombicMayHaveSymmetryForCmm(checker *C) {
	newFormula, err := formula.NewBuilder().
		Rhombic().
		LatticeHeight(1.0).
		AddWavePacket(suite.baseWavePacket).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacket.Multiplier()).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						suite.baseWavePacket.Terms()[0].PowerM,
						suite.baseWavePacket.Terms()[0].PowerN,
					),
				).
				Build(),
		).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacket.Multiplier()).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						suite.baseWavePacket.Terms()[0].PowerM*-1,
						suite.baseWavePacket.Terms()[0].PowerN*-1,
					),
				).
				Build(),
		).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacket.Multiplier()).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						suite.baseWavePacket.Terms()[0].PowerN*-1,
						suite.baseWavePacket.Terms()[0].PowerM*-1,
					),
				).
				Build(),
		).
		Build()

	checker.Assert(err, IsNil)

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 3)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.Cm)
	checker.Assert(newFormula.SymmetriesFound()[2], Equals, formula.Cmm)
}

type RhombicCreatedWithDesiredSymmetry struct {
	baseWavePacket      *formula.WavePacket
	wallpaperMultiplier complex128
}

var _ = Suite(&RhombicCreatedWithDesiredSymmetry{})

func (suite *RhombicCreatedWithDesiredSymmetry) SetUpTest(checker *C) {
	suite.baseWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(1, 0)).
		AddTerm(
			formula.NewTermWithMultiplierAndPowers(
				complex(1, 0),
				1,
				-2,
			),
		).
		Build()
	suite.wallpaperMultiplier = complex(1, 0)
}

func (suite *RhombicCreatedWithDesiredSymmetry) TestCreateWallpaperWithCm(checker *C) {
	newFormula, err := formula.NewBuilder().
		Rhombic().
		LatticeHeight(1.0).
		AddWavePacket(suite.baseWavePacket).
		DesiredSymmetry(formula.Cm).
		Build()

	checker.Assert(err, IsNil)

	checker.Assert(newFormula.WavePackets(), HasLen, 2)
	checker.Assert(newFormula.WavePackets()[0].Terms(), HasLen, 2)

	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerN, Equals, suite.baseWavePacket.Terms()[0].PowerM)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerM, Equals, suite.baseWavePacket.Terms()[0].PowerN)
}

func (suite *RhombicCreatedWithDesiredSymmetry) TestCreateWallpaperWithCmm(checker *C) {
	newFormula, err := formula.NewBuilder().
		Rhombic().
		LatticeHeight(1.0).
		AddWavePacket(suite.baseWavePacket).
		DesiredSymmetry(formula.Cmm).
		Build()

	checker.Assert(err, IsNil)

	checker.Assert(newFormula.WavePackets(), HasLen, 4)
	checker.Assert(newFormula.WavePackets()[0].Terms(), HasLen, 2)

	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerN, Equals, suite.baseWavePacket.Terms()[0].PowerN*-1)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerM, Equals, suite.baseWavePacket.Terms()[0].PowerM*-1)
	checker.Assert(newFormula.WavePackets()[2].Terms()[0].PowerN, Equals, suite.baseWavePacket.Terms()[0].PowerM)
	checker.Assert(newFormula.WavePackets()[2].Terms()[0].PowerM, Equals, suite.baseWavePacket.Terms()[0].PowerN)
	checker.Assert(newFormula.WavePackets()[3].Terms()[0].PowerN, Equals, suite.baseWavePacket.Terms()[0].PowerM*-1)
	checker.Assert(newFormula.WavePackets()[3].Terms()[0].PowerM, Equals, suite.baseWavePacket.Terms()[0].PowerN*-1)
}

func (suite *RhombicCreatedWithDesiredSymmetry) TestWhenOtherDesiredSymmetry_BuilderReturnsError(checker *C) {
	newFormula, err := formula.NewBuilder().
		Rhombic().
		LatticeHeight(1.0).
		AddWavePacket(suite.baseWavePacket).
		DesiredSymmetry(formula.P2).
		Build()

	checker.Assert(err, ErrorMatches, "rhombic lattice can apply these desired symmetries: P1, Cm, Cmm")
	checker.Assert(reflect.TypeOf(newFormula).String(), Equals, "*formula.Identity")
}
