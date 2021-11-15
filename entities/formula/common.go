package formula

import (
	"errors"
	"fmt"
	"math"
)

// ConvertToLatticeCoordinates changes the coordinate to match the axes defined by the latticeVectors.
func ConvertToLatticeCoordinates(cartesianPoint complex128, latticeVectors []complex128) complex128 {
	vector1 := latticeVectors[0]
	vector2 := latticeVectors[1]
	swapVectorsDuringCalculation := real(vector1) < 1e-6

	if swapVectorsDuringCalculation == true {
		vector1 = latticeVectors[1]
		vector2 = latticeVectors[0]
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

func vectorIsZero(vector complex128) bool {
	return real(vector) == 0 && imag(vector) == 0
}

// vectorsAreCollinear returns true if both vectors are perfectly lined up
func vectorsAreCollinear(vector1 complex128, vector2 complex128) bool {
	absoluteValueDotProduct := math.Abs((real(vector1) * real(vector2)) + (imag(vector1) * imag(vector2)))
	lengthOfVector1 := math.Sqrt((real(vector1) * real(vector1)) + (imag(vector1) * imag(vector1)))
	lengthOfVector2 := math.Sqrt((real(vector2) * real(vector2)) + (imag(vector2) * imag(vector2)))

	tolerance := 1e-8
	return math.Abs(absoluteValueDotProduct-(lengthOfVector1*lengthOfVector2)) < tolerance
}

// ValidateLatticeVectors returns an error if the lattice vectors are invalid.
func ValidateLatticeVectors(latticeVectors []complex128) error {
	if vectorIsZero(latticeVectors[0]) || vectorIsZero(latticeVectors[1]) {
		return errors.New(`lattice vectors cannot be (0,0)`)
	}
	if vectorsAreCollinear(latticeVectors[0], latticeVectors[1]) {
		return fmt.Errorf(
			`vectors cannot be collinear: (%f,%f) and \(%f,%f)`,
			real(latticeVectors[0]),
			imag(latticeVectors[0]),
			real(latticeVectors[1]),
			imag(latticeVectors[1]),
		)
	}
	return nil
}