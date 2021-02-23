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
		Context("Create a function of linked powers", func() {
			It("Creates a formula with two pairs", func() {
				arbitraryFormula := formula.SymmetryFormula{
					PairedCoefficients : []*formula.CoefficientPairs{
						{
							Scale: complex(1, 0),
							PowerN: 1,
							PowerM: 0,
						},
						{
							Scale: complex(3, 0),
							PowerN: 0,
							PowerM: 1,
						},
					},
				}
				result := arbitraryFormula.Calculate(complex(2,1))
				Expect(real(result)).To(BeNumerically("~", 8))
				Expect(imag(result)).To(BeNumerically("~", -2))
			})
		})
})