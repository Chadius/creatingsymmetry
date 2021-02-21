package mathutility_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/cmplx"
	"wallpaper/entities/mathutility"
)

var _ = Describe("Testing minimum/maximum boundaries", func() {
	Context("Scaling between two ranges, proportionally", func() {
		It("gives min value if value is less than old range min", func() {
			Expect(mathutility.ScaleValueBetweenTwoRanges(
				-200,
				-100.0,
				100.0,
				0,
				200.0)).To(
				BeNumerically("~", 0))
		})
		It("gives min of new range if value is equal to old range min", func() {
			Expect(mathutility.ScaleValueBetweenTwoRanges(
				-100.0,
				-100.0,
				100.0,
				0,
				200.0)).To(
				BeNumerically("~", 0))
		})
		It("gives max of new range if value is greater to old range max", func() {
			Expect(mathutility.ScaleValueBetweenTwoRanges(
				9001,
				-100.0,
				100.0,
				0,
				200.0)).To(
				BeNumerically("~", 200))
		})
		It("gives max of new range if value is equal to old range max", func() {
			Expect(mathutility.ScaleValueBetweenTwoRanges(
				100,
				-100.0,
				100.0,
				0,
				200.0)).To(
				BeNumerically("~", 200))
		})
		It("gives proportionally scaled value of new range", func() {
			Expect(mathutility.ScaleValueBetweenTwoRanges(
				0,
				-100.0,
				100.0,
				0,
				200.0)).To(
				BeNumerically("~", 100))
		})
	})

	Context("Bounding box calculation", func() {
		It("Calculating bounding box for dataset", func() {
			lotsOfComplexNumbers := []complex128{
				complex(0, 0),
				complex(10, 0),
				complex(0, -100),
				complex(0, -100.2),
				complex(0, 0),
				complex(-10, 0),
				complex(0, 25),
				complex(-100, 0),
				complex(0, 25.5),
				complex(9000.1, 0),
			}

			min, max := mathutility.GetBoundingBox(lotsOfComplexNumbers)

			Expect(real(min)).To(BeNumerically("~", -100))
			Expect(imag(min)).To(BeNumerically("~", -100.2))
			Expect(real(max)).To(BeNumerically("~", 9000.1))
			Expect(imag(max)).To(BeNumerically("~", 25.5))
		})
		It("Bounding box is origin if there are no numbers", func() {
			min, max := mathutility.GetBoundingBox([]complex128{})

			Expect(real(min)).To(BeNumerically("~", 0))
			Expect(imag(min)).To(BeNumerically("~", 0))
			Expect(real(max)).To(BeNumerically("~", 0))
			Expect(imag(max)).To(BeNumerically("~", 0))
		})
		It("Bounding box ignores infinity values", func() {
			lotsOfComplexNumbers := []complex128{
				complex(-10, -10),
				complex(10, 10),
				cmplx.Inf(),
				-1 * cmplx.Inf(),
			}
			min, max := mathutility.GetBoundingBox(lotsOfComplexNumbers)

			Expect(real(min)).To(BeNumerically("~", -10))
			Expect(imag(min)).To(BeNumerically("~", -10))
			Expect(real(max)).To(BeNumerically("~", 10))
			Expect(imag(max)).To(BeNumerically("~", 10))
		})
	})
})