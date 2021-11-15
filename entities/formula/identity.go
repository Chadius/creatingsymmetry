package formula

// Identity formulas will transform points by returning the same coordinates.
type Identity struct {}

// Calculate returns the coordinate, the Identity formula returns the given input.
func (i *Identity) Calculate(coordinate complex128) complex128 {
	return coordinate
}

// WavePackets returns an empty array, this type of formula does not use WavePackets.
func (i *Identity) WavePackets() []WavePacket {
	return nil
}

// FormulaLevelTerms returns an empty list, Identity formulas do not have terms.
func (i *Identity) FormulaLevelTerms() []Term {
	return nil
}

// LatticeVectors returns an empty list, this formula does not use them
func (i *Identity) LatticeVectors() []complex128 {
	return nil
}