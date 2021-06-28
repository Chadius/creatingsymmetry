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

// NewSquareWallpaperFormulaWithSymmetry will try to create a new Hexagonal RhombicWallpaperFormula WavePacket
//   with the desired Terms, Multiplier and Symmetry.
func NewSquareWallpaperFormulaWithSymmetry(terms []*formula.EisensteinFormulaTerm, wallpaperMultiplier complex128, desiredSymmetry Symmetry) (*SquareWallpaperFormula, error) {
	newWavePackets := []*WavePacket{}
	for _, term := range terms {
		newWavePackets = append(
			newWavePackets,
			&WavePacket{
				Terms:      []*formula.EisensteinFormulaTerm{term},
				Multiplier: term.Multiplier,
			},
		)

		newWavePackets = addNewWavePacketsBasedOnSymmetry(term, desiredSymmetry, newWavePackets)
	}

	newBaseWallpaper := &SquareWallpaperFormula{
		Formula: &WallpaperFormula{
			WavePackets: newWavePackets,
			Multiplier:  wallpaperMultiplier,
		},
	}
	newBaseWallpaper.SetUp()
	return newBaseWallpaper, nil
}

// HasSymmetry returns true if the WavePackets involved form symmetry.
func (squareWaveFormula *SquareWallpaperFormula) HasSymmetry(desiredSymmetry Symmetry) bool {
	if desiredSymmetry == P4 {
		return true
	}

	return HasSymmetry(squareWaveFormula.Formula.WavePackets, desiredSymmetry, map[Symmetry][]coefficient.Relationship {
		P4m: {coefficient.PlusMPlusN},
		P4g: {coefficient.PlusMPlusNMaybeFlipScale},
	})
}
