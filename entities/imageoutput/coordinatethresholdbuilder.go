package imageoutput

// CoordinateThresholdBuilderOptions contains options used to make a RectangularCoordinateThreshold.
type CoordinateThresholdBuilderOptions struct {
	minimumX float64
	maximumX float64
	maximumY float64
	minimumY float64
}

// CoordinateFilterBuilder creates a CoordinateThresholdBuilderOptions with default values.
//   Can be chained with other class functions. Call Build() to create the
//   final object.
func CoordinateFilterBuilder() *CoordinateThresholdBuilderOptions {
	return &CoordinateThresholdBuilderOptions{
		minimumX: 0.0,
		maximumX: 0.0,
		maximumY: 0.0,
		minimumY: 0.0,
	}
}

// WithMinimumX sets the minimum transformedX value for the filter.
func (e *CoordinateThresholdBuilderOptions) WithMinimumX(xMin float64) *CoordinateThresholdBuilderOptions {
	e.minimumX = xMin
	return e
}

// WithMaximumX sets the maximum transformedX value for the filter.
func (e *CoordinateThresholdBuilderOptions) WithMaximumX(xMax float64) *CoordinateThresholdBuilderOptions {
	e.maximumX = xMax
	return e
}

// WithMinimumY sets the minimum transformedY value for the filter.
func (e *CoordinateThresholdBuilderOptions) WithMinimumY(yMin float64) *CoordinateThresholdBuilderOptions {
	e.minimumY = yMin
	return e
}

// WithMaximumY sets the maximum transformedY value for the filter.
func (e *CoordinateThresholdBuilderOptions) WithMaximumY(yMax float64) *CoordinateThresholdBuilderOptions {
	e.maximumY = yMax
	return e
}

// Build uses the builder options to create a power.
func (e *CoordinateThresholdBuilderOptions) Build() *RectangularCoordinateThreshold {
	return &RectangularCoordinateThreshold{
		minimumX: e.minimumX,
		maximumX: e.maximumX,
		minimumY: e.minimumY,
		maximumY: e.maximumY,
	}
}
