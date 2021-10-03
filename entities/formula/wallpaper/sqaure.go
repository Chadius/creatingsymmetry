package wallpaper

import (
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	"github.com/Chadius/creating-symmetry/entities/formula/latticevector"
)

// createVectorsForSquareWallpaper creates two vectors of a fixed shape and size.
func createVectorsForSquareWallpaper(formula *Formula) error {
	formula.Lattice = &latticevector.Pair{
		XLatticeVector: complex(1, 0),
		YLatticeVector: complex(0, 1),
	}
	return nil
}

func lockCoefficientPairsForSquareWallpaper(formula *Formula) {
	formula.lockEisensteinTermsBasedOnRelationship([]coefficient.Relationship{
		coefficient.PlusMMinusN,
		coefficient.MinusNMinusM,
		coefficient.MinusMPlusN,
	})
}

func checksForSymmetryForSquareType(formula *Formula, targetSymmetry Symmetry) bool {
	if targetSymmetry == P4 {
		return true
	}

	return HasSymmetry(formula.WavePackets, targetSymmetry, map[Symmetry][]coefficient.Relationship{
		P4m: {coefficient.PlusMPlusN},
		P4g: {coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum},
	})
}
