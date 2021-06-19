package wavepacket

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/coefficient"
	"wallpaper/entities/utility"
)

// RhombicWallpaperFormulaMarshalled can be marshalled into a Rhombic formula
type RhombicWallpaperFormulaMarshalled struct {
	Formula *WallpaperFormulaMarshalled `json:"formula" yaml:"formula"`
	LatticeHeight float64  `json:"lattice_height" yaml:"lattice_height"`
}

// RhombicWallpaperFormula helps transform one point to a 2D wallpaper pattern that uses the Rhombic lattice.
//   The underlying lattice forms 4 lines of equal length. Common forms are Squares and Diamonds.
type RhombicWallpaperFormula struct {
	Formula *WallpaperFormula
	LatticeHeight float64
}

// SetUp will create the rhombic RhombicWallpaperFormula using the given LatticeHeight.
func (rhombic *RhombicWallpaperFormula) SetUp() {

	rhombic.Formula.Lattice = &formula.LatticeVectorPair{
		XLatticeVector: complex(0.5, rhombic.LatticeHeight),
		YLatticeVector: complex(0.5, rhombic.LatticeHeight * -1),
	}

	rhombic.Formula.SetUp(
		[]coefficient.Relationship{
			coefficient.PlusMPlusN,
		},
	)
}

// Calculate applies the formula to the complex number z.
//  It modifies the formula's result to track the contribution per term
//  As well as the final numerical result.
func (rhombic *RhombicWallpaperFormula) Calculate(z complex128) *formula.CalculationResultForFormula {
	return rhombic.Formula.Calculate(z)
}

// NewRhombicWallpaperFormulaFromJSON reads the data and returns a formula term from it.
func NewRhombicWallpaperFormulaFromJSON(data []byte) (*RhombicWallpaperFormula, error) {
	return newRhombicWallpaperFormulaFromDatastream(data, json.Unmarshal)
}

// NewRhombicWallpaperFormulaFromYAML reads the data and returns a formula term from it.
func NewRhombicWallpaperFormulaFromYAML(data []byte) (*RhombicWallpaperFormula, error) {
	return newRhombicWallpaperFormulaFromDatastream(data, yaml.Unmarshal)
}

//newRhombicWallpaperFormulaFromDatastream consumes a given bytestream and tries to create a new object from it.
func newRhombicWallpaperFormulaFromDatastream(data []byte, unmarshal utility.UnmarshalFunc) (*RhombicWallpaperFormula, error) {
	var unmarshalError error
	var rhombicFormulaMarshalled RhombicWallpaperFormulaMarshalled
	unmarshalError = unmarshal(data, &rhombicFormulaMarshalled)

	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return NewRhombicWallpaperFormulaFromMarshalObject(rhombicFormulaMarshalled), nil
}

// NewRhombicWallpaperFormulaFromMarshalObject uses a marshalled object to create a new object.
func NewRhombicWallpaperFormulaFromMarshalObject(marshalObject RhombicWallpaperFormulaMarshalled) *RhombicWallpaperFormula {
	formula := NewWallpaperFormulaFromMarshalObject(*marshalObject.Formula)
	return &RhombicWallpaperFormula{
		Formula:       formula,
		LatticeHeight: marshalObject.LatticeHeight,
	}
}