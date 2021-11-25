package formula

import (
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
)

// WavePacket for Waves mathematically creates repeating, cyclical mathematical patterns
//   in 2D space, similar to waves on the ocean.
type WavePacket struct {
	terms      []Term
	multiplier complex128
}

// Multiplier is a getter.
func (wavePacket WavePacket) Multiplier() complex128 {
	return wavePacket.multiplier
}

// Terms is a getter.
func (wavePacket WavePacket) Terms() []Term {
	return wavePacket.terms
}

// Calculate takes the complex number zInLatticeCoordinates and processes it using the mathematical terms.
func (wavePacket WavePacket) Calculate(zInLatticeCoordinates complex128) complex128 {
	sum := complex(0,0)

	for _, term := range wavePacket.Terms() {
		termContribution := term.CalculateInLatticeCoordinates(zInLatticeCoordinates)
		sum += termContribution
	}
	return sum * wavePacket.Multiplier()
}

// GetWavePacketRelationship returns a list of relationships that all of the wave packets conform to.
func GetWavePacketRelationship(wavePacket1, wavePacket2 WavePacket) []coefficient.Relationship {
	termForWavePacket1 := wavePacket1.Terms()[0]
	termForWavePacket2 := wavePacket2.Terms()[0]

	return GetAllPossibleTermRelationships3(termForWavePacket1, termForWavePacket2, wavePacket1.Multiplier(), wavePacket2.Multiplier())
}

// ContainsRelationship returns true if the target relationship is in the list of relationships.
func ContainsRelationship(relationships []coefficient.Relationship, target coefficient.Relationship) bool {
	for _, thisRelationship := range relationships {
		if target == thisRelationship {
			return true
		}
	}
	return false
}

// GetAllPossibleTermRelationships3 returns a list of relationships that all terms conform to.
func GetAllPossibleTermRelationships3(
	term1, term2 Term,
	term1Multiplier complex128,
	term2Multiplier complex128,
) []coefficient.Relationship {
	foundRelationships := []coefficient.Relationship{}

	relationshipsToTest := []coefficient.Relationship{
		coefficient.PlusNPlusM,
		coefficient.PlusMPlusN,
		coefficient.MinusNMinusM,
		coefficient.MinusMMinusN,
		coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum,
		coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum,
		coefficient.PlusMMinusSumNAndM,
		coefficient.MinusSumNAndMPlusN,
		coefficient.PlusMMinusN,
		coefficient.MinusMPlusN,
		coefficient.PlusNMinusM,
		coefficient.MinusNPlusM,
		coefficient.PlusNMinusMNegateMultiplierIfOddPowerN,
		coefficient.MinusNPlusMNegateMultiplierIfOddPowerN,
		coefficient.PlusNMinusMNegateMultiplierIfOddPowerSum,
		coefficient.MinusNPlusMNegateMultiplierIfOddPowerSum,
	}

	for _, relationshipToTest := range relationshipsToTest {
		if SatisfiesRelationship(term1, term2, term1Multiplier, term2Multiplier, relationshipToTest) {
			foundRelationships = append(foundRelationships, relationshipToTest)
		}
	}

	return foundRelationships
}

// SatisfiesRelationship sees if the all terms match the coefficient relationship.
func SatisfiesRelationship(term1, term2 Term, term1Multiplier, term2Multiplier complex128, relationship coefficient.Relationship) bool {
	if !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && !termMultipliersAreNegated(term1Multiplier, term2Multiplier) {
		return false
	}

	type TwoTermChecker func(term1, term2 Term, term1Multiplier, term2Multiplier complex128) bool
	relationshipCheckerByRelationship := map[coefficient.Relationship]TwoTermChecker{
		coefficient.PlusMPlusN:                                satisfiesRelationshipPlusMPlusN,
		coefficient.MinusNMinusM:                              satisfiesRelationshipMinusNMinusM,
		coefficient.PlusNPlusM:                                satisfiesRelationshipPlusNPlusM,
		coefficient.MinusMMinusN:                              satisfiesRelationshipMinusMMinusN,
		coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum:   satisfiesRelationshipPlusMPlusNMaybeFlipScale,
		coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum: satisfiesRelationshipMinusMMinusNMaybeFlipScale,
		coefficient.PlusMMinusSumNAndM:                        satisfiesRelationshipPlusMMinusSumNAndM,
		coefficient.MinusSumNAndMPlusN:                        satisfiesRelationshipMinusSumNAndMPlusN,
		coefficient.PlusMMinusN:                               satisfiesRelationshipPlusMMinusN,
		coefficient.MinusMPlusN:                               satisfiesRelationshipMinusMPlusN,
		coefficient.PlusNMinusM:                               satisfiesRelationshipPlusNMinusM,
		coefficient.MinusNPlusM:                               satisfiesRelationshipMinusNPlusM,
		coefficient.PlusNMinusMNegateMultiplierIfOddPowerN:    satisfiesRelationshipPlusNMinusMNegateMultiplierIfOddPowerN,
		coefficient.PlusNMinusMNegateMultiplierIfOddPowerSum:  satisfiesRelationshipPlusNMinusMNegateMultiplierIfOddPowerSum,
		coefficient.MinusNPlusMNegateMultiplierIfOddPowerN:    satisfiesRelationshipMinusNPlusMNegateMultiplierIfOddPowerN,
		coefficient.MinusNPlusMNegateMultiplierIfOddPowerSum:  satisfiesRelationshipMinusNPlusMNegateMultiplierIfOddPowerSum,
	}
	relationshipChecker := relationshipCheckerByRelationship[relationship]
	return relationshipChecker(term1, term2, term1Multiplier, term2Multiplier)
}

