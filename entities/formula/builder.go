package formula

// Builder is used to create formula objects.
type Builder struct {
	formulaType string
	formulaLevelTerms []Term
}

// NewBuilder returns a new object used to Build Formula objects.
func NewBuilder() *Builder {
	return &Builder{
		formulaType: "identity",
		formulaLevelTerms: []Term{},
	}
}

// Rosette sets the formula as a rosette formula.
func (b *Builder) Rosette() *Builder {
	b.formulaType = "rosette"
	return b
}

// Frieze sets the formula as a frieze formula.
func (b *Builder) Frieze() *Builder {
	b.formulaType = "frieze"
	return b
}

// AddTerm adds a term to the formula.
func (b *Builder) AddTerm(term *Term) *Builder {
	b.formulaLevelTerms = append(b.formulaLevelTerms, *term)
	return b
}

// Build creates a new Formula object.
func (b *Builder) Build() Arbitrary {
	if b.formulaType == "rosette" {
		return NewRosetteFormula(b.formulaLevelTerms)
	}
	if b.formulaType == "frieze" {
		return NewFriezeFormula(b.formulaLevelTerms)
	}
	return &Identity{}
}
