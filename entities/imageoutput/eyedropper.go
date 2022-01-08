package imageoutput

import (
	"github.com/Chadius/creating-symmetry/entities/mathutility"
	"image"
	"image/color"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . Eyedropper
// Eyedropper uses a set of coordinates to choose colors from a source image.
type Eyedropper interface {
	ConvertCoordinatesToColors(collection *CoordinateCollection) *[]color.Color
}

// RectangularEyedropper will sample transformed coordinates against a rectangular portion of the source image.
type RectangularEyedropper struct {
	leftBoundary   int
	rightBoundary  int
	topBoundary    int
	bottomBoundary int
	sourceImage    image.Image
}

// LeftSide returns the left side of the boundary.
func (e *RectangularEyedropper) LeftSide() int {
	return e.leftBoundary
}

// RightSide returns the right side of the boundary.
func (e *RectangularEyedropper) RightSide() int {
	return e.rightBoundary
}

// TopSide returns the top side of the boundary.
func (e *RectangularEyedropper) TopSide() int {
	return e.topBoundary
}

// BottomSide returns the bottom side of the boundary.
func (e *RectangularEyedropper) BottomSide() int {
	return e.bottomBoundary
}

//Image returns the source image
func (e *RectangularEyedropper) Image() image.Image {
	return e.sourceImage
}

// ConvertCoordinatesToColors uses the collection of coordinates, maps it to the eyedropper range,
//   and samples the color in the source image at that location.
//   if the coordinate is mapped outside the source image, it will turn transparent.
func (e *RectangularEyedropper) ConvertCoordinatesToColors(collection *CoordinateCollection) *[]color.Color {
	var convertedColors []color.Color
	convertedColors = []color.Color{}
	e.mapCoordinatesToEyedropperBoundary(collection)

	for _, coordinate := range *collection.Coordinates() {
		var newColor color.NRGBA
		if coordinate.HasMappedCoordinate() {
			mappedCoordinateX, mappedCoordinateY := coordinate.MappedCoordinate()
			sourceColorR, sourceColorG, sourceColorB, sourceColorA := (e.Image()).At(int(mappedCoordinateX), int(mappedCoordinateY)).RGBA()
			newColor = color.NRGBA{
				R: uint8(sourceColorR >> 8),
				G: uint8(sourceColorG >> 8),
				B: uint8(sourceColorB >> 8),
				A: uint8(sourceColorA >> 8),
			}
		} else {
			newColor = color.NRGBA{
				R: uint8(0 >> 8),
				G: uint8(0 >> 8),
				B: uint8(0 >> 8),
				A: uint8(0 >> 8),
			}
		}
		convertedColors = append(convertedColors, newColor)
	}
	return &convertedColors
}

// mapCoordinatesToEyedropperBoundary maps each coordinate from its minimum and maximum to the eyedropper's boundary.
//   Only coordinates that satisfied their filter will be updated.
func (e *RectangularEyedropper) mapCoordinatesToEyedropperBoundary(collection *CoordinateCollection) {
	collectionMinimumX := collection.MinimumTransformedX()
	collectionMaximumX := collection.MaximumTransformedX()
	collectionMinimumY := collection.MinimumTransformedY()
	collectionMaximumY := collection.MaximumTransformedY()

	for _, coordinate := range *collection.Coordinates() {
		if !coordinate.SatisfiesFilter() {
			continue
		}
		if !coordinate.CanBeCompared() {
			continue
		}

		eyedropperX := mathutility.ScaleValueBetweenTwoRanges(
			coordinate.TransformedX(),
			collectionMinimumX,
			collectionMaximumX,
			float64(e.LeftSide()),
			float64(e.RightSide()),
		)

		eyedropperY := mathutility.ScaleValueBetweenTwoRanges(
			coordinate.TransformedY(),
			collectionMinimumY,
			collectionMaximumY,
			float64(e.TopSide()),
			float64(e.BottomSide()),
		)

		coordinate.StoreMappedCoordinate(eyedropperX, eyedropperY)
	}
}
