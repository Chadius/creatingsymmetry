package imageoutput

import (
	"github.com/Chadius/creating-symmetry/entities/mathutility"
	"image"
	"image/color"
)

// Eyedropper will sample transformed coordinates against a subrange of the source image.
type Eyedropper struct {
	leftBoundary   int
	rightBoundary  int
	topBoundary    int
	bottomBoundary int
	sourceImage    *image.Image
}

// LeftSide returns the left side of the boundary.
func (e *Eyedropper) LeftSide() int {
	return e.leftBoundary
}

// RightSide returns the right side of the boundary.
func (e *Eyedropper) RightSide() int {
	return e.rightBoundary
}

// TopSide returns the top side of the boundary.
func (e *Eyedropper) TopSide() int {
	return e.topBoundary
}

// BottomSide returns the bottom side of the boundary.
func (e *Eyedropper) BottomSide() int {
	return e.bottomBoundary
}

//Image returns the source image
func (e *Eyedropper) Image() *image.Image {
	return e.sourceImage
}

// ConvertCoordinatesToColors uses the collection of coordinates, maps it to the eyedropper range,
//   and samples the color in the source image at that location.
//   if the coordinate is mapped outside the source image, it will turn transparent.
func (e *Eyedropper) ConvertCoordinatesToColors(collection *CoordinateCollection) *[]color.Color {
	var convertedColors []color.Color
	convertedColors = []color.Color{}
	e.MapCoordinatesToEyedropperBoundary(collection)

	for _, coordinate := range *collection.Coordinates() {
		var newColor color.NRGBA
		if coordinate.HasMappedCoordinate() {
			mappedCoordinateX, mappedCoordinateY := coordinate.MappedCoordinate()
			sourceColorR, sourceColorG, sourceColorB, sourceColorA := (*e.Image()).At(int(mappedCoordinateX), int(mappedCoordinateY)).RGBA()
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

// MapCoordinatesToEyedropperBoundary maps each coordinate from its minimum and maximum to the eyedropper's boundary.
//   Only coordinates that satisfied their filter will be updated.
func (e *Eyedropper) MapCoordinatesToEyedropperBoundary(collection *CoordinateCollection) {
	collectionMinimumX := collection.MinimumX()
	collectionMaximumX := collection.MaximumX()
	collectionMinimumY := collection.MinimumY()
	collectionMaximumY := collection.MaximumY()

	for _, coordinate := range *collection.Coordinates() {
		if !coordinate.SatisfiesFilter() {
			continue
		}
		if !coordinate.CanBeCompared() {
			continue
		}

		eyedropperX := mathutility.ScaleValueBetweenTwoRanges(
			coordinate.X(),
			collectionMinimumX,
			collectionMaximumX,
			float64(e.LeftSide()),
			float64(e.RightSide()),
		)

		eyedropperY := mathutility.ScaleValueBetweenTwoRanges(
			coordinate.Y(),
			collectionMinimumY,
			collectionMaximumY,
			float64(e.TopSide()),
			float64(e.BottomSide()),
		)

		coordinate.StoreMappedCoordinate(eyedropperX, eyedropperY)
	}
}
