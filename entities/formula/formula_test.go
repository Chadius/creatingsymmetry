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
			Terms: []*formula.ZExponentialFormulaTerm{
				{
					Scale: complex(3, 0),
					PowerN: 1,
					PowerM: 0,
					IgnoreComplexConjugate: false,
					CoefficientPairs: formula.LockedCoefficientPair{
						Multiplier: 1,
						OtherCoefficientRelationships: []formula.CoefficientRelationship{
							formula.PlusMPlusN,
						},
					},
				},
			},
		}
		result := rosetteFormula.Calculate(complex(2,1))
		total := result.Total
		Expect(real(total)).To(BeNumerically("~", 12))
		Expect(imag(total)).To(BeNumerically("~", 0))
	})
	Context("Analyze Rosettes for symmetry", func() {
		It("Can determine there is multifold symmetry with 1 term", func() {
			rosetteFormula := formula.RosetteFormula{
				Terms: []*formula.ZExponentialFormulaTerm{
					{
						Scale: complex(1, 0),
						PowerN: 6,
						PowerM: 0,
						IgnoreComplexConjugate: false,
						CoefficientPairs: formula.LockedCoefficientPair{
							Multiplier: 1,
							OtherCoefficientRelationships: []formula.CoefficientRelationship{
								formula.PlusMPlusN,
							},
						},
					},
				},
			}
			symmetriesDetected := rosetteFormula.AnalyzeForSymmetry()
			Expect(symmetriesDetected.Multifold).To(Equal(6))
		})
		It("Multifold symmetry is always a positive value", func() {
			rosetteFormula := formula.RosetteFormula{
				Terms: []*formula.ZExponentialFormulaTerm{
					{
						Scale: complex(1, 0),
						PowerN: -6,
						PowerM: 0,
						IgnoreComplexConjugate: false,
						CoefficientPairs: formula.LockedCoefficientPair{
							Multiplier: 1,
							OtherCoefficientRelationships: []formula.CoefficientRelationship{
								formula.PlusMPlusN,
							},
						},
					},
				},
			}
			symmetriesDetected := rosetteFormula.AnalyzeForSymmetry()
			Expect(symmetriesDetected.Multifold).To(Equal(6))
		})
		It("Multifold symmetry uses the greatest common denominator of all elements", func() {
			rosetteFormula := formula.RosetteFormula{
				Terms: []*formula.ZExponentialFormulaTerm{
					{
						Scale: complex(1, 0),
						PowerN: -6,
						PowerM: 0,
						IgnoreComplexConjugate: false,
						CoefficientPairs: formula.LockedCoefficientPair{
							Multiplier: 1,
							OtherCoefficientRelationships: []formula.CoefficientRelationship{
								formula.PlusMPlusN,
							},
						},
					},
					{
						Scale: complex(1, 0),
						PowerN: -8,
						PowerM: 4,
						IgnoreComplexConjugate: false,
						CoefficientPairs: formula.LockedCoefficientPair{
							Multiplier: 1,
							OtherCoefficientRelationships: []formula.CoefficientRelationship{
								formula.PlusMPlusN,
							},
						},
					},
					{
						Scale: complex(1, 0),
						PowerN: 2,
						PowerM: 0,
						IgnoreComplexConjugate: false,
						CoefficientPairs: formula.LockedCoefficientPair{
							Multiplier: 1,
							OtherCoefficientRelationships: []formula.CoefficientRelationship{
								formula.PlusMPlusN,
							},
						},
					},
				},
			}
			symmetriesDetected := rosetteFormula.AnalyzeForSymmetry()
			Expect(symmetriesDetected.Multifold).To(Equal(2))
		})
	})
	It("Can determine the contribution by each term of a Rosette formula", func() {
		rosetteFormula := formula.RosetteFormula{
			Terms: []*formula.ZExponentialFormulaTerm{
				{
					Scale: complex(3, 0),
					PowerN: 1,
					PowerM: 0,
					IgnoreComplexConjugate: false,
					CoefficientPairs: formula.LockedCoefficientPair{
						Multiplier: 1,
						OtherCoefficientRelationships: []formula.CoefficientRelationship{
							formula.PlusMPlusN,
						},
					},
				},
			},
		}
		result := rosetteFormula.Calculate(complex(2,1))
		Expect(result.ContributionByTerm).To(HaveLen(1))
		contributionByFirstTerm := result.ContributionByTerm[0]
		Expect(real(contributionByFirstTerm)).To(BeNumerically("~", 12))
		Expect(imag(contributionByFirstTerm)).To(BeNumerically("~", 0))
	})

	Context("Terms that involve z^n * zConj^m", func() {
		It("Can make a z to the n exponential formula", func() {

			form := formula.ZExponentialFormulaTerm{
				Scale:                  complex(3, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: true,
			}
			total := form.Calculate(complex(3,2))
			Expect(real(total)).To(BeNumerically("~", 15))
			Expect(imag(total)).To(BeNumerically("~", 36))
		})
		It("Can make a z to the n exponential formula using locked pairs", func() {
			form := formula.ZExponentialFormulaTerm{
				Scale:                  complex(3, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: true,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: -1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.PlusMPlusN,
					},
				},
			}
			total := form.Calculate(complex(3,2))
			Expect(real(total)).To(BeNumerically("~", 12))
			Expect(imag(total)).To(BeNumerically("~", 36))
		})
		It("Can make a z to the n exponential formula using a complex conjugate", func() {
			form := formula.ZExponentialFormulaTerm{
				Scale:                  complex(3, 0),
				PowerN:                 2,
				PowerM:                 1,
				IgnoreComplexConjugate: false,
			}
			total := form.Calculate(complex(3,2))
			Expect(real(total)).To(BeNumerically("~", 117))
			Expect(imag(total)).To(BeNumerically("~", 78))
		})
	})

	Context("Coefficient Relationships", func() {
		It("Can keep coefficients in same order", func() {
			form := formula.ZExponentialFormulaTerm{
				Scale:                  complex(1, 0),
				PowerN:                 1,
				PowerM:                 0,
				IgnoreComplexConjugate: true,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: -1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.PlusNPlusM,
					},
				},
			}
			total := form.Calculate(complex(3,2))
			Expect(real(total)).To(BeNumerically("~", 0))
			Expect(imag(total)).To(BeNumerically("~", 0))
		})
		It("Can swap coefficients", func() {
			form := formula.ZExponentialFormulaTerm{
				Scale:                  complex(1, 0),
				PowerN:                 1,
				PowerM:                 0,
				IgnoreComplexConjugate: true,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: -1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.PlusMPlusN,
					},
				},
			}
			total := form.Calculate(complex(3,2))
			Expect(real(total)).To(BeNumerically("~", 2))
			Expect(imag(total)).To(BeNumerically("~", 2))
		})
	})

	Context("Terms that involve e^(inz) * e^(-imzConj)", func() {
		It("Can calculate a formula that uses Euler and complex numbers", func() {
			form := formula.EulerFormulaTerm{
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
			form := formula.EulerFormulaTerm{
				Scale:                  complex(3, 0),
				PowerN:                 2,
				PowerM:                 0,
				IgnoreComplexConjugate: true,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: 1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.PlusMPlusN,
					},
				},
			}
			result := form.Calculate(complex(math.Pi / 6.0,1))
			Expect(real(result)).To(BeNumerically("~", 3 * ((math.Exp(-2) * 1.0 / 2.0) + 1.0)))
			Expect(imag(result)).To(BeNumerically("~", 3 * math.Exp(-2) * math.Sqrt(3.0) / 2.0))
		})
		It("Can calculate a formula that uses the complex conjugate", func() {
			form := formula.EulerFormulaTerm{
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
			Terms: []*formula.EulerFormulaTerm{
				{
					Scale: complex(2, 0),
					PowerN: 1,
					PowerM: 0,
					IgnoreComplexConjugate: false,
					CoefficientPairs: formula.LockedCoefficientPair{
						Multiplier: 1,
						OtherCoefficientRelationships: []formula.CoefficientRelationship{
							formula.PlusMPlusN,
						},
					},
				},
			},
		}
		result := friezeFormula.Calculate(complex(math.Pi/6, 1))
		total := result.Total

		expectedResult := complex(math.Exp(-1), 0) * complex(math.Sqrt(3) * 2, 0)
		Expect(real(total)).To(BeNumerically("~", real(expectedResult)))
		Expect(imag(total)).To(BeNumerically("~", imag(expectedResult)))
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

	Context("Analyze Friezes for symmetry", func() {
		It("Knows when a pattern is p211", func() {
			friezeFormula := formula.FriezeFormula{
				Terms: []*formula.EulerFormulaTerm{
					{
						Scale: complex(1, 0),
						PowerN: 2,
						PowerM: 0,
						IgnoreComplexConjugate: false,
						CoefficientPairs: formula.LockedCoefficientPair{
							Multiplier: 1,
							OtherCoefficientRelationships: []formula.CoefficientRelationship{
								formula.MinusNMinusM,
							},
						},
					},
				},
			}
			symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
			Expect(symmetriesDetected.P211).To(BeTrue())
		})
		It("Knows when a pattern is p1m1", func() {
			friezeFormula := formula.FriezeFormula{
				Terms: []*formula.EulerFormulaTerm{
					{
						Scale: complex(1, 0),
						PowerN: 2,
						PowerM: 0,
						IgnoreComplexConjugate: false,
						CoefficientPairs: formula.LockedCoefficientPair{
							Multiplier: 1,
							OtherCoefficientRelationships: []formula.CoefficientRelationship{
								formula.PlusMPlusN,
							},
						},
					},
				},
			}
			symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
			Expect(symmetriesDetected.P1m1).To(BeTrue())
		})
		It("Knows when a pattern is p11m", func() {
			friezeFormula := formula.FriezeFormula{
				Terms: []*formula.EulerFormulaTerm{
					{
						Scale: complex(1, 0),
						PowerN: 2,
						PowerM: 0,
						IgnoreComplexConjugate: false,
						CoefficientPairs: formula.LockedCoefficientPair{
							Multiplier: 1,
							OtherCoefficientRelationships: []formula.CoefficientRelationship{
								formula.MinusMMinusN,
							},
						},
					},
				},
			}
			symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
			Expect(symmetriesDetected.P11m).To(BeTrue())
		})
		It("Knows when a pattern is p11g", func() {
			friezeFormula := formula.FriezeFormula{
				Terms: []*formula.EulerFormulaTerm{
					{
						Scale: complex(1, 0),
						PowerN: 2,
						PowerM: 1,
						IgnoreComplexConjugate: false,
						CoefficientPairs: formula.LockedCoefficientPair{
							Multiplier: 1,
							OtherCoefficientRelationships: []formula.CoefficientRelationship{
								formula.MinusMMinusNMaybeFlipScale,
							},
						},
					},
				},
			}
			symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
			Expect(symmetriesDetected.P11g).To(BeTrue())
		})
		It("Knows when a pattern is p11m if a p11g pattern has even sum powers", func() {
			friezeFormula := formula.FriezeFormula{
				Terms: []*formula.EulerFormulaTerm{
					{
						Scale: complex(1, 0),
						PowerN: 2,
						PowerM: 0,
						IgnoreComplexConjugate: false,
						CoefficientPairs: formula.LockedCoefficientPair{
							Multiplier: 1,
							OtherCoefficientRelationships: []formula.CoefficientRelationship{
								formula.MinusMMinusNMaybeFlipScale,
							},
						},
					},
				},
			}
			symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
			Expect(symmetriesDetected.P11m).To(BeTrue())
		})
		It("Knows when a pattern is p2mm", func() {
			friezeFormula := formula.FriezeFormula{
				Terms: []*formula.EulerFormulaTerm{
					{
						Scale: complex(1, 0),
						PowerN: 2,
						PowerM: 0,
						IgnoreComplexConjugate: false,
						CoefficientPairs: formula.LockedCoefficientPair{
							Multiplier: 1,
							OtherCoefficientRelationships: []formula.CoefficientRelationship{
								formula.MinusNMinusM,
								formula.PlusMPlusN,
								formula.MinusMMinusN,
							},
						},
					},
				},
			}
			symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
			Expect(symmetriesDetected.P2mm).To(BeTrue())
		})
		It("Knows when a pattern is p2mg", func() {
			friezeFormula := formula.FriezeFormula{
				Terms: []*formula.EulerFormulaTerm{
					{
						Scale: complex(1, 0),
						PowerN: 2,
						PowerM: -1,
						IgnoreComplexConjugate: false,
						CoefficientPairs: formula.LockedCoefficientPair{
							Multiplier: 1,
							OtherCoefficientRelationships: []formula.CoefficientRelationship{
								formula.MinusNMinusM,
								formula.PlusMPlusNMaybeFlipScale,
								formula.MinusMMinusNMaybeFlipScale,
							},
						},
					},
				},
			}
			symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
			Expect(symmetriesDetected.P2mg).To(BeTrue())
		})
		It("Knows when a pattern is p2mm if a p2mg pattern has even sum powers", func() {
			friezeFormula := formula.FriezeFormula{
				Terms: []*formula.EulerFormulaTerm{
					{
						Scale: complex(1, 0),
						PowerN: 2,
						PowerM: 0,
						IgnoreComplexConjugate: false,
						CoefficientPairs: formula.LockedCoefficientPair{
							Multiplier: 1,
							OtherCoefficientRelationships: []formula.CoefficientRelationship{
								formula.MinusNMinusM,
								formula.PlusMPlusNMaybeFlipScale,
								formula.MinusMMinusNMaybeFlipScale,
							},
						},
					},
				},
			}
			symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
			Expect(symmetriesDetected.P2mm).To(BeTrue())
		})
		It("Knows when a pattern is p111", func() {
			friezeFormula := formula.FriezeFormula{
				Terms: []*formula.EulerFormulaTerm{
					{
						Scale: complex(1, 0),
						PowerN: 2,
						PowerM: 0,
						IgnoreComplexConjugate: false,
						CoefficientPairs: formula.LockedCoefficientPair{},
					},
				},
			}
			symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
			Expect(symmetriesDetected.P111).To(BeTrue())
		})
		It("Knows when a pattern is p111 because complex coordinates are ignored", func() {
			friezeFormula := formula.FriezeFormula{
				Terms: []*formula.EulerFormulaTerm{
					{
						Scale: complex(1, 0),
						PowerN: 2,
						PowerM: 0,
						IgnoreComplexConjugate: true,
						CoefficientPairs: formula.LockedCoefficientPair{
							Multiplier: 1,
							OtherCoefficientRelationships: []formula.CoefficientRelationship{
								formula.MinusNMinusM,
							},
						},
					},
				},
			}
			symmetriesDetected := friezeFormula.AnalyzeForSymmetry()
			Expect(symmetriesDetected.P111).To(BeTrue())
			Expect(symmetriesDetected.P211).To(BeFalse())
		})
	})
	It("Can determine the contribution by each term of a Frieze formula", func() {
		friezeFormula := formula.FriezeFormula{
			Terms: []*formula.EulerFormulaTerm{
				{
					Scale: complex(2, 0),
					PowerN: 1,
					PowerM: 0,
					IgnoreComplexConjugate: false,
					CoefficientPairs: formula.LockedCoefficientPair{
						Multiplier: 1,
						OtherCoefficientRelationships: []formula.CoefficientRelationship{
							formula.PlusMPlusN,
						},
					},
				},
			},
		}
		result := friezeFormula.Calculate(complex(math.Pi/6, 1))

		Expect(result.ContributionByTerm).To(HaveLen(1))
		contributionByFirstTerm := result.ContributionByTerm[0]

		expectedResult := complex(math.Exp(-1), 0) * complex(math.Sqrt(3) * 2, 0)
		Expect(real(contributionByFirstTerm)).To(BeNumerically("~", real(expectedResult)))
		Expect(imag(contributionByFirstTerm)).To(BeNumerically("~", imag(expectedResult)))
	})
})