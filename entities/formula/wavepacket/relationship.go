package wavepacket

// Relationship explains how the Wave packet coefficients are related.
//   Plus means *1, Minus means *-1
//   If N appears first the powers then power N is applied to the number and power M to the complex conjugate.
//   If M appears first the powers then power M is applied to the number and power N to the complex conjugate.
//	 MaybeFlipScale will multiply the scale by -1 if N + M is odd.
type Relationship struct {
	NoRelation bool
	PlusMPlusN bool
	MinusNMinusM bool
	MinusMMinusN bool
	MinusNMinusMPlusMPlusNMinusMMinusN bool
}

// FindWaveRelationships returns the relationship explaining how this set of wavepacket formulas are related.
func FindWaveRelationships(waveFormulas []*Formula) Relationship {

	formulaRelationship := Relationship{}

	if len(waveFormulas) < 2 {
		formulaRelationship.NoRelation = true
		return formulaRelationship
	}

	if len(waveFormulas) == 2 {
		return getRelationshipBetweenTwoFormulas(waveFormulas[0], waveFormulas[1])
	}

	if len(waveFormulas) == 4 {
		relation01 := getRelationshipBetweenTwoFormulas(waveFormulas[0], waveFormulas[1])
		relation02 := getRelationshipBetweenTwoFormulas(waveFormulas[0], waveFormulas[2])
		relation03 := getRelationshipBetweenTwoFormulas(waveFormulas[0], waveFormulas[3])

		foundMinusNMinusM := relation01.MinusNMinusM || relation02.MinusNMinusM || relation03.MinusNMinusM
		foundPlusMPlusN := relation01.PlusMPlusN || relation02.PlusMPlusN || relation03.PlusMPlusN
		foundMinusMMinusN := relation01.MinusMMinusN || relation02.MinusMMinusN || relation03.MinusMMinusN

		if foundMinusMMinusN && foundPlusMPlusN && foundMinusNMinusM {
			formulaRelationship.MinusNMinusMPlusMPlusNMinusMMinusN = true
			return formulaRelationship
		}
	}

	formulaRelationship.NoRelation = true
	return formulaRelationship
}

func getRelationshipBetweenTwoFormulas(pair1, pair2 *Formula) Relationship {
	firstPowerNtoSecondPowerMRatio := 0
	if pair2.Terms[0].PowerM != 0 {
		firstPowerNtoSecondPowerMRatio = pair1.Terms[0].PowerN / pair2.Terms[0].PowerM
	}

	firstPowerMtoSecondPowerNRatio := 0
	if pair2.Terms[0].PowerN != 0 {
		firstPowerMtoSecondPowerNRatio = pair1.Terms[0].PowerM / pair2.Terms[0].PowerN
	}

	firstPowerNtoSecondPowerNRatio := 0
	if pair2.Terms[0].PowerN != 0 {
		firstPowerNtoSecondPowerNRatio = pair1.Terms[0].PowerN / pair2.Terms[0].PowerN
	}

	firstPowerMtoSecondPowerMRatio := 0
	if pair2.Terms[0].PowerM != 0 {
		firstPowerMtoSecondPowerMRatio = pair1.Terms[0].PowerM / pair2.Terms[0].PowerM
	}

	relationBetween := Relationship{}
	relationFound := false

	if firstPowerNtoSecondPowerMRatio == 1 && firstPowerMtoSecondPowerNRatio == 1{
		relationBetween.PlusMPlusN = true
		relationFound = true
	}

	if firstPowerNtoSecondPowerNRatio == -1 && firstPowerMtoSecondPowerMRatio == -1 {
		relationBetween.MinusNMinusM = true
		relationFound = true
	}

	if firstPowerNtoSecondPowerMRatio == -1 && firstPowerMtoSecondPowerNRatio == -1 {
		relationBetween.MinusMMinusN = true
		relationFound = true
	}

	if !relationFound {
		relationBetween.NoRelation = true
	}
	return relationBetween
}