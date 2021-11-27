package formula_test

import (
	"github.com/Chadius/creating-symmetry/entities/formula"
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	"github.com/Chadius/creating-symmetry/entities/utility"
	. "gopkg.in/check.v1"
	"math"
)

type FriezeTest struct{}

var _ = Suite(&FriezeTest{})

func (suite *FriezeTest) TestFriezeFormulaTransforms(checker *C) {
	FriezeFormula, _ := formula.NewBuilder().Frieze().AddTerm(
		formula.NewTermBuilder().
			Multiplier(complex(2, 0)).
			PowerN(1).
			PowerM(0).
			AddCoefficientRelationship(coefficient.PlusMPlusN).
			Build(),
	).Build()

	basePoint := complex(math.Pi/6, 1)
	transformedPoint := FriezeFormula.Calculate(basePoint)

	expectedResult := complex(math.Exp(-1), 0) * complex(math.Sqrt(3)*2, 0)
	checker.Assert(real(transformedPoint), utility.NumericallyCloseEnough{}, real(expectedResult), 1e-6)
	checker.Assert(imag(transformedPoint), utility.NumericallyCloseEnough{}, imag(expectedResult), 1e-6)
}
