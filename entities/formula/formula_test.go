package formula_test

import (
	. "gopkg.in/check.v1"
	"testing"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/coefficient"
)

func Test(t *testing.T) { TestingT(t) }

type CommonFormulaFeatures struct {
}

var _ = Suite(&CommonFormulaFeatures{})

func (suite *CommonFormulaFeatures) SetUpTest(checker *C) {
}

func (suite *CommonFormulaFeatures) TestPlusNPlusM(checker *C) {
	power1, power2, scale := formula.SetCoefficientsBasedOnRelationship(1, 2, complex(3, 4), coefficient.PlusNPlusM)
	checker.Assert(power1, Equals, 1)
	checker.Assert(power2, Equals, 2)
	checker.Assert(real(scale), Equals, 3.0)
	checker.Assert(imag(scale), Equals, 4.0)
}

func (suite *CommonFormulaFeatures) TestPlusMPlusN(checker *C) {
	power1, power2, scale := formula.SetCoefficientsBasedOnRelationship(1, 2, complex(3, 4), coefficient.PlusMPlusN)
	checker.Assert(power1, Equals, 2)
	checker.Assert(power2, Equals, 1)
	checker.Assert(real(scale), Equals, 3.0)
	checker.Assert(imag(scale), Equals, 4.0)
}

func (suite *CommonFormulaFeatures) TestMinusNMinusM(checker *C) {
	power1, power2, scale := formula.SetCoefficientsBasedOnRelationship(1, 2, complex(3, 4), coefficient.MinusNMinusM)
	checker.Assert(power1, Equals, -1)
	checker.Assert(power2, Equals, -2)
	checker.Assert(real(scale), Equals, 3.0)
	checker.Assert(imag(scale), Equals, 4.0)
}

func (suite *CommonFormulaFeatures) TestMinusMMinusN(checker *C) {
	power1, power2, scale := formula.SetCoefficientsBasedOnRelationship(1, 2, complex(3, 4), coefficient.MinusMMinusN)
	checker.Assert(power1, Equals, -2)
	checker.Assert(power2, Equals, -1)
	checker.Assert(real(scale), Equals, 3.0)
	checker.Assert(imag(scale), Equals, 4.0)
}

func (suite *CommonFormulaFeatures) TestPlusMPlusNMaybeFlipScaleOdd(checker *C) {
	power1, power2, scale := formula.SetCoefficientsBasedOnRelationship(1, 2, complex(3, 4), coefficient.PlusMPlusNMaybeFlipScale)
	checker.Assert(power1, Equals, 2)
	checker.Assert(power2, Equals, 1)
	checker.Assert(real(scale), Equals, -3.0)
	checker.Assert(imag(scale), Equals, -4.0)
}

func (suite *CommonFormulaFeatures) TestPlusMPlusNMaybeFlipScaleEven(checker *C) {
	power1, power2, scale := formula.SetCoefficientsBasedOnRelationship(0, 2, complex(3, 4), coefficient.PlusMPlusNMaybeFlipScale)
	checker.Assert(power1, Equals, 2)
	checker.Assert(power2, Equals, 0)
	checker.Assert(real(scale), Equals, 3.0)
	checker.Assert(imag(scale), Equals, 4.0)
}

func (suite *CommonFormulaFeatures) TestMinusMMinusNMaybeFlipScaleOdd(checker *C) {
	power1, power2, scale := formula.SetCoefficientsBasedOnRelationship(1, 2, complex(3, 4), coefficient.MinusMMinusNMaybeFlipScale)
	checker.Assert(power1, Equals, -2)
	checker.Assert(power2, Equals, -1)
	checker.Assert(real(scale), Equals, -3.0)
	checker.Assert(imag(scale), Equals, -4.0)
}

func (suite *CommonFormulaFeatures) TestMinusMMinusNMaybeFlipScaleEven(checker *C) {
	power1, power2, scale := formula.SetCoefficientsBasedOnRelationship(0, 2, complex(3, 4), coefficient.MinusMMinusNMaybeFlipScale)
	checker.Assert(power1, Equals, -2)
	checker.Assert(power2, Equals, 0)
	checker.Assert(real(scale), Equals, 3.0)
	checker.Assert(imag(scale), Equals, 4.0)
}

func (suite *CommonFormulaFeatures) TestLockedCoefficientPairYAML(checker *C) {
	yamlByteStream := []byte(`multiplier: 1
relationships:
  - "+N+M"
  - "-M-NF"
`)
	coefficientPair, err := formula.NewLockedCoefficientPairFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(coefficientPair.Multiplier, Equals, 1.0)
	checker.Assert(coefficientPair.OtherCoefficientRelationships, HasLen, 2)
	checker.Assert(coefficientPair.OtherCoefficientRelationships[0], Equals, coefficient.Relationship(coefficient.PlusNPlusM))
	checker.Assert(coefficientPair.OtherCoefficientRelationships[1], Equals, coefficient.Relationship(coefficient.MinusMMinusNMaybeFlipScale))
}
