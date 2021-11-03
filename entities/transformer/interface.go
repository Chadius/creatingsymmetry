package transformer

import (
	"github.com/Chadius/creating-symmetry/entities/command"
	"github.com/Chadius/creating-symmetry/entities/imageoutput"
	"image"
)

// Transformer turns one image stream into another
type Transformer interface {
	Transform(setting *Settings) *image.NRGBA
}

// Settings are required to transform a given image.
type Settings struct {
	PatternViewportXMin float64
	PatternViewportXMax float64
	PatternViewportYMin float64
	PatternViewportYMax float64
	InputImage image.Image
	Formula *command.CreateSymmetryPattern
	CoordinateThreshold imageoutput.CoordinateThreshold
	Eyedropper imageoutput.Eyedropper
	OutputWidth int
	OutputHeight int
}