package formula

// Builder is used to create formula objects.
type Builder struct {
	formulaType string
	formulaLevelTerms []Term
	latticeWidth float64
	latticeHeight float64
	wavePackets []WavePacket
	desiredSymmetry Symmetry
}

// NewBuilder returns a new object used to Build Formula objects.
func NewBuilder() *Builder {
	return &Builder{
		formulaType: "identity",
		formulaLevelTerms: []Term{},
		latticeWidth: 0.0,
		latticeHeight: 0.0,
		wavePackets: []WavePacket{},
		desiredSymmetry: P1,
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

// LatticeWidth sets the lattice width for wallpaper based patterns.
func (b *Builder) LatticeWidth(latticeWidth float64) *Builder {
	b.latticeWidth = latticeWidth
	return b
}

// Rectangular sets the formula as a rectangular formula.
func (b *Builder) Rectangular() *Builder {
	b.formulaType = "rectangular"
	return b
}

// Square sets the formula as a square formula.
func (b *Builder) Square() *Builder {
	b.formulaType = "square"
	return b
}

// Hexagonal sets the formula as a hexagonal formula.
func (b *Builder) Hexagonal() *Builder {
	b.formulaType = "hexagonal"
	return b
}

// Rhombic sets the formula as a rhombic formula.
func (b *Builder) Rhombic() *Builder {
	b.formulaType = "rhombic"
	return b
}

// Generic sets the formula as a generic formula.
func (b *Builder) Generic() *Builder {
	b.formulaType = "generic"
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

// DesiredSymmetry sets the desired symmetry
func (b *Builder)DesiredSymmetry(symmetry Symmetry) *Builder {
	b.desiredSymmetry = symmetry
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
	if b.formulaType == "square" {
		return NewSquareFormula(b.wavePackets)
	}
	if b.formulaType == "hexagonal" {
		return NewHexagonalFormula(b.wavePackets, b.desiredSymmetry)
	}
	if b.formulaType == "rhombic" {
		formula, err := NewRhombicFormula(b.wavePackets, b.latticeHeight)
		if formula == nil {
			return &Identity{}, err
		}
		return formula, err
	}
	if b.formulaType == "generic" {
		formula, err := NewGenericFormula(b.wavePackets, b.latticeWidth, b.latticeHeight)
		if formula == nil {
			return &Identity{}, err
		}
		return formula, err
	}
	return &Identity{}, nil
}
