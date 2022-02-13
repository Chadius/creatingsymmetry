package formula

import (
	"encoding/json"
	"github.com/chadius/creatingsymmetry/entities/utility"
	"gopkg.in/yaml.v2"
)

// Builder is used to create formula objects.
type Builder struct {
	formulaType       string
	formulaLevelTerms []Term
	latticeWidth      float64
	latticeHeight     float64
	wavePackets       []WavePacket
	desiredSymmetry   Symmetry
}

// NewBuilder returns a new object used to Build Formula objects.
func NewBuilder() *Builder {
	return &Builder{
		formulaType:       "identity",
		formulaLevelTerms: []Term{},
		latticeWidth:      0.0,
		latticeHeight:     0.0,
		wavePackets:       []WavePacket{},
		desiredSymmetry:   P1,
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
func (b *Builder) DesiredSymmetry(symmetry Symmetry) *Builder {
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
		formula, err := NewRectangularFormula(b.wavePackets, b.latticeHeight, b.desiredSymmetry)
		if formula == nil {
			return &Identity{}, err
		}
		return formula, err
	}
	if b.formulaType == "square" {
		formula, err := NewSquareFormula(b.wavePackets, b.desiredSymmetry)
		if formula == nil {
			return &Identity{}, err
		}
		return formula, err
	}
	if b.formulaType == "hexagonal" {
		formula, err := NewHexagonalFormula(b.wavePackets, b.desiredSymmetry)
		if formula == nil {
			return &Identity{}, err
		}
		return formula, err
	}
	if b.formulaType == "rhombic" {
		formula, err := NewRhombicFormula(b.wavePackets, b.latticeHeight, b.desiredSymmetry)
		if formula == nil {
			return &Identity{}, err
		}
		return formula, err
	}
	if b.formulaType == "generic" {
		formula, err := NewGenericFormula(b.wavePackets, b.latticeWidth, b.latticeHeight, b.desiredSymmetry)
		if formula == nil {
			return &Identity{}, err
		}
		return formula, err
	}
	return &Identity{}, nil
}

// UsingYAMLData updates the builder, given data
func (b *Builder) UsingYAMLData(data []byte) *Builder {
	return b.usingByteStream(data, yaml.Unmarshal)
}

// UsingJSONData updates the builder, given data
func (b *Builder) UsingJSONData(data []byte) *Builder {
	return b.usingByteStream(data, json.Unmarshal)
}

// BuilderOptionMarshal is a flattened representation of all Builder options.
type BuilderOptionMarshal struct {
	Type            string              `json:"type" yaml:"type"`
	Terms           []TermMarshal       `json:"terms" yaml:"terms"`
	LatticeWidth    float64             `json:"lattice_width" yaml:"lattice_width"`
	LatticeHeight   float64             `json:"lattice_height" yaml:"lattice_height"`
	WavePackets     []WavePacketMarshal `json:"wave_packets" yaml:"wave_packets"`
	DesiredSymmetry Symmetry            `json:"desired_symmetry" yaml:"desired_symmetry"`
}

func (b *Builder) usingByteStream(data []byte, unmarshal utility.UnmarshalFunc) *Builder {
	var unmarshalError error
	var marshaledOptions BuilderOptionMarshal

	unmarshalError = unmarshal(data, &marshaledOptions)

	if unmarshalError != nil {
		return b
	}

	return b.WithMarshalOptions(marshaledOptions)
}

// WithMarshalOptions uses the options to populate the builder.
func (b *Builder) WithMarshalOptions(marshaledOptions BuilderOptionMarshal) *Builder {
	if marshaledOptions.Type == "rosette" {
		b.Rosette()
	}

	if marshaledOptions.Type == "frieze" {
		b.Frieze()
	}

	if marshaledOptions.Type == "generic" {
		b.Generic()
	}

	if marshaledOptions.Type == "rectangular" {
		b.Rectangular()
	}

	if marshaledOptions.Type == "rhombic" {
		b.Rhombic()
	}

	if marshaledOptions.Type == "hexagonal" {
		b.Hexagonal()
	}

	if marshaledOptions.Type == "square" {
		b.Square()
	}

	b.LatticeWidth(marshaledOptions.LatticeWidth).
		LatticeHeight(marshaledOptions.LatticeHeight)

	for _, wavePacketMarshal := range marshaledOptions.WavePackets {
		b.AddWavePacket(
			NewWavePacketBuilder().WithMarshalOptions(wavePacketMarshal).Build(),
		)
	}

	if marshaledOptions.DesiredSymmetry != "" {
		b.DesiredSymmetry(marshaledOptions.DesiredSymmetry)
	}

	for _, termMarshal := range marshaledOptions.Terms {
		newTerm := NewTermBuilder().WithMarshalOptions(termMarshal).Build()
		b.AddTerm(newTerm)
	}

	return b
}
