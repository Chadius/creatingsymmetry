package coefficient

// Pairing notes the multiplier and the powers applied to a formula term.
type Pairing struct {
	PowerN		int
	PowerM		int
	Multiplier	complex128
}

// GenerateCoefficientSets creates a list of locked coefficient sets (powers and multipliers)
//  based on a given list of relationships.
func (pairing Pairing) GenerateCoefficientSets(relationships []Relationship) []*Pairing {
	pairs := []*Pairing{}

	multiplierMultiplierBasedOnPowerNPlusPowerMIsEvenOrOdd := 1.0
	if (pairing.PowerN + pairing.PowerM) % 2 != 0 {
		multiplierMultiplierBasedOnPowerNPlusPowerMIsEvenOrOdd = -1.0
	}

	pairingByRelationship := map[Relationship]*Pairing{
		PlusNPlusM: {
			PowerN: pairing.PowerN,
			PowerM: pairing.PowerM,
			Multiplier: pairing.Multiplier,
		},
		PlusMPlusN: {
			PowerN: pairing.PowerM,
			PowerM: pairing.PowerN,
			Multiplier: pairing.Multiplier,
		},
		PlusMPlusNMaybeFlipScale: {
			PowerN:     pairing.PowerM,
			PowerM:     pairing.PowerN,
			Multiplier: complex(
				multiplierMultiplierBasedOnPowerNPlusPowerMIsEvenOrOdd * real(pairing.Multiplier),
				multiplierMultiplierBasedOnPowerNPlusPowerMIsEvenOrOdd * imag(pairing.Multiplier),
			),
		},
		MinusNMinusM: {
			PowerN:     -1 * pairing.PowerN,
			PowerM:     -1 * pairing.PowerM,
			Multiplier: pairing.Multiplier,
		},
		MinusMMinusN: {
			PowerN:     -1 * pairing.PowerM,
			PowerM:     -1 * pairing.PowerN,
			Multiplier: pairing.Multiplier,
		},
		MinusMMinusNMaybeFlipScale: {
			PowerN:     -1 * pairing.PowerM,
			PowerM:     -1 * pairing.PowerN,
			Multiplier: complex(
				multiplierMultiplierBasedOnPowerNPlusPowerMIsEvenOrOdd * real(pairing.Multiplier),
				multiplierMultiplierBasedOnPowerNPlusPowerMIsEvenOrOdd * imag(pairing.Multiplier),
			),
		},
	}

	for _, relationship := range relationships {
		pairWithSameRelationship := pairingByRelationship[relationship]
		newPair := &Pairing{
			PowerN: pairWithSameRelationship.PowerN,
			PowerM: pairWithSameRelationship.PowerM,
			Multiplier: pairWithSameRelationship.Multiplier,
		}
		pairs = append(pairs, newPair)
	}

	return pairs
}

// Relationship relates how a pair of coordinates should be applied.
type Relationship string

// Relationship s determine the order and sign of powers n and m.
//   Plus means *1, Minus means *-1
//   If N appears first the powers then power N is applied to the number and power M to the complex conjugate.
//   If M appears first the powers then power M is applied to the number and power N to the complex conjugate.
//	 MaybeFlipScale will multiply the scale by -1 if N + M is odd.
const (
	PlusNPlusM                 Relationship = "+N+M"
	PlusMPlusN                              = "+M+N"
	MinusNMinusM                            = "-N-M"
	MinusMMinusN                            = "-M-N"
	PlusMPlusNMaybeFlipScale                = "+M+NF"
	MinusMMinusNMaybeFlipScale              = "-M-NF"
)

