package exponential_test

import (
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	"github.com/Chadius/creating-symmetry/entities/oldformula/exponential"
	"github.com/Chadius/creating-symmetry/entities/utility"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type ExponentialTerm struct{}

var _ = Suite(&ExponentialTerm{})

func (suite *ExponentialTerm) SetUpTest(checker *C) {}

func (suite *ExponentialTerm) TestCreateTermFromYAML(checker *C) {
	yamlByteStream := []byte(`
multiplier:
  real: -1.0
  imaginary: 2e-2
power_n: 12
power_m: -10
ignore_complex_conjugate: true
coefficient_relationships:
  - -M-N
  - +M+NF(N+M)
`)

	term, err := exponential.NewTermFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(real(term.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(term.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
	checker.Assert(term.PowerN, Equals, 12)
	checker.Assert(term.PowerM, Equals, -10)
	checker.Assert(term.IgnoreComplexConjugate, Equals, true)
	checker.Assert(term.CoefficientRelationships, HasLen, 2)
	checker.Assert(term.CoefficientRelationships[0], Equals, coefficient.Relationship(coefficient.MinusMMinusN))
	checker.Assert(term.CoefficientRelationships[1], Equals, coefficient.Relationship(coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum))
}

func (suite *ExponentialTerm) TestCreateTermFromJSON(checker *C) {
	jsonByteStream := []byte(`{
				"multiplier": {
					"real": -1.0,
					"imaginary": 2e-2
				},
				"power_n": 12,
				"power_m": -10,
				"ignore_complex_conjugate": true,
				"coefficient_relationships": ["-M-N", "+M+NF(N+M)"]
			}`)
	term, err := exponential.NewTermFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(real(term.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(term.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
	checker.Assert(term.PowerN, Equals, 12)
	checker.Assert(term.PowerM, Equals, -10)
	checker.Assert(term.IgnoreComplexConjugate, Equals, true)
	checker.Assert(term.CoefficientRelationships, HasLen, 2)
	checker.Assert(term.CoefficientRelationships[0], Equals, coefficient.Relationship(coefficient.MinusMMinusN))
	checker.Assert(term.CoefficientRelationships[1], Equals, coefficient.Relationship(coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum))
}
