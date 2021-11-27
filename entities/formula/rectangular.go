package formula

import (
	"errors"
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
)

// Rectangular formulas will transform points by returning the same coordinates.
type Rectangular struct {
	latticeVectors []complex128
	wavePackets    []WavePacket
}

// NewRectangularFormula returns a new formula object.
func NewRectangularFormula(packets []WavePacket, latticeHeight float64, desiredSymmetry Symmetry) (*Rectangular, error) {
	if latticeHeight == 0.0 {
		return nil, errors.New("rectangular lattice must specify height")
	}

	if desiredSymmetry != "" &&
		desiredSymmetry != P1 &&
		desiredSymmetry != Pm &&
		desiredSymmetry != Pg &&
		desiredSymmetry != Pmm &&
		desiredSymmetry != Pmg &&
		desiredSymmetry != Pgg {
		return nil, errors.New("rectangular lattice can apply these desired symmetries: P1, Pm, Pg, Pmm, Pmg, Pgg")
	}

	wavePacketsWithDesiredSymmetry := createNewWavePacketsBasedOnDesiredSymmetry(packets, desiredSymmetry)

	return &Rectangular{
			wavePackets: wavePacketsWithDesiredSymmetry,
			latticeVectors: []complex128{
				complex(1, 0),
				complex(0, latticeHeight),
			},
		},
		nil
}

// WavePackets returns the wave packets used.
func (r *Rectangular) WavePackets() []WavePacket {
	return r.wavePackets
}

// Calculate transforms the coordinate using the Rectangular lattice's wave packets.
func (r *Rectangular) Calculate(coordinate complex128) complex128 {
	return CalculateCoordinateUsingWavePackets(coordinate, r.LatticeVectors(), r.WavePackets())
}

// FormulaLevelTerms returns an empty list, Rectangular formulas do not have base-level terms.
func (r *Rectangular) FormulaLevelTerms() []Term {
	return nil
}

// LatticeVectors returns the lattice vectors used to create the rectangle.
func (r *Rectangular) LatticeVectors() []complex128 {
	return r.latticeVectors
}

// SymmetriesFound returns all symmetries found in this pattern.
func (r *Rectangular) SymmetriesFound() []Symmetry {
	symmetriesFound := []Symmetry{P1}

	expectedCoefficientsBySymmetry := map[Symmetry][]coefficient.Relationship{
		Pm: {coefficient.PlusNMinusM},
		Pg: {coefficient.PlusNMinusMNegateMultiplierIfOddPowerN},
		Pmm: {
			coefficient.PlusNMinusM,
			coefficient.MinusNMinusM,
			coefficient.MinusNPlusM,
		},
		Pmg: {
			coefficient.MinusNMinusM,
			coefficient.PlusNMinusMNegateMultiplierIfOddPowerN,
			coefficient.MinusNPlusMNegateMultiplierIfOddPowerN,
		},
		Pgg: {
			coefficient.MinusNMinusM,
			coefficient.PlusNMinusMNegateMultiplierIfOddPowerSum,
			coefficient.MinusNPlusMNegateMultiplierIfOddPowerSum,
		},
	}

	order := []Symmetry{Pm, Pg, Pmm, Pmg, Pgg}
	for _, symmetryType := range order {
		coefficients := expectedCoefficientsBySymmetry[symmetryType]
		if HasSymmetry(r.WavePackets(), symmetryType, map[Symmetry][]coefficient.Relationship{
			symmetryType: coefficients,
		}) {
			symmetriesFound = append(symmetriesFound, symmetryType)
		}
	}

	return symmetriesFound
}
