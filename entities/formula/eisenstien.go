package formula

import (
	"encoding/json"
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	"github.com/Chadius/creating-symmetry/entities/utility"
	"gopkg.in/yaml.v2"
	"math"
	"math/cmplx"
)

// EisensteinFormulaTermMarshal can be marshaled and converted to a EisensteinFormulaTerm
type EisensteinFormulaTermMarshal struct {
	PowerN int `json:"power_n" yaml:"power_n"`
	PowerM int `json:"power_m" yaml:"power_m"`
}

// EisensteinFormulaTerm defines the shape of a lattice, a 2D structure that remains consistent
//    in wallpaper symmetry.
type EisensteinFormulaTerm struct {
	PowerN int
	PowerM int
}

// PowerSumIsEven returns true if the sum of the term powers is divisible by 2.
func (term EisensteinFormulaTerm) PowerSumIsEven() bool {
	return (term.PowerM+term.PowerN)%2 == 0
}

// Calculate uses the Eisenstein formula on the complex number z.
// Calculate(z) = e ^ (2 PI i * (nX + mY))
//  where n amd m are PowerN and PowerM,
//  and X and Y are the real and imag parts of (zInLatticeCoordinates)
func (term EisensteinFormulaTerm) Calculate(zInLatticeCoordinates complex128) complex128 {
	powerMultiplier := (float64(term.PowerN) * real(zInLatticeCoordinates)) +
		(float64(term.PowerM) * imag(zInLatticeCoordinates))
	expo := cmplx.Exp(complex(0, 2.0*math.Pi*powerMultiplier))
	return expo
}

// NewEisensteinFormulaTermFromJSON reads the data and returns a formula term from it.
func NewEisensteinFormulaTermFromJSON(data []byte) (*EisensteinFormulaTerm, error) {
	return newEisensteinFormulaTermFromDatastream(data, json.Unmarshal)
}

// NewEisensteinFormulaTermFromYAML reads the data and returns a formula term from it.
func NewEisensteinFormulaTermFromYAML(data []byte) (*EisensteinFormulaTerm, error) {
	return newEisensteinFormulaTermFromDatastream(data, yaml.Unmarshal)
}

//newEisensteinFormulaTermFromDatastream consumes a given bytestream and tries to create a new object from it.
func newEisensteinFormulaTermFromDatastream(data []byte, unmarshal utility.UnmarshalFunc) (*EisensteinFormulaTerm, error) {
	var unmarshalError error
	var formulaTermMarshal EisensteinFormulaTermMarshal
	unmarshalError = unmarshal(data, &formulaTermMarshal)

	if unmarshalError != nil {
		return nil, unmarshalError
	}

	formulaTerm := NewEisensteinFormulaTermFromMarshalObject(formulaTermMarshal)
	return formulaTerm, nil
}

// NewEisensteinFormulaTermFromMarshalObject converts the marshaled intermediary object into a usable object.
func NewEisensteinFormulaTermFromMarshalObject(marshalObject EisensteinFormulaTermMarshal) *EisensteinFormulaTerm {
	return &EisensteinFormulaTerm{
		PowerN: marshalObject.PowerN,
		PowerM: marshalObject.PowerM,
	}
}

// GetAllPossibleTermRelationships returns a list of relationships that all of the terms conform to.
func GetAllPossibleTermRelationships(
	term1, term2 *EisensteinFormulaTerm,
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
func SatisfiesRelationship(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128, relationship coefficient.Relationship) bool {
	if !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && !termMultipliersAreNegated(term1Multiplier, term2Multiplier) {
		return false
	}

	type TwoTermChecker func(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool
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

func satisfiesRelationshipPlusMPlusN(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term1.PowerN == term2.PowerM && term1.PowerM == term2.PowerN
}

func satisfiesRelationshipMinusNMinusM(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term1.PowerN == -1*term2.PowerN && term1.PowerM == -1*term2.PowerM
}

func satisfiesRelationshipPlusNPlusM(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term1.PowerN == term2.PowerN && term1.PowerM == term2.PowerM
}

func satisfiesRelationshipMinusNPlusM(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == -1*term1.PowerN && term2.PowerM == term1.PowerM
}

func satisfiesRelationshipMinusMMinusN(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term1.PowerN == -1*term2.PowerM && term1.PowerM == -1*term2.PowerN
}

func satisfiesRelationshipPlusMPlusNMaybeFlipScale(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerSumIsEven() && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if !term1.PowerSumIsEven() && !termMultipliersAreNegated(term1Multiplier, term2Multiplier) {
		return false
	}

	return term1.PowerN == term2.PowerM && term1.PowerM == term2.PowerN
}

func satisfiesRelationshipMinusMMinusNMaybeFlipScale(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerSumIsEven() && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if !term1.PowerSumIsEven() && !termMultipliersAreNegated(term1Multiplier, term2Multiplier) {
		return false
	}

	return term1.PowerN == -1*term2.PowerM && term1.PowerM == -1*term2.PowerN
}
func satisfiesRelationshipPlusMMinusSumNAndM(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == term1.PowerM && term2.PowerM == -1*(term1.PowerN+term1.PowerM)
}
func satisfiesRelationshipMinusSumNAndMPlusN(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == -1*(term1.PowerN+term1.PowerM) && term2.PowerM == term1.PowerN
}
func satisfiesRelationshipPlusMMinusN(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == term1.PowerM && term2.PowerM == -1*term1.PowerN
}
func satisfiesRelationshipMinusMPlusN(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == -1*term1.PowerM && term2.PowerM == term1.PowerN
}
func satisfiesRelationshipPlusNMinusM(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	return termMultipliersAreTheSame(term1Multiplier, term2Multiplier) && term2.PowerN == term1.PowerN && term2.PowerM == -1*term1.PowerM
}
func satisfiesRelationshipPlusNMinusMNegateMultiplierIfOddPowerN(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerN%2 == 0 && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if term1.PowerN%2 != 0 && termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	return term2.PowerN == term1.PowerN && term2.PowerM == -1*term1.PowerM
}
func satisfiesRelationshipPlusNMinusMNegateMultiplierIfOddPowerSum(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerSumIsEven() && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if !term1.PowerSumIsEven() && termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	return term2.PowerN == term1.PowerN && term2.PowerM == -1*term1.PowerM
}

func satisfiesRelationshipMinusNPlusMNegateMultiplierIfOddPowerN(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerN%2 == 0 && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if term1.PowerN%2 != 0 && termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	return term2.PowerN == -1*term1.PowerN && term2.PowerM == term1.PowerM
}
func satisfiesRelationshipMinusNPlusMNegateMultiplierIfOddPowerSum(term1, term2 *EisensteinFormulaTerm, term1Multiplier, term2Multiplier complex128) bool {
	if term1.PowerSumIsEven() && !termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	if !term1.PowerSumIsEven() && termMultipliersAreTheSame(term1Multiplier, term2Multiplier) {
		return false
	}

	return term2.PowerN == -1*term1.PowerN && term2.PowerM == term1.PowerM
}
