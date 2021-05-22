package wave

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
	WavePackets []*Formula
	Multiplier complex128
}

// SetUp initializes all of the needed wallpaper terms.
func (squareWaveFormula *SquareWallpaperFormula) SetUp() {
	for _, wavePacket := range squareWaveFormula.WavePackets {
		baseCoefficientPairing := coefficient.Pairing{
			PowerN: wavePacket.Terms[0].PowerN,
			PowerM: wavePacket.Terms[0].PowerM,
		}

		newPairings := baseCoefficientPairing.GenerateCoefficientSets([]coefficient.Relationship{
			coefficient.PlusMMinusN,
			coefficient.MinusNMinusM,
			coefficient.MinusMPlusN,
		})

		baseXVector := complex(1, 0)
		baseYVector := complex(0, 1)
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
func (squareWaveFormula *SquareWallpaperFormula) Calculate(z complex128) *formula.CalculationResultForFormula {

	result := &formula.CalculationResultForFormula{
		Total: complex(0,0),
		ContributionByTerm: []complex128{},
	}

	for _, wavePacket := range squareWaveFormula.WavePackets {
		termContribution := wavePacket.Calculate(z)
		result.Total += termContribution.Total * 1.0/4.0
		result.ContributionByTerm = append(result.ContributionByTerm, termContribution.Total)
	}
	result.Total *= squareWaveFormula.Multiplier

	return result
}

// NewSquareWallpaperFormulaFromJSON reads the data and returns a formula term from it.
func NewSquareWallpaperFormulaFromJSON(data []byte) (*SquareWallpaperFormula, error) {
	return newSquareWallpaperFormulaFromDatastream(data, json.Unmarshal)
}

// NewSquareWallpaperFormulaFromYAML reads the data and returns a formula term from it.
func NewSquareWallpaperFormulaFromYAML(data []byte) (*SquareWallpaperFormula, error) {
	return newSquareWallpaperFormulaFromDatastream(data, yaml.Unmarshal)
}

//newSquareWallpaperFormulaFromDatastream consumes a given bytestream and tries to create a new object from it.
func newSquareWallpaperFormulaFromDatastream(data []byte, unmarshal utility.UnmarshalFunc) (*SquareWallpaperFormula, error) {
	var unmarshalError error
	var formulaMarshal WallpaperFormulaMarshalled
	unmarshalError = unmarshal(data, &formulaMarshal)

	if unmarshalError != nil {
		return nil, unmarshalError
	}

	formulaTerm := NewSquareWallpaperFormulaFromMarshalObject(formulaMarshal)
	return formulaTerm, nil
}

// NewSquareWallpaperFormulaFromMarshalObject uses a marshalled object to create a new object.
func NewSquareWallpaperFormulaFromMarshalObject(marshalObject WallpaperFormulaMarshalled) *SquareWallpaperFormula {
	wavePackets := []*Formula{}
	for _,packet := range marshalObject.WavePackets {
		newWavePacket := NewWaveFormulaFromMarshalObject(*packet)
		wavePackets = append(wavePackets, newWavePacket)
	}

	return &SquareWallpaperFormula{
		WavePackets: wavePackets,
		Multiplier:	complex(marshalObject.Multiplier.Real, marshalObject.Multiplier.Imaginary),
	}
}
