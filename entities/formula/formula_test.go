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
})