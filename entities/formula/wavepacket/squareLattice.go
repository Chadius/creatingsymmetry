package wavepacket

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/coefficient"
	"wallpaper/entities/utility"
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
	return newSquareWallpaperFormulaFromDatastream(data, json.Unmarshal)
}

// NewSquareWallpaperFormulaFromYAML reads the data and returns a formula term from it.
func NewSquareWallpaperFormulaFromYAML(data []byte) (*SquareWallpaperFormula, error) {
	return newSquareWallpaperFormulaFromDatastream(data, yaml.Unmarshal)
}

// NewSquareWallpaperFormulaFromMarshalObject uses a marshalled object to create a new object.
func NewSquareWallpaperFormulaFromMarshalObject(marshalObject WallpaperFormulaMarshalled) *SquareWallpaperFormula {
	formula := NewWallpaperFormulaFromMarshalObject(marshalObject)

	if marshalObject.DesiredSymmetry != "" {
		wallpaper, err := NewSquareWallpaperFormulaWithSymmetry(
			formula.WavePackets[0].Terms,
			formula.Multiplier,
			Symmetry(marshalObject.DesiredSymmetry),
		)

		if err != nil {
			return nil
		}
		return wallpaper
	}

	return &SquareWallpaperFormula{
		Formula:       formula,
	}
}

//newSquareWallpaperFormulaFromDatastream consumes a given bytestream and tries to create a new object from it.
func newSquareWallpaperFormulaFromDatastream(data []byte, unmarshal utility.UnmarshalFunc) (*SquareWallpaperFormula, error) {
	var unmarshalError error
	var hexFormulaMarshalled WallpaperFormulaMarshalled
	unmarshalError = unmarshal(data, &hexFormulaMarshalled)

	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return NewSquareWallpaperFormulaFromMarshalObject(hexFormulaMarshalled), nil
}

// NewSquareWallpaperFormulaWithSymmetry will try to create a new Hexagonal RhombicWallpaperFormula WavePacket
//   with the desired Terms, Multiplier and Symmetry.
func NewSquareWallpaperFormulaWithSymmetry(terms []*formula.EisensteinFormulaTerm, wallpaperMultiplier complex128, desiredSymmetry Symmetry) (*SquareWallpaperFormula, error) {
	newWavePackets := []*WavePacket{}
	for _, term := range terms {
		if real(term.Multiplier) == 0 && imag(term.Multiplier) == 0 {
			term.Multiplier = complex(1, 0)
		}

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
