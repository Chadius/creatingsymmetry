package command

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"wallpaper/entities/formula/frieze"
	"wallpaper/entities/formula/rosette"
	"wallpaper/entities/formula/wavepacket"
	"wallpaper/entities/utility"
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

// CreateWallpaperCommand records the desired command to generate.
type CreateWallpaperCommand struct {
	SampleSpace				  ComplexNumberCorners               `json:"sample_space" yaml:"sample_space"`
	OutputImageSize			  WidthHeightDimensions              `json:"output_size" yaml:"output_size"`
	SampleSourceFilename	  string                                `json:"sample_source_filename" yaml:"sample_source_filename"`
	OutputFilename			  string                              `json:"output_filename" yaml:"output_filename"`
	ColorValueSpace			  ComplexNumberCorners               `json:"color_value_space" yaml:"color_value_space"`
	RosetteFormula			  *rosette.Formula                    `json:"rosette_formula" yaml:"rosette_formula"`
	FriezeFormula			  *frieze.Formula                      `json:"frieze_formula" yaml:"frieze_formula"`
	HexagonalWallpaperFormula *wavepacket.HexagonalWallpaperFormula `json:"hexagonal_wallpaper_formula" yaml:"hexagonal_wallpaper_formula"`
	SquareWallpaperFormula *wavepacket.SquareWallpaperFormula       `json:"square_wallpaper_formula" yaml:"square_wallpaper_formula"`
	RhombicWallpaperFormula *wavepacket.RhombicWallpaperFormula            `json:"rhombic_wallpaper_formula" yaml:"rhombic_wallpaper_formula"`
	RectangularWallpaperFormula *wavepacket.RectangularWallpaperFormula            `json:"rectangular_wallpaper_formula" yaml:"rectangular_wallpaper_formula"`
	GenericWallpaperFormula *wavepacket.GenericWallpaperFormula            `json:"generic_wallpaper_formula" yaml:"generic_wallpaper_formula"`
}

// CreateWallpaperCommandMarshal can be marshaled and converted to a CreateWallpaperCommand
type CreateWallpaperCommandMarshal struct {
	SampleSpace				ComplexNumberCorners                  `json:"sample_space" yaml:"sample_space"`
	OutputImageSize			WidthHeightDimensions                 `json:"output_size" yaml:"output_size"`
	SampleSourceFilename	string                                   `json:"sample_source_filename" yaml:"sample_source_filename"`
	OutputFilename			string                                 `json:"output_filename" yaml:"output_filename"`
	ColorValueSpace			ComplexNumberCorners                  `json:"color_value_space" yaml:"color_value_space"`
	RosetteFormula			*rosette.MarshaledFormula              `json:"rosette_formula" yaml:"rosette_formula"`
	FriezeFormula			*frieze.MarshaledFormula                `json:"frieze_formula" yaml:"frieze_formula"`
	HexagonalWallpaperFormula *wavepacket.WallpaperFormulaMarshalled `json:"hexagonal_wallpaper_formula" yaml:"hexagonal_wallpaper_formula"`
	SquareWallpaperFormula *wavepacket.WallpaperFormulaMarshalled    `json:"square_wallpaper_formula" yaml:"square_wallpaper_formula"`
	RhombicWallpaperFormula *wavepacket.RhombicWallpaperFormulaMarshalled       `json:"rhombic_wallpaper_formula" yaml:"rhombic_wallpaper_formula"`
	RectangularWallpaperFormula *wavepacket.RectangularWallpaperFormulaMarshalled            `json:"rectangular_wallpaper_formula" yaml:"rectangular_wallpaper_formula"`
	GenericWallpaperFormula *wavepacket.GenericWallpaperFormulaMarshalled            `json:"generic_wallpaper_formula" yaml:"generic_wallpaper_formula"`
}

// NewCreateWallpaperCommandFromYAML reads the data and returns a CreateWallpaperCommand from it.
func NewCreateWallpaperCommandFromYAML(data []byte) (*CreateWallpaperCommand, error) {
	return newCreateWallpaperCommandFromDatastream(data, yaml.Unmarshal)
}

// NewCreateWallpaperCommandFromJSON reads the data and returns a CreateWallpaperCommand from it.
func NewCreateWallpaperCommandFromJSON(data []byte) (*CreateWallpaperCommand, error) {
	return newCreateWallpaperCommandFromDatastream(data, json.Unmarshal)
}

// newCreateWallpaperCommandFromDatastream consumes a given bytestream and tries to create a new object from it.
func newCreateWallpaperCommandFromDatastream(data []byte, unmarshal utility.UnmarshalFunc) (*CreateWallpaperCommand, error) {
	var unmarshalError error
	var commandToCreateMarshal CreateWallpaperCommandMarshal
	unmarshalError = unmarshal(data, &commandToCreateMarshal)

	if unmarshalError != nil {
		return nil, unmarshalError
	}

	commandToCreate := &CreateWallpaperCommand{
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

	if commandToCreateMarshal.HexagonalWallpaperFormula != nil {
		commandToCreate.HexagonalWallpaperFormula = wavepacket.NewHexagonalWallpaperFormulaFromMarshalObject(*commandToCreateMarshal.HexagonalWallpaperFormula)
	}

	if commandToCreateMarshal.SquareWallpaperFormula != nil {
		commandToCreate.SquareWallpaperFormula = wavepacket.NewSquareWallpaperFormulaFromMarshalObject(*commandToCreateMarshal.SquareWallpaperFormula)
	}

	if commandToCreateMarshal.RhombicWallpaperFormula != nil {
		commandToCreate.RhombicWallpaperFormula = wavepacket.NewRhombicWallpaperFormulaFromMarshalObject(*commandToCreateMarshal.RhombicWallpaperFormula)
	}

	if commandToCreateMarshal.RectangularWallpaperFormula != nil {
		commandToCreate.RectangularWallpaperFormula = wavepacket.NewRectangularWallpaperFormulaFromMarshalObject(*commandToCreateMarshal.RectangularWallpaperFormula)
	}

	if commandToCreateMarshal.GenericWallpaperFormula != nil {
		commandToCreate.GenericWallpaperFormula = wavepacket.NewGenericWallpaperFormulaFromMarshalObject(*commandToCreateMarshal.GenericWallpaperFormula)
	}

	return commandToCreate, nil
}