package formula

// Identity formulas will transform points by returning the same coordinates.
type Identity struct {}

// Calculate TODO
func (i *Identity) Calculate(coordinate complex128) complex128 {
	return coordinate
}

// FormulaLevelTerms returns an empty list, Identity formulas do not have terms.
func (i *Identity) FormulaLevelTerms() []Term {
	return nil
}