package formula_test

import (
	"github.com/Chadius/creating-symmetry/entities/formula"
	"github.com/Chadius/creating-symmetry/entities/utility"
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"
	"reflect"
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
				Multiplier(complex(1, 0)).
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

func (suite *RectangularWallpaper) TestCalculationOfPoints(checker *C) {
	calculation := suite.newFormula.Calculate(complex(0.75, -0.25))

	expectedAnswer := cmplx.Exp(complex(0, math.Pi*7/2))
	checker.Assert(real(calculation), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(calculation), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}

type RectangularWallpaperHasSymmetryTest struct {
	baseWavePacketWithEvenPowerNAndOddPowerSum *formula.WavePacket
	baseWavePacketWithOddPowerNAndEvenPowerSum *formula.WavePacket
	wallpaperMultiplier                        complex128
}

var _ = Suite(&RectangularWallpaperHasSymmetryTest{})

func (suite *RectangularWallpaperHasSymmetryTest) SetUpTest(checker *C) {
	suite.baseWavePacketWithEvenPowerNAndOddPowerSum = formula.NewWavePacketBuilder().
		AddTerm(
			formula.NewTermWithMultiplierAndPowers(
				complex(1, 0),
				8,
				-3,
			),
		).
		Multiplier(complex(1, 0)).
		Build()

	suite.baseWavePacketWithOddPowerNAndEvenPowerSum = formula.NewWavePacketBuilder().
		AddTerm(
			formula.NewTermWithMultiplierAndPowers(
				complex(1, 0),
				7,
				-3,
			),
		).
		Multiplier(complex(1, 0)).
		Build()

	suite.wallpaperMultiplier = complex(1, 0)
}

func (suite *RectangularWallpaperHasSymmetryTest) TestRectangularHasNoSymmetry(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Rectangular().
		LatticeHeight(1.5).
		AddWavePacket(suite.baseWavePacketWithEvenPowerNAndOddPowerSum).
		Build()

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 1)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
}

func (suite *RectangularWallpaperHasSymmetryTest) TestRectangularMayHaveSymmetryForPm(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Rectangular().
		LatticeHeight(1.5).
		AddWavePacket(suite.baseWavePacketWithEvenPowerNAndOddPowerSum).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier()).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier(),
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms()[0].PowerN,
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms()[0].PowerM*-1,
					),
				).
				Build(),
		).
		Build()

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 3)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.Pm)
	checker.Assert(newFormula.SymmetriesFound()[2], Equals, formula.Pg)
}

func (suite *RectangularWallpaperHasSymmetryTest) TestRectangularMayHaveSymmetryForPg(checker *C) {
	newFormulaWithEvenPowerN, _ := formula.NewBuilder().
		Rectangular().
		LatticeHeight(1.5).
		AddWavePacket(suite.baseWavePacketWithEvenPowerNAndOddPowerSum).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier()).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier(),
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms()[0].PowerN,
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms()[0].PowerM*-1,
					),
				).
				Build(),
		).
		Build()

	checker.Assert(newFormulaWithEvenPowerN.SymmetriesFound(), HasLen, 3)
	checker.Assert(newFormulaWithEvenPowerN.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormulaWithEvenPowerN.SymmetriesFound()[1], Equals, formula.Pm)
	checker.Assert(newFormulaWithEvenPowerN.SymmetriesFound()[2], Equals, formula.Pg)

	newFormulaWithOddPowerN, _ := formula.NewBuilder().
		Rectangular().
		LatticeHeight(1.5).
		AddWavePacket(suite.baseWavePacketWithOddPowerNAndEvenPowerSum).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier() * -1).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier()*-1,
						suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN,
						suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM*-1,
					),
				).
				Build(),
		).
		Build()

	checker.Assert(newFormulaWithOddPowerN.SymmetriesFound(), HasLen, 2)
	checker.Assert(newFormulaWithOddPowerN.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormulaWithOddPowerN.SymmetriesFound()[1], Equals, formula.Pg)
}

func (suite *RectangularWallpaperHasSymmetryTest) TestPmmAndPmgWithEvenPowerN(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Rectangular().
		LatticeHeight(1.5).
		AddWavePacket(suite.baseWavePacketWithEvenPowerNAndOddPowerSum).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier()).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms()[0].PowerN*-1,
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms()[0].PowerM*-1,
					),
				).
				Build(),
		).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier()).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms()[0].PowerN,
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms()[0].PowerM*-1,
					),
				).
				Build(),
		).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier()).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms()[0].PowerN*-1,
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms()[0].PowerM,
					),
				).
				Build(),
		).
		Build()

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 5)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.Pm)
	checker.Assert(newFormula.SymmetriesFound()[2], Equals, formula.Pg)
	checker.Assert(newFormula.SymmetriesFound()[3], Equals, formula.Pmm)
	checker.Assert(newFormula.SymmetriesFound()[4], Equals, formula.Pmg)
}

func (suite *RectangularWallpaperHasSymmetryTest) TestPmgWithOddPowerN(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Rectangular().
		LatticeHeight(1.5).
		AddWavePacket(suite.baseWavePacketWithOddPowerNAndEvenPowerSum).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier()).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN*-1,
						suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM*-1,
					),
				).
				Build(),
		).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier() * -1).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN,
						suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM*-1,
					),
				).
				Build(),
		).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier() * -1).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN*-1,
						suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM,
					),
				).
				Build(),
		).
		Build()

	checker.Assert(newFormula.SymmetriesFound(), HasLen, 3)
	checker.Assert(newFormula.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormula.SymmetriesFound()[1], Equals, formula.Pg)
	checker.Assert(newFormula.SymmetriesFound()[2], Equals, formula.Pmg)
}

func (suite *RectangularWallpaperHasSymmetryTest) TestPgg(checker *C) {
	newFormulaWithOddPowerSum, _ := formula.NewBuilder().
		Rectangular().
		LatticeHeight(1.5).
		AddWavePacket(suite.baseWavePacketWithEvenPowerNAndOddPowerSum).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier()).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms()[0].PowerN*-1,
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms()[0].PowerM*-1,
					),
				).
				Build(),
		).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier() * -1).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms()[0].PowerN,
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms()[0].PowerM*-1,
					),
				).
				Build(),
		).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Multiplier() * -1).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms()[0].PowerN*-1,
						suite.baseWavePacketWithEvenPowerNAndOddPowerSum.Terms()[0].PowerM,
					),
				).
				Build(),
		).
		Build()

	checker.Assert(newFormulaWithOddPowerSum.SymmetriesFound(), HasLen, 2)
	checker.Assert(newFormulaWithOddPowerSum.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormulaWithOddPowerSum.SymmetriesFound()[1], Equals, formula.Pgg)

	newFormulaWithEvenPowerSum, _ := formula.NewBuilder().
		Rectangular().
		LatticeHeight(1.5).
		AddWavePacket(suite.baseWavePacketWithOddPowerNAndEvenPowerSum).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier()).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN*-1,
						suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM*-1,
					),
				).
				Build(),
		).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier()).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN,
						suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM*-1,
					),
				).
				Build(),
		).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier()).
				AddTerm(
					formula.NewTermWithMultiplierAndPowers(
						complex(1, 0),
						suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN*-1,
						suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM,
					),
				).
				Build(),
		).
		Build()

	checker.Assert(newFormulaWithEvenPowerSum.SymmetriesFound(), HasLen, 4)
	checker.Assert(newFormulaWithEvenPowerSum.SymmetriesFound()[0], Equals, formula.P1)
	checker.Assert(newFormulaWithEvenPowerSum.SymmetriesFound()[1], Equals, formula.Pm)
	checker.Assert(newFormulaWithEvenPowerSum.SymmetriesFound()[2], Equals, formula.Pmm)
	checker.Assert(newFormulaWithEvenPowerSum.SymmetriesFound()[3], Equals, formula.Pgg)
}

type RectangularCreatedWithDesiredSymmetry struct {
	baseWavePacketWithOddPowerNAndEvenPowerSum *formula.WavePacket
}

var _ = Suite(&RectangularCreatedWithDesiredSymmetry{})

func (suite *RectangularCreatedWithDesiredSymmetry) SetUpTest(checker *C) {
	suite.baseWavePacketWithOddPowerNAndEvenPowerSum = formula.NewWavePacketBuilder().
		Multiplier(complex(1, 0)).
		AddTerm(
			formula.NewTermWithMultiplierAndPowers(
				complex(1, 0),
				7,
				-3,
			),
		).
		Build()
}

func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateWallpaperWithPm(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Rectangular().
		LatticeHeight(1.5).
		AddWavePacket(suite.baseWavePacketWithOddPowerNAndEvenPowerSum).
		DesiredSymmetry(formula.Pm).
		Build()

	checker.Assert(newFormula.WavePackets(), HasLen, 2)
	checker.Assert(newFormula.WavePackets()[0].Terms(), HasLen, 1)

	checker.Assert(newFormula.WavePackets()[1].Multiplier(), Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier())
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM*-1)

	foundPmSymmetry := false
	for _, symmetry := range newFormula.SymmetriesFound() {
		if symmetry == formula.Pm {
			foundPmSymmetry = true
		}
	}
	checker.Assert(foundPmSymmetry, Equals, true)
}

func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateWallpaperWithPg(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Rectangular().
		LatticeHeight(1.5).
		AddWavePacket(suite.baseWavePacketWithOddPowerNAndEvenPowerSum).
		DesiredSymmetry(formula.Pg).
		Build()

	checker.Assert(newFormula.WavePackets(), HasLen, 2)
	checker.Assert(newFormula.WavePackets()[0].Terms(), HasLen, 1)

	checker.Assert(newFormula.WavePackets()[1].Multiplier(), Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier()*-1)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM*-1)

	foundPgSymmetry := false
	for _, symmetry := range newFormula.SymmetriesFound() {
		if symmetry == formula.Pg {
			foundPgSymmetry = true
		}
	}
	checker.Assert(foundPgSymmetry, Equals, true)
}

func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateWallpaperWithPmm(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Rectangular().
		LatticeHeight(1.5).
		AddWavePacket(suite.baseWavePacketWithOddPowerNAndEvenPowerSum).
		DesiredSymmetry(formula.Pmm).
		Build()

	checker.Assert(newFormula.WavePackets(), HasLen, 4)
	checker.Assert(newFormula.WavePackets()[0].Terms(), HasLen, 1)

	checker.Assert(newFormula.WavePackets()[1].Multiplier(), Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier())
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN*-1)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM*-1)

	checker.Assert(newFormula.WavePackets()[2].Multiplier(), Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier())
	checker.Assert(newFormula.WavePackets()[2].Terms()[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN*-1)
	checker.Assert(newFormula.WavePackets()[2].Terms()[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM)

	checker.Assert(newFormula.WavePackets()[3].Multiplier(), Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier())
	checker.Assert(newFormula.WavePackets()[3].Terms()[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN)
	checker.Assert(newFormula.WavePackets()[3].Terms()[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM*-1)

	foundPmmSymmetry := false
	for _, symmetry := range newFormula.SymmetriesFound() {
		if symmetry == formula.Pmm {
			foundPmmSymmetry = true
		}
	}
	checker.Assert(foundPmmSymmetry, Equals, true)
}

func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateWallpaperWithPmg(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Rectangular().
		LatticeHeight(1.5).
		AddWavePacket(suite.baseWavePacketWithOddPowerNAndEvenPowerSum).
		DesiredSymmetry(formula.Pmg).
		Build()

	checker.Assert(newFormula.WavePackets(), HasLen, 4)
	checker.Assert(newFormula.WavePackets()[0].Terms(), HasLen, 1)

	checker.Assert(newFormula.WavePackets()[1].Multiplier(), Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier())
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN*-1)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM*-1)

	checker.Assert(newFormula.WavePackets()[2].Multiplier(), Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier()*-1)
	checker.Assert(newFormula.WavePackets()[2].Terms()[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN*-1)
	checker.Assert(newFormula.WavePackets()[2].Terms()[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM)

	checker.Assert(newFormula.WavePackets()[3].Multiplier(), Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier()*-1)
	checker.Assert(newFormula.WavePackets()[3].Terms()[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN)
	checker.Assert(newFormula.WavePackets()[3].Terms()[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM*-1)

	foundPmgSymmetry := false
	for _, symmetry := range newFormula.SymmetriesFound() {
		if symmetry == formula.Pmg {
			foundPmgSymmetry = true
		}
	}
	checker.Assert(foundPmgSymmetry, Equals, true)
}

func (suite *RectangularCreatedWithDesiredSymmetry) TestCreateWallpaperWithPgg(checker *C) {
	newFormula, _ := formula.NewBuilder().
		Rectangular().
		LatticeHeight(1.5).
		AddWavePacket(suite.baseWavePacketWithOddPowerNAndEvenPowerSum).
		DesiredSymmetry(formula.Pgg).
		Build()

	checker.Assert(newFormula.WavePackets(), HasLen, 4)
	checker.Assert(newFormula.WavePackets()[0].Terms(), HasLen, 1)

	checker.Assert(newFormula.WavePackets()[1].Multiplier(), Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier())
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN*-1)
	checker.Assert(newFormula.WavePackets()[1].Terms()[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM*-1)

	checker.Assert(newFormula.WavePackets()[2].Multiplier(), Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier())
	checker.Assert(newFormula.WavePackets()[2].Terms()[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN*-1)
	checker.Assert(newFormula.WavePackets()[2].Terms()[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM)

	checker.Assert(newFormula.WavePackets()[3].Multiplier(), Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Multiplier())
	checker.Assert(newFormula.WavePackets()[3].Terms()[0].PowerN, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerN)
	checker.Assert(newFormula.WavePackets()[3].Terms()[0].PowerM, Equals, suite.baseWavePacketWithOddPowerNAndEvenPowerSum.Terms()[0].PowerM*-1)

	foundPggSymmetry := false
	for _, symmetry := range newFormula.SymmetriesFound() {
		if symmetry == formula.Pgg {
			foundPggSymmetry = true
		}
	}
	checker.Assert(foundPggSymmetry, Equals, true)
}

func (suite *RectangularCreatedWithDesiredSymmetry) TestWhenOtherDesiredSymmetry_BuilderReturnsError(checker *C) {
	newFormula, err := formula.NewBuilder().
		Rectangular().
		LatticeHeight(1.5).
		AddWavePacket(suite.baseWavePacketWithOddPowerNAndEvenPowerSum).
		DesiredSymmetry(formula.P2).
		Build()

	checker.Assert(err, ErrorMatches, "rectangular lattice can apply these desired symmetries: P1, Pm, Pg, Pmm, Pmg, Pgg")
	checker.Assert(reflect.TypeOf(newFormula).String(), Equals, "*formula.Identity")
}