func termMultipliersAreTheSame(term1Multiplier, term2Multiplier complex128) bool {
	return real(term1Multiplier) == real(term2Multiplier) && imag(term1Multiplier) == imag(term2Multiplier)
}

func termMultipliersAreNegated(term1Multiplier, term2Multiplier complex128) bool {
	return real(term1Multiplier) == -1*real(term2Multiplier) && imag(term1Multiplier) == -1*imag(term2Multiplier)
}

func satisfiesRelationshipPlusMPlusN(term1, term2 Term, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term1.PowerN == term2.PowerM && term1.PowerM == term2.PowerN
}

func satisfiesRelationshipMinusNMinusM(term1, term2 Term, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term1.PowerN == -1*term2.PowerN && term1.PowerM == -1*term2.PowerM
}

func satisfiesRelationshipPlusNPlusM(term1, term2 Term, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term1.PowerN == term2.PowerN && term1.PowerM == term2.PowerM
}

func satisfiesRelationshipMinusNPlusM(term1, term2 Term, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == -1*term1.PowerN && term2.PowerM == term1.PowerM
}

func satisfiesRelationshipMinusMMinusN(term1, term2 Term, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term1.PowerN == -1*term2.PowerM && term1.PowerM == -1*term2.PowerN
}

func satisfiesRelationshipPlusMPlusNMaybeFlipScale(term1, term2 Term, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerSumIsEven() && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if !term1.PowerSumIsEven() && !termMultipliersAreNegated(term1Multiplier, term2Multiplier) {
		return false
	}

	return term1.PowerN == term2.PowerM && term1.PowerM == term2.PowerN
}

func satisfiesRelationshipMinusMMinusNMaybeFlipScale(term1, term2 Term, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerSumIsEven() && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if !term1.PowerSumIsEven() && !termMultipliersAreNegated(term1Multiplier, term2Multiplier) {
		return false
	}

	return term1.PowerN == -1*term2.PowerM && term1.PowerM == -1*term2.PowerN
}
func satisfiesRelationshipPlusMMinusSumNAndM(term1, term2 Term, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == term1.PowerM && term2.PowerM == -1*(term1.PowerN+term1.PowerM)
}
func satisfiesRelationshipMinusSumNAndMPlusN(term1, term2 Term, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == -1*(term1.PowerN+term1.PowerM) && term2.PowerM == term1.PowerN
}
func satisfiesRelationshipPlusMMinusN(term1, term2 Term, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == term1.PowerM && term2.PowerM == -1*term1.PowerN
}
func satisfiesRelationshipMinusMPlusN(term1, term2 Term, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == -1*term1.PowerM && term2.PowerM == term1.PowerN
}
func satisfiesRelationshipPlusNMinusM(term1, term2 Term, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == term1.PowerN && term2.PowerM == -1*term1.PowerM
}
func satisfiesRelationshipPlusNMinusMNegateMultiplierIfOddPowerN(term1, term2 Term, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerN%2 == 0 && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if term1.PowerN%2 != 0 && termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	return term2.PowerN == term1.PowerN && term2.PowerM == -1*term1.PowerM
}
func satisfiesRelationshipPlusNMinusMNegateMultiplierIfOddPowerSum(term1, term2 Term, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerSumIsEven() && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if !term1.PowerSumIsEven() && termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	return term2.PowerN == term1.PowerN && term2.PowerM == -1*term1.PowerM
}

func satisfiesRelationshipMinusNPlusMNegateMultiplierIfOddPowerN(term1, term2 Term, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerN%2 == 0 && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if term1.PowerN%2 != 0 && termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	return term2.PowerN == -1*term1.PowerN && term2.PowerM == term1.PowerM
}
func satisfiesRelationshipMinusNPlusMNegateMultiplierIfOddPowerSum(term1, term2 Term, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerSumIsEven() && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if !term1.PowerSumIsEven() && termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	return term2.PowerN == -1*term1.PowerN && term2.PowerM == term1.PowerM
}

// WavePacketBuilder is used to create new WavePackets.
type WavePacketBuilder struct {
	multiplier complex128
	terms []Term
}

// NewWavePacketBuilder returns a blank WavePacketBuilder.
func NewWavePacketBuilder() *WavePacketBuilder {
	return &WavePacketBuilder{
		multiplier: complex(0,0),
		terms: []Term{},
	}
}

// Multiplier sets the wave packet's multiplier.
func (w *WavePacketBuilder) Multiplier(multiplier complex128) *WavePacketBuilder {
	w.multiplier = multiplier
	return w
}

// AddTerm adds a term to the formula.
func (w *WavePacketBuilder) AddTerm(term *Term) *WavePacketBuilder {
	w.terms = append(w.terms, *term)
	return w
}

// Build returns a new WavePacket.
func (w *WavePacketBuilder) Build() *WavePacket {
	return &WavePacket{
		terms:      w.terms,
		multiplier: w.multiplier,
	}
}