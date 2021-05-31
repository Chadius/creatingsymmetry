package wavepacket

import (
	"errors"
	"math"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/coefficient"
)

// HexagonalWallpaperFormula uses waves that create a 3 rotation symmetry.
//  Each term will be rotated 3 times and averaged by 1/3.
type HexagonalWallpaperFormula struct {
	Formula *WallpaperFormula
}

// SetUp initializes all of the needed wallpaper terms.
func (hexWaveFormula *HexagonalWallpaperFormula) SetUp() {
	hexWaveFormula.Formula.Lattice = &formula.LatticeVectorPair{
		XLatticeVector: complex(1, 0),
		YLatticeVector: complex(-0.5, math.Sqrt(3.0)/2.0),
	}
	hexWaveFormula.Formula.SetUp(
		[]coefficient.Relationship{
			coefficient.PlusMMinusSumNAndM,
			coefficient.MinusSumNAndMPlusN,
		},
	)
}

// Calculate applies the formula to the complex number z.
//  It modifies the formula's result to track the contribution per term
//  As well as the final numerical result.
func (hexWaveFormula *HexagonalWallpaperFormula) Calculate(z complex128) *formula.CalculationResultForFormula {
	return hexWaveFormula.Formula.Calculate(z)
}

// FindSymmetries returns an object with a bunch of symmetries.
func (hexWaveFormula *HexagonalWallpaperFormula) FindSymmetries() *Symmetry {
	foundSymmetries := Symmetry{
		P3: true,
	}

	symmetryFound := FindWaveRelationships(hexWaveFormula.Formula.WavePackets)
	if symmetryFound.PlusMPlusN {
		foundSymmetries.P31m = true
	}
	if symmetryFound.MinusMMinusN {
		foundSymmetries.P3m1 = true
	}
	if symmetryFound.MinusNMinusM {
		foundSymmetries.P6 = true
	}
	if symmetryFound.MinusNMinusMPlusMPlusNMinusMMinusN {
		foundSymmetries.P6m = true
	}
	return &foundSymmetries
}

// NewHexagonalWallpaperFormulaFromJSON reads the data and returns a formula term from it.
func NewHexagonalWallpaperFormulaFromJSON(data []byte) (*HexagonalWallpaperFormula, error) {
	formula, err := NewWallpaperFormulaFromJSON(data)

	if err != nil {
		return nil, err
	}

	return &HexagonalWallpaperFormula{
		Formula: formula,
	}, nil
}

// NewHexagonalWallpaperFormulaFromYAML reads the data and returns a formula term from it.
func NewHexagonalWallpaperFormulaFromYAML(data []byte) (*HexagonalWallpaperFormula, error) {
	formula, err := NewWallpaperFormulaFromYAML(data)

	if err != nil {
		return nil, err
	}

	return &HexagonalWallpaperFormula{
		Formula: formula,
	}, nil
}

// NewHexagonalWallpaperFormulaFromMarshalObject uses a marshalled object to create a new object.
func NewHexagonalWallpaperFormulaFromMarshalObject(marshalObject WallpaperFormulaMarshalled) *HexagonalWallpaperFormula {
	return &HexagonalWallpaperFormula{
		Formula: NewWallpaperFormulaFromMarshalObject(marshalObject),
	}
}

// NewHexagonalWallpaperFormulaWithSymmetry will try to create a new Hexagonal Wallpaper Formula
//   with the desired Terms, Multiplier and Symmetry.
func NewHexagonalWallpaperFormulaWithSymmetry(terms []*formula.EisensteinFormulaTerm, wallpaperMultiplier complex128, desiredSymmetry *Symmetry) (*HexagonalWallpaperFormula, error) {
	err := checkForIncompatibleHexagonalSymmetries(desiredSymmetry)
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

	newBaseWallpaper := &HexagonalWallpaperFormula{
		Formula: &WallpaperFormula{
			WavePackets: newWavePackets,
			Multiplier:  wallpaperMultiplier,
		},
	}
	newBaseWallpaper.SetUp()
	return newBaseWallpaper, nil
}

func checkForIncompatibleHexagonalSymmetries(desiredSymmetry *Symmetry) error {
	desiredSymmetryAlreadySet := false
	if desiredSymmetry.P3m1 {
		if desiredSymmetryAlreadySet {
			return errors.New("invalid desired symmetry")
		}
		desiredSymmetryAlreadySet = true
	}
	if desiredSymmetry.P31m {
		if desiredSymmetryAlreadySet {
			return errors.New("invalid desired symmetry")
		}
		desiredSymmetryAlreadySet = true
	}
	if desiredSymmetry.P6 {
		if desiredSymmetryAlreadySet {
			return errors.New("invalid desired symmetry")
		}
		desiredSymmetryAlreadySet = true
	}
	if desiredSymmetry.P6m {
		if desiredSymmetryAlreadySet {
			return errors.New("invalid desired symmetry")
		}
		desiredSymmetryAlreadySet = true
	}
	return nil
}