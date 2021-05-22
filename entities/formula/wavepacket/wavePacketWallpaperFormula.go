package wavepacket

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/coefficient"
	"wallpaper/entities/utility"
)

// WallpaperFormulaMarshalled can be marshalled into Wave Packet formulas
type WallpaperFormulaMarshalled struct {
	WavePackets []*FormulaMarshalable			`json:"wave_packets" yaml:"wave_packets"`
	Multiplier utility.ComplexNumberForMarshal	`json:"multiplier" yaml:"multiplier"`
}

// WallpaperFormula uses wave packets that enforce rotation symmetry.
type WallpaperFormula struct {
	WavePackets []*Formula
	Multiplier complex128
}

// SetUp adds the locked Eisenstein terms to the formula.
//   SetUp assumes only the first term of each Wave Packet was given.
//   The type of wallpaper supplies the paired coefficient relationships
//     as well as the base vectors.
func (wallpaperFormula *WallpaperFormula) SetUp(
	lockedRelationships []coefficient.Relationship,
	baseXVector complex128,
	baseYVector complex128,
	) {
	for _, wavePacket := range wallpaperFormula.WavePackets {
		baseCoefficientPairing := coefficient.Pairing{
			PowerN: wavePacket.Terms[0].PowerN,
			PowerM: wavePacket.Terms[0].PowerM,
		}

		newPairings := baseCoefficientPairing.GenerateCoefficientSets(lockedRelationships)

		wavePacket.Terms[0].XLatticeVector = complex(real(baseXVector), imag(baseXVector))
		wavePacket.Terms[0].YLatticeVector = complex(real(baseYVector), imag(baseYVector))

		for _, newCoefficientPair := range newPairings {
			xVector := baseXVector
			yVector := baseYVector

			if newCoefficientPair.NegateMultiplier == true {
				xVector *= -1
				yVector *= -1
			}
			newEisenstein := &formula.EisensteinFormulaTerm{
				XLatticeVector: xVector,
				YLatticeVector: yVector,
				PowerN:         newCoefficientPair.PowerN,
				PowerM:         newCoefficientPair.PowerM,
			}
			wavePacket.Terms = append(wavePacket.Terms, newEisenstein)
		}
	}
}

// Calculate takes the complex number z and processes it using the mathematical terms.
func (wallpaperFormula *WallpaperFormula) Calculate(z complex128) *formula.CalculationResultForFormula {

	result := &formula.CalculationResultForFormula{
		Total: complex(0,0),
		ContributionByTerm: []complex128{},
	}

	for _, wavePacket := range wallpaperFormula.WavePackets {
		termContribution := wavePacket.Calculate(z)
		result.Total += termContribution.Total / complex(float64(len(wavePacket.Terms)), 0)
		result.ContributionByTerm = append(result.ContributionByTerm, termContribution.Total)
	}
	result.Total *= wallpaperFormula.Multiplier

	return result
}

// NewWallpaperFormulaFromYAML reads the data and returns a formula from it.
func NewWallpaperFormulaFromYAML(data []byte) (*WallpaperFormula, error) {
	return newWallpaperFormulaFromDatastream(data, yaml.Unmarshal)
}

// NewWallpaperFormulaFromJSON reads the data and returns a formula from it.
func NewWallpaperFormulaFromJSON(data []byte) (*WallpaperFormula, error) {
	return newWallpaperFormulaFromDatastream(data, json.Unmarshal)
}

//newWallpaperFormulaFromDatastream consumes a given bytestream and tries to create a new object from it.
func newWallpaperFormulaFromDatastream(data []byte, unmarshal utility.UnmarshalFunc) (*WallpaperFormula, error) {
	var unmarshalError error
	var formulaMarshal WallpaperFormulaMarshalled
	unmarshalError = unmarshal(data, &formulaMarshal)

	if unmarshalError != nil {
		return nil, unmarshalError
	}

	formula := NewWallpaperFormulaFromMarshalObject(formulaMarshal)
	return formula, nil
}

// NewWallpaperFormulaFromMarshalObject uses a marshalled object to create a new Wave Formula.
func NewWallpaperFormulaFromMarshalObject(marshalObject WallpaperFormulaMarshalled) *WallpaperFormula {
	wavePackets := []*Formula{}
	for _,packet := range marshalObject.WavePackets {
		newWavePacket := NewWaveFormulaFromMarshalObject(*packet)
		wavePackets = append(wavePackets, newWavePacket)
	}

	return &WallpaperFormula{
		WavePackets: wavePackets,
		Multiplier:  complex(marshalObject.Multiplier.Real, marshalObject.Multiplier.Imaginary),
	}
}
