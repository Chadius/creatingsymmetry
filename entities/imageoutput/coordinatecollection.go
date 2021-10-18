package imageoutput

type CoordinateCollection struct {
	coordinates *[]complex128
}

// Coordinates returns the collection of coordinates.
func (c *CoordinateCollection) Coordinates() *[]complex128 {
	return c.coordinates
}

// MinimumX returns the lowest x coordinate in the collection.
func (c *CoordinateCollection) MinimumX() float64 {
	minimumX := real((*c.coordinates)[0])
	for _, coordinate := range *c.coordinates {
		if real(coordinate) < minimumX {
			minimumX = real(coordinate)
		}
	}
	return minimumX
}

// MaximumX returns the greatest x coordinate in the collection.
func (c *CoordinateCollection) MaximumX() float64 {
	maximumX := real((*c.coordinates)[0])
	for _, coordinate := range *c.coordinates {
		if real(coordinate) > maximumX {
			maximumX = real(coordinate)
		}
	}
	return maximumX
}

// MinimumY returns the lowest y coordinate in the collection.
func (c *CoordinateCollection) MinimumY() float64 {
	minimumY := imag((*c.coordinates)[0])
	for _, coordinate := range *c.coordinates {
		if imag(coordinate) < minimumY {
			minimumY = imag(coordinate)
		}
	}
	return minimumY
}

// MaximumY returns the greatest y coordinate in the collection.
func (c *CoordinateCollection) MaximumY() float64 {
	maximumY := imag((*c.coordinates)[0])
	for _, coordinate := range *c.coordinates {
		if imag(coordinate) > maximumY {
			maximumY = imag(coordinate)
		}
	}
	return maximumY
}