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
	identityFormula := formula.NewBuilder().Build()
	checker.Assert(reflect.TypeOf(identityFormula).String(), Equals, "*formula.Identity")
}

func (b *BuilderTest) TestRosetteFormula(checker *C) {
	rosetteFormula := formula.NewBuilder().
		Rosette().
		AddTerm(
			formula.NewTermBuilder().Build(),
		).
		Build()
	checker.Assert(reflect.TypeOf(rosetteFormula).String(), Equals, "*formula.Rosette")
	checker.Assert(rosetteFormula.FormulaLevelTerms(), HasLen, 1)
}

func (b *BuilderTest) TestFriezeFormula(checker *C) {
	rosetteFormula := formula.NewBuilder().
		Frieze().
		AddTerm(
			formula.NewTermBuilder().Build(),
		).
		Build()
	checker.Assert(reflect.TypeOf(rosetteFormula).String(), Equals, "*formula.Frieze")
	checker.Assert(rosetteFormula.FormulaLevelTerms(), HasLen, 1)
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