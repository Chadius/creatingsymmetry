package wavepacket

import (
	"errors"
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

// NewSquareWallpaperFormulaWithSymmetry will try to create a new Hexagonal Wallpaper Formula
//   with the desired Terms, Multiplier and Symmetry.
func NewSquareWallpaperFormulaWithSymmetry(terms []*formula.EisensteinFormulaTerm, wallpaperMultiplier complex128, desiredSymmetry *Symmetry) (*SquareWallpaperFormula, error) {
	err := checkForIncompatibleSquareSymmetries(terms, desiredSymmetry)
	if err != nil {
		return nil, err
	}

	newWavePackets := []*Formula{}
	for _, term := range terms {
		newWavePackets = append(
			newWavePackets,
			&Formula{
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


func checkForIncompatibleSquareSymmetries(terms []*formula.EisensteinFormulaTerm, desiredSymmetry *Symmetry) error {
	atLeastOneTermPowerSumIsOdd := false
	for _, term := range terms {
		powerSumIsOdd := (term.PowerN + term.PowerM) % 2 != 0
		if powerSumIsOdd {
			atLeastOneTermPowerSumIsOdd = true
			break
		}
	}

	if desiredSymmetry.P4g && desiredSymmetry.P4m && atLeastOneTermPowerSumIsOdd {
		return errors.New("invalid desired symmetry")
	}
	return nil
}