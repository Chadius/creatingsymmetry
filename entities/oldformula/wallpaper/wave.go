package wallpaper

import (
	"encoding/json"
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	"github.com/Chadius/creating-symmetry/entities/oldformula/eisenstien"
	"github.com/Chadius/creating-symmetry/entities/oldformula/result"
	"github.com/Chadius/creating-symmetry/entities/utility"
	"gopkg.in/yaml.v2"
)

// Marshal can be marshaled and converted to a EisensteinFormulaTerm
type Marshal struct {
	Terms      []*eisenstien.EisensteinFormulaTermMarshal `json:"terms" yaml:"terms"`
	Multiplier utility.ComplexNumberForMarshal            `json:"multiplier" yaml:"multiplier"`
}

// WavePacket for Waves mathematically creates repeating, cyclical mathematical patterns
//   in 2D space, similar to waves on the ocean.
type WavePacket struct {
	Terms      []*eisenstien.EisensteinFormulaTerm
	Multiplier complex128
}

// Calculate takes the complex number zInLatticeCoordinates and processes it using the mathematical terms.
func (waveFormula WavePacket) Calculate(zInLatticeCoordinates complex128) *result.CalculationResultForFormula {
	result := &result.CalculationResultForFormula{
		Total:              complex(0, 0),
		ContributionByTerm: []complex128{},
	}

	for _, term := range waveFormula.Terms {
		termContribution := term.Calculate(zInLatticeCoordinates)
		result.Total += termContribution
		result.ContributionByTerm = append(result.ContributionByTerm, termContribution)
	}
	result.Total *= waveFormula.Multiplier

	return result
}

// NewWaveFormulaFromJSON reads the data and returns a oldformula term from it.
func NewWaveFormulaFromJSON(data []byte) (*WavePacket, error) {
	return newWaveFormulaFromDatastream(data, json.Unmarshal)
}

// NewWaveFormulaFromYAML reads the data and returns a oldformula term from it.
func NewWaveFormulaFromYAML(data []byte) (*WavePacket, error) {
	return newWaveFormulaFromDatastream(data, yaml.Unmarshal)
}

//newWaveFormulaFromDatastream consumes a given bytestream and tries to create a new object from it.
func newWaveFormulaFromDatastream(data []byte, unmarshal utility.UnmarshalFunc) (*WavePacket, error) {
	var unmarshalError error
	var formulaMarshal Marshal
	unmarshalError = unmarshal(data, &formulaMarshal)

	if unmarshalError != nil {
		return nil, unmarshalError
	}

	formulaTerm := NewWaveFormulaFromMarshalObject(formulaMarshal)
	return formulaTerm, nil
}

// NewWaveFormulaFromMarshalObject converts the marshaled intermediary object into a usable object.
func NewWaveFormulaFromMarshalObject(marshalObject Marshal) *WavePacket {
	formulaTerms := []*eisenstien.EisensteinFormulaTerm{}
	for _, term := range marshalObject.Terms {
		newEisenstein := eisenstien.NewEisensteinFormulaTermFromMarshalObject(*term)
		formulaTerms = append(formulaTerms, newEisenstein)
	}

	return &WavePacket{
		Terms:      formulaTerms,
		Multiplier: complex(marshalObject.Multiplier.Real, marshalObject.Multiplier.Imaginary),
	}
}

// GetWavePacketRelationship returns a list of relationships that all of the wave packets conform to.
func GetWavePacketRelationship(wavePacket1, wavePacket2 *WavePacket) []coefficient.Relationship {
	if wavePacket1 == nil || wavePacket2 == nil {
		return []coefficient.Relationship{}
	}

	return GetAllPossibleTermRelationships(wavePacket1.Terms[0], wavePacket2.Terms[0], wavePacket1.Multiplier, wavePacket2.Multiplier)
}

// GetAllPossibleTermRelationships returns a list of relationships that all of the terms conform to.
func GetAllPossibleTermRelationships(
	term1, term2 *eisenstien.EisensteinFormulaTerm,
	term1Multiplier complex128,
	term2Multiplier complex128,
) []coefficient.Relationship {
	if term1 == nil || term2 == nil {
		return []coefficient.Relationship{}
	}

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

// SatisfiesRelationship sees if the all of terms match the coefficient relationship.
func SatisfiesRelationship(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128, relationship coefficient.Relationship) bool {
	if !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && !termMultipliersAreNegated(term1Multiplier, term2Multiplier) {
		return false
	}

	type TwoTermChecker func(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool
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

func satisfiesRelationshipPlusMPlusN(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term1.PowerN == term2.PowerM && term1.PowerM == term2.PowerN
}

func satisfiesRelationshipMinusNMinusM(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term1.PowerN == -1*term2.PowerN && term1.PowerM == -1*term2.PowerM
}

func satisfiesRelationshipPlusNPlusM(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term1.PowerN == term2.PowerN && term1.PowerM == term2.PowerM
}

func satisfiesRelationshipMinusNPlusM(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == -1*term1.PowerN && term2.PowerM == term1.PowerM
}

func satisfiesRelationshipMinusMMinusN(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term1.PowerN == -1*term2.PowerM && term1.PowerM == -1*term2.PowerN
}

func satisfiesRelationshipPlusMPlusNMaybeFlipScale(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerSumIsEven() && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if !term1.PowerSumIsEven() && !termMultipliersAreNegated(term1Multiplier, term2Multiplier) {
		return false
	}

	return term1.PowerN == term2.PowerM && term1.PowerM == term2.PowerN
}

func satisfiesRelationshipMinusMMinusNMaybeFlipScale(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerSumIsEven() && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if !term1.PowerSumIsEven() && !termMultipliersAreNegated(term1Multiplier, term2Multiplier) {
		return false
	}

	return term1.PowerN == -1*term2.PowerM && term1.PowerM == -1*term2.PowerN
}
func satisfiesRelationshipPlusMMinusSumNAndM(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == term1.PowerM && term2.PowerM == -1*(term1.PowerN+term1.PowerM)
}
func satisfiesRelationshipMinusSumNAndMPlusN(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == -1*(term1.PowerN+term1.PowerM) && term2.PowerM == term1.PowerN
}
func satisfiesRelationshipPlusMMinusN(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == term1.PowerM && term2.PowerM == -1*term1.PowerN
}
func satisfiesRelationshipMinusMPlusN(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == -1*term1.PowerM && term2.PowerM == term1.PowerN
}
func satisfiesRelationshipPlusNMinusM(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == term1.PowerN && term2.PowerM == -1*term1.PowerM
}
func satisfiesRelationshipPlusNMinusMNegateMultiplierIfOddPowerN(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerN%2 == 0 && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if term1.PowerN%2 != 0 && termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	return term2.PowerN == term1.PowerN && term2.PowerM == -1*term1.PowerM
}
func satisfiesRelationshipPlusNMinusMNegateMultiplierIfOddPowerSum(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerSumIsEven() && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if !term1.PowerSumIsEven() && termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	return term2.PowerN == term1.PowerN && term2.PowerM == -1*term1.PowerM
}

func satisfiesRelationshipMinusNPlusMNegateMultiplierIfOddPowerN(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerN%2 == 0 && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if term1.PowerN%2 != 0 && termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	return term2.PowerN == -1*term1.PowerN && term2.PowerM == term1.PowerM
}
func satisfiesRelationshipMinusNPlusMNegateMultiplierIfOddPowerSum(term1, term2 *eisenstien.EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerSumIsEven() && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if !term1.PowerSumIsEven() && termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	return term2.PowerN == -1*term1.PowerN && term2.PowerM == term1.PowerM
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

// CanWavePacketsBeGroupedAmongCoefficientRelationships returns true if the WavePackets involved satisfy the relationships.
func CanWavePacketsBeGroupedAmongCoefficientRelationships(wavePackets []*WavePacket, desiredRelationships []coefficient.Relationship) bool {
	wavePacketsMatched := []bool{}
	for range wavePackets {
		wavePacketsMatched = append(wavePacketsMatched, false)
	}

	for indexA, wavePacketA := range wavePackets {
		relationshipWasFound := map[coefficient.Relationship]bool{}
		for _, r := range desiredRelationships {
			relationshipWasFound[r] = false
		}

		if wavePacketsMatched[indexA] == true {
			continue
		}

		for offsetB, wavePacketB := range wavePackets[indexA+1:] {
			relationshipsFound := GetWavePacketRelationship(
				wavePacketA,
				wavePacketB,
			)

			for _, relationshipToLookFor := range desiredRelationships {
				if ContainsRelationship(relationshipsFound, relationshipToLookFor) {
					wavePacketsMatched[indexA+offsetB+1] = true
					relationshipWasFound[relationshipToLookFor] = true
					break
				}
			}
		}

		for _, relationshipFound := range relationshipWasFound {
			if relationshipFound != true {
				return false
			}
		}
		wavePacketsMatched[indexA] = true
	}

	return true
}

// HasSymmetry returns true if the WavePackets involved form the desired symmetry.
func HasSymmetry(wavePackets []*WavePacket, desiredSymmetry Symmetry, desiredSymmetryToCoefficients map[Symmetry][]coefficient.Relationship) bool {
	numberOfWavePackets := len(wavePackets)
	if numberOfWavePackets < 2 || numberOfWavePackets%2 == 1 {
		return false
	}

	coefficientsToFind := desiredSymmetryToCoefficients[desiredSymmetry]

	if coefficientsToFind == nil {
		return false
	}

	return CanWavePacketsBeGroupedAmongCoefficientRelationships(wavePackets, coefficientsToFind)
}
