package wallpaper

import (
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	"github.com/Chadius/creating-symmetry/entities/formula/latticevector"
)

func createVectorsForGenericWallpaper(formula *Formula) error {
	formula.Lattice = &latticevector.Pair{
		XLatticeVector: complex(1, 0),
		YLatticeVector: complex(formula.LatticeSize.Width, formula.LatticeSize.Height),
	}

	return nil
}

func checksForSymmetryForGenericType(formula *Formula, targetSymmetry Symmetry) bool {
	return HasSymmetry(formula.WavePackets, targetSymmetry, map[Symmetry][]coefficient.Relationship{
		P2: {coefficient.MinusNMinusM},
	})
}
