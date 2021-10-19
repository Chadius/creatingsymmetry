package imageoutput

import "math"

type CoordinateCollection struct {
	coordinates *[]*MappedCoordinate
}

// Coordinates returns the collection of coordinates.
func (c *CoordinateCollection) Coordinates() *[]*MappedCoordinate {
	return c.coordinates
}

// MinimumX returns the lowest x coordinate in the collection.
func (c *CoordinateCollection) MinimumX() float64 {
	foundCandidate := false
	minimumX := math.NaN()
	for _, coordinate := range *c.coordinates {
		coordinateIsValid := !coordinate.IsAtInfinity() && coordinate.SatisfiesFilter()
		if coordinateIsValid && (!foundCandidate || (coordinate.X() < minimumX)) {
			minimumX = coordinate.X()
			foundCandidate = true
		}
	}
	return minimumX
}

// MaximumX returns the greatest x coordinate in the collection.
func (c *CoordinateCollection) MaximumX() float64 {
	foundCandidate := false
	maximumX := math.NaN()
	for _, coordinate := range *c.coordinates {
		coordinateIsValid := !coordinate.IsAtInfinity() && coordinate.SatisfiesFilter()
		if coordinateIsValid && (!foundCandidate || coordinate.X() > maximumX) {
			maximumX = coordinate.X()
			foundCandidate = true
		}
	}
	return maximumX
}

// MinimumY returns the lowest y coordinate in the collection.
func (c *CoordinateCollection) MinimumY() float64 {
	foundCandidate := false
	minimumY := math.NaN()
	for _, coordinate := range *c.coordinates {
		coordinateIsValid := !coordinate.IsAtInfinity() && coordinate.SatisfiesFilter()
		if coordinateIsValid && (!foundCandidate || coordinate.Y() < minimumY) {
			minimumY = coordinate.Y()
			foundCandidate = true
		}
	}
	return minimumY
}

// MaximumY returns the greatest y coordinate in the collection.
func (c *CoordinateCollection) MaximumY() float64 {
	foundCandidate := false
	maximumY := math.NaN()
	for _, coordinate := range *c.coordinates {
		coordinateIsValid := !coordinate.IsAtInfinity() && coordinate.SatisfiesFilter()
		if coordinateIsValid && (!foundCandidate || coordinate.Y() > maximumY) {
			maximumY = coordinate.Y()
			foundCandidate = true
		}
	}
	return maximumY
}