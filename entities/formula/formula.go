package formula

import (
	"gopkg.in/yaml.v2"
	"wallpaper/entities/formula/coefficient"
	"wallpaper/entities/utility"
)

// LockedCoefficientPair describes how to create a new Term based on the current one.
type LockedCoefficientPair struct {
	Multiplier                    float64                    `json:"multiplier" yaml:"multiplier"`
	OtherCoefficientRelationships []coefficient.Relationship `json:"relationships" yaml:"relationships"`
}

// NewLockedCoefficientPairFromYAML reads the data and returns a LockedCoefficientPair.
func NewLockedCoefficientPairFromYAML(data []byte) (*LockedCoefficientPair, error) {
	return newLockedCoefficientPairFromDatastream(data, yaml.Unmarshal)
}

// newLockedCoefficientPairFromDatastream consumes a given bytestream and tries to create a new object from it.
func newLockedCoefficientPairFromDatastream(data []byte, unmarshal utility.UnmarshalFunc) (*LockedCoefficientPair, error) {
	var unmarshalError error
	var lockedCoefficientPairToCreate LockedCoefficientPair
	unmarshalError = unmarshal(data, &lockedCoefficientPairToCreate)

	if unmarshalError != nil {
		return nil, unmarshalError
	}
	return &lockedCoefficientPairToCreate, nil
}

// CalculationResultForFormula shows the results of a calculation
type CalculationResultForFormula struct {
	Total                 complex128
	ContributionByTerm []complex128
}