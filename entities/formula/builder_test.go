package formula_test

import (
	"github.com/Chadius/creating-symmetry/entities/formula"
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	. "gopkg.in/check.v1"
	"reflect"
)

type BuilderTest struct {}

var _ = Suite(&BuilderTest{})

func (b *BuilderTest) TestIdentityFormula(checker *C) {
	identityFormula, _ := formula.NewBuilder().Build()
	checker.Assert(reflect.TypeOf(identityFormula).String(), Equals, "*formula.Identity")
}

func (b *BuilderTest) TestRosetteFormula(checker *C) {
	rosetteFormula, _ := formula.NewBuilder().
		Rosette().
		AddTerm(
			formula.NewTermBuilder().Build(),
		).
		Build()
	checker.Assert(reflect.TypeOf(rosetteFormula).String(), Equals, "*formula.Rosette")
	checker.Assert(rosetteFormula.FormulaLevelTerms(), HasLen, 1)
}

func (b *BuilderTest) TestFriezeFormula(checker *C) {
	rosetteFormula, _ := formula.NewBuilder().
		Frieze().
		AddTerm(
			formula.NewTermBuilder().Build(),
		).
		Build()
	checker.Assert(reflect.TypeOf(rosetteFormula).String(), Equals, "*formula.Frieze")
	checker.Assert(rosetteFormula.FormulaLevelTerms(), HasLen, 1)
}

func (b *BuilderTest) TestRectangularFormula(checker *C) {
	rectangularFormula, _ := formula.NewBuilder().
		Rectangular().
		LatticeHeight(0.5).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1,0)).
				AddTerm(
					formula.NewTermBuilder().PowerN(1).PowerM(-2).Build(),
				).
				Build(),
		).
		Build()

	checker.Assert(reflect.TypeOf(rectangularFormula).String(), Equals, "*formula.Rectangular")
	checker.Assert(rectangularFormula.WavePackets(), HasLen, 1)
}

func (b *BuilderTest) TestWhenNoLatticeHeight_ThenRectangularFormulaReturnsError(checker *C) {
	rectangularFormula, err := formula.NewBuilder().
		Rectangular().
		Build()

	checker.Assert(err, ErrorMatches, "rectangular lattice must specify height")
	checker.Assert(reflect.TypeOf(rectangularFormula).String(), Equals, "*formula.Identity")
}

func (b *BuilderTest) TestSquareFormula(checker *C) {
	squareFormula, _ := formula.NewBuilder().
		Square().
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1,0)).
				AddTerm(
					formula.NewTermBuilder().PowerN(1).PowerM(-2).Build(),
				).
				Build(),
		).
		Build()

	checker.Assert(reflect.TypeOf(squareFormula).String(), Equals, "*formula.Square")
	checker.Assert(squareFormula.WavePackets(), HasLen, 1)
}

func (b *BuilderTest) TestHexagonalFormula(checker *C) {
	hexagonFormula, _ := formula.NewBuilder().
		Hexagonal().
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1,0)).
				AddTerm(
					formula.NewTermBuilder().PowerN(1).PowerM(-2).Build(),
				).
				Build(),
		).
		Build()

	checker.Assert(reflect.TypeOf(hexagonFormula).String(), Equals, "*formula.Hexagonal")
	checker.Assert(hexagonFormula.WavePackets(), HasLen, 1)
}

func (b *BuilderTest) TestHexagonalFormulaWithDesiredSymmetry(checker *C) {
	hexagonFormula, _ := formula.NewBuilder().
		Hexagonal().
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1,0)).
				AddTerm(
					formula.NewTermBuilder().PowerN(1).PowerM(-2).Build(),
				).
				Build(),
		).
		DesiredSymmetry(formula.P31m).
		Build()

	checker.Assert(hexagonFormula.SymmetriesFound(), HasLen, 1)
	checker.Assert(hexagonFormula.SymmetriesFound()[0], Equals, formula.P31m)
}

func (b *BuilderTest) TestRhombicFormula(checker *C) {
	rhombicFormula, _ := formula.NewBuilder().
		Rhombic().
		LatticeHeight(0.5).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1,0)).
				AddTerm(
					formula.NewTermBuilder().PowerN(1).PowerM(-2).Build(),
				).
				Build(),
		).
		Build()

	checker.Assert(reflect.TypeOf(rhombicFormula).String(), Equals, "*formula.Rhombic")
	checker.Assert(rhombicFormula.WavePackets(), HasLen, 1)
}

func (b *BuilderTest) TestWhenNoLatticeHeight_RhombicFormulaReturnsError(checker *C) {
	rhombicFormula, err := formula.NewBuilder().
		Rhombic().
		Build()

	checker.Assert(err, ErrorMatches, "rhombic lattice must specify height")
	checker.Assert(reflect.TypeOf(rhombicFormula).String(), Equals, "*formula.Identity")
}

func (b *BuilderTest) TestGenericFormula(checker *C) {
	genericFormula, _ := formula.NewBuilder().
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

	checker.Assert(reflect.TypeOf(genericFormula).String(), Equals, "*formula.Generic")
	checker.Assert(genericFormula.WavePackets(), HasLen, 1)
}

func (b *BuilderTest) TestWhenNoLatticeHeight_GenericFormulaReturnsError(checker *C) {
	genericFormula, err := formula.NewBuilder().
		Generic().
		Build()

	checker.Assert(err, ErrorMatches, "generic lattice must specify dimensions")
	checker.Assert(reflect.TypeOf(genericFormula).String(), Equals, "*formula.Identity")
}

type TermBuilderTest struct {}

var _ = Suite(&TermBuilderTest{})

func (t *TermBuilderTest) TestCreateTerm(checker *C) {
	term := formula.NewTermBuilder().
		Multiplier(complex(2e-3, -5e7)).
		PowerN(-11).
		PowerM(13).
		IgnoreComplexConjugate().
		AddCoefficientRelationship(coefficient.PlusMPlusN).
		AddCoefficientRelationship(coefficient.MinusMMinusN).
		Build()
	checker.Assert(term.Multiplier, Equals, complex(2e-3, -5e7))
	checker.Assert(term.PowerN, Equals, -11)
	checker.Assert(term.PowerM, Equals, 13)
	checker.Assert(term.IgnoreComplexConjugate, Equals, true)
	checker.Assert(term.CoefficientRelationships, HasLen, 2)
	checker.Assert(term.CoefficientRelationships[0], Equals, coefficient.PlusMPlusN)
	checker.Assert(term.CoefficientRelationships[1], Equals, coefficient.MinusMMinusN)
}