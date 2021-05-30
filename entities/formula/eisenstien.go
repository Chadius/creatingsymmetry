package formula

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"math"
	"math/cmplx"
	"wallpaper/entities/utility"
)

// EisensteinFormulaTermMarshal can be marshaled and converted to a EisensteinFormulaTerm
type EisensteinFormulaTermMarshal struct {
	PowerN					int								`json:"power_n" yaml:"power_n"`
	PowerM					int								`json:"power_m" yaml:"power_m"`
	Multiplier utility.ComplexNumberForMarshal	`json:"multiplier" yaml:"multiplier"`
}

// EisensteinFormulaTerm defines the shape of a lattice, a 2D structure that remains consistent
//    in wallpaper symmetry.
type EisensteinFormulaTerm struct {
	PowerN					int
	PowerM					int
	Multiplier 		complex128
}

// Calculate uses the Eisenstein formula on the complex number z.
// Calculate(z) = e ^ (2 PI i * (nX + mY))
//  where n amd m are PowerN and PowerM,
//  and X and Y are the real and imag parts of (zInLatticeCoordinates)
func(term EisensteinFormulaTerm)Calculate(zInLatticeCoordinates complex128) complex128 {
	powerMultiplier := (float64(term.PowerN) * real(zInLatticeCoordinates)) +
		(float64(term.PowerM) * imag(zInLatticeCoordinates))
	expo := cmplx.Exp(complex(0, 2.0 * math.Pi * powerMultiplier))
	return expo * term.Multiplier
}

// NewEisensteinFormulaTermFromJSON reads the data and returns a formula term from it.
func NewEisensteinFormulaTermFromJSON(data []byte) (*EisensteinFormulaTerm, error) {
	return newEisensteinFormulaTermFromDatastream(data, json.Unmarshal)
}

// NewEisensteinFormulaTermFromYAML reads the data and returns a formula term from it.
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
		PowerN:                 marshalObject.PowerN,
		PowerM:                 marshalObject.PowerM,
		Multiplier: complex(marshalObject.Multiplier.Real, marshalObject.Multiplier.Imaginary),
	}
}
