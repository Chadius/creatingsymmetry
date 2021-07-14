package wavepacket

import (
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/coefficient"
	"wallpaper/entities/utility"
)

// WallpaperFormulaMarshalled can be marshalled into Wave Packet formulas
type WallpaperFormulaMarshalled struct {
	WavePackets []*Marshal                     `json:"wave_packets" yaml:"wave_packets"`
	Multiplier utility.ComplexNumberForMarshal `json:"multiplier" yaml:"multiplier"`
	Lattice *formula.LatticeVectorPairMarshal  `json:"lattice" yaml:"lattice"`
	DesiredSymmetry string `json:"desired_symmetry" yaml:"desired_symmetry"`
}

// WallpaperFormula uses wave packets that enforce rotation symmetry.
type WallpaperFormula struct {
	WavePackets []*WavePacket
	Multiplier complex128
	Lattice *formula.LatticeVectorPair
}

// SetUp adds locked Eisenstein terms to the formula based on the relationships.
//  Note there is NO way to change the multipliers.
func (wallpaperFormula *WallpaperFormula) SetUp(
	lockedRelationships []coefficient.Relationship,
	) {
	for _, wavePacket := range wallpaperFormula.WavePackets {
		baseCoefficientPairing := coefficient.Pairing{
			PowerN: wavePacket.Terms[0].PowerN,
			PowerM: wavePacket.Terms[0].PowerM,
		}

		newPairings := baseCoefficientPairing.GenerateCoefficientSets(lockedRelationships)

		for _, newCoefficientPair := range newPairings {
			newEisenstein := &formula.EisensteinFormulaTerm{
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

	zInLatticeCoordinates := wallpaperFormula.Lattice.ConvertToLatticeCoordinates(z)

	for _, wavePacket := range wallpaperFormula.WavePackets {
		termContribution := wavePacket.Calculate(zInLatticeCoordinates)
		result.Total += termContribution.Total / complex(float64(len(wavePacket.Terms)), 0)
		result.ContributionByTerm = append(result.ContributionByTerm, termContribution.Total)
	}
	result.Total *= wallpaperFormula.Multiplier

	return result
}

// NewWallpaperFormulaFromMarshalObject uses a marshalled object to create a new Wave WavePacket.
func NewWallpaperFormulaFromMarshalObject(marshalObject WallpaperFormulaMarshalled) *WallpaperFormula {
	wavePackets := []*WavePacket{}
	for _,packet := range marshalObject.WavePackets {
		newWavePacket := NewWaveFormulaFromMarshalObject(*packet)
		wavePackets = append(wavePackets, newWavePacket)
	}

	return &WallpaperFormula{
		WavePackets: wavePackets,
		Multiplier:  complex(marshalObject.Multiplier.Real, marshalObject.Multiplier.Imaginary),
	}
}
