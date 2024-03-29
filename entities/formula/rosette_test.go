package formula_test

import (
	"github.com/chadius/creatingsymmetry/entities/formula"
	"github.com/chadius/creatingsymmetry/entities/formula/coefficient"
	"github.com/chadius/creatingsymmetry/entities/utility"
	. "gopkg.in/check.v1"
)

type RosetteTest struct{}

var _ = Suite(&RosetteTest{})

func (suite *RosetteTest) TestRosetteFormulaTransforms(checker *C) {
	rosetteFormula, _ := formula.NewBuilder().Rosette().AddTerm(
		formula.NewTermBuilder().
			Multiplier(complex(3, 0)).
			PowerN(1).
			PowerM(0).
			AddCoefficientRelationship(coefficient.PlusMPlusN).
			Build(),
	).Build()

	basePoint := complex(2, 1)
	transformedPoint := rosetteFormula.Calculate(basePoint)
	checker.Assert(real(transformedPoint), utility.NumericallyCloseEnough{}, 12, 1e-6)
	checker.Assert(imag(transformedPoint), utility.NumericallyCloseEnough{}, 0, 1e-6)
}
