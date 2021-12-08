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
}

func (suite *BuilderMakeTermUsingDataStream) TestMakeTermWithDefaultMultiplier(checker *C) {
	yamlByteStream := []byte(`
`)
	newTerm := formula.NewTermBuilder().UsingYAMLData(yamlByteStream).Build()
	checker.Assert(newTerm.Multiplier, Equals, complex(1, 0))
}

func (suite *BuilderMakeTermUsingDataStream) TestMakeTermWithSpecificMultiplier(checker *C) {
	yamlByteStream := []byte(`
multiplier:
  real: 2
  imaginary: 3e-1
`)
	newTerm := formula.NewTermBuilder().UsingYAMLData(yamlByteStream).Build()
	checker.Assert(newTerm.Multiplier, Equals, complex(2, 3e-1))
}

func (suite *BuilderMakeTermUsingDataStream) TestMakeTermIgnoreComplexConjugate(checker *C) {
	yamlByteStream := []byte(`
ignore_complex_conjugate: true
`)
	newTerm := formula.NewTermBuilder().UsingYAMLData(yamlByteStream).Build()
	checker.Assert(newTerm.IgnoreComplexConjugate, Equals, true)
}

func (suite *BuilderMakeTermUsingDataStream) TestMakeTermWithCoefficientRelationships(checker *C) {
	yamlByteStream := []byte(`
coefficient_relationships: ["-N-M", "-N+MF(N+M)", "-(N+M)+N"]
`)
	newTerm := formula.NewTermBuilder().UsingYAMLData(yamlByteStream).Build()
	checker.Assert(newTerm.CoefficientRelationships, HasLen, 3)
	checker.Assert(newTerm.CoefficientRelationships[0], Equals, coefficient.MinusNMinusM)
	checker.Assert(newTerm.CoefficientRelationships[1], Equals, coefficient.MinusNPlusMNegateMultiplierIfOddPowerSum)
	checker.Assert(newTerm.CoefficientRelationships[2], Equals, coefficient.MinusSumNAndMPlusN)
}

func (suite *BuilderMakeTermUsingDataStream) TestMakeTermWithJSON(checker *C) {
	jsonByteStream := []byte(`{
	"multiplier": {
		"real": 1.2e2,
		"imaginary": 4.3e-5
	},
	"power_n": 5,
	"power_m": 7,
	"ignore_complex_conjugate": true,
	"coefficient_relationships": ["+N+M", "-N+M"]
}`)
	newTerm := formula.NewTermBuilder().UsingJSONData(jsonByteStream).Build()
	checker.Assert(newTerm, NotNil)
	checker.Assert(newTerm.Multiplier, Equals, complex(1.2e2, 4.3e-5))
	checker.Assert(newTerm.PowerN, Equals, 5)
	checker.Assert(newTerm.PowerM, Equals, 7)
	checker.Assert(newTerm.IgnoreComplexConjugate, Equals, true)
	checker.Assert(newTerm.CoefficientRelationships, HasLen, 2)
	checker.Assert(newTerm.CoefficientRelationships[0], Equals, coefficient.PlusNPlusM)
	checker.Assert(newTerm.CoefficientRelationships[1], Equals, coefficient.MinusNPlusM)
}
