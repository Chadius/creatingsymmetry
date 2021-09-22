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
	MinX	float64	`json:"minx" yaml:"minx"`
	MinY	float64	`json:"miny" yaml:"miny"`
	MaxX	float64	`json:"maxx" yaml:"maxx"`
	MaxY	float64	`json:"maxy" yaml:"maxy"`
}

// WidthHeightDimensions is a width + height combination.
type WidthHeightDimensions struct {
	Width 	int `json:"width" yaml:"width"`
	Height	int `json:"height" yaml:"height"`
}

// CreateSymmetryPattern records the desired command to generate.
type CreateSymmetryPattern struct {
	SampleSpace				  ComplexNumberCorners               `json:"sample_space" yaml:"sample_space"`
	OutputImageSize			  WidthHeightDimensions              `json:"output_size" yaml:"output_size"`
	SampleSourceFilename	  string                                `json:"sample_source_filename" yaml:"sample_source_filename"`
	OutputFilename			  string                              `json:"output_filename" yaml:"output_filename"`
	ColorValueSpace			  ComplexNumberCorners               `json:"color_value_space" yaml:"color_value_space"`

	RosetteFormula			  *rosette.Formula                    `json:"rosette_formula" yaml:"rosette_formula"`
	FriezeFormula			  *frieze.Formula                      `json:"frieze_formula" yaml:"frieze_formula"`
	LatticePattern *wallpaper.Formula `json:"lattice_pattern" yaml:"lattice_pattern"`
}

// CreateWallpaperCommandMarshal can be marshaled and converted to a CreateSymmetryPattern
type CreateWallpaperCommandMarshal struct {
	SampleSpace				ComplexNumberCorners                  `json:"sample_space" yaml:"sample_space"`
	OutputImageSize			WidthHeightDimensions                 `json:"output_size" yaml:"output_size"`
	SampleSourceFilename	string                                   `json:"sample_source_filename" yaml:"sample_source_filename"`
	OutputFilename			string                                 `json:"output_filename" yaml:"output_filename"`
	ColorValueSpace			ComplexNumberCorners                  `json:"color_value_space" yaml:"color_value_space"`

	RosetteFormula			*rosette.MarshaledFormula              `json:"rosette_formula" yaml:"rosette_formula"`
	FriezeFormula			*frieze.MarshaledFormula                `json:"frieze_formula" yaml:"frieze_formula"`
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
		SampleSpace:          commandToCreateMarshal.SampleSpace,
		OutputImageSize:      commandToCreateMarshal.OutputImageSize,
		SampleSourceFilename: commandToCreateMarshal.SampleSourceFilename,
		OutputFilename:       commandToCreateMarshal.OutputFilename,
		ColorValueSpace:      commandToCreateMarshal.ColorValueSpace,
	}

	if commandToCreateMarshal.RosetteFormula != nil {
		commandToCreate.RosetteFormula  = rosette.NewRosetteFormulaFromMarshalObject(*commandToCreateMarshal.RosetteFormula)
	}

	if commandToCreateMarshal.FriezeFormula != nil {
		commandToCreate.FriezeFormula  = frieze.NewFriezeFormulaFromMarshalObject(*commandToCreateMarshal.FriezeFormula)
	}

	if commandToCreateMarshal.LatticePattern != nil {
		commandToCreate.LatticePattern = wallpaper.NewFormulaFromMarshalObject(*commandToCreateMarshal.LatticePattern)
	}

	return commandToCreate, nil
}