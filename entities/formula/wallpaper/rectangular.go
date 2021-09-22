package wallpaper

import (
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	"github.com/Chadius/creating-symmetry/entities/formula/latticevector"
)

// createVectorsForRectangularWallpaper creates two vectors of a fixed shape and size.
func createVectorsForRectangularWallpaper(formula *Formula) error {
	formula.Lattice = &latticevector.Pair{
		XLatticeVector: complex(1, 0),
		YLatticeVector: complex(0, formula.LatticeSize.Height),
	}
	return nil
}

func checksForSymmetryForRectangularType(formula *Formula, targetSymmetry Symmetry) bool {
	return HasSymmetry(formula.WavePackets, targetSymmetry, map[Symmetry][]coefficient.Relationship {
		Pm: {coefficient.PlusNMinusM},
		Pg: {coefficient.PlusNMinusMNegateMultiplierIfOddPowerN},
		Pmm: {
			coefficient.PlusNMinusM,
			coefficient.MinusNMinusM,
			coefficient.MinusNPlusM,
		},
		Pmg: {
			coefficient.MinusNMinusM,
			coefficient.PlusNMinusMNegateMultiplierIfOddPowerN,
			coefficient.MinusNPlusMNegateMultiplierIfOddPowerN,
		},
		Pgg: {
			coefficient.MinusNMinusM,
			coefficient.PlusNMinusMNegateMultiplierIfOddPowerSum,
			coefficient.MinusNPlusMNegateMultiplierIfOddPowerSum,
		},
	})
}