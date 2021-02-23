package formula_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"wallpaper/entities/formula"
)

var _ = Describe("Common formula formats", func() {
		It("Can calculate z^n * ~z^m functions", func() {
			z := complex(2,1)
			calculated := formula.CalculateExponentPairOnNumberAndConjugate(z, 1, 0)
			Expect(real(calculated)).To(BeNumerically("~", 2))
			Expect(imag(calculated)).To(BeNumerically("~", 1))
		})
		It("Knows how to pair coefficients", func() {
			swapOrderFormula := formula.RecipeFormula{
				Coefficients: []*formula.CoefficientPairs{
					{
						Scale: complex(3, 0),
						PowerN: 1,
						PowerM: 0,
					},
				},
				Relationships: []formula.CoefficientRelationship{
					formula.PlusNPlusM,
					formula.PlusMPlusN,
				},
			}
			result := swapOrderFormula.Calculate(complex(2,1))
			Expect(real(result)).To(BeNumerically("~", 12))
			Expect(imag(result)).To(BeNumerically("~", 0))
		})
})