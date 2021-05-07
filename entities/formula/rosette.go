package formula

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"math/cmplx"
	"wallpaper/entities/formula/coefficient"
	"wallpaper/entities/utility"
)

// ZExponentialFormulaTermMarshalable can be marshaled and converted to a ZExponentialFormulaTerm
type ZExponentialFormulaTermMarshalable struct {
	Multiplier					utility.ComplexNumberForMarshal	`json:"multiplier" yaml:"multiplier"`
	PowerN						int								`json:"power_n" yaml:"power_n"`
	PowerM						int								`json:"power_m" yaml:"power_m"`
	IgnoreComplexConjugate		bool							`json:"ignore_complex_conjugate" yaml:"ignore_complex_conjugate"`
	CoefficientRelationships	[]coefficient.Relationship		`json:"coefficient_relationships" yaml:"coefficient_relationships"`
}

// ZExponentialFormulaTerm describes a formula of the form Multiplier * z^PowerN * zConjugate^PowerM.
type ZExponentialFormulaTerm struct {
	Multiplier				complex128
	PowerN					int
	PowerM					int
	// IgnoreComplexConjugate will make sure zConjugate is not used in this calculation
	//    (effectively setting it to 1 + 0i)
	IgnoreComplexConjugate	bool
	// CoefficientRelationships has a list of locked coefficient pairings. These locks are
	//   used to generate similar locked terms. Relationships affect PowerN, PowerM and Multiplier.
	CoefficientRelationships	[]coefficient.Relationship
}

// NewZExponentialFormulaTermFromYAML reads the data and returns a formula term from it.
func NewZExponentialFormulaTermFromYAML(data []byte) (*ZExponentialFormulaTerm, error) {
	return newZExponentialFormulaTermFromDatastream(data, yaml.Unmarshal)
}

// NewZExponentialFormulaTermFromJSON reads the data and returns a formula term from it.
func NewZExponentialFormulaTermFromJSON(data []byte) (*ZExponentialFormulaTerm, error) {
	return newZExponentialFormulaTermFromDatastream(data, json.Unmarshal)
}

// newZExponentialFormulaTermFromDatastream consumes a given bytestream and tries to create a new object from it.
func newZExponentialFormulaTermFromDatastream(data []byte, unmarshal utility.UnmarshalFunc) (*ZExponentialFormulaTerm, error) {
	var unmarshalError error
	var formulaTermMarshal ZExponentialFormulaTermMarshalable
	unmarshalError = unmarshal(data, &formulaTermMarshal)

	if unmarshalError != nil {
		return nil, unmarshalError
	}

	formulaTerm := newZExponentialFormulaTermFromMarshalObject(formulaTermMarshal)
	return formulaTerm, nil
}

func newZExponentialFormulaTermFromMarshalObject(marshalObject ZExponentialFormulaTermMarshalable) *ZExponentialFormulaTerm {
	return &ZExponentialFormulaTerm{
		Multiplier:             	complex(marshalObject.Multiplier.Real, marshalObject.Multiplier.Imaginary),
		PowerN:                 	marshalObject.PowerN,
		PowerM:                 	marshalObject.PowerM,
		IgnoreComplexConjugate: 	marshalObject.IgnoreComplexConjugate,
		CoefficientRelationships:	marshalObject.CoefficientRelationships,
	}
}

// Calculate returns the result of using the formula on the given complex number.
func (term ZExponentialFormulaTerm) Calculate(z complex128) complex128 {
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
		sum += CalculateExponentTerm(z, relationshipSet.PowerN, relationshipSet.PowerM, multiplier, term.IgnoreComplexConjugate)
	}
	return sum
}

// RosetteFormula uses a collection of z^m terms to calculate results.
//    This transforms the input into a circular pattern rotating around the
//    origin.
type RosetteFormula struct {
	Terms []*ZExponentialFormulaTerm
}

// Calculate applies the Rosette formula to the complex number z.
func (r RosetteFormula) Calculate(z complex128) *CalculationResultForFormula {
	result := &CalculationResultForFormula{
		Total: complex(0,0),
		ContributionByTerm: []complex128{},
	}

	for _, term := range r.Terms {
		termResult := term.Calculate(z)
		result.Total += termResult
		result.ContributionByTerm = append(result.ContributionByTerm, termResult)
	}

	return result
}

// RosetteSymmetry notes the kinds of symmetries the rosette formula contains.
type RosetteSymmetry struct {
	Multifold int
}

// AnalyzeForSymmetry analyzes the formula for symmetries.
func (r RosetteFormula) AnalyzeForSymmetry() *RosetteSymmetry {
	symmetriesFound := &RosetteSymmetry{
		Multifold: 1,
	}

	r.calculateMultifoldSymmetry(symmetriesFound)
	return symmetriesFound
}

func (r RosetteFormula) calculateMultifoldSymmetry(symmetriesFound *RosetteSymmetry) {
	termPowerDifferences := []int{}

	for _, term := range r.Terms {
		powerDifference := term.PowerN - term.PowerM
		if powerDifference < 0 {
			powerDifference *= -1
		}
		termPowerDifferences = append(termPowerDifferences, powerDifference)
	}

	if len(termPowerDifferences) == 1 {
		symmetriesFound.Multifold = termPowerDifferences[0]
	} else if len(termPowerDifferences) > 1 {
		var currentGreatestCommonDenominator int
		for index := range termPowerDifferences {
			if index >= len(termPowerDifferences) - 1 {
				break
			}
			currentGreatestCommonDenominator = getGreatestCommonDenominator(
				termPowerDifferences[index],
				termPowerDifferences[index + 1])
		}
		symmetriesFound.Multifold = currentGreatestCommonDenominator
	}
}

// getGreatestCommonDenominator finds the largest integer that divides into
//   integers a and b, leaving 0 behind.
func getGreatestCommonDenominator(a, b int) int {
	if a == b {
		return a
	}

	larger := a
	smaller := b

	remainder := larger % smaller
	if remainder == 0 {
		return smaller
	}
	return getGreatestCommonDenominator(smaller, remainder)
}

// CalculateExponentTerm calculates (z^power * zConj^conjugatePower)
//   where z is a complex number, zConj is the complex conjugate
//   and power and conjugatePower are integers.
func CalculateExponentTerm(z complex128, power1, power2 int, scale complex128, ignoreComplexConjugate bool) complex128 {
	zRaisedToN := cmplx.Pow(z, complex(float64(power1), 0))
	if ignoreComplexConjugate {
		return zRaisedToN * scale
	}

	complexConjugate := complex(real(z), -1 * imag(z))
	complexConjugateRaisedToM := cmplx.Pow(complexConjugate, complex(float64(power2), 0))
	return zRaisedToN * complexConjugateRaisedToM * scale
}

// NewRosetteFormulaFromYAML reads the data and returns a RosetteFormula from it.
func NewRosetteFormulaFromYAML(data []byte) (*RosetteFormula, error) {
	return newRosetteFormulaFromDatastream(data, yaml.Unmarshal)
}

// NewRosetteFormulaFromJSON reads the data and returns a RosetteFormula from it.
func NewRosetteFormulaFromJSON(data []byte) (*RosetteFormula, error) {
	return newRosetteFormulaFromDatastream(data, json.Unmarshal)
}

// RosetteFormulaMarshalable can be marshaled and mapped to a RosetteFormula object.
type RosetteFormulaMarshalable struct {
	Terms []*ZExponentialFormulaTermMarshalable
}

// newRosetteFormulaFromDatastream consumes a given bytestream and tries to create a new object from it.
func newRosetteFormulaFromDatastream(data []byte, unmarshal utility.UnmarshalFunc) (*RosetteFormula, error) {
	var unmarshalError error
	var rosetteFormulaMarshal RosetteFormulaMarshalable
	unmarshalError = unmarshal(data, &rosetteFormulaMarshal)

	if unmarshalError != nil {
		return nil, unmarshalError
	}

	rosetteFormula := NewRosetteFormulaFromMarshalObject(rosetteFormulaMarshal)
	return rosetteFormula, nil
}

// NewRosetteFormulaFromMarshalObject converts the marshalled object to a usable one.
func NewRosetteFormulaFromMarshalObject(marshalObject RosetteFormulaMarshalable) *RosetteFormula {
	terms := []*ZExponentialFormulaTerm{}
	for _, termMarshal := range marshalObject.Terms {
		newTerm := newZExponentialFormulaTermFromMarshalObject(*termMarshal)
		terms = append(terms, newTerm)
	}
	return &RosetteFormula{Terms: terms}
}
