package wallpaper

import (
	"encoding/json"
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	eisensteinFormula "github.com/Chadius/creating-symmetry/entities/oldformula/eisenstien"
	"github.com/Chadius/creating-symmetry/entities/oldformula/latticevector"
	"github.com/Chadius/creating-symmetry/entities/oldformula/result"
	"github.com/Chadius/creating-symmetry/entities/utility"
	"gopkg.in/yaml.v2"
)

// DimensionsMarshal tracks the width and height
type DimensionsMarshal struct {
	Width  float64 `json:"width" yaml:"width"`
	Height float64 `json:"height" yaml:"height"`
}

// Dimensions tracks the width and height
type Dimensions struct {
	Width  float64
	Height float64
}

// FormulaMarshal can be created from data streams and used to create Formula objects.
type FormulaMarshal struct {
	LatticeType     string                          `json:"lattice_type" yaml:"lattice_type"`
	LatticeSize     *DimensionsMarshal              `json:"lattice_size" yaml:"lattice_size"`
	Multiplier      utility.ComplexNumberForMarshal `json:"multiplier" yaml:"multiplier"`
	WavePackets     []*Marshal                      `json:"wave_packets" yaml:"wave_packets"`
	DesiredSymmetry string                          `json:"desired_symmetry" yaml:"desired_symmetry"`
}

// LatticeType notes the shape of the underlying lattice
type LatticeType string

const (
	// Generic lattice can use any 4 sided shape
	Generic LatticeType = "generic"
	// Hexagonal lattice will have 3 way rotational symmetry
	Hexagonal LatticeType = "hexagonal"
	// Rectangular lattice will be aligned along TransformedX & TransformedY axis, but they may not be the same size.
	Rectangular LatticeType = "rectangular"
	// Rhombic lattices all have the same size but may not be aligned along axes (and may not be at right angles)
	Rhombic LatticeType = "rhombic"
	// Square lattice will be aligned along TransformedX & TransformedY axis and be the same size.
	Square LatticeType = "square"
)

// Formula stores the information needed to create wallpapers using a Lattice.
type Formula struct {
	LatticeType     LatticeType
	LatticeSize     *Dimensions
	Lattice         *latticevector.Pair
	Multiplier      complex128
	WavePackets     []*WavePacket
	DesiredSymmetry Symmetry
}

// NewFormulaFromYAML returns a new Formula from the given YAML.
func NewFormulaFromYAML(data []byte) (*Formula, error) {
	return newFormulaFromDatastream(data, yaml.Unmarshal)
}

// NewFormulaFromJSON returns a new Formula from the given JSON.
func NewFormulaFromJSON(data []byte) (*Formula, error) {
	return newFormulaFromDatastream(data, json.Unmarshal)
}

func newFormulaFromDatastream(data []byte, unmarshal utility.UnmarshalFunc) (*Formula, error) {
	var unmarshalError error
	var marshal FormulaMarshal
	unmarshalError = unmarshal(data, &marshal)

	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return NewFormulaFromMarshalObject(marshal), nil
}

// NewFormulaFromMarshalObject converts a marshaled oldformula into a oldformula object
func NewFormulaFromMarshalObject(marshaledFormula FormulaMarshal) *Formula {
	wavePackets := []*WavePacket{}
	for _, packet := range marshaledFormula.WavePackets {
		newWavePacket := NewWaveFormulaFromMarshalObject(*packet)
		wavePackets = append(wavePackets, newWavePacket)
	}

	desiredSymmetry := P1
	if marshaledFormula.DesiredSymmetry != "" {
		desiredSymmetry = Symmetry(marshaledFormula.DesiredSymmetry)
	}

	latticeWidth := 0.0
	latticeHeight := 0.0
	if marshaledFormula.LatticeSize != nil {
		latticeWidth = marshaledFormula.LatticeSize.Width
		latticeHeight = marshaledFormula.LatticeSize.Height
	}

	return &Formula{
		LatticeType:     LatticeType(marshaledFormula.LatticeType),
		LatticeSize:     &Dimensions{Width: latticeWidth, Height: latticeHeight},
		Lattice:         &latticevector.Pair{XLatticeVector: complex(0, 0), YLatticeVector: complex(0, 0)},
		Multiplier:      complex(marshaledFormula.Multiplier.Real, marshaledFormula.Multiplier.Imaginary),
		WavePackets:     wavePackets,
		DesiredSymmetry: desiredSymmetry,
	}
}

// Setup creates lattice vectors and locked in Eisenstein pairs based on the Lattice Type.
//  modifies the given Formula.
//  returns any errors (returns nil if there are no errors)
func (formula *Formula) Setup() error {
	vectorErr := formula.createVectors()
	if vectorErr != nil {
		return vectorErr
	}

	formula.satisfyDesiredSymmetry()

	formula.lockEisensteinTerms()

	return nil
}

