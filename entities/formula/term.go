package formula

import (
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
)

// Term objects help shape the calculation of every formula.
type Term struct {
	Multiplier complex128
	PowerN     int
	PowerM     int
	IgnoreComplexConjugate bool
	CoefficientRelationships []coefficient.Relationship
}

// TermBuilder is used to create formula objects.
type TermBuilder struct {
	multiplier complex128
	powerN int
	powerM int
	coefficientRelationships []coefficient.Relationship
	ignoreComplexConjugate bool
}

// NewTermBuilder returns a new object used to build Term objects.
func NewTermBuilder() *TermBuilder {
	return &TermBuilder{
		multiplier: complex(0,0),
		powerN:  0,
		powerM:  0,
		coefficientRelationships: []coefficient.Relationship{},
		ignoreComplexConjugate: false,
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

// AddCoefficientRelationship sets the field.
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
	return &Term{
		Multiplier: t.multiplier,
		PowerN: t.powerN,
		PowerM: t.powerM,
		IgnoreComplexConjugate: t.ignoreComplexConjugate,
		CoefficientRelationships: t.coefficientRelationships,
	}
}

