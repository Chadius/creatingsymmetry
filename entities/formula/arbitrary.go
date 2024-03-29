package formula

// Arbitrary formulas have all of these methods.
type Arbitrary interface {
	Calculate(coordinate complex128) complex128
	FormulaLevelTerms() []Term
	WavePackets() []WavePacket
	LatticeVectors() []complex128
	SymmetriesFound() []Symmetry
}
