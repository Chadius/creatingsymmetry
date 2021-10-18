package imageoutput

import (
	"github.com/Chadius/creating-symmetry/entities/mathutility"
	"image"
	"image/color"
)

type Eyedropper struct {
	leftBoundary int
	rightBoundary int
	topBoundary int
	bottomBoundary int
	sourceImage *image.Image
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

// Image returns the source image
func (e *Eyedropper) Image() *image.Image {
	return e.sourceImage
}

// ConvertCoordinatesToColors uses the collection of coordinates, maps it to the eyedropper range,
//   and samples the color in the source image at that location.
//   if the coordinate is mapped outside the source image, it will turn transparent.
func (e *Eyedropper) ConvertCoordinatesToColors(collection *CoordinateCollection) *[]color.Color {
	var convertedColors []color.Color
	convertedColors = []color.Color{}

	//println(imag((*collection.Coordinates())[1]))

	//eyedropperY := mathutility.ScaleValueBetweenTwoRanges(
	//	imag((*collection.Coordinates())[0]),
	//	collection.MinimumY(),
	//	collection.MaximumY(),
	//	float64(e.TopSide()),
	//	float64(e.BottomSide()),
	//)
	//println("---")
	//
	//value := imag((*collection.Coordinates())[1])
	//oldRangeMin := collection.MinimumY()
	//oldRangeMax := collection.MaximumY()
	//newRangeMin := float64(e.TopSide())
	//newRangeMax := float64(e.BottomSide())
	//distanceAcrossOldRange := oldRangeMax - oldRangeMin
	//println(distanceAcrossOldRange)
	//valueDistanceAcrossOldRange := value - oldRangeMin
	//println(valueDistanceAcrossOldRange)
	//ratioAcrossRange := valueDistanceAcrossOldRange / distanceAcrossOldRange
	//println(ratioAcrossRange)
	//distanceAcrossNewRange := newRangeMax - newRangeMin
	//println(distanceAcrossNewRange)
	//println ((ratioAcrossRange * distanceAcrossNewRange) + newRangeMin)

	for _, coordinate := range *collection.Coordinates() {
		newColor := e.convertCoordinateToColor(coordinate, collection)
		convertedColors = append(convertedColors, newColor)
	}

	return &convertedColors
}

func (e *Eyedropper) convertCoordinateToColor(coordinate complex128, collection *CoordinateCollection) color.Color {
	eyedropperX := mathutility.ScaleValueBetweenTwoRanges(
		real(coordinate),
		collection.MinimumX(),
		collection.MaximumX(),
		float64(e.LeftSide()),
		float64(e.RightSide()),
	)

	eyedropperY := mathutility.ScaleValueBetweenTwoRanges(
		imag(coordinate),
		collection.MinimumY(),
		collection.MaximumY(),
		float64(e.TopSide()),
		float64(e.BottomSide()),
	)

	outOfBoundsColor := color.NRGBA{
		R: uint8(0 >> 8),
		G: uint8(0 >> 8),
		B: uint8(0 >> 8),
		A: uint8(0 >> 8),
	}

	if eyedropperX < 0 {
		return outOfBoundsColor
	}

	if eyedropperX > float64((*e.Image()).Bounds().Size().X) {
		return outOfBoundsColor
	}

	if eyedropperY < 0 {
		return outOfBoundsColor
	}

	if eyedropperY > float64((*e.Image()).Bounds().Size().Y) {
		return outOfBoundsColor
	}

	sourceColorR, sourceColorG, sourceColorB, sourceColorA := (*e.Image()).At(int(eyedropperX), int(eyedropperY)).RGBA()
	return color.NRGBA{
		R: uint8(sourceColorR >> 8),
		G: uint8(sourceColorG >> 8),
		B: uint8(sourceColorB >> 8),
		A: uint8(sourceColorA >> 8),
	}
}