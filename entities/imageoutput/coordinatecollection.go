package imageoutput

type CoordinateCollection struct {
	coordinates *[]*MappedCoordinate
}

// Coordinates returns the collection of coordinates.
func (c *CoordinateCollection) Coordinates() *[]*MappedCoordinate {
	return c.coordinates
}

// MinimumX returns the lowest x coordinate in the collection.
func (c *CoordinateCollection) MinimumX() float64 {
	minimumX := (*c.coordinates)[0].X()
	for _, coordinate := range *c.coordinates {
		if !coordinate.IsAtInfinity() && !coordinate.IsFiltered() && coordinate.X() < minimumX {
			minimumX = coordinate.X()
		}
	}
	return minimumX
}

// MaximumX returns the greatest x coordinate in the collection.
func (c *CoordinateCollection) MaximumX() float64 {
	maximumX := (*c.coordinates)[0].X()
	for _, coordinate := range *c.coordinates {
		if !coordinate.IsAtInfinity() && !coordinate.IsFiltered() && coordinate.X() > maximumX {
			maximumX = coordinate.X()
		}
	}
	return maximumX
}

// MinimumY returns the lowest y coordinate in the collection.
func (c *CoordinateCollection) MinimumY() float64 {
	minimumY := (*c.coordinates)[0].Y()
	for _, coordinate := range *c.coordinates {
		if !coordinate.IsAtInfinity() && !coordinate.IsFiltered() && coordinate.Y() < minimumY {
			minimumY = coordinate.Y()
		}
	}
	return minimumY
}

// MaximumY returns the greatest y coordinate in the collection.
func (c *CoordinateCollection) MaximumY() float64 {
	maximumY := (*c.coordinates)[0].Y()
	for _, coordinate := range *c.coordinates {
		if !coordinate.IsAtInfinity() && !coordinate.IsFiltered() && coordinate.Y() > maximumY {
			maximumY = coordinate.Y()
		}
	}
	return maximumY
}