package formula

import (
	"encoding/json"
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	"github.com/Chadius/creating-symmetry/entities/utility"
	"gopkg.in/yaml.v2"
	"math"
	"math/cmplx"
)

// Term objects help shape the calculation of every formula.
type Term struct {
	Multiplier               complex128
	PowerN                   int
	PowerM                   int
	IgnoreComplexConjugate   bool
	CoefficientRelationships []coefficient.Relationship
}

// NewTerm returns a new Term object.
func NewTerm(multiplier complex128, powerN, powerM int, ignoreComplexConjugate bool, coefficientRelationships []coefficient.Relationship) *Term {
	return &Term{
		Multiplier:               multiplier,
		PowerN:                   powerN,
		PowerM:                   powerM,
		IgnoreComplexConjugate:   ignoreComplexConjugate,
		CoefficientRelationships: coefficientRelationships,
	}
}

// NewTermWithMultiplierAndPowers returns a new Term object using just the multiplier, PowerN and PowerM.
func NewTermWithMultiplierAndPowers(multiplier complex128, powerN, powerM int) *Term {
	return &Term{
		Multiplier: multiplier,
		PowerN:     powerN,
		PowerM:     powerM,
	}
}

// CalculateInLatticeCoordinates (z) = e ^ (2 PI i * (nX + mY))
//  where n amd m are PowerN and PowerM,
//  and TransformedX and TransformedY are the real and imag parts of (zInLatticeCoordinates)
func (term Term) CalculateInLatticeCoordinates(zInLatticeCoordinates complex128) complex128 {
	powerMultiplier := (float64(term.PowerN) * real(zInLatticeCoordinates)) +
		(float64(term.PowerM) * imag(zInLatticeCoordinates))
	expo := cmplx.Exp(complex(0, 2.0*math.Pi*powerMultiplier))
	return expo
}

// PowerSumIsEven returns true if the sum of the term powers is divisible by 2.
func (term Term) PowerSumIsEven() bool {
	return (term.PowerM+term.PowerN)%2 == 0
}

// TermBuilder is used to create formula objects.
type TermBuilder struct {
	multiplier               complex128
	powerN                   int
	powerM                   int
	coefficientRelationships []coefficient.Relationship
	ignoreComplexConjugate   bool
}

// NewTermBuilder returns a new object used to build Term objects.
func NewTermBuilder() *TermBuilder {
	return &TermBuilder{
		multiplier:               complex(1, 0),
		powerN:                   0,
		powerM:                   0,
		coefficientRelationships: []coefficient.Relationship{},
		ignoreComplexConjugate:   false,
	}
}

// Multiplier sets the field.
func (t *TermBuilder) Multiplier(multiplier complex128) *TermBuilder {
	t.multiplier = multiplier
	return t
}

// PowerN sets the field.
func (t *TermBuilder) PowerN(coefficient int) *TermBuilder {
	t.powerN = coefficient
	return t
}

// PowerM sets the field.
func (t *TermBuilder) PowerM(coefficient int) *TermBuilder {
	t.powerM = coefficient
	return t
}

// AddCoefficientRelationship adds the coefficient.
func (t *TermBuilder) AddCoefficientRelationship(coefficient coefficient.Relationship) *TermBuilder {
	t.coefficientRelationships = append(t.coefficientRelationships, coefficient)
	return t
}

// IgnoreComplexConjugate sets the field.
func (t *TermBuilder) IgnoreComplexConjugate() *TermBuilder {
	t.ignoreComplexConjugate = true
	return t
}

// Build creates a new Term object.
func (t *TermBuilder) Build() *Term {
	return NewTerm(t.multiplier, t.powerN, t.powerM, t.ignoreComplexConjugate, t.coefficientRelationships)
}

// UsingYAMLData updates the term, given data
func (t *TermBuilder) UsingYAMLData(data []byte) *TermBuilder {
	return t.usingByteStream(data, yaml.Unmarshal)
}

// UsingJSONData updates the term, given data
func (t *TermBuilder) UsingJSONData(data []byte) *TermBuilder {
	return t.usingByteStream(data, json.Unmarshal)
}

// TermMarshal is a representation of a term object
type TermMarshal struct {
	Multiplier               *utility.ComplexNumberForMarshal `json:"multiplier" yaml:"multiplier"`
	PowerN                   int                              `json:"power_n" yaml:"power_n"`
	PowerM                   int                              `json:"power_m" yaml:"power_m"`
	CoefficientRelationships []coefficient.Relationship       `json:"coefficient_relationships" yaml:"coefficient_relationships"`
	IgnoreComplexConjugate   bool                             `json:"ignore_complex_conjugate" yaml:"ignore_complex_conjugate"`
}

func (t *TermBuilder) usingByteStream(data []byte, unmarshal utility.UnmarshalFunc) *TermBuilder {
	var unmarshalError error
	var marshaledOptions TermMarshal

	unmarshalError = unmarshal(data, &marshaledOptions)

	if unmarshalError != nil {
		return t
	}
	return t.WithMarshalOptions(marshaledOptions)
}

func (t *TermBuilder) WithMarshalOptions(marshaledOptions TermMarshal) *TermBuilder {
	t.PowerN(marshaledOptions.PowerN)
	t.PowerM(marshaledOptions.PowerM)

	if marshaledOptions.Multiplier != nil {
		t.Multiplier(complex(marshaledOptions.Multiplier.Real, marshaledOptions.Multiplier.Imaginary))
	} else {
		t.Multiplier(complex(1, 0))
	}

	if marshaledOptions.IgnoreComplexConjugate {
		t.IgnoreComplexConjugate()
	}

	for _, coefficient := range marshaledOptions.CoefficientRelationships {
		t.AddCoefficientRelationship(coefficient)
	}

	return t
}
