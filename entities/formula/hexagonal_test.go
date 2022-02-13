package formula_test

import (
	"github.com/chadius/creatingsymmetry/entities/formula"
	"github.com/chadius/creatingsymmetry/entities/utility"
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"
	"reflect"
)

type HexagonalWallpaper struct {
	newFormula formula.Arbitrary
}

var _ = Suite(&HexagonalWallpaper{})

func (suite *HexagonalWallpaper) SetUpTest(checker *C) {
	suite.newFormula, _ = formula.NewBuilder().
		Hexagonal().
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

func (suite *HexagonalWallpaper) TestSetupCreatesLatticeVectors(checker *C) {
	checker.Assert(real(suite.newFormula.LatticeVectors()[0]), utility.NumericallyCloseEnough{}, 1, 1e-6)
	checker.Assert(imag(suite.newFormula.LatticeVectors()[0]), utility.NumericallyCloseEnough{}, 0, 1e-6)

	checker.Assert(real(suite.newFormula.LatticeVectors()[1]), utility.NumericallyCloseEnough{}, -0.5, 1e-6)
	checker.Assert(imag(suite.newFormula.LatticeVectors()[1]), utility.NumericallyCloseEnough{}, math.Sqrt(3.0)/2.0, 1e-6)
}

func (suite *HexagonalWallpaper) TestSetupAddsLockedPairs(checker *C) {
	checker.Assert(suite.newFormula.WavePackets()[0].Terms(), HasLen, 3)
	checker.Assert(suite.newFormula.WavePackets()[0].Terms()[1].PowerN, Equals, -2)
	checker.Assert(suite.newFormula.WavePackets()[0].Terms()[1].PowerM, Equals, 1)
	checker.Assert(suite.newFormula.WavePackets()[0].Terms()[2].PowerN, Equals, 1)
	checker.Assert(suite.newFormula.WavePackets()[0].Terms()[2].PowerM, Equals, 1)
}

func (suite *HexagonalWallpaper) TestCalculationOfPoints(checker *C) {
	calculation := suite.newFormula.Calculate(complex(math.Sqrt(3), -1*math.Sqrt(3)))

	expectedAnswer := (cmplx.Exp(complex(0, 2*math.Pi*(3+math.Sqrt(3)))) +
		cmplx.Exp(complex(0, 2*math.Pi*(-2*math.Sqrt(3)))) +
		cmplx.Exp(complex(0, 2*math.Pi*(-3+math.Sqrt(3))))) / 3

	checker.Assert(real(calculation), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(calculation), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}

type HexagonalWallpaperHasSymmetryTest struct {
	baseWavePacket      *formula.WavePacket
	wallpaperMultiplier complex128
}

var _ = Suite(&HexagonalWallpaperHasSymmetryTest{})

func (suite *HexagonalWallpaperHasSymmetryTest) SetUpTest(checker *C) {
	suite.baseWavePacket = formula.NewWavePacketBuilder().
		AddTerm(formula.NewTermWithMultiplierAndPowers(complex(1, 0), 8, -3)).
		Build()

	suite.wallpaperMultiplier = complex(1, 0)
}

func (suite *HexagonalWallpaperHasSymmetryTest) TestHexagonalWillAlwaysHaveP3(checker *C) {
	newFormula, err := formula.NewBuilder().
		Hexagonal().
		AddWavePacket(suite.baseWavePacket).
		Build()

	checker.Assert(err, IsNil)
	checker.Assert(newFormula.WavePackets(), HasLen, 1)
	checker.Assert(newFormula.SymmetriesFound(), HasLen, 2)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.P3)
}

func (suite *HexagonalWallpaperHasSymmetryTest) TestHexagonalMayHaveSymmetryForP31m(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Hexagonal().
		AddWavePacket(suite.baseWavePacket).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacket.Multiplier()).
				AddTerm(formula.NewTermWithMultiplierAndPowers(
					complex(1, 0),
					suite.baseWavePacket.Terms()[0].PowerM,
					suite.baseWavePacket.Terms()[0].PowerN,
				),
				).
				Build(),
		).
		Build()

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 3)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.P3)
	checker.Assert(newFormula.SymmetriesFound()[2], Equals, formula.P31m)
}

func (suite *HexagonalWallpaperHasSymmetryTest) TestHexagonalMayHaveSymmetryForP3m1(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Hexagonal().
		AddWavePacket(suite.baseWavePacket).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacket.Multiplier()).
				AddTerm(formula.NewTermWithMultiplierAndPowers(
					complex(1, 0),
					suite.baseWavePacket.Terms()[0].PowerM*-1,
					suite.baseWavePacket.Terms()[0].PowerN*-1,
				),
				).
				Build(),
		).
		Build()

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 3)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.P3)
	checker.Assert(newFormula.SymmetriesFound()[2], Equals, formula.P3m1)
}

func (suite *HexagonalWallpaperHasSymmetryTest) TestHexagonalMayHaveSymmetryForP6(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Hexagonal().
		AddWavePacket(suite.baseWavePacket).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacket.Multiplier()).
				AddTerm(formula.NewTermWithMultiplierAndPowers(
					complex(1, 0),
					suite.baseWavePacket.Terms()[0].PowerN*-1,
					suite.baseWavePacket.Terms()[0].PowerM*-1,
				),
				).
				Build(),
		).
		Build()

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 3)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.P3)
	checker.Assert(newFormula.SymmetriesFound()[2], Equals, formula.P6)
}

func (suite *HexagonalWallpaperHasSymmetryTest) TestHexagonalMayHaveSymmetryForP6m(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Hexagonal().
		AddWavePacket(suite.baseWavePacket).
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
		AddWavePacket(
			formula.NewWavePacketBuilder().
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
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						suite.baseWavePacket.Terms()[0].PowerM*-1,
						suite.baseWavePacket.Terms()[0].PowerN*-1,
					),
				).
				Build(),
		).
		Build()

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 6)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.P3)
	checker.Assert(newFormula.SymmetriesFound()[2], Equals, formula.P31m)
	checker.Assert(newFormula.SymmetriesFound()[3], Equals, formula.P3m1)
	checker.Assert(newFormula.SymmetriesFound()[4], Equals, formula.P6)
	checker.Assert(newFormula.SymmetriesFound()[5], Equals, formula.P6m)
}

type HexagonalCreatedWithDesiredSymmetry struct {
	baseWavePacket      *formula.WavePacket
	wallpaperMultiplier complex128
}

var _ = Suite(&HexagonalCreatedWithDesiredSymmetry{})

func (suite *HexagonalCreatedWithDesiredSymmetry) SetUpTest(checker *C) {
	suite.baseWavePacket =
		formula.NewWavePacketBuilder().
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

func (suite *HexagonalCreatedWithDesiredSymmetry) TestCreateWallpaperWithP31m(checker *C) {
	newFormula, err := formula.NewBuilder().
		Hexagonal().
		AddWavePacket(suite.baseWavePacket).
		DesiredSymmetry(formula.P31m).
		Build()

	checker.Assert(err, IsNil)

	checker.Assert(newFormula.WavePackets(), HasLen, 2)
	checker.Assert(newFormula.WavePackets()[0].Terms(), HasLen, 3)

	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerN, Equals, -2)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerM, Equals, 1)

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 3)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.P3)
	checker.Assert(newFormula.SymmetriesFound()[2], Equals, formula.P31m)
}

func (suite *HexagonalCreatedWithDesiredSymmetry) TestCreateWallpaperWithP3m1(checker *C) {
	newFormula, err := formula.NewBuilder().
		Hexagonal().
		AddWavePacket(suite.baseWavePacket).
		DesiredSymmetry(formula.P3m1).
		Build()

	checker.Assert(err, IsNil)

	checker.Assert(newFormula.WavePackets(), HasLen, 2)
	checker.Assert(newFormula.WavePackets()[0].Terms(), HasLen, 3)

	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerN, Equals, 2)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerM, Equals, -1)

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 3)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.P3)
	checker.Assert(newFormula.SymmetriesFound()[2], Equals, formula.P3m1)
}

func (suite *HexagonalCreatedWithDesiredSymmetry) TestCreateWallpaperWithP6(checker *C) {
	newFormula, err := formula.NewBuilder().
		Hexagonal().
		AddWavePacket(suite.baseWavePacket).
		DesiredSymmetry(formula.P6).
		Build()

	checker.Assert(err, IsNil)
	checker.Assert(newFormula.WavePackets(), HasLen, 2)
	checker.Assert(newFormula.WavePackets()[0].Terms(), HasLen, 3)

	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerN, Equals, -1)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerM, Equals, 2)

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 3)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.P3)
	checker.Assert(newFormula.SymmetriesFound()[2], Equals, formula.P6)
}

func (suite *HexagonalCreatedWithDesiredSymmetry) TestCreateWallpaperWithP6m(checker *C) {
	newFormula, err := formula.NewBuilder().
		Hexagonal().
		AddWavePacket(suite.baseWavePacket).
		DesiredSymmetry(formula.P6m).
		Build()

	checker.Assert(err, IsNil)
	checker.Assert(newFormula.WavePackets(), HasLen, 4)
	checker.Assert(newFormula.WavePackets()[0].Terms(), HasLen, 3)

	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerN, Equals, -1)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerM, Equals, 2)

	checker.Assert(newFormula.WavePackets()[2].Terms()[0].PowerN, Equals, -2)
	checker.Assert(newFormula.WavePackets()[2].Terms()[0].PowerM, Equals, 1)

	checker.Assert(newFormula.WavePackets()[3].Terms()[0].PowerN, Equals, 2)
	checker.Assert(newFormula.WavePackets()[3].Terms()[0].PowerM, Equals, -1)

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 6)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.P3)
	checker.Assert(newFormula.SymmetriesFound()[2], Equals, formula.P31m)
	checker.Assert(newFormula.SymmetriesFound()[3], Equals, formula.P3m1)
	checker.Assert(newFormula.SymmetriesFound()[4], Equals, formula.P6)
	checker.Assert(newFormula.SymmetriesFound()[5], Equals, formula.P6m)
}

func (suite *HexagonalCreatedWithDesiredSymmetry) TestWhenOtherDesiredSymmetry_BuilderReturnsError(checker *C) {
	newFormula, err := formula.NewBuilder().
		Hexagonal().
		AddWavePacket(suite.baseWavePacket).
		DesiredSymmetry(formula.P2).
		Build()

	checker.Assert(err, ErrorMatches, "hexagonal lattice can apply these desired symmetries: P1, P3, P31m, P3m1, P6, P6m")
	checker.Assert(reflect.TypeOf(newFormula).String(), Equals, "*formula.Identity")
}
