package imageoutput

type CoordinateFilter struct {
	minimumX float64
	maximumX float64
	minimumY float64
	maximumY float64
}

// MinimumX returns the minimum x value for the filter.
func (c *CoordinateFilter) MinimumX() float64 {
	return c.minimumX
}

// MaximumX returns the maximum x value for the filter.
func (c *CoordinateFilter) MaximumX() float64 {
	return c.maximumX
}

// MinimumY returns the minimum y value for the filter.
func (c *CoordinateFilter) MinimumY() float64 {
	return c.minimumY
}

// MaximumY returns the maximum y value for the filter.
func (c *CoordinateFilter) MaximumY() float64 {
	return c.maximumY
}

// FilterAndMarkMappedCoordinate checks if the coordinate satisfies the filter. Then it marks the coordinate if it was filtered out.
func (c *CoordinateFilter) FilterAndMarkMappedCoordinate(coordinate *MappedCoordinate) {
	if coordinate.IsAtInfinity() {
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