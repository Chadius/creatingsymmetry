package imageoutput

type CoordinateFilterFactoryOptions struct {
	minimumX float64
	maximumX float64
	maximumY float64
	minimumY float64
}

// CoordinateFilterFactory creates a CoordinateFilterFactoryOptions with default values.
//   Can be chained with other class functions. Call Build() to create the
//   final object.
func CoordinateFilterFactory() *CoordinateFilterFactoryOptions {
	return &CoordinateFilterFactoryOptions{
		minimumX: 0.0,
		maximumX: 0.0,
		maximumY: 0.0,
		minimumY: 0.0,
	}
}

// WithMinimumX sets the minimum x value for the filter.
func (e *CoordinateFilterFactoryOptions) WithMinimumX(xMin float64) *CoordinateFilterFactoryOptions {
	e.minimumX = xMin
	return e
}

// WithMaximumX sets the maximum x value for the filter.
func (e *CoordinateFilterFactoryOptions) WithMaximumX(xMax float64) *CoordinateFilterFactoryOptions {
	e.maximumX = xMax
	return e
}

// WithMinimumY sets the minimum y value for the filter.
func (e *CoordinateFilterFactoryOptions) WithMinimumY(yMin float64) *CoordinateFilterFactoryOptions {
	e.minimumY = yMin
	return e
}

// WithMaximumY sets the maximum y value for the filter.
func (e *CoordinateFilterFactoryOptions) WithMaximumY(yMax float64) *CoordinateFilterFactoryOptions {
	e.maximumY = yMax
	return e
}

// Build uses the PowerFactoryOptions to create a power.
func (e *CoordinateFilterFactoryOptions) Build() *CoordinateFilter {
	return &CoordinateFilter{
		minimumX: e.minimumX,
		maximumX: e.maximumX,
		minimumY: e.minimumY,
		maximumY: e.maximumY,
	}
}


