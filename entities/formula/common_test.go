package formula_test

import (
	"github.com/Chadius/creating-symmetry/entities/formula"
	"github.com/Chadius/creating-symmetry/entities/utility"
	. "gopkg.in/check.v1"
)

type LatticeVectorSuite struct{}

var _ = Suite(&LatticeVectorSuite{})

func (suite *LatticeVectorSuite) TestVectorCannotBeZero(checker *C) {
	badLatticeFormula := []complex128{
		complex(0, 0),
		complex(0, 1),
	}
	err := formula.ValidateLatticeVectors(badLatticeFormula)
	checker.Assert(err, ErrorMatches, "lattice vectors cannot be \\(0,0\\)")
}

func (suite *LatticeVectorSuite) TestVectorsCannotBeCollinear(checker *C) {
	badLatticeFormula := []complex128{
		complex(1, 1),
		complex(-2, -2),
	}
	err := formula.ValidateLatticeVectors(badLatticeFormula)
	checker.Assert(err, ErrorMatches, "vectors cannot be collinear: (.*,.*) and (.*,.*)")
}

func (suite *LatticeVectorSuite) TestGoodLatticeVectorsAreValid(checker *C) {
	squareLatticeFormula := []complex128{
		complex(1, 0),
		complex(0, 1),
	}
	err := formula.ValidateLatticeVectors(squareLatticeFormula)
	checker.Assert(err, IsNil)
}

func (suite *LatticeVectorSuite) TestConvertToLatticeVector(checker *C) {
	squareLatticeVectors := []complex128{
		complex(1, 0),
		complex(0, 1),
	}

	latticeCoordinate := formula.ConvertToLatticeCoordinates(complex(1.0, 2.0), squareLatticeVectors)
	checker.Assert(real(latticeCoordinate), utility.NumericallyCloseEnough{}, 1.0, 1e-6)
	checker.Assert(imag(latticeCoordinate), utility.NumericallyCloseEnough{}, 2.0, 1e-6)
}

func (suite *LatticeVectorSuite) TestConvertToLatticeVectorNonPerpendicularVectors(checker *C) {
	squareLatticeVectors := []complex128{
		complex(0.5, 1),
		complex(0.5, -1),
	}

	latticeCoordinate := formula.ConvertToLatticeCoordinates(complex(0.75, -0.25), squareLatticeVectors)
	checker.Assert(real(latticeCoordinate), utility.NumericallyCloseEnough{}, 0.625, 1e-6)
	checker.Assert(imag(latticeCoordinate), utility.NumericallyCloseEnough{}, 0.875, 1e-6)
}

func (suite *LatticeVectorSuite) TestConvertToLatticeVectorEvenIfFirstVectorHasZeroRealComponent(checker *C) {
	squareLatticeFormulaWithFlippedVectors := []complex128{
		complex(0, 1),
		complex(1, 0),
	}

	latticeCoordinate := formula.ConvertToLatticeCoordinates(complex(1.0, 2.0), squareLatticeFormulaWithFlippedVectors)
	checker.Assert(real(latticeCoordinate), utility.NumericallyCloseEnough{}, 2.0, 1e-6)
	checker.Assert(imag(latticeCoordinate), utility.NumericallyCloseEnough{}, 1.0, 1e-6)
}
