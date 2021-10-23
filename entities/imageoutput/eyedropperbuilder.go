package imageoutput

import "image"

// EyedropperBuilderOptions stores the options used to build an eyedropper.
type EyedropperBuilderOptions struct {
	leftSide    int
	rightSide   int
	bottomSide  int
	topSide     int
	sourceImage *image.Image
}

// EyedropperBuilder creates a EyedropperBuilderOptions with default values.
//   Can be chained with other class functions. Call Build() to create the
//   final object.
func EyedropperBuilder() *EyedropperBuilderOptions {
	return &EyedropperBuilderOptions{
		leftSide:    0,
		rightSide:   0,
		bottomSide:  0,
		topSide:     0,
		sourceImage: nil,
	}
}

// WithLeftSide sets the left boundary.
func (e *EyedropperBuilderOptions) WithLeftSide(xMin int) *EyedropperBuilderOptions {
	e.leftSide = xMin
	return e
}

// WithRightSide sets the right boundary.
func (e *EyedropperBuilderOptions) WithRightSide(xMax int) *EyedropperBuilderOptions {
	e.rightSide = xMax
	return e
}

// WithTopSide sets the top boundary.
func (e *EyedropperBuilderOptions) WithTopSide(yMin int) *EyedropperBuilderOptions {
	e.topSide = yMin
	return e
}

// WithBottomSide sets the bottom boundary.
func (e *EyedropperBuilderOptions) WithBottomSide(yMax int) *EyedropperBuilderOptions {
	e.bottomSide = yMax
	return e
}

// WithImage sets the bottom boundary.
func (e *EyedropperBuilderOptions) WithImage(sourceImage *image.Image) *EyedropperBuilderOptions {
	e.sourceImage = sourceImage
	return e
}

// Build uses the builder options to create a power.
func (e *EyedropperBuilderOptions) Build() *Eyedropper {
	return &Eyedropper{
		leftBoundary:   e.leftSide,
		rightBoundary:  e.rightSide,
		topBoundary:    e.topSide,
		bottomBoundary: e.bottomSide,
		sourceImage:    e.sourceImage,
	}
}
