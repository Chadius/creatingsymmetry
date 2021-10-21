package command

import (
	"encoding/json"
	"github.com/Chadius/creating-symmetry/entities/formula/frieze"
	"github.com/Chadius/creating-symmetry/entities/formula/rosette"
	"github.com/Chadius/creating-symmetry/entities/formula/wallpaper"
	"github.com/Chadius/creating-symmetry/entities/utility"
	"gopkg.in/yaml.v2"
)

// ComplexNumberCorners notes the sides of a rectangle drawn in the complex space.
type ComplexNumberCorners struct {
	XMin float64 `json:"x_min" yaml:"x_min"`
	YMin float64 `json:"y_min" yaml:"y_min"`
	XMax float64 `json:"x_max" yaml:"x_max"`
	YMax float64 `json:"y_max" yaml:"y_max"`
}

// CreateSymmetryPattern records the desired command to generate.
type CreateSymmetryPattern struct {
	PatternViewport     ComplexNumberCorners `json:"pattern_viewport" yaml:"pattern_viewport"`
	CoordinateThreshold ComplexNumberCorners `json:"coordinate_threshold" yaml:"coordinate_threshold"`

	RosetteFormula *rosette.Formula   `json:"rosette_formula" yaml:"rosette_formula"`
	FriezeFormula  *frieze.Formula    `json:"frieze_formula" yaml:"frieze_formula"`
	LatticePattern *wallpaper.Formula `json:"lattice_pattern" yaml:"lattice_pattern"`
}

// CreateWallpaperCommandMarshal can be marshaled and converted to a CreateSymmetryPattern
type CreateWallpaperCommandMarshal struct {
	PatternViewport     ComplexNumberCorners `json:"pattern_viewport" yaml:"pattern_viewport"`
	CoordinateThreshold ComplexNumberCorners `json:"coordinate_threshold" yaml:"coordinate_threshold"`

	RosetteFormula *rosette.MarshaledFormula `json:"rosette_formula" yaml:"rosette_formula"`
	FriezeFormula  *frieze.MarshaledFormula  `json:"frieze_formula" yaml:"frieze_formula"`
	LatticePattern *wallpaper.FormulaMarshal `json:"lattice_pattern" yaml:"lattice_pattern"`
}

// NewCreateWallpaperCommandFromYAML reads the data and returns a CreateSymmetryPattern from it.
func NewCreateWallpaperCommandFromYAML(data []byte) (*CreateSymmetryPattern, error) {
	return newCreateWallpaperCommandFromDatastream(data, yaml.Unmarshal)
}

// NewCreateWallpaperCommandFromJSON reads the data and returns a CreateSymmetryPattern from it.
func NewCreateWallpaperCommandFromJSON(data []byte) (*CreateSymmetryPattern, error) {
	return newCreateWallpaperCommandFromDatastream(data, json.Unmarshal)
}

// newCreateWallpaperCommandFromDatastream consumes a given bytestream and tries to create a new object from it.
func newCreateWallpaperCommandFromDatastream(data []byte, unmarshal utility.UnmarshalFunc) (*CreateSymmetryPattern, error) {
	var unmarshalError error
	var commandToCreateMarshal CreateWallpaperCommandMarshal
	unmarshalError = unmarshal(data, &commandToCreateMarshal)

	if unmarshalError != nil {
		return nil, unmarshalError
	}

	commandToCreate := &CreateSymmetryPattern{
		PatternViewport:     commandToCreateMarshal.PatternViewport,
		CoordinateThreshold: commandToCreateMarshal.CoordinateThreshold,
	}

	if commandToCreateMarshal.RosetteFormula != nil {
		commandToCreate.RosetteFormula = rosette.NewRosetteFormulaFromMarshalObject(*commandToCreateMarshal.RosetteFormula)
	}

	if commandToCreateMarshal.FriezeFormula != nil {
		commandToCreate.FriezeFormula = frieze.NewFriezeFormulaFromMarshalObject(*commandToCreateMarshal.FriezeFormula)
	}

	if commandToCreateMarshal.LatticePattern != nil {
		commandToCreate.LatticePattern = wallpaper.NewFormulaFromMarshalObject(*commandToCreateMarshal.LatticePattern)
	}

	return commandToCreate, nil
}
