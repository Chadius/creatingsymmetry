package frieze

import (
	"gopkg.in/yaml.v2"
	"math/cmplx"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/coefficient"
	"wallpaper/entities/formula/exponential"
	"wallpaper/entities/utility"
)

// Formula is used to generate frieze patterns.
type Formula struct {
	Terms []*exponential.Term
}

// Calculate applies the Frieze formula to the complex number z.
func (friezeFormula Formula) Calculate(z complex128) *formula.CalculationResultForFormula {
	result := &formula.CalculationResultForFormula{
		Total: complex(0,0),
		ContributionByTerm: []complex128{},
	}

	for _, term := range friezeFormula.Terms {
		termResult := friezeFormula.calculateTerm(term, z)
		result.Total += termResult
		result.ContributionByTerm = append(result.ContributionByTerm, termResult)
	}

	return result
}

func (friezeFormula *Formula) calculateTerm(term *exponential.Term, z complex128) complex128 {
	sum := complex(0.0,0.0)

	coefficientRelationships := []coefficient.Relationship{coefficient.PlusNPlusM}
	coefficientRelationships = append(coefficientRelationships, term.CoefficientRelationships...)
	coefficientSets := coefficient.Pairing{
		PowerN:     term.PowerN,
		PowerM:     term.PowerM,
	}.GenerateCoefficientSets(coefficientRelationships)

	for _, relationshipSet := range coefficientSets {
		multiplier := term.Multiplier
		if relationshipSet.NegateMultiplier == true {
			multiplier *= -1
		}
		sum += CalculateEulerTerm(z, relationshipSet.PowerN, relationshipSet.PowerM, multiplier, term.IgnoreComplexConjugate)
	}
	return sum
}

// Symmetry notes the kinds of symmetries the formula contains.
type Symmetry struct {
	P111 bool
	P11m bool
	P211 bool
	P1m1 bool
	P11g bool
	P2mm bool
	P2mg bool
}

//AnalyzeForSymmetry scans the formula and returns a list of symmetries.
func (friezeFormula Formula) AnalyzeForSymmetry() *Symmetry {
	symmetriesFound := &Symmetry{
		P111: true,
		P11m: true,
		P211: true,
		P1m1: true,
		P11g: true,
		P2mm: true,
		P2mg: true,
	}
	for _, term := range friezeFormula.Terms {
		if term.IgnoreComplexConjugate {
			symmetriesFound.P211 = false
			symmetriesFound.P1m1 = false
			symmetriesFound.P11g = false
			symmetriesFound.P11m = false
			symmetriesFound.P2mm = false
			symmetriesFound.P2mg = false
		}

		powerSumIsEven := (term.PowerN + term.PowerM) % 2 == 0

		containsMinusNMinusM := coefficientRelationshipsIncludes(term.CoefficientRelationships, coefficient.MinusNMinusM)
		containsMinusMMinusN := coefficientRelationshipsIncludes(term.CoefficientRelationships, coefficient.MinusMMinusN)
		containsPlusMPlusN := coefficientRelationshipsIncludes(term.CoefficientRelationships, coefficient.PlusMPlusN)

		containsMinusMMinusNAndPowerSumIsOdd := coefficientRelationshipsIncludes(term.CoefficientRelationships, coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum) && !powerSumIsEven
		containsPlusMPlusNAndPowerSumIsOdd := coefficientRelationshipsIncludes(term.CoefficientRelationships, coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum) && !powerSumIsEven

		containsMinusMMinusNAndPowerSumIsEven := coefficientRelationshipsIncludes(term.CoefficientRelationships, coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum) && powerSumIsEven
		containsPlusMPlusNAndPowerSumIsEven := coefficientRelationshipsIncludes(term.CoefficientRelationships, coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum) && powerSumIsEven

		if !containsMinusNMinusM {
			symmetriesFound.P211 = false
		}
		if !containsPlusMPlusN {
			symmetriesFound.P1m1 = false
		}
		if !containsMinusMMinusNAndPowerSumIsOdd {
			symmetriesFound.P11g = false
		}
		if !(containsMinusMMinusN || containsMinusMMinusNAndPowerSumIsEven) {
			symmetriesFound.P11m = false
		}
		if !(
			containsMinusNMinusM &&
				(containsPlusMPlusN || containsPlusMPlusNAndPowerSumIsEven) &&
				(containsMinusMMinusN || containsMinusMMinusNAndPowerSumIsEven)) {
			symmetriesFound.P2mm = false
		}
		if !(containsMinusNMinusM && containsPlusMPlusNAndPowerSumIsOdd && containsMinusMMinusNAndPowerSumIsOdd) {
			symmetriesFound.P2mg = false
		}
	}

	return symmetriesFound
}

// CalculateEulerTerm calculates e^(i*n*z) * e^(-i*m*zConj)
func CalculateEulerTerm(z complex128, power1, power2 int, scale complex128, ignoreComplexConjugate bool) complex128 {
	eRaisedToTheNZi := cmplx.Exp(complex(0,1) * z * complex(float64(power1), 0))
	if ignoreComplexConjugate {
		return eRaisedToTheNZi * scale
	}

	complexConjugate := complex(real(z), -1 * imag(z))
	eRaisedToTheNegativeMZConji := cmplx.Exp(complexConjugate * complex(0, -1 * float64(power2)))
	return eRaisedToTheNZi * eRaisedToTheNegativeMZConji * scale
}

func coefficientRelationshipsIncludes(relationships []coefficient.Relationship, relationshipToFind coefficient.Relationship) bool {
	for _, relationship := range relationships {
		if relationship == relationshipToFind {
			return true
		}
	}
	return false
}

// NewFriezeFormulaFromYAML reads the data and returns a Frieze formula from it.
func NewFriezeFormulaFromYAML(data []byte) (*Formula, error) {
	return newFriezeFormulaFromDatastream(data, yaml.Unmarshal)
}

// NewFriezeFormulaFromJSON reads the data and returns a Frieze formula from it.
func NewFriezeFormulaFromJSON(data []byte) (*Formula, error) {
	return newFriezeFormulaFromDatastream(data, yaml.Unmarshal)
}

// MarshaledFormula can be marshaled and can be converted into a Formula.
type MarshaledFormula struct {
	Terms []*exponential.TermMarshalable
}

// newFriezeFormulaFromDatastream consumes a given bytestream and tries to create a new object from it.
func newFriezeFormulaFromDatastream(data []byte, unmarshal utility.UnmarshalFunc) (*Formula, error) {
	var unmarshalError error
	var friezeFormulaMarshal MarshaledFormula
	unmarshalError = unmarshal(data, &friezeFormulaMarshal)

	if unmarshalError != nil {
		return nil, unmarshalError
	}

	friezeFormula := NewFriezeFormulaFromMarshalObject(friezeFormulaMarshal)
	return friezeFormula, nil
}

// NewFriezeFormulaFromMarshalObject converts the marshaled object into a Formula.
func NewFriezeFormulaFromMarshalObject(marshalObject MarshaledFormula) *Formula {
	terms := []*exponential.Term{}
	for _, termMarshal := range marshalObject.Terms {
		newTerm := exponential.NewTermFromMarshalObject(*termMarshal)
		terms = append(terms, newTerm)
	}
	return &Formula{Terms: terms}
}
