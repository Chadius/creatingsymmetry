package formula_test

import (
	"github.com/Chadius/creating-symmetry/entities/formula"
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	. "gopkg.in/check.v1"
)

type TermBuilderTest struct{}

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

type BuilderMakeTermUsingDataStream struct{}

var _ = Suite(&BuilderMakeTermUsingDataStream{})

func (suite *BuilderMakeTermUsingDataStream) TestMakeTermWithPowers(checker *C) {
	yamlByteStream := []byte(`
  power_n: 3
  power_m: 1
`)
	newTerm := formula.NewTermBuilder().UsingYAMLData(yamlByteStream).Build()
	checker.Assert(newTerm, NotNil)
	checker.Assert(newTerm.PowerN, Equals, 3)
	checker.Assert(newTerm.PowerM, Equals, 1)

	//term := formula.FormulaLevelTerms()[0]
	// Make sure multipliers are 1, 0 by default
	//PowerN                   int
	//PowerM                   int
	//IgnoreComplexConjugate   bool
	//CoefficientRelationships []coefficient.Relationship
}

// TODO Add tests for making Terms from YAML
// TODO Make a new Term builder test file