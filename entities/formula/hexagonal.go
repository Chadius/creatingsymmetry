package formula

import (
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	"math"
)

// Hexagonal formulas will transform points by returning the same coordinates.
type Hexagonal struct {
	latticeVectors []complex128
	wavePackets []WavePacket
}

// NewHexagonalFormula returns a new formula object.
func NewHexagonalFormula(packets []WavePacket) (*Hexagonal, error) {
	packetsWithLockedCoefficients := lockTermsBasedOnRelationship(
		[]coefficient.Relationship{
			coefficient.PlusMMinusSumNAndM,
			coefficient.MinusSumNAndMPlusN,
		},
		packets)

	return &Hexagonal{
		wavePackets: packetsWithLockedCoefficients,
		latticeVectors: []complex128{
			complex(1, 0),
			complex(-0.5, math.Sqrt(3.0)/2.0),
		},
	},
	nil
}

// WavePackets returns the wave packets used.
func (r *Hexagonal) WavePackets() []WavePacket {
	return r.wavePackets
}

// Calculate transforms the coordinate using the Hexagonal lattice's wave packets.
func (r *Hexagonal) Calculate(coordinate complex128) complex128 {
	return CalculateCoordinateUsingWavePackets(coordinate, r.LatticeVectors(), r.WavePackets())
}

// FormulaLevelTerms returns an empty list, Hexagonal formulas do not have base-level terms.
func (r *Hexagonal) FormulaLevelTerms() []Term {
	return nil
}

// LatticeVectors returns the lattice vectors used to create the rectangle.
func (r *Hexagonal) LatticeVectors() []complex128 {
	return r.latticeVectors
}