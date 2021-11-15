package formula

import (
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
)

// Square formulas will transform points by returning the same coordinates.
type Square struct {
	latticeVectors []complex128
	wavePackets []WavePacket
}

// NewSquareFormula returns a new formula object.
func NewSquareFormula(packets []WavePacket) (*Square, error) {
	packetsWithLockedCoefficients := lockTermsBasedOnRelationship(
		[]coefficient.Relationship{
			coefficient.PlusMMinusN,
			coefficient.MinusNMinusM,
			coefficient.MinusMPlusN,
		}, 
		packets)
	
	squareWallpaperFormula := &Square{
		wavePackets: packetsWithLockedCoefficients,
		latticeVectors: []complex128{
			complex(1, 0),
			complex(0, 1),
		},
	}
	
	return squareWallpaperFormula, nil
}

// WavePackets returns the wave packets used.
func (r *Square) WavePackets() []WavePacket {
	return r.wavePackets
}

// Calculate transforms the coordinate using the Square lattice's wave packets.
func (r *Square) Calculate(coordinate complex128) complex128 {
	result := complex(0,0)
	return result
/*	zInLatticeCoordinates := ConvertToLatticeCoordinates(coordinate, r.LatticeVectors())

	for _, wavePacket := range r.WavePackets() {
		termContribution := wavePacket.Calculate(zInLatticeCoordinates)
		result += termContribution / complex(float64(len(wavePacket.Terms)), 0)
	}

	return result
*/}

// FormulaLevelTerms returns an empty list, Square formulas do not have base-level terms.
func (r *Square) FormulaLevelTerms() []Term {
	return nil
}

// LatticeVectors returns the lattice vectors used to create the rectangle.
func (r *Square) LatticeVectors() []complex128 {
	return r.latticeVectors
}