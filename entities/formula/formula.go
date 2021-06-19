package formula

import (
	"errors"
	"fmt"
	"math"
	"wallpaper/entities/utility"
)

// CalculationResultForFormula shows the results of a calculation
type CalculationResultForFormula struct {
	Total				complex128
	ContributionByTerm	[]complex128
}

// LatticeVectorPairMarshal can be marshaled and converted to a LatticeVectorPair
type LatticeVectorPairMarshal struct {
	XLatticeVector			utility.ComplexNumberForMarshal	`json:"x_lattice_vector" yaml:"x_lattice_vector"`
	YLatticeVector			utility.ComplexNumberForMarshal	`json:"y_lattice_vector" yaml:"y_lattice_vector"`
}

// LatticeVectorPair defines the shape of the wallpaper lattice.
type LatticeVectorPair struct {
	XLatticeVector			complex128
	YLatticeVector			complex128
}

func vectorIsZero(vector complex128) bool {
	return real(vector) == 0 && imag(vector) == 0
}

// vectorsAreCollinear returns true if both vectors are perfectly lined up
func vectorsAreCollinear(vector1 complex128, vector2 complex128) bool {
	absoluteValueDotProduct := math.Abs((real(vector1) * real(vector2)) + (imag(vector1) * imag(vector2)))
	lengthOfVector1 := math.Sqrt((real(vector1) * real(vector1)) + (imag(vector1) * imag(vector1)))
	lengthOfVector2 := math.Sqrt((real(vector2) * real(vector2)) + (imag(vector2) * imag(vector2)))

	tolerance := 1e-8
	return math.Abs(absoluteValueDotProduct - lengthOfVector1 * lengthOfVector2) < tolerance
}

// Validate returns an error if this is an invalid formula.
func(lattice LatticeVectorPair)Validate() error {
	if vectorIsZero(lattice.XLatticeVector) || vectorIsZero(lattice.YLatticeVector) {
		return errors.New(`lattice vectors cannot be (0,0)`)
	}
	if vectorsAreCollinear(lattice.XLatticeVector, lattice.YLatticeVector) {
		return fmt.Errorf(
			`vectors cannot be collinear: (%f,%f) and \(%f,%f)`,
			real(lattice.XLatticeVector),
			imag(lattice.XLatticeVector),
			real(lattice.YLatticeVector),
			imag(lattice.YLatticeVector),
		)
	}
	return nil
}

// ConvertToLatticeCoordinates converts a point from cartesian coordinates to the lattice coordinates
func (lattice LatticeVectorPair) ConvertToLatticeCoordinates(cartesianPoint complex128) complex128 {

	vector1 := lattice.XLatticeVector
	vector2 := lattice.YLatticeVector
	swapVectorsDuringCalculation := real(vector1) < 1e-6

	if swapVectorsDuringCalculation == true {
		vector1 = lattice.YLatticeVector
		vector2 = lattice.XLatticeVector
	}

	scalarForVector2Numerator := (real(vector1) * imag(cartesianPoint)) - (imag(vector1) * real(cartesianPoint))
	scalarForVector2Denominator := (real(vector1) * imag(vector2)) - (imag(vector1) * real(vector2))
	scalarForVector2 := scalarForVector2Numerator / scalarForVector2Denominator

	scalarForVector1Numerator := real(cartesianPoint) - (scalarForVector2 * real(vector2))
	scalarForVector1Denominator := real(vector1)
	scalarForVector1 := scalarForVector1Numerator / scalarForVector1Denominator

	if swapVectorsDuringCalculation {
		return complex(scalarForVector2, scalarForVector1)
	}

	return complex(scalarForVector1, scalarForVector2)
}
