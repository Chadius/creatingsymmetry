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
	}

	suite.oddSumPair = &coefficient.Pairing{
		PowerN: 1,
		PowerM: 2,
	}
}

func (suite *CoefficientPairFeatures) TestPlusNPlusM(checker *C) {
	newSets := suite.evenSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.PlusNPlusM,
	})

	checker.Assert(newSets, HasLen, 1)
	checker.Assert(newSets[0].PowerN, Equals, 1)
	checker.Assert(newSets[0].PowerM, Equals, 3)
	checker.Assert(newSets[0].NegateMultiplier, Equals, false)
}

func (suite *CoefficientPairFeatures) TestPlusMPlusN(checker *C) {
	newSets := suite.evenSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.PlusMPlusN,
	})

	checker.Assert(newSets, HasLen, 1)
	checker.Assert(newSets[0].PowerN, Equals, 3)
	checker.Assert(newSets[0].PowerM, Equals, 1)
	checker.Assert(newSets[0].NegateMultiplier, Equals, false)
}

func (suite *CoefficientPairFeatures) TestReturnsMultiples(checker *C) {
	newSets := suite.evenSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.PlusNPlusM,
		coefficient.PlusMPlusN,
	})

	checker.Assert(newSets, HasLen, 2)
	checker.Assert(newSets[0].PowerN, Equals, 1)
	checker.Assert(newSets[0].PowerM, Equals, 3)
	checker.Assert(newSets[0].NegateMultiplier, Equals, false)

	checker.Assert(newSets[1].PowerN, Equals, 3)
	checker.Assert(newSets[1].PowerM, Equals, 1)
	checker.Assert(newSets[1].NegateMultiplier, Equals, false)
}

func (suite *CoefficientPairFeatures) TestPlusMPlusNMaybeFlipScale(checker *C) {
	newSets := suite.oddSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.PlusMPlusNMaybeFlipScale,
	})

	checker.Assert(newSets, HasLen, 1)
	checker.Assert(newSets[0].PowerN, Equals, 2)
	checker.Assert(newSets[0].PowerM, Equals, 1)
	checker.Assert(newSets[0].NegateMultiplier, Equals, true)
}

func (suite *CoefficientPairFeatures) TestMinusNMinusM(checker *C) {
	newSets := suite.evenSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.MinusNMinusM,
	})

	checker.Assert(newSets, HasLen, 1)
	checker.Assert(newSets[0].PowerN, Equals, -1)
	checker.Assert(newSets[0].PowerM, Equals, -3)
	checker.Assert(newSets[0].NegateMultiplier, Equals, false)
}

func (suite *CoefficientPairFeatures) TestMinusMMinusN(checker *C) {
	newSets := suite.evenSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.MinusMMinusN,
	})

	checker.Assert(newSets, HasLen, 1)
	checker.Assert(newSets[0].PowerN, Equals, -3)
	checker.Assert(newSets[0].PowerM, Equals, -1)
	checker.Assert(newSets[0].NegateMultiplier, Equals, false)
}

func (suite *CoefficientPairFeatures) TestMinusMMinusNMaybeFlipScale(checker *C) {
	newSets := suite.oddSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.MinusMMinusNMaybeFlipScale,
	})

	checker.Assert(newSets, HasLen, 1)
	checker.Assert(newSets[0].PowerN, Equals, -2)
	checker.Assert(newSets[0].PowerM, Equals, -1)
	checker.Assert(newSets[0].NegateMultiplier, Equals, true)
}

func (suite *CoefficientPairFeatures) TestPlusMMinusSumNAndM(checker *C) {
	newSets := suite.evenSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.PlusMMinusSumNAndM,
	})

	checker.Assert(newSets, HasLen, 1)
	checker.Assert(newSets[0].PowerN, Equals, 3)
	checker.Assert(newSets[0].PowerM, Equals, -(1 + 3))
	checker.Assert(newSets[0].NegateMultiplier, Equals, false)
}

func (suite *CoefficientPairFeatures) TestMinusSumNAndMPlusN(checker *C) {
	newSets := suite.evenSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.MinusSumNAndMPlusN,
	})

	checker.Assert(newSets, HasLen, 1)
	checker.Assert(newSets[0].PowerN, Equals, -(1+3))
	checker.Assert(newSets[0].PowerM, Equals, 3)
	checker.Assert(newSets[0].NegateMultiplier, Equals, false)
}

func (suite *CoefficientPairFeatures) TestPlusMMinusN(checker *C) {
	newSets := suite.evenSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.PlusMMinusN,
	})

	checker.Assert(newSets, HasLen, 1)
	checker.Assert(newSets[0].PowerN, Equals, 3)
	checker.Assert(newSets[0].PowerM, Equals, -1)
	checker.Assert(newSets[0].NegateMultiplier, Equals, false)
}

func (suite *CoefficientPairFeatures) TestMinusMPlusN(checker *C) {
	newSets := suite.evenSumPair.GenerateCoefficientSets([]coefficient.Relationship{
		coefficient.MinusMPlusN,
	})

	checker.Assert(newSets, HasLen, 1)
	checker.Assert(newSets[0].PowerN, Equals, -3)
	checker.Assert(newSets[0].PowerM, Equals, 1)
	checker.Assert(newSets[0].NegateMultiplier, Equals, false)
}
