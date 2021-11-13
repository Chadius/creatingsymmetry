package eisenstien

import (
	"encoding/json"
	"github.com/Chadius/creating-symmetry/entities/utility"
	"gopkg.in/yaml.v2"
	"math"
	"math/cmplx"
)

// EisensteinFormulaTermMarshal can be marshaled and converted to a EisensteinFormulaTerm
type EisensteinFormulaTermMarshal struct {
	PowerN int `json:"power_n" yaml:"power_n"`
	PowerM int `json:"power_m" yaml:"power_m"`
}

// EisensteinFormulaTerm defines the shape of a lattice, a 2D structure that remains consistent
//    in wallpaper symmetry.
type EisensteinFormulaTerm struct {
	PowerN int
	PowerM int
}

// PowerSumIsEven returns true if the sum of the term powers is divisible by 2.
func (term EisensteinFormulaTerm) PowerSumIsEven() bool {
	return (term.PowerM+term.PowerN)%2 == 0
}

// Calculate uses the Eisenstein oldformula on the complex number z.
// Calculate(z) = e ^ (2 PI i * (nX + mY))
//  where n amd m are PowerN and PowerM,
//  and TransformedX and TransformedY are the real and imag parts of (zInLatticeCoordinates)
func (term EisensteinFormulaTerm) Calculate(zInLatticeCoordinates complex128) complex128 {
	powerMultiplier := (float64(term.PowerN) * real(zInLatticeCoordinates)) +
		(float64(term.PowerM) * imag(zInLatticeCoordinates))
	expo := cmplx.Exp(complex(0, 2.0*math.Pi*powerMultiplier))
	return expo
}

// NewEisensteinFormulaTermFromJSON reads the data and returns a oldformula term from it.
func NewEisensteinFormulaTermFromJSON(data []byte) (*EisensteinFormulaTerm, error) {
	return newEisensteinFormulaTermFromDatastream(data, json.Unmarshal)
}

// NewEisensteinFormulaTermFromYAML reads the data and returns a oldformula term from it.
func NewEisensteinFormulaTermFromYAML(data []byte) (*EisensteinFormulaTerm, error) {
	return newEisensteinFormulaTermFromDatastream(data, yaml.Unmarshal)
}

//newEisensteinFormulaTermFromDatastream consumes a given bytestream and tries to create a new object from it.
func newEisensteinFormulaTermFromDatastream(data []byte, unmarshal utility.UnmarshalFunc) (*EisensteinFormulaTerm, error) {
	var unmarshalError error
	var formulaTermMarshal EisensteinFormulaTermMarshal
	unmarshalError = unmarshal(data, &formulaTermMarshal)

	if unmarshalError != nil {
		return nil, unmarshalError
	}

	formulaTerm := NewEisensteinFormulaTermFromMarshalObject(formulaTermMarshal)
	return formulaTerm, nil
}

// NewEisensteinFormulaTermFromMarshalObject converts the marshaled intermediary object into a usable object.
func NewEisensteinFormulaTermFromMarshalObject(marshalObject EisensteinFormulaTermMarshal) *EisensteinFormulaTerm {
	return &EisensteinFormulaTerm{
		PowerN: marshalObject.PowerN,
		PowerM: marshalObject.PowerM,
	}
}

