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

// Build uses the PowerFactoryOptions to create a power.
func (c *CoordinateCollectionFactoryOptions) Build() *CoordinateCollection {
	return &CoordinateCollection{
		coordinates: c.coordinates,
	}
}


