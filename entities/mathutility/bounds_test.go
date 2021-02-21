package mathutility_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"wallpaper/entities/mathutility"
)

var _ = Describe("Linear interpolation between scales", func() {
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