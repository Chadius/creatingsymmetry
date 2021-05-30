package wavepacket

import (
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/coefficient"
)

// SquareWallpaperFormula uses waves that create a 4 rotation symmetry.
//  Each term will be rotated 4 times and averaged by 1/4.
//  The two vectors have the same length and are perpendicular to each other.
type SquareWallpaperFormula struct {
	Formula *WallpaperFormula
}

// SetUp initializes all of the needed wallpaper terms.
func (squareWaveFormula *SquareWallpaperFormula) SetUp() {
	squareWaveFormula.Formula.Lattice = &formula.LatticeVectorPair{
		XLatticeVector: complex(1, 0),
		YLatticeVector: complex(0, 1),
	}
	squareWaveFormula.Formula.SetUp(
		[]coefficient.Relationship{
			coefficient.PlusMMinusN,
			coefficient.MinusNMinusM,
			coefficient.MinusMPlusN,
		},
	)
}

// Calculate applies the formula to the complex number z.
//  It modifies the formula's result to track the contribution per term
//  As well as the final numerical result.
func (squareWaveFormula *SquareWallpaperFormula) Calculate(z complex128) *formula.CalculationResultForFormula {
	return squareWaveFormula.Formula.Calculate(z)
}

// FindSymmetries returns an object that tracks all of the symmetries found
//  in this formula.
func (squareWaveFormula *SquareWallpaperFormula) FindSymmetries() *Symmetry {
	foundSymmetries := Symmetry{
		P4: true,
	}

	symmetryFound := FindWaveRelationships(squareWaveFormula.Formula.WavePackets)
	if symmetryFound.PlusMPlusN {
		foundSymmetries.P4m = true
	}
	if symmetryFound.MaybeNegateBasedOnSumPlusMPlusN {
		foundSymmetries.P4g = true
	}

	return &foundSymmetries
}

// NewSquareWallpaperFormulaFromJSON reads the data and returns a formula term from it.
func NewSquareWallpaperFormulaFromJSON(data []byte) (*SquareWallpaperFormula, error) {
	formula, err := NewWallpaperFormulaFromJSON(data)

	if err != nil {
		return nil, err
	}

	return &SquareWallpaperFormula{
		Formula: formula,
	}, nil
}

// NewSquareWallpaperFormulaFromYAML reads the data and returns a formula term from it.
func NewSquareWallpaperFormulaFromYAML(data []byte) (*SquareWallpaperFormula, error) {
	formula, err := NewWallpaperFormulaFromYAML(data)

	if err != nil {
		return nil, err
	}

	return &SquareWallpaperFormula{
		Formula: formula,
	}, nil
}
// NewSquareWallpaperFormulaFromMarshalObject uses a marshalled object to create a new object.
func NewSquareWallpaperFormulaFromMarshalObject(marshalObject WallpaperFormulaMarshalled) *SquareWallpaperFormula {
	return &SquareWallpaperFormula{
		Formula: NewWallpaperFormulaFromMarshalObject(marshalObject),
	}
}
