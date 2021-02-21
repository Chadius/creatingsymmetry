package mathutility

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