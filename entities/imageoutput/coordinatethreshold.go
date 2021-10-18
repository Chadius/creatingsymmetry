package imageoutput

// CoordinateThreshold defines a range in which coordinates will be kept.
type CoordinateThreshold struct {
	minimumX float64
	maximumX float64
	minimumY float64
	maximumY float64
}

// MinimumX returns the minimum x value for the filter.
func (c *CoordinateThreshold) MinimumX() float64 {
	return c.minimumX
}

// MaximumX returns the maximum x value for the filter.
func (c *CoordinateThreshold) MaximumX() float64 {
	return c.maximumX
}

// MinimumY returns the minimum y value for the filter.
func (c *CoordinateThreshold) MinimumY() float64 {
	return c.minimumY
}

// MaximumY returns the maximum y value for the filter.
func (c *CoordinateThreshold) MaximumY() float64 {
	return c.maximumY
}

// FilterAndMarkMappedCoordinate checks if the coordinate satisfies the filter.
//   Then it marks the coordinate if it satisfied the filtered out.
func (c *CoordinateThreshold) FilterAndMarkMappedCoordinate(coordinate *MappedCoordinate) {
	if !coordinate.CanBeCompared() {
		return
	}

	if coordinate.X() < c.MinimumX() {
		return
	}

	if coordinate.X() > c.MaximumX() {
		return
	}

	if coordinate.Y() < c.MinimumY() {
		return
	}

	if coordinate.Y() > c.MaximumY() {
		return
	}
	coordinate.MarkAsSatisfyingFilter()
}

// FilterAndMarkMappedCoordinateCollection checks all coordinates against the filter.
//   Then it marks each coordinate if it satisfied the filter.
func (c *CoordinateThreshold) FilterAndMarkMappedCoordinateCollection(collection *CoordinateCollection) {
	for _, coordinateToFiler := range *collection.Coordinates() {
		c.FilterAndMarkMappedCoordinate(coordinateToFiler)
	}
}
