package wave

// Relationship explains how the Wave packet coefficients are related.
type Relationship string

// WaveRelationship s determine the order and sign of powers n and m.
//   Plus means *1, Minus means *-1
//   If N appears first the powers then power N is applied to the number and power M to the complex conjugate.
//   If M appears first the powers then power M is applied to the number and power N to the complex conjugate.
//	 MaybeFlipScale will multiply the scale by -1 if N + M is odd.
const (
	NoRelation Relationship = ""
	PlusMPlusN                         Relationship = "+M+N"
	MinusNMinusM                       Relationship = "-N-M"
	MinusMMinusN                       Relationship = "-M-N"
	MinusNMinusMPlusMPlusNMinusMMinusN Relationship = "-N-M+M+N-M-N"
)

// FindWaveRelationships returns the relationship explaining how this set of wave formulas are related.
func FindWaveRelationships(waveFormulas []*Formula) Relationship {

	if len(waveFormulas) < 2 {
		return ""
	}

	if len(waveFormulas) == 2 {
		return getRelationshipBetweenTwoFormulas(waveFormulas[0], waveFormulas[1])
	}

	if len(waveFormulas) == 4 {
		relation01 := getRelationshipBetweenTwoFormulas(waveFormulas[0], waveFormulas[1])
		relation02 := getRelationshipBetweenTwoFormulas(waveFormulas[0], waveFormulas[2])
		relation03 := getRelationshipBetweenTwoFormulas(waveFormulas[0], waveFormulas[3])

		foundMinusNMinusM := relation01 == MinusNMinusM || relation02 == MinusNMinusM || relation03 == MinusNMinusM
		foundPlusMPlusN := relation01 == PlusMPlusN || relation02 == PlusMPlusN || relation03 == PlusMPlusN
		foundMinusMMinusN := relation01 == MinusMMinusN || relation02 == MinusMMinusN || relation03 == MinusMMinusN

		if foundMinusMMinusN && foundPlusMPlusN && foundMinusNMinusM {
			return MinusNMinusMPlusMPlusNMinusMMinusN
		}
	}

	return NoRelation
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

	if firstPowerNtoSecondPowerMRatio == 1 && firstPowerMtoSecondPowerNRatio == 1{
		return PlusMPlusN
	}

	if firstPowerNtoSecondPowerNRatio == -1 && firstPowerMtoSecondPowerMRatio == -1 {
		return MinusNMinusM
	}

	if firstPowerNtoSecondPowerMRatio == -1 && firstPowerMtoSecondPowerNRatio == -1 {
		return MinusMMinusN
	}

	return NoRelation
}