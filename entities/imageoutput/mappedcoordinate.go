package imageoutput

import "math"

type MappedCoordinate struct {
	x float64
	y float64
	wasFiltered bool
	hasMappedCoordinates bool
	mappedCoordinateX float64
	mappedCoordinateY float64
}

// NewMappedCoordinate returns a new mapped coordinate at the given x and y location.
func NewMappedCoordinate(x, y float64) *MappedCoordinate {
	return &MappedCoordinate{
		x: x,
		y: y,
		wasFiltered: false,
	}
}

// X returns the X coordinate.
func (m *MappedCoordinate) X() float64 {
	return m.x
}

// Y returns the Y coordinate.
func (m *MappedCoordinate) Y() float64 {
	return m.y
}

// IsAtInfinity returns true if either x or y coordinate is at infinity.
func (m *MappedCoordinate) IsAtInfinity() bool {
	return math.IsInf(m.X(),0) || math.IsInf(m.Y(),0)
}

// MarkAsFiltered marks this coordinate as filtered.
func (m *MappedCoordinate) MarkAsFiltered() {
	m.wasFiltered = true
}

// IsFiltered returns the filtered status.
func (m *MappedCoordinate) IsFiltered() bool {
	return m.wasFiltered
}

// HasMappedCoordinate returns true if this coordinate stored another mapped coordinate
func (m *MappedCoordinate) HasMappedCoordinate() bool {
	return m.hasMappedCoordinates
}

// StoreMappedCoordinate sets the coordinate's mapped coordinates.
func (m *MappedCoordinate) StoreMappedCoordinate(complexCoordinates complex128) {
	m.mappedCoordinateX = real(complexCoordinates)
	m.mappedCoordinateY = imag(complexCoordinates)
	m.hasMappedCoordinates = true
}

// MappedCoordinate returns the stored mapped coordinates.
func (m *MappedCoordinate) MappedCoordinate() (float64, float64) {
	return m.mappedCoordinateX, m.mappedCoordinateY
}
