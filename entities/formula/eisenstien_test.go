package formula_test

import (
	. "gopkg.in/check.v1"
	"testing"
	"wallpaper/entities/formula"
	"wallpaper/entities/utility"
)

func Test(t *testing.T) { TestingT(t) }

type EisensteinFormulaSuite struct {}

var _ = Suite(&EisensteinFormulaSuite{})

func (suite *EisensteinFormulaSuite) SetUpTest(checker *C) {
}

func (suite *EisensteinFormulaSuite) TestCalculateEisensteinTermForGivenPoint(checker *C) {
	squareLatticePair := formula.LatticeVectorPair{
		XLatticeVector: complex(1, 0),
		YLatticeVector: complex(0, 1),
	}

	latticeCoordinate := squareLatticePair.ConvertToLatticeCoordinates(complex(1.0,1.5))

	eisenstein := formula.EisensteinFormulaTerm{
		PowerN: 1,
		PowerM: 1,
		Multiplier: complex(0.5,0),
	}
	calculatedCoordinate := eisenstein.Calculate(latticeCoordinate)

	checker.Assert(real(calculatedCoordinate), utility.NumericallyCloseEnough{}, -0.5, 1e-6)
	checker.Assert(imag(calculatedCoordinate), utility.NumericallyCloseEnough{}, 0.0, 1e-6)
}

func (suite *EisensteinFormulaSuite) TestCreateFormulaWithJSON(checker *C) {
	jsonByteStream := []byte(`{
				"power_n": 12,
				"power_m": -10,
				"multiplier": {
					"real": -1.0,
					"imaginary": 2e-2
				}
			}`)
	term, err := formula.NewEisensteinFormulaTermFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(term.PowerN, Equals, 12)
	checker.Assert(term.PowerM, Equals, -10)
	checker.Assert(real(term.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(term.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
}

func (suite *EisensteinFormulaSuite) TestCreateFormulaWithYAML(checker *C) {
	yamlByteStream := []byte(`
power_n: 12
power_m: -10
multiplier:
  real: -1.0
  imaginary: 2e-2
`)
	term, err := formula.NewEisensteinFormulaTermFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(term.PowerN, Equals, 12)
	checker.Assert(term.PowerM, Equals, -10)
	checker.Assert(real(term.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(term.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
}