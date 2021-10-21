package imageoutput

type CoordinateCollectionFactoryOptions struct {
	coordinates *[]*MappedCoordinate
}

// CoordinateCollectionFactory creates a CoordinateCollectionFactoryOptions with default values.
//   Can be chained with other class functions. Call Build() to create the
//   final object.
func CoordinateCollectionFactory() *CoordinateCollectionFactoryOptions {
	return &CoordinateCollectionFactoryOptions{
		coordinates: nil,
	}
}

// WithCoordinates sets the coordinates stored in the collection.
func (c *CoordinateCollectionFactoryOptions) WithCoordinates(coordinates *[]*MappedCoordinate) *CoordinateCollectionFactoryOptions {
	c.coordinates = coordinates
	return c
}

// WithComplexNumbers sets the coordinates stored in the collection using complex numbers.
//  The real portion is used as the x coordinate.
//  The imaginary portion is used as the y coordinate.
func (c *CoordinateCollectionFactoryOptions) WithComplexNumbers(complexNumbers *[]complex128) *CoordinateCollectionFactoryOptions {
	newCoordinates := []*MappedCoordinate{}
	for _, complexNumber := range *complexNumbers {
		newCoordinates = append(newCoordinates, NewMappedCoordinate(
			real(complexNumber),
			imag(complexNumber),
		))
	}
	c.coordinates = &newCoordinates
	return c
}

// Build uses the PowerFactoryOptions to create a power.
func (c *CoordinateCollectionFactoryOptions) Build() *CoordinateCollection {
	return &CoordinateCollection{
		coordinates: c.coordinates,
	}
}


