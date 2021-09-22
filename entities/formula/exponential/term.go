package exponential

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	"github.com/Chadius/creating-symmetry/entities/utility"
)

// TermMarshalable is an object that can be easily marshaled to and from data streams.
type TermMarshalable struct {
	Multiplier					utility.ComplexNumberForMarshal	`json:"multiplier" yaml:"multiplier"`
	PowerN						int								`json:"power_n" yaml:"power_n"`
	PowerM						int								`json:"power_m" yaml:"power_m"`
	IgnoreComplexConjugate		bool							`json:"ignore_complex_conjugate" yaml:"ignore_complex_conjugate"`
	CoefficientRelationships	[]coefficient.Relationship		`json:"coefficient_relationships" yaml:"coefficient_relationships"`
}

// RosetteFriezeTerm is used in Friezes and Rosettes, applying different calculations to them.
type RosetteFriezeTerm struct {
	Multiplier					complex128
	PowerN						int
	PowerM						int
	// IgnoreComplexConjugate will make sure zConjugate is not used in this calculation
	//    (effectively setting it to 1 + 0i)
	IgnoreComplexConjugate		bool
	// CoefficientRelationships has a list of locked coefficient pairings. These locks are
	//   used to generate similar locked terms. Relationships affect PowerN, PowerM and Multiplier.
	CoefficientRelationships	[]coefficient.Relationship
}

// NewTermFromYAML reads the data and returns a formula term from it.
func NewTermFromYAML(data []byte) (*RosetteFriezeTerm, error) {
	return newTermFromDatastream(data, yaml.Unmarshal)
}

// NewTermFromJSON reads the data and returns a formula term from it.
func NewTermFromJSON(data []byte) (*RosetteFriezeTerm, error) {
	return newTermFromDatastream(data, json.Unmarshal)
}

//newTermFromDatastream consumes a given bytestream and tries to create a new object from it.
func newTermFromDatastream(data []byte, unmarshal utility.UnmarshalFunc) (*RosetteFriezeTerm, error) {
	var unmarshalError error
	var formulaTermMarshal TermMarshalable
	unmarshalError = unmarshal(data, &formulaTermMarshal)

	if unmarshalError != nil {
		return nil, unmarshalError
	}

	formulaTerm := NewTermFromMarshalObject(formulaTermMarshal)
	return formulaTerm, nil
}

// NewTermFromMarshalObject creates an object from the marshaled object.
func NewTermFromMarshalObject(marshalObject TermMarshalable) *RosetteFriezeTerm {
	return &RosetteFriezeTerm{
		Multiplier:             	complex(marshalObject.Multiplier.Real, marshalObject.Multiplier.Imaginary),
		PowerN:                 	marshalObject.PowerN,
		PowerM:                 	marshalObject.PowerM,
		IgnoreComplexConjugate:		marshalObject.IgnoreComplexConjugate,
		CoefficientRelationships:	marshalObject.CoefficientRelationships,
	}
}
