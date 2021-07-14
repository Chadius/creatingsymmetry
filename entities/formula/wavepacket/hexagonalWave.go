package wavepacket

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"math"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/coefficient"
	"wallpaper/entities/utility"
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

// HasSymmetry returns true if the WavePackets involved form symmetry.
func (hexWaveFormula *HexagonalWallpaperFormula) HasSymmetry(desiredSymmetry Symmetry) bool {
	if desiredSymmetry == P3 {
		return true
	}

	return HasSymmetry(hexWaveFormula.Formula.WavePackets, desiredSymmetry, map[Symmetry][]coefficient.Relationship {
		P31m: {coefficient.PlusMPlusN},
		P3m1: {coefficient.MinusMMinusN},
		P6: {coefficient.MinusNMinusM},
		P6m: {
			coefficient.MinusNMinusM,
			coefficient.MinusMMinusN,
			coefficient.PlusMPlusN,
		},
	})
}

// NewHexagonalWallpaperFormulaFromJSON reads the data and returns a formula term from it.
func NewHexagonalWallpaperFormulaFromJSON(data []byte) (*HexagonalWallpaperFormula, error) {
	return newHexagonalWallpaperFormulaFromDatastream(data, json.Unmarshal)
}

// NewHexagonalWallpaperFormulaFromYAML reads the data and returns a formula term from it.
func NewHexagonalWallpaperFormulaFromYAML(data []byte) (*HexagonalWallpaperFormula, error) {
	return newHexagonalWallpaperFormulaFromDatastream(data, yaml.Unmarshal)
}

//newHexagonalWallpaperFormulaFromDatastream consumes a given bytestream and tries to create a new object from it.
func newHexagonalWallpaperFormulaFromDatastream(data []byte, unmarshal utility.UnmarshalFunc) (*HexagonalWallpaperFormula, error) {
	var unmarshalError error
	var hexFormulaMarshalled WallpaperFormulaMarshalled
	unmarshalError = unmarshal(data, &hexFormulaMarshalled)

	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return NewHexagonalWallpaperFormulaFromMarshalObject(hexFormulaMarshalled), nil
}

// NewHexagonalWallpaperFormulaFromMarshalObject uses a marshalled object to create a new object.
func NewHexagonalWallpaperFormulaFromMarshalObject(marshalObject WallpaperFormulaMarshalled) *HexagonalWallpaperFormula {
	formula := NewWallpaperFormulaFromMarshalObject(marshalObject)

	if marshalObject.DesiredSymmetry != "" {
		wallpaper, err := NewHexagonalWallpaperFormulaWithSymmetry(
			formula.WavePackets[0].Terms,
			formula.Multiplier,
			Symmetry(marshalObject.DesiredSymmetry),
		)

		if err != nil {
			return nil
		}
		return wallpaper
	}

	return &HexagonalWallpaperFormula{
		Formula:       formula,
	}
}

// NewHexagonalWallpaperFormulaWithSymmetry will try to create a new Hexagonal Wallpaper
//   with the desired Terms, Multiplier and Symmetry.
func NewHexagonalWallpaperFormulaWithSymmetry(terms []*formula.EisensteinFormulaTerm, wallpaperMultiplier complex128, desiredSymmetry Symmetry) (*HexagonalWallpaperFormula, error) {
	newWavePackets := []*WavePacket{}
	for _, term := range terms {
		newWavePackets = append(
			newWavePackets,
			&WavePacket{
				Terms:      []*formula.EisensteinFormulaTerm{term},
				Multiplier: wallpaperMultiplier,
			},
		)

		newWavePackets = addNewWavePacketsBasedOnSymmetry(term, wallpaperMultiplier, desiredSymmetry, newWavePackets)
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
