package wavepacket

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"wallpaper/entities/formula"
	"wallpaper/entities/utility"

	//"wallpaper/entities/utility"
)

// GenericWallpaperFormulaMarshalled can be marshalled into a Generic formula
type GenericWallpaperFormulaMarshalled struct {
	Formula *WallpaperFormulaMarshalled `json:"formula" yaml:"formula"`
	VectorWidth float64  `json:"vector_width" yaml:"vector_width"`
	VectorHeight float64  `json:"vector_height" yaml:"vector_height"`
}

// GenericWallpaperFormula helps transform one point to a 2D wallpaper pattern that uses the Generic lattice.
//   The underlying lattice has 1 vector that is 1 unit horizontal. The second vector uses VectorWidth and VectorHeight.
//   VectorHeight cannot be 0.
type GenericWallpaperFormula struct {
	Formula *WallpaperFormula
	VectorWidth float64
	VectorHeight float64
}

// SetUp will create the Generic GenericWallpaperFormula using the given LatticeHeight.
func (Generic *GenericWallpaperFormula) SetUp() error {

	Generic.Formula.Lattice = &formula.LatticeVectorPair{
		XLatticeVector: complex(1, 0),
		YLatticeVector: complex(Generic.VectorWidth, Generic.VectorHeight),
	}

	err := Generic.Formula.Lattice.Validate()
	if err != nil {
		return err
	}

	return nil
}

//Calculate applies the formula to the complex number z.
// It modifies the formula's result to track the contribution per term
// As well as the final numerical result.
func (Generic *GenericWallpaperFormula) Calculate(z complex128) *formula.CalculationResultForFormula {
	return Generic.Formula.Calculate(z)
}

//// HasSymmetry returns true if the WavePackets involved form symmetry.
//func (Generic *GenericWallpaperFormula) HasSymmetry(desiredSymmetry Symmetry) bool {
//	return HasSymmetry(Generic.Formula.WavePackets, desiredSymmetry, map[Symmetry][]coefficient.Relationship {
//		Pm: {coefficient.PlusNMinusM},
//		Pg: {coefficient.PlusNMinusMNegateMultiplierIfOddPowerN},
//		Pmm: {
//			coefficient.PlusNMinusM,
//			coefficient.MinusNMinusM,
//			coefficient.MinusNPlusM,
//		},
//		Pmg: {
//			coefficient.MinusNMinusM,
//			coefficient.PlusNMinusMNegateMultiplierIfOddPowerN,
//			coefficient.MinusNPlusMNegateMultiplierIfOddPowerN,
//		},
//		Pgg: {
//			coefficient.MinusNMinusM,
//			coefficient.PlusNMinusMNegateMultiplierIfOddPowerSum,
//			coefficient.MinusNPlusMNegateMultiplierIfOddPowerSum,
//		},
//	})
//}

// NewGenericWallpaperFormulaFromJSON reads the data and returns a formula term from it.
func NewGenericWallpaperFormulaFromJSON(data []byte) (*GenericWallpaperFormula, error) {
	return newGenericWallpaperFormulaFromDatastream(data, json.Unmarshal)
}

// NewGenericWallpaperFormulaFromYAML reads the data and returns a formula term from it.
func NewGenericWallpaperFormulaFromYAML(data []byte) (*GenericWallpaperFormula, error) {
	return newGenericWallpaperFormulaFromDatastream(data, yaml.Unmarshal)
}

//newGenericWallpaperFormulaFromDatastream consumes a given bytestream and tries to create a new object from it.
func newGenericWallpaperFormulaFromDatastream(data []byte, unmarshal utility.UnmarshalFunc) (*GenericWallpaperFormula, error) {
	var unmarshalError error
	var GenericFormulaMarshalled GenericWallpaperFormulaMarshalled
	unmarshalError = unmarshal(data, &GenericFormulaMarshalled)

	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return NewGenericWallpaperFormulaFromMarshalObject(GenericFormulaMarshalled), nil
}

// NewGenericWallpaperFormulaFromMarshalObject uses a marshalled object to create a new object.
func NewGenericWallpaperFormulaFromMarshalObject(marshalObject GenericWallpaperFormulaMarshalled) *GenericWallpaperFormula {
	formula := NewWallpaperFormulaFromMarshalObject(*marshalObject.Formula)

	//if marshalObject.Formula.DesiredSymmetry != "" {
	//	wallpaper, err := NewGenericWallpaperFormulaWithSymmetry(
	//		formula.WavePackets[0].Terms,
	//		formula.Multiplier,
	//		marshalObject.LatticeHeight,
	//		Symmetry(marshalObject.Formula.DesiredSymmetry),
	//	)
	//
	//	if err != nil {
	//		return nil
	//	}
	//	return wallpaper
	//}

	return &GenericWallpaperFormula{
		Formula:       formula,
		VectorWidth: marshalObject.VectorWidth,
		VectorHeight: marshalObject.VectorHeight,
	}
}

//// NewGenericWallpaperFormulaWithSymmetry will try to create a new GenericWallpaperFormula WavePacket
////   with the desired Terms, Multiplier and Symmetry.
//func NewGenericWallpaperFormulaWithSymmetry(terms []*formula.EisensteinFormulaTerm, wallpaperMultiplier complex128, latticeHeight float64, desiredSymmetry Symmetry) (*GenericWallpaperFormula, error) {
//	newWavePackets := []*WavePacket{}
//	for _, term := range terms {
//		newWavePackets = append(
//			newWavePackets,
//			&WavePacket{
//				Terms:      []*formula.EisensteinFormulaTerm{term},
//				Multiplier: wallpaperMultiplier,
//			},
//		)
//
//		newWavePackets = addNewWavePacketsBasedOnSymmetry(term, wallpaperMultiplier, desiredSymmetry, newWavePackets)
//	}
//
//	newBaseWallpaper := &GenericWallpaperFormula{
//		Formula: &WallpaperFormula{
//			WavePackets: newWavePackets,
//			Multiplier:  wallpaperMultiplier,
//		},
//		LatticeHeight: latticeHeight,
//	}
//	newBaseWallpaper.SetUp()
//	return newBaseWallpaper, nil
//}
