package wallpaper

import (
	"github.com/Chadius/creating-symmetry/entities/oldformula/coefficient"
	"github.com/Chadius/creating-symmetry/entities/oldformula/latticevector"
	"math"
)

// createVectorsForHexagonalWallpaper creates two vectors of a fixed shape and size.
func createVectorsForHexagonalWallpaper(formula *Formula) error {
	formula.Lattice = &latticevector.Pair{
		XLatticeVector: complex(1, 0),
		YLatticeVector: complex(-0.5, math.Sqrt(3.0)/2.0),
	}
	return nil
}

func lockCoefficientPairsForHexagonalWallpaper(formula *Formula) {
	formula.lockEisensteinTermsBasedOnRelationship([]coefficient.Relationship{
		coefficient.PlusMMinusSumNAndM,
		coefficient.MinusSumNAndMPlusN,
	})
}

func checksForSymmetryForHexagonalType(formula *Formula, targetSymmetry Symmetry) bool {
	if targetSymmetry == P3 {
		return true
	}

	return HasSymmetry(formula.WavePackets, targetSymmetry, map[Symmetry][]coefficient.Relationship{
		P31m: {coefficient.PlusMPlusN},
		P3m1: {coefficient.MinusMMinusN},
		P6:   {coefficient.MinusNMinusM},
		P6m: {
			coefficient.MinusNMinusM,
			coefficient.MinusMMinusN,
			coefficient.PlusMPlusN,
		},
	})
}
