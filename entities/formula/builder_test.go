package formula_test

import (
	"github.com/Chadius/creating-symmetry/entities/formula"
	. "gopkg.in/check.v1"
	"reflect"
)

type BuilderTest struct{}

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
				Multiplier(complex(1, 0)).
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

func (b *BuilderTest) TestRectangularFormulaWithDesiredSymmetry(checker *C) {
	rectangularFormula, _ := formula.NewBuilder().
		Rectangular().
		LatticeHeight(0.5).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
				AddTerm(
					formula.NewTermBuilder().PowerN(1).PowerM(-2).Build(),
				).
				Build(),
		).
		DesiredSymmetry(formula.Pg).
		Build()

	foundPgSymmetry := false
	for _, symmetry := range rectangularFormula.SymmetriesFound() {
		if symmetry == formula.Pg {
			foundPgSymmetry = true
		}
	}
	checker.Assert(foundPgSymmetry, Equals, true)
}

func (b *BuilderTest) TestSquareFormula(checker *C) {
	squareFormula, _ := formula.NewBuilder().
		Square().
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
				AddTerm(
					formula.NewTermBuilder().PowerN(1).PowerM(-2).Build(),
				).
				Build(),
		).
		Build()

	checker.Assert(reflect.TypeOf(squareFormula).String(), Equals, "*formula.Square")
	checker.Assert(squareFormula.WavePackets(), HasLen, 1)
}

func (b *BuilderTest) TestSquareFormulaWithDesiredSymmetry(checker *C) {
	squareFormula, _ := formula.NewBuilder().
		Square().
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
				AddTerm(
					formula.NewTermBuilder().PowerN(-2).PowerM(1).Build(),
				).
				Build(),
		).
		DesiredSymmetry(formula.P4m).
		Build()

	foundP4mSymmetry := false
	for _, symmetry := range squareFormula.SymmetriesFound() {
		if symmetry == formula.P4m {
			foundP4mSymmetry = true
		}
	}
	checker.Assert(foundP4mSymmetry, Equals, true)
}

func (b *BuilderTest) TestHexagonalFormula(checker *C) {
	hexagonFormula, _ := formula.NewBuilder().
		Hexagonal().
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
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
				Multiplier(complex(1, 0)).
				AddTerm(
					formula.NewTermBuilder().PowerN(1).PowerM(-2).Build(),
				).
				Build(),
		).
		DesiredSymmetry(formula.P31m).
		Build()

	foundP31mSymmetry := false
	for _, symmetry := range hexagonFormula.SymmetriesFound() {
		if symmetry == formula.P31m {
			foundP31mSymmetry = true
		}
	}
	checker.Assert(foundP31mSymmetry, Equals, true)
}

func (b *BuilderTest) TestRhombicFormula(checker *C) {
	rhombicFormula, _ := formula.NewBuilder().
		Rhombic().
		LatticeHeight(0.5).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
				AddTerm(
					formula.NewTermBuilder().PowerN(1).PowerM(-2).Build(),
				).
				Build(),
		).
		Build()

	checker.Assert(reflect.TypeOf(rhombicFormula).String(), Equals, "*formula.Rhombic")
	checker.Assert(rhombicFormula.WavePackets(), HasLen, 1)
}

func (b *BuilderTest) TestRhombicFormulaWithDesiredSymmetry(checker *C) {
	squareFormula, _ := formula.NewBuilder().
		Rhombic().
		LatticeHeight(0.5).
		AddWavePacket(
			formula.NewWavePacketBuilder().
				Multiplier(complex(1, 0)).
				AddTerm(
					formula.NewTermBuilder().PowerN(1).PowerM(-2).Build(),
				).
				Build(),
		).
		DesiredSymmetry(formula.Cm).
		Build()

	foundCmSymmetry := false
	for _, symmetry := range squareFormula.SymmetriesFound() {
		if symmetry == formula.Cm {
			foundCmSymmetry = true
		}
	}
	checker.Assert(foundCmSymmetry, Equals, true)
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
				Multiplier(complex(1, 0)).
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

type BuilderMakeFormulaUsingDataStream struct{}

var _ = Suite(&BuilderMakeFormulaUsingDataStream{})

func (suite *BuilderMakeFormulaUsingDataStream) TestMakeRosetteFormulaWithYAML(checker *C) {
	yamlByteStream := []byte(`
type: rosette
terms:
  -
    power_n: 3
    power_m: 1
`)
	newFormula, _ := formula.NewBuilder().UsingYAMLData(yamlByteStream).Build()
	checker.Assert(newFormula, NotNil)
	checker.Assert(reflect.TypeOf(newFormula).String(), Equals, "*formula.Rosette")
	checker.Assert(newFormula.FormulaLevelTerms(), HasLen, 1)

	//term := formula.FormulaLevelTerms()[0]
	// Make sure multipliers are 1, 0 by default
	//PowerN                   int
	//PowerM                   int
	//IgnoreComplexConjugate   bool
	//CoefficientRelationships []coefficient.Relationship
}

// TODO Add tests for making Terms from YAML
// TODO Make a new Term builder test file