func (formula *Formula) createVectors() error {
	type VectorCreator func(formula *Formula) error

	vectorCreatorBasedOnLatticeType := map[LatticeType]VectorCreator{
		Generic:     createVectorsForGenericWallpaper,
		Hexagonal:   createVectorsForHexagonalWallpaper,
		Rhombic:     createVectorsForRhombicWallpaper,
		Square:      createVectorsForSquareWallpaper,
		Rectangular: createVectorsForRectangularWallpaper,
	}

	customErr := vectorCreatorBasedOnLatticeType[formula.LatticeType](formula)
	if customErr != nil {
		return customErr
	}

	return formula.Lattice.Validate()
}

// lockEisensteinTerms creates eisenstein Terms based on the LatticeType
func (formula *Formula) lockEisensteinTerms() {
	if formula.LatticeType == Generic || formula.LatticeType == Rectangular {
		return
	}

	type EisensteinCreator func(formula *Formula)

	lockedCoefficientPairsBasedOnLatticeType := map[LatticeType]EisensteinCreator{
		Hexagonal: lockCoefficientPairsForHexagonalWallpaper,
		Rhombic:   lockCoefficientPairsForRhombicWallpaper,
		Square:    lockCoefficientPairsForSquareWallpaper,
	}

	lockedCoefficientPairsBasedOnLatticeType[formula.LatticeType](formula)
}

// lockEisensteinTermsBasedOnRelationship adds locked Eisenstein terms to the oldformula based on the relationships.
func (formula *Formula) lockEisensteinTermsBasedOnRelationship(
	lockedRelationships []coefficient.Relationship,
) {
	for _, wavePacket := range formula.WavePackets {
		baseCoefficientPairing := coefficient.Pairing{
			PowerN: wavePacket.Terms[0].PowerN,
			PowerM: wavePacket.Terms[0].PowerM,
		}

		newPairings := baseCoefficientPairing.GenerateCoefficientSets(lockedRelationships)

		for _, newCoefficientPair := range newPairings {
			newEisenstein := &eisensteinFormula.EisensteinFormulaTerm{
				PowerN: newCoefficientPair.PowerN,
				PowerM: newCoefficientPair.PowerM,
			}
			wavePacket.Terms = append(wavePacket.Terms, newEisenstein)
		}
	}
}

// satisfyDesiredSymmetry creates WavePackets to satisfy DesiredSymmetry.
func (formula *Formula) satisfyDesiredSymmetry() {
	newWavePackets := []*WavePacket{}
	for _, existingWavePacket := range formula.WavePackets {
		newWavePackets = append(newWavePackets, existingWavePacket)
	}

	for _, existingWavePacket := range formula.WavePackets {
		newWavePackets = addNewWavePacketsBasedOnSymmetry(existingWavePacket.Terms[0], existingWavePacket.Multiplier, formula.DesiredSymmetry, newWavePackets)
	}

	formula.WavePackets = newWavePackets
}

// HasSymmetry returns true if the WavePackets involved form symmetry.
func (formula *Formula) HasSymmetry(targetSymmetry Symmetry) bool {
	if targetSymmetry == P1 {
		return true
	}

	type SymmetryChecker func(formula *Formula, targetSymmetry Symmetry) bool

	checksForSymmetryBasedOnLatticeType := map[LatticeType]SymmetryChecker{
		Generic:     checksForSymmetryForGenericType,
		Hexagonal:   checksForSymmetryForHexagonalType,
		Rhombic:     checksForSymmetryForRhombicType,
		Square:      checksForSymmetryForSquareType,
		Rectangular: checksForSymmetryForRectangularType,
	}

	return checksForSymmetryBasedOnLatticeType[formula.LatticeType](formula, targetSymmetry)
}

//Calculate applies the oldformula to the complex number z.
// It modifies the oldformula's result to track the contribution per term
// As well as the final numerical result.
func (formula *Formula) Calculate(z complex128) *result.CalculationResultForFormula {
	result := &result.CalculationResultForFormula{
		Total:              complex(0, 0),
		ContributionByTerm: []complex128{},
	}

	zInLatticeCoordinates := formula.Lattice.ConvertToLatticeCoordinates(z)

	for _, wavePacket := range formula.WavePackets {
		termContribution := wavePacket.Calculate(zInLatticeCoordinates)
		result.Total += termContribution.Total / complex(float64(len(wavePacket.Terms)), 0)
		result.ContributionByTerm = append(result.ContributionByTerm, termContribution.Total)
	}
	result.Total *= formula.Multiplier

	return result
}
