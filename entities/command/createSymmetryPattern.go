package command

import (
	"encoding/json"
	"github.com/Chadius/creating-symmetry/entities/formula"
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

// PixelCorners note the sides of a rectangle in integer space.
type PixelCorners struct {
	LeftSide   int `json:"left" yaml:"left"`
	RightSide  int `json:"right" yaml:"right"`
	TopSide    int `json:"top" yaml:"top"`
	BottomSide int `json:"bottom" yaml:"bottom"`
}

// CreateSymmetryPattern records the desired command to generate.
type CreateSymmetryPattern struct {
	PatternViewport     ComplexNumberCorners `json:"pattern_viewport" yaml:"pattern_viewport"`
	CoordinateThreshold ComplexNumberCorners `json:"coordinate_threshold" yaml:"coordinate_threshold"`
	Eyedropper          *PixelCorners        `json:"eyedropper" yaml:"eyedropper"`

	Formula formula.Arbitrary `json:"formula" yaml:"formula"`
}

// CreateWallpaperCommandMarshal can be marshaled and converted to a CreateSymmetryPattern
type CreateWallpaperCommandMarshal struct {
	PatternViewport     ComplexNumberCorners `json:"pattern_viewport" yaml:"pattern_viewport"`
	CoordinateThreshold ComplexNumberCorners `json:"coordinate_threshold" yaml:"coordinate_threshold"`
	Eyedropper          *PixelCorners        `json:"eyedropper" yaml:"eyedropper"`

	Formula *formula.BuilderOptionMarshal `json:"formula" yaml:"formula"`
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
		Eyedropper:          commandToCreateMarshal.Eyedropper,
	}

	if commandToCreateMarshal.Formula != nil {
		commandToCreate.Formula, _ = formula.NewBuilder().WithMarshalOptions(*commandToCreateMarshal.Formula).Build()
	}

	return commandToCreate, nil
}
