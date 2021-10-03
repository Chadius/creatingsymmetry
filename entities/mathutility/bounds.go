package mathutility

import "math/cmplx"

// ScaleValueBetweenTwoRanges translates value (exists between min1 and max2)
//   to somewhere between min2 and max2. Linear scale is assumed.
//   So if value is 25% between min1 and max1, then this will return
//     a float64 that is 25% between min2 and max2.
func ScaleValueBetweenTwoRanges(value, oldRangeMin, oldRangeMax, newRangeMin, newRangeMax float64) float64 {
	if value <= oldRangeMin {
		return newRangeMin
	}

	if value >= oldRangeMax {
		return newRangeMax
	}

	distanceAcrossOldRange := oldRangeMax - oldRangeMin
	valueDistanceAcrossOldRange := value - oldRangeMin
	ratioAcrossRange := valueDistanceAcrossOldRange / distanceAcrossOldRange
	distanceAcrossNewRange := newRangeMax - newRangeMin
	return (ratioAcrossRange * distanceAcrossNewRange) + newRangeMin
}

// GetBoundingBox returns two complex numbers that contain all of the non-infinity numbers inside.
//   The first number is the minimum (all numbers have a real & imaginary component greater than or equal to it)
//   The second number is the maximum (all numbers have a real & imaginary component less than or equal to it)
//   If numbers is an empty slice, returns (0+0i, 0+0i)
func GetBoundingBox(numbers []complex128) (complex128, complex128) {
	xMin, yMin, xMax, yMax := 0.0, 0.0, 0.0, 0.0
	for _, number := range numbers {
		if cmplx.IsInf(number) {
			continue
		}

		if real(number) < xMin {
			xMin = real(number)
		}
		if real(number) > xMax {
			xMax = real(number)
		}

		if imag(number) < yMin {
			yMin = imag(number)
		}
		if imag(number) > yMax {
			yMax = imag(number)
		}
	}
	return complex(xMin, yMin), complex(xMax, yMax)
}
