package formula

// Identity formulas will transform points by returning the same coordinates.
type Identity struct {}

func (i *Identity) Calculate(coordinate complex128) complex128 {
	return coordinate
}