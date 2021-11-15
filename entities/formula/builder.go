package formula

// Builder is used to create formula objects.
type Builder struct {
	formulaType string
	formulaLevelTerms []Term
	latticeHeight float64
	wavePackets []WavePacket
}

// NewBuilder returns a new object used to Build Formula objects.
func NewBuilder() *Builder {
	return &Builder{
		formulaType: "identity",
		formulaLevelTerms: []Term{},
		latticeHeight: 0.0,
		wavePackets: []WavePacket{},
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

// LatticeHeight sets the lattice height for wallpaper based patterns.
func (b *Builder) LatticeHeight(latticeHeight float64) *Builder {
	b.latticeHeight = latticeHeight
	return b
}

// Rectangular sets the formula as a frieze formula.
func (b *Builder) Rectangular() *Builder {
	b.formulaType = "rectangular"
	return b
}

// AddTerm adds a term to the formula.
func (b *Builder) AddTerm(term *Term) *Builder {
	b.formulaLevelTerms = append(b.formulaLevelTerms, *term)
	return b
}

// AddWavePacket adds a wave packet to the formula.
func (b *Builder) AddWavePacket(packet *WavePacket) *Builder {
	b.wavePackets = append(b.wavePackets, *packet)
	return b
}

// Build creates a new Formula object.
func (b *Builder) Build() (Arbitrary, error) {
	if b.formulaType == "rosette" {
		return NewRosetteFormula(b.formulaLevelTerms)
	}
	if b.formulaType == "frieze" {
		return NewFriezeFormula(b.formulaLevelTerms)
	}
	if b.formulaType == "rectangular" {
		formula, err := NewRectangularFormula(b.wavePackets, b.latticeHeight)
		if formula == nil {
			return &Identity{}, err
		}
		return formula, err
	}
	return &Identity{}, nil
}
