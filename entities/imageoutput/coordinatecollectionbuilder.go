package imageoutput

// CoordinateCollectionBuilderOptions records options used to build a CoordinateCollection.
type CoordinateCollectionBuilderOptions struct {
	coordinates *[]*MappedCoordinate
}

// CoordinateCollectionBuilder creates a CoordinateCollectionBuilderOptions with default values.
//   Can be chained with other class functions. Call Build() to create the
//   final object.
func CoordinateCollectionBuilder() *CoordinateCollectionBuilderOptions {
	return &CoordinateCollectionBuilderOptions{
		coordinates: nil,
	}
}

// WithCoordinates sets the coordinates stored in the collection.
func (c *CoordinateCollectionBuilderOptions) WithCoordinates(coordinates *[]*MappedCoordinate) *CoordinateCollectionBuilderOptions {
	c.coordinates = coordinates
	return c
}

// WithComplexNumbers sets the coordinates stored in the collection using complex numbers.
//  The real portion is used as the transformedX coordinate.
//  The imaginary portion is used as the transformedY coordinate.
func (c *CoordinateCollectionBuilderOptions) WithComplexNumbers(complexNumbers *[]complex128) *CoordinateCollectionBuilderOptions {
	newCoordinates := []*MappedCoordinate{}
	for _, complexNumber := range *complexNumbers {
		newCoordinates = append(newCoordinates, NewMappedCoordinateUsingTransformedCoordinates(
			real(complexNumber),
			imag(complexNumber),
		))
	}
	c.coordinates = &newCoordinates
	return c
}

// Build uses the builder options to create a power.
func (c *CoordinateCollectionBuilderOptions) Build() *CoordinateCollection {
	return &CoordinateCollection{
		coordinates: c.coordinates,
	}
}
