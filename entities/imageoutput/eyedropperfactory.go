package imageoutput

import "image"

type EyedropperFactoryOptions struct {
	leftSide int
	rightSide int
	bottomSide int
	topSide int
	sourceImage *image.Image
}

// EyedropperFactory creates a EyedropperFactoryOptions with default values.
//   Can be chained with other class functions. Call Build() to create the
//   final object.
func EyedropperFactory() *EyedropperFactoryOptions {
	return &EyedropperFactoryOptions{
		leftSide: 0,
		rightSide: 0,
		bottomSide: 0,
		topSide: 0,
		sourceImage: nil,
	}
}

// WithLeftSide sets the left boundary.
func (e *EyedropperFactoryOptions) WithLeftSide(xMin int) *EyedropperFactoryOptions {
	e.leftSide = xMin
	return e
}

// WithRightSide sets the right boundary.
func (e *EyedropperFactoryOptions) WithRightSide(xMax int) *EyedropperFactoryOptions {
	e.rightSide = xMax
	return e
}

// WithTopSide sets the top boundary.
func (e *EyedropperFactoryOptions) WithTopSide(yMin int) *EyedropperFactoryOptions {
	e.topSide = yMin
	return e
}

// WithBottomSide sets the bottom boundary.
func (e *EyedropperFactoryOptions) WithBottomSide(yMax int) *EyedropperFactoryOptions {
	e.bottomSide = yMax
	return e
}

// WithImage sets the bottom boundary.
func (e *EyedropperFactoryOptions) WithImage(sourceImage *image.Image) *EyedropperFactoryOptions {
	e.sourceImage = sourceImage
	return e
}

// Build uses the PowerFactoryOptions to create a power.
func (e *EyedropperFactoryOptions) Build() *Eyedropper {
	return &Eyedropper{
		leftBoundary: e.leftSide,
		rightBoundary: e.rightSide,
		topBoundary: e.topSide,
		bottomBoundary: e.bottomSide,
		sourceImage: e.sourceImage,
	}
}

