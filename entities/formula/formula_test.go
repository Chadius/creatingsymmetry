package formula_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
	"wallpaper/entities/formula"
)

var _ = Describe("Common formula formats", func() {
	It("Can calculate a Rosette formula", func() {
		rosetteFormula := formula.RosetteFormula{
			Elements: []*formula.ZExponentialFormulaElement{
				{
					Scale: complex(3, 0),
					PowerN: 1,
					PowerM: 0,
					IgnoreComplexConjugate: false,
					LockedCoefficientPairs: []*formula.LockedCoefficientPair{
						{
							Multiplier: 1,
							OtherCoefficientRelationships: []formula.CoefficientRelationship{
								formula.PlusMPlusN,
							},
						},
					},
				},
			},
		}
		result := rosetteFormula.Calculate(complex(2,1))
		Expect(real(result)).To(BeNumerically("~", 12))
		Expect(imag(result)).To(BeNumerically("~", 0))
	})

	Context("Terms that involve z^n * zConj^m", func() {
		It("Can make a z to the n exponential formula", func() {

			form := formula.ZExponentialFormulaElement{
				Scale:                  complex(3, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: true,
			}
			result := form.Calculate(complex(3,2))
			Expect(real(result)).To(BeNumerically("~", 15))
			Expect(imag(result)).To(BeNumerically("~", 36))
		})
		It("Can make a z to the n exponential formula using locked pairs", func() {
			form := formula.ZExponentialFormulaElement{
				Scale:                  complex(3, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: true,
				LockedCoefficientPairs: []*formula.LockedCoefficientPair{
					{
						Multiplier: -1,
						OtherCoefficientRelationships: []formula.CoefficientRelationship{
							formula.PlusMPlusN,
						},
					},
				},
			}
			result := form.Calculate(complex(3,2))
			Expect(real(result)).To(BeNumerically("~", 12))
			Expect(imag(result)).To(BeNumerically("~", 36))
		})
		It("Can make a z to the n exponential formula using a complex conjugate", func() {
			form := formula.ZExponentialFormulaElement{
				Scale:                  complex(3, 0),
				PowerN:                 2,
				PowerM:                 1,
				IgnoreComplexConjugate: false,
			}
			result := form.Calculate(complex(3,2))
			Expect(real(result)).To(BeNumerically("~", 117))
			Expect(imag(result)).To(BeNumerically("~", 78))
		})
	})

	Context("Coefficient Relationships", func() {
		It("Can keep coefficients in same order", func() {
			form := formula.ZExponentialFormulaElement{
				Scale:                  complex(1, 0),
				PowerN:                 1,
				PowerM:                 0,
				IgnoreComplexConjugate: true,
				LockedCoefficientPairs: []*formula.LockedCoefficientPair{
					{
						Multiplier: -1,
						OtherCoefficientRelationships: []formula.CoefficientRelationship{
							formula.PlusNPlusM,
						},
					},
				},
			}
			result := form.Calculate(complex(3,2))
			Expect(real(result)).To(BeNumerically("~", 0))
			Expect(imag(result)).To(BeNumerically("~", 0))
		})
		It("Can swap coefficients", func() {
			form := formula.ZExponentialFormulaElement{
				Scale:                  complex(1, 0),
				PowerN:                 1,
				PowerM:                 0,
				IgnoreComplexConjugate: true,
				LockedCoefficientPairs: []*formula.LockedCoefficientPair{
					{
						Multiplier: -1,
						OtherCoefficientRelationships: []formula.CoefficientRelationship{
							formula.PlusMPlusN,
						},
					},
				},
			}
			result := form.Calculate(complex(3,2))
			Expect(real(result)).To(BeNumerically("~", 2))
			Expect(imag(result)).To(BeNumerically("~", 2))
		})
	})

	Context("Terms that involve e^(inz) * e^(-imzConj)", func() {
		It("Can calculate a formula that uses Euler and complex numbers", func() {
			form := formula.EulerFormulaElement{
				Scale:                  complex(3, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: true,
			}
			result := form.Calculate(complex(math.Pi / 6.0,1))
			Expect(real(result)).To(BeNumerically("~", 3 * math.Exp(-2) * 1.0 / 2.0))
			Expect(imag(result)).To(BeNumerically("~", 3 * math.Exp(-2) * math.Sqrt(3.0) / 2.0))
		})
		It("Can calculate a formula that uses locked coefficient pairs", func() {
			form := formula.EulerFormulaElement{
				Scale:                  complex(3, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: true,
				LockedCoefficientPairs: []*formula.LockedCoefficientPair{
					{
						Multiplier: 1,
						OtherCoefficientRelationships: []formula.CoefficientRelationship{
							formula.PlusMPlusN,
						},
					},
				},
			}
			result := form.Calculate(complex(math.Pi / 6.0,1))
			Expect(real(result)).To(BeNumerically("~", 3 * ((math.Exp(-2) * 1.0 / 2.0) + 1.0)))
			Expect(imag(result)).To(BeNumerically("~", 3 * math.Exp(-2) * math.Sqrt(3.0) / 2.0))
		})
		It("Can calculate a formula that uses the complex conjugate", func() {
			form := formula.EulerFormulaElement{
				Scale:                  complex(3, 0),
				PowerN:                 2,
				PowerM:                 1,
				IgnoreComplexConjugate: false,
			}
			result := form.Calculate(complex(math.Pi / 6.0,2))
			Expect(real(result)).To(BeNumerically("~", 3 * math.Exp(-6) * math.Sqrt(3.0) / 2.0))
			Expect(imag(result)).To(BeNumerically("~", 3 * math.Exp(-6) * 1.0 / 2.0))
		})
	})

	It("Can calculate a Frieze formula", func() {
		friezeFormula := formula.FriezeFormula{
			Elements: []*formula.EulerFormulaElement{
				{
					Scale: complex(2, 0),
					PowerN: 1,
					PowerM: 0,
					IgnoreComplexConjugate: false,
					LockedCoefficientPairs: []*formula.LockedCoefficientPair{
						{
							Multiplier: 1,
							OtherCoefficientRelationships: []formula.CoefficientRelationship{
								formula.PlusMPlusN,
							},
						},
					},
				},
			},
		}
		result := friezeFormula.Calculate(complex(math.Pi/6, 1))

		expectedResult := complex(math.Exp(-1), 0) * complex(math.Sqrt(3) * 2, 0)
		Expect(real(result)).To(BeNumerically("~", real(expectedResult)))
		Expect(imag(result)).To(BeNumerically("~", imag(expectedResult)))
	})

	Context("Can determine coefficients and scale based on coefficient relationship", func() {
		It("PlusNPlusM reuses given coefficients", func() {
			power1, power2, scale := formula.SetCoefficientsBasedOnRelationship(1, 2, complex(3, 4), formula.PlusNPlusM)
			Expect(power1).To(Equal(1))
			Expect(power2).To(Equal(2))
			Expect(real(scale)).To(BeNumerically("~", 3))
			Expect(imag(scale)).To(BeNumerically("~", 4))
		})
		It("PlusMPlusN swaps given coefficients", func() {
			power1, power2, scale := formula.SetCoefficientsBasedOnRelationship(1, 2, complex(3, 4), formula.PlusMPlusN)
			Expect(power1).To(Equal(2))
			Expect(power2).To(Equal(1))
			Expect(real(scale)).To(BeNumerically("~", 3))
			Expect(imag(scale)).To(BeNumerically("~", 4))
		})
		It("MinusNMinusM reuses given coefficients", func() {
			power1, power2, scale := formula.SetCoefficientsBasedOnRelationship(1, 2, complex(3, 4), formula.MinusNMinusM)
			Expect(power1).To(Equal(-1))
			Expect(power2).To(Equal(-2))
			Expect(real(scale)).To(BeNumerically("~", 3))
			Expect(imag(scale)).To(BeNumerically("~", 4))
		})
		It("MinusMMinusN swaps given coefficients", func() {
			power1, power2, scale := formula.SetCoefficientsBasedOnRelationship(1, 2, complex(3, 4), formula.MinusMMinusN)
			Expect(power1).To(Equal(-2))
			Expect(power2).To(Equal(-1))
			Expect(real(scale)).To(BeNumerically("~", 3))
			Expect(imag(scale)).To(BeNumerically("~", 4))
		})
		It("PlusMPlusNMaybeFlipScale changes the scale of odd summed powers ", func() {
			power1, power2, scale := formula.SetCoefficientsBasedOnRelationship(1, 2, complex(3, 4), formula.PlusMPlusNMaybeFlipScale)
			Expect(power1).To(Equal(2))
			Expect(power2).To(Equal(1))
			Expect(real(scale)).To(BeNumerically("~", -3))
			Expect(imag(scale)).To(BeNumerically("~", -4))
		})
		It("PlusMPlusNMaybeFlipScale does not change the scale of even summed powers ", func() {
			power1, power2, scale := formula.SetCoefficientsBasedOnRelationship(0, 2, complex(3, 4), formula.PlusMPlusNMaybeFlipScale)
			Expect(power1).To(Equal(2))
			Expect(power2).To(Equal(0))
			Expect(real(scale)).To(BeNumerically("~", 3))
			Expect(imag(scale)).To(BeNumerically("~", 4))
		})
		It("MinusMMinusNMaybeFlipScale changes the scale of odd summed powers ", func() {
			power1, power2, scale := formula.SetCoefficientsBasedOnRelationship(1, 2, complex(3, 4), formula.MinusMMinusNMaybeFlipScale)
			Expect(power1).To(Equal(-2))
			Expect(power2).To(Equal(-1))
			Expect(real(scale)).To(BeNumerically("~", -3))
			Expect(imag(scale)).To(BeNumerically("~", -4))
		})
		It("MinusMMinusNMaybeFlipScale does not change the scale of even summed powers ", func() {
			power1, power2, scale := formula.SetCoefficientsBasedOnRelationship(0, 2, complex(3, 4), formula.MinusMMinusNMaybeFlipScale)
			Expect(power1).To(Equal(-2))
			Expect(power2).To(Equal(0))
			Expect(real(scale)).To(BeNumerically("~", 3))
			Expect(imag(scale)).To(BeNumerically("~", 4))
		})
	})
})