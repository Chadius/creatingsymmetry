package coefficient_test

import (
	. "gopkg.in/check.v1"
	"testing"
	"wallpaper/entities/formula/coefficient"
)

func Test(t *testing.T) { TestingT(t) }

type CoefficientPairFeatures struct {
	evenSumPair *coefficient.Pairing
	oddSumPair *coefficient.Pairing
}

var _ = Suite(&CoefficientPairFeatures{})

func (suite *CoefficientPairFeatures) SetUpTest(checker *C) {
	suite.evenSumPair = &coefficient.Pairing{
		PowerN: 1,
		PowerM: 3,
		Multiplier: complex(2.5, 1),
	}

	suite.oddSumPair = &coefficient.Pairing{
		PowerN: 1,
		PowerM: 2,
		Multiplier: complex(2.5, 1),
	}
}

func (suite *CoefficientPairFeatures) TestPlusNPlusM(checker *C) {
	newSets := suite.evenSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.PlusNPlusM,
	})

	checker.Assert(newSets, HasLen, 1)
	checker.Assert(newSets[0].PowerN, Equals, 1)
	checker.Assert(newSets[0].PowerM, Equals, 3)
	checker.Assert(real(newSets[0].Multiplier), Equals, 2.5)
	checker.Assert(imag(newSets[0].Multiplier), Equals, 1.0)
}

func (suite *CoefficientPairFeatures) TestPlusMPlusN(checker *C) {
	newSets := suite.evenSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.PlusMPlusN,
	})

	checker.Assert(newSets, HasLen, 1)
	checker.Assert(newSets[0].PowerN, Equals, 3)
	checker.Assert(newSets[0].PowerM, Equals, 1)
	checker.Assert(real(newSets[0].Multiplier), Equals, 2.5)
	checker.Assert(imag(newSets[0].Multiplier), Equals, 1.0)
}

func (suite *CoefficientPairFeatures) TestReturnsMultiples(checker *C) {
	newSets := suite.evenSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.PlusNPlusM,
		coefficient.PlusMPlusN,
	})

	checker.Assert(newSets, HasLen, 2)
	checker.Assert(newSets[0].PowerN, Equals, 1)
	checker.Assert(newSets[0].PowerM, Equals, 3)
	checker.Assert(real(newSets[0].Multiplier), Equals, 2.5)
	checker.Assert(imag(newSets[0].Multiplier), Equals, 1.0)

	checker.Assert(newSets[1].PowerN, Equals, 3)
	checker.Assert(newSets[1].PowerM, Equals, 1)
	checker.Assert(real(newSets[0].Multiplier), Equals, 2.5)
	checker.Assert(imag(newSets[0].Multiplier), Equals, 1.0)
}

func (suite *CoefficientPairFeatures) TestPlusMPlusNMaybeFlipScale(checker *C) {
	newSets := suite.oddSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.PlusMPlusNMaybeFlipScale,
	})

	checker.Assert(newSets, HasLen, 1)
	checker.Assert(newSets[0].PowerN, Equals, 2)
	checker.Assert(newSets[0].PowerM, Equals, 1)
	checker.Assert(real(newSets[0].Multiplier), Equals, -2.5)
	checker.Assert(imag(newSets[0].Multiplier), Equals, -1.0)
}

func (suite *CoefficientPairFeatures) TestMinusNMinusM(checker *C) {
	newSets := suite.evenSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.MinusNMinusM,
	})

	checker.Assert(newSets, HasLen, 1)
	checker.Assert(newSets[0].PowerN, Equals, -1)
	checker.Assert(newSets[0].PowerM, Equals, -3)
	checker.Assert(real(newSets[0].Multiplier), Equals, 2.5)
	checker.Assert(imag(newSets[0].Multiplier), Equals, 1.0)
}

func (suite *CoefficientPairFeatures) TestMinusMMinusN(checker *C) {
	newSets := suite.evenSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.MinusMMinusN,
	})

	checker.Assert(newSets, HasLen, 1)
	checker.Assert(newSets[0].PowerN, Equals, -3)
	checker.Assert(newSets[0].PowerM, Equals, -1)
	checker.Assert(real(newSets[0].Multiplier), Equals, 2.5)
	checker.Assert(imag(newSets[0].Multiplier), Equals, 1.0)
}

func (suite *CoefficientPairFeatures) TestMinusMMinusNMaybeFlipScale(checker *C) {
	newSets := suite.oddSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.MinusMMinusNMaybeFlipScale,
	})

	checker.Assert(newSets, HasLen, 1)
	checker.Assert(newSets[0].PowerN, Equals, -2)
	checker.Assert(newSets[0].PowerM, Equals, -1)
	checker.Assert(real(newSets[0].Multiplier), Equals, -2.5)
	checker.Assert(imag(newSets[0].Multiplier), Equals, -1.0)
}