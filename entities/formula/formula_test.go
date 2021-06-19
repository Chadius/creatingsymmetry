package formula_test

import (
	. "gopkg.in/check.v1"
	"wallpaper/entities/formula"
	"wallpaper/entities/utility"
)

type LatticeVectorSuite struct {}

var _ = Suite(&LatticeVectorSuite{})

func (suite *LatticeVectorSuite) TestVectorCannotBeZero(checker *C) {
	badLatticeFormula := formula.LatticeVectorPair{
		XLatticeVector: complex(0, 0),
		YLatticeVector: complex(0, 1),
	}
	err := badLatticeFormula.Validate()
	checker.Assert(err, ErrorMatches, "lattice vectors cannot be \\(0,0\\)")
}

func (suite *LatticeVectorSuite) TestVectorsCannotBeCollinear(checker *C) {
	badLatticeFormula := formula.LatticeVectorPair{
		XLatticeVector: complex(1, 1),
		YLatticeVector: complex(-2, -2),
	}
	err := badLatticeFormula.Validate()
	checker.Assert(err, ErrorMatches, "vectors cannot be collinear: (.*,.*) and (.*,.*)")
}

func (suite *LatticeVectorSuite) TestGoodLatticeVectorsAreValid(checker *C) {
	squareLatticeFormula := formula.LatticeVectorPair{
		XLatticeVector: complex(1, 0),
		YLatticeVector: complex(0, 1),
	}
	err := squareLatticeFormula.Validate()
	checker.Assert(err, IsNil)
}

func (suite *LatticeVectorSuite) TestConvertToLatticeVector(checker *C) {
	squareLatticeFormula := formula.LatticeVectorPair{
		XLatticeVector: complex(1, 0),
		YLatticeVector: complex(0, 1),
	}

	latticeCoordinate := squareLatticeFormula.ConvertToLatticeCoordinates(complex(1.0,2.0))
	checker.Assert(real(latticeCoordinate), utility.NumericallyCloseEnough{}, 1.0, 1e-6)
	checker.Assert(imag(latticeCoordinate), utility.NumericallyCloseEnough{}, 2.0, 1e-6)
}

func (suite *LatticeVectorSuite) TestConvertToLatticeVectorNonPerpendicularVectors(checker *C) {
	squareLatticeFormula := formula.LatticeVectorPair{
		XLatticeVector: complex(0.5, 1),
		YLatticeVector: complex(0.5, -1),
	}

	latticeCoordinate := squareLatticeFormula.ConvertToLatticeCoordinates(complex(0.75, -0.25))
	checker.Assert(real(latticeCoordinate), utility.NumericallyCloseEnough{}, 0.625, 1e-6)
	checker.Assert(imag(latticeCoordinate), utility.NumericallyCloseEnough{}, 0.875, 1e-6)
}

func (suite *LatticeVectorSuite) TestConvertToLatticeVectorEvenIfFirstVectorHasZeroRealComponent(checker *C) {
	squareLatticeFormulaWithFlippedVectors := formula.LatticeVectorPair{
		XLatticeVector: complex(0, 1),
		YLatticeVector: complex(1, 0),
	}

	latticeCoordinate := squareLatticeFormulaWithFlippedVectors.ConvertToLatticeCoordinates(complex(1.0,2.0))
	checker.Assert(real(latticeCoordinate), utility.NumericallyCloseEnough{}, 2.0, 1e-6)
	checker.Assert(imag(latticeCoordinate), utility.NumericallyCloseEnough{}, 1.0, 1e-6)
}
