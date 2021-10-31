package imageoutput

import "math"

// MappedCoordinate stores the journey of an individual coordinate.
type MappedCoordinate struct {
	outputImageX         int
	outputImageY         int
	transformedX         float64
	transformedY         float64
	satisfiedFilter      bool
	hasMappedCoordinates bool
	mappedCoordinateX    float64
	mappedCoordinateY    float64
}

// NewMappedCoordinateUsingOutputImageCoordinates returns a new mapped coordinate at the given outputImageX and outputImageY location.
func NewMappedCoordinateUsingOutputImageCoordinates(outputImageX, outputImageY int) *MappedCoordinate {
	return &MappedCoordinate{
		outputImageX: outputImageX,
		outputImageY: outputImageY,
	}
}

// NewMappedCoordinateUsingTransformedCoordinates returns a new mapped coordinate at the given transformedX and transformedY location.
func NewMappedCoordinateUsingTransformedCoordinates(transformedX, transformedY float64) *MappedCoordinate {
	return &MappedCoordinate{
		transformedX:    transformedX,
		transformedY:    transformedY,
		satisfiedFilter: false,
	}
}

// OutputImageX returns the OutputImageX coordinate.
func (m *MappedCoordinate) OutputImageX() int {
	return m.outputImageX
}

// OutputImageY returns the OutputImageY coordinate.
func (m *MappedCoordinate) OutputImageY() int {
	return m.outputImageY
}

// TransformedX returns the TransformedX coordinate.
func (m *MappedCoordinate) TransformedX() float64 {
	return m.transformedX
}

// TransformedY returns the TransformedY coordinate.
func (m *MappedCoordinate) TransformedY() float64 {
	return m.transformedY
}

// UpdateTransformedCoordinates will update transformedX and transformedY coordinates.
func (m *MappedCoordinate) UpdateTransformedCoordinates(x, y float64) {
	m.transformedX = x
	m.transformedY = y
}

// PatternViewportX returns the PatternViewportX coordinate.
func (m *MappedCoordinate) PatternViewportX() float64 {
	return m.transformedX
}

// PatternViewportY returns the PatternViewportY coordinate.
func (m *MappedCoordinate) PatternViewportY() float64 {
	return m.transformedY
}

// UpdatePatternViewportCoordinates will update transformedX and transformedY coordinates.
func (m *MappedCoordinate) UpdatePatternViewportCoordinates(x, y float64) {
	m.transformedX = x
	m.transformedY = y
}

// CanBeCompared returns true if either transformedX and transformedY coordinate can be compared.
//   This means neither are Infinity nor NaN.
func (m *MappedCoordinate) CanBeCompared() bool {
	return !(math.IsInf(m.TransformedX(), 0) ||
		math.IsInf(m.TransformedY(), 0) ||
		math.IsNaN(m.TransformedX()) ||
		math.IsNaN(m.TransformedY()))
}

// MarkAsSatisfyingFilter marks this coordinate as satisfying the filter.
func (m *MappedCoordinate) MarkAsSatisfyingFilter() {
	m.satisfiedFilter = true
}

// SatisfiesFilter returns the filtered status.
func (m *MappedCoordinate) SatisfiesFilter() bool {
	return m.satisfiedFilter
}

// HasMappedCoordinate returns true if this coordinate stored another mapped coordinate
func (m *MappedCoordinate) HasMappedCoordinate() bool {
	return m.hasMappedCoordinates
}

// StoreMappedCoordinate sets the coordinate's mapped coordinates.
func (m *MappedCoordinate) StoreMappedCoordinate(x, y float64) {
	m.mappedCoordinateX = x
	m.mappedCoordinateY = y
	m.hasMappedCoordinates = true
}

// MappedCoordinate returns the stored mapped coordinates.
func (m *MappedCoordinate) MappedCoordinate() (float64, float64) {
	return m.mappedCoordinateX, m.mappedCoordinateY
}
