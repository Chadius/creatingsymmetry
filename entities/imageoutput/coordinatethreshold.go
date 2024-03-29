package imageoutput

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . CoordinateThreshold
// CoordinateThreshold looks at a CoordinateCollection and determines which coordinates will be kept.
type CoordinateThreshold interface {
	FilterAndMarkMappedCoordinateCollection(collection *CoordinateCollection)
}

// NullCoordinateThreshold is a default object where all coordinates with rational coordinates satisfy the filter.
type NullCoordinateThreshold struct{}

// FilterAndMarkMappedCoordinateCollection marks comparable coordinates as satisfying the filter.
func (c *NullCoordinateThreshold) FilterAndMarkMappedCoordinateCollection(collection *CoordinateCollection) {
	for _, coordinateToFilter := range *collection.Coordinates() {
		if coordinateToFilter.CanBeCompared() {
			coordinateToFilter.MarkAsSatisfyingFilter()
		}
	}
}

// RectangularCoordinateThreshold defines a rectangular range in which coordinates will be kept.
type RectangularCoordinateThreshold struct {
	minimumX float64
	maximumX float64
	minimumY float64
	maximumY float64
}

// MinimumX returns the minimum transformedX value for the filter.
func (c *RectangularCoordinateThreshold) MinimumX() float64 {
	return c.minimumX
}

// MaximumX returns the maximum transformedX value for the filter.
func (c *RectangularCoordinateThreshold) MaximumX() float64 {
	return c.maximumX
}

// MinimumY returns the minimum transformedY value for the filter.
func (c *RectangularCoordinateThreshold) MinimumY() float64 {
	return c.minimumY
}

// MaximumY returns the maximum transformedY value for the filter.
func (c *RectangularCoordinateThreshold) MaximumY() float64 {
	return c.maximumY
}

// filterAndMarkMappedCoordinate checks if the coordinate satisfies the filter.
//   Then it marks the coordinate if it satisfied the filtered out.
func (c *RectangularCoordinateThreshold) filterAndMarkMappedCoordinate(coordinate *MappedCoordinate) {
	if !coordinate.CanBeCompared() {
		return
	}

	if coordinate.TransformedX() < c.MinimumX() {
		return
	}

	if coordinate.TransformedX() > c.MaximumX() {
		return
	}

	if coordinate.TransformedY() < c.MinimumY() {
		return
	}

	if coordinate.TransformedY() > c.MaximumY() {
		return
	}
	coordinate.MarkAsSatisfyingFilter()
}

// FilterAndMarkMappedCoordinateCollection checks all coordinates against the filter.
//   Then it marks each coordinate if it satisfied the filter.
func (c *RectangularCoordinateThreshold) FilterAndMarkMappedCoordinateCollection(collection *CoordinateCollection) {
	for _, coordinateToFilter := range *collection.Coordinates() {
		c.filterAndMarkMappedCoordinate(coordinateToFilter)
	}
}
