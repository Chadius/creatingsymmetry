package imageoutput

import "math"

// CoordinateCollection holds an array of coordinates as they turn into symmetry patterns.
type CoordinateCollection struct {
	coordinates *[]*MappedCoordinate
}

// Coordinates returns the collection of coordinates.
func (c *CoordinateCollection) Coordinates() *[]*MappedCoordinate {
	return c.coordinates
}

// MinimumTransformedX returns the lowest TransformedX coordinate in the collection.
func (c *CoordinateCollection) MinimumTransformedX() float64 {
	foundCandidate := false
	minimumX := math.NaN()
	for _, coordinate := range *c.coordinates {
		coordinateIsValid := coordinate.CanBeCompared() && coordinate.SatisfiesFilter()
		if coordinateIsValid && (!foundCandidate || (coordinate.TransformedX() < minimumX)) {
			minimumX = coordinate.TransformedX()
			foundCandidate = true
		}
	}
	return minimumX
}

// MaximumTransformedX returns the greatest TransformedX coordinate in the collection.
func (c *CoordinateCollection) MaximumTransformedX() float64 {
	foundCandidate := false
	maximumX := math.NaN()
	for _, coordinate := range *c.coordinates {
		coordinateIsValid := coordinate.CanBeCompared() && coordinate.SatisfiesFilter()
		if coordinateIsValid && (!foundCandidate || coordinate.TransformedX() > maximumX) {
			maximumX = coordinate.TransformedX()
			foundCandidate = true
		}
	}
	return maximumX
}

// MinimumTransformedY returns the lowest TransformedY coordinate in the collection.
func (c *CoordinateCollection) MinimumTransformedY() float64 {
	foundCandidate := false
	minimumY := math.NaN()
	for _, coordinate := range *c.coordinates {
		coordinateIsValid := coordinate.CanBeCompared() && coordinate.SatisfiesFilter()
		if coordinateIsValid && (!foundCandidate || coordinate.TransformedY() < minimumY) {
			minimumY = coordinate.TransformedY()
			foundCandidate = true
		}
	}
	return minimumY
}

// MaximumTransformedY returns the greatest TransformedY coordinate in the collection.
func (c *CoordinateCollection) MaximumTransformedY() float64 {
	foundCandidate := false
	maximumY := math.NaN()
	for _, coordinate := range *c.coordinates {
		coordinateIsValid := coordinate.CanBeCompared() && coordinate.SatisfiesFilter()
		if coordinateIsValid && (!foundCandidate || coordinate.TransformedY() > maximumY) {
			maximumY = coordinate.TransformedY()
			foundCandidate = true
		}
	}
	return maximumY
}
