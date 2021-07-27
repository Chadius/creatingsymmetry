package wallpaper

import (
	"wallpaper/entities/formula/coefficient"
	"wallpaper/entities/formula/latticevector"
)

// createVectorsForRhombicWallpaper creates two vectors of a fixed shape and size.
func createVectorsForRhombicWallpaper(formula *Formula) error {
	formula.Lattice = &latticevector.Pair{
		XLatticeVector: complex(0.5, formula.LatticeSize.Height),
		YLatticeVector: complex(0.5, formula.LatticeSize.Height * -1),
	}
	return nil
}

func lockCoefficientPairsForRhombicWallpaper(formula *Formula) {
	formula.lockEisensteinTermsBasedOnRelationship([]coefficient.Relationship{
		coefficient.PlusMPlusN,
	})
}

func checksForSymmetryForRhombicType(formula *Formula, targetSymmetry Symmetry) bool {
	return HasSymmetry(formula.WavePackets, targetSymmetry, map[Symmetry][]coefficient.Relationship {
		Cm: {coefficient.PlusMPlusN},
		Cmm: {
			coefficient.MinusNMinusM,
			coefficient.MinusMMinusN,
			coefficient.PlusMPlusN,
		},
	})
}