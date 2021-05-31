package wavepacket_test

import (
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/wavepacket"
	"wallpaper/entities/utility"
)

type SquareLatticeWallpaper struct {
	squareWavePacket *wavepacket.SquareWallpaperFormula
}

var _ = Suite(&SquareLatticeWallpaper{})

func (suite *SquareLatticeWallpaper) SetUpTest(checker *C) {
	suite.squareWavePacket = &wavepacket.SquareWallpaperFormula{
		Formula: &wavepacket.WallpaperFormula{
			WavePackets: []*wavepacket.Formula{
				{
					Terms:[]*formula.EisensteinFormulaTerm{
						{
							PowerN:         1,
							PowerM:         -2,
						},
					},
					Multiplier: complex(1, 0),
				},
			},
			Multiplier: complex(1, 0),
		},
	}
}

func (suite *SquareLatticeWallpaper) TestSquareLatticeWallpaperImpliesAveragedLockedTerms(checker *C) {
	suite.squareWavePacket.SetUp()
	calculation := suite.squareWavePacket.Calculate(complex(2, 0.5))
	total := calculation.Total

	expectedAnswer :=
		(
			cmplx.Exp(complex(0, 2 * math.Pi)) +
			cmplx.Exp(complex(0, 2 * math.Pi * -3.5)) +
			cmplx.Exp(complex(0, 2 * math.Pi * -1)) +
			cmplx.Exp(complex(0, 2 * math.Pi * 4.5)))/4

	checker.Assert(real(total), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(total), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}

func (suite *SquareLatticeWallpaper) TestUnmarshalFromJSON(checker *C) {
	jsonByteStream := []byte(`{
				"multiplier": {
					"real": -1.0,
					"imaginary": 2e-2
				},
				"wave_packets": [
					{
						"multiplier": {
							"real": -1.0,
							"imaginary": 2e-2
						},
						"terms": [
							{
								"power_n": 12,
								"power_m": -10
							}
						]
					}
				]
			}`)
	squareFormula, err := wavepacket.NewSquareWallpaperFormulaFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(real(squareFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(squareFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
	checker.Assert(squareFormula.Formula.WavePackets, HasLen, 1)
	checker.Assert(squareFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -10)
}

func (suite *SquareLatticeWallpaper) TestUnmarshalFromYAML(checker *C) {
	yamlByteStream := []byte(`
multiplier:
 real: -1.0
 imaginary: 2e-2
wave_packets:
 -
   multiplier:
     real: -1.0
     imaginary: 2e-2
   terms:
     -
       power_n: 12
       power_m: -10
`)
	squareFormula, err := wavepacket.NewSquareWallpaperFormulaFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(real(squareFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(squareFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
	checker.Assert(squareFormula.Formula.WavePackets, HasLen, 1)
	checker.Assert(squareFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -10)
}

type SquareWaveSymmetry struct {
	baseWavePacket *wavepacket.Formula
}

var _ = Suite(&SquareWaveSymmetry{})

func (suite *SquareWaveSymmetry) SetUpTest(checker *C) {
	suite.baseWavePacket = &wavepacket.Formula{
		Terms:[]*formula.EisensteinFormulaTerm{
			{
				PowerN:         8,
				PowerM:         -3,
			},
		},
		Multiplier: complex(1, 0),
	}
}

func (suite *SquareWaveSymmetry) TestEverySquareWallpaperHasP4Symmetry(checker *C) {
	squareFormula := wavepacket.SquareWallpaperFormula{
		Formula: &wavepacket.WallpaperFormula{
			WavePackets: []*wavepacket.Formula{
				suite.baseWavePacket,
			},
			Multiplier: complex(1, 0),
		},
	}

	squareFormula.SetUp()

	symmetriesFound := squareFormula.FindSymmetries()
	checker.Assert(symmetriesFound.P4, Equals, true)
	checker.Assert(symmetriesFound.P4m, Equals, false)
	checker.Assert(symmetriesFound.P4g, Equals, false)
}

func (suite *SquareWaveSymmetry) TestP4m(checker *C) {
	squareFormula := wavepacket.SquareWallpaperFormula{
		Formula: &wavepacket.WallpaperFormula{
			WavePackets: []*wavepacket.Formula{
				suite.baseWavePacket,
				{
					Terms: []*formula.EisensteinFormulaTerm{
						{
							PowerN:         suite.baseWavePacket.Terms[0].PowerM,
							PowerM:         suite.baseWavePacket.Terms[0].PowerN,
							Multiplier: suite.baseWavePacket.Multiplier,
						},
					},
					Multiplier: suite.baseWavePacket.Multiplier,
				},
			},
			Multiplier: complex(1, 0),
		},
	}

	squareFormula.SetUp()

	symmetriesFound := squareFormula.FindSymmetries()
	checker.Assert(symmetriesFound.P4, Equals, true)
	checker.Assert(symmetriesFound.P4m, Equals, true)
	checker.Assert(symmetriesFound.P4g, Equals, false)
}

func (suite *SquareWaveSymmetry) TestP4gWithOddSumPowers(checker *C) {
	squareFormula := wavepacket.SquareWallpaperFormula{
		Formula: &wavepacket.WallpaperFormula{
			WavePackets: []*wavepacket.Formula{
				suite.baseWavePacket,
				{
					Terms: []*formula.EisensteinFormulaTerm{
						{
							PowerN:         suite.baseWavePacket.Terms[0].PowerM,
							PowerM:         suite.baseWavePacket.Terms[0].PowerN,
							Multiplier: suite.baseWavePacket.Multiplier,
						},
					},
					Multiplier: -1.0 * suite.baseWavePacket.Multiplier,
				},
			},
			Multiplier: complex(1, 0),
		},
	}

	squareFormula.SetUp()

	symmetriesFound := squareFormula.FindSymmetries()
	checker.Assert(symmetriesFound.P4, Equals, true)
	checker.Assert(symmetriesFound.P4m, Equals, false)
	checker.Assert(symmetriesFound.P4g, Equals, true)
}

func (suite *SquareWaveSymmetry) TestP4gWithEvenSumPowers(checker *C) {
	squareFormula := wavepacket.SquareWallpaperFormula{
		Formula: &wavepacket.WallpaperFormula{
			WavePackets: []*wavepacket.Formula{
				{
					Terms: []*formula.EisensteinFormulaTerm{
						{
							PowerN: 8,
							PowerM: -4,
							Multiplier: suite.baseWavePacket.Multiplier,
						},
					},
					Multiplier: suite.baseWavePacket.Multiplier,
				},
				{
					Terms: []*formula.EisensteinFormulaTerm{
						{
							PowerN:         -4,
							PowerM:         8,
							Multiplier: suite.baseWavePacket.Multiplier,
						},
					},
					Multiplier: suite.baseWavePacket.Multiplier,
				},
			},
			Multiplier: complex(1, 0),
		},
	}

	squareFormula.SetUp()

	symmetriesFound := squareFormula.FindSymmetries()
	checker.Assert(symmetriesFound.P4, Equals, true)
	checker.Assert(symmetriesFound.P4m, Equals, true)
	checker.Assert(symmetriesFound.P4g, Equals, true)
}

type SquareCreatedWithDesiredSymmetry struct {
	singleEisensteinOddSumFormulaTerm []*formula.EisensteinFormulaTerm
	singleEisensteinEvenSumFormulaTerm []*formula.EisensteinFormulaTerm
	wallpaperMultiplier complex128
}

var _ = Suite(&SquareCreatedWithDesiredSymmetry{})

func (suite *SquareCreatedWithDesiredSymmetry) SetUpTest (checker *C) {
	suite.singleEisensteinOddSumFormulaTerm = []*formula.EisensteinFormulaTerm{
		{
			PowerM: -2,
			PowerN: 1,
			Multiplier: complex(1, 0),
		},
	}
	suite.singleEisensteinEvenSumFormulaTerm = []*formula.EisensteinFormulaTerm{
		{
			PowerM: 3,
			PowerN: -1,
			Multiplier: complex(1, 0),
		},
	}
	suite.wallpaperMultiplier = complex(1, 0)
}

func (suite *SquareCreatedWithDesiredSymmetry) SymmetryCheck(checker *C, squareFormula *wavepacket.SquareWallpaperFormula , desiredSymmetry wavepacket.Symmetry) {
	symmetriesFound := squareFormula.FindSymmetries()
	checker.Assert(symmetriesFound.P4, Equals, desiredSymmetry.P4)
	checker.Assert(symmetriesFound.P4m, Equals, desiredSymmetry.P4m)
	checker.Assert(symmetriesFound.P4g, Equals, desiredSymmetry.P4g)
}

func (suite *SquareCreatedWithDesiredSymmetry) TestNoSymmetryDoesNotChangePattern(checker *C) {
	squareFormula, err := wavepacket.NewSquareWallpaperFormulaWithSymmetry(
		suite.singleEisensteinOddSumFormulaTerm,
		suite.wallpaperMultiplier,
		&wavepacket.Symmetry{},
	)

	checker.Assert(err, IsNil)
	checker.Assert(squareFormula.Formula.WavePackets, HasLen, 1)
	checker.Assert(squareFormula.Formula.WavePackets[0].Terms, HasLen, 4)

	suite.SymmetryCheck(checker, squareFormula, wavepacket.Symmetry{
		P4:   true,
	})
}

func (suite *SquareCreatedWithDesiredSymmetry) TestCreateWallpaperWithP4m(checker *C) {
	squareFormula, err := wavepacket.NewSquareWallpaperFormulaWithSymmetry(
		suite.singleEisensteinOddSumFormulaTerm,
		suite.wallpaperMultiplier,
		&wavepacket.Symmetry{
			P4m: true,
		},
	)

	checker.Assert(err, IsNil)
	checker.Assert(squareFormula.Formula.WavePackets, HasLen, 2)
	checker.Assert(squareFormula.Formula.WavePackets[0].Terms, HasLen, 4)

	checker.Assert(squareFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, 1)
	checker.Assert(squareFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, -2)

	suite.SymmetryCheck(checker, squareFormula, wavepacket.Symmetry{
		P4:   true,
		P4m:   true,
	})
}

func (suite *SquareCreatedWithDesiredSymmetry) TestCreateWallpaperWithP4gAndOddSumPowers(checker *C) {
	squareFormula, errOddSum := wavepacket.NewSquareWallpaperFormulaWithSymmetry(
		suite.singleEisensteinOddSumFormulaTerm,
		suite.wallpaperMultiplier,
		&wavepacket.Symmetry{
			P4g: true,
		},
	)

	checker.Assert(errOddSum, IsNil)
	checker.Assert(squareFormula.Formula.WavePackets, HasLen, 2)
	checker.Assert(squareFormula.Formula.WavePackets[0].Terms, HasLen, 4)

	checker.Assert(real(squareFormula.Formula.WavePackets[1].Multiplier), utility.NumericallyCloseEnough{}, real(suite.wallpaperMultiplier) * -1, 1e-6)
	checker.Assert(imag(squareFormula.Formula.WavePackets[1].Multiplier), utility.NumericallyCloseEnough{}, imag(suite.wallpaperMultiplier) * -1, 1e-6)
	checker.Assert(squareFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, 1)
	checker.Assert(squareFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, -2)

	suite.SymmetryCheck(checker, squareFormula, wavepacket.Symmetry{
		P4:   true,
		P4g:   true,
	})
}

func (suite *SquareCreatedWithDesiredSymmetry) TestCreateWallpaperWithP4gAndEvenSumPowers(checker *C) {
	squareFormula, errEvenSum := wavepacket.NewSquareWallpaperFormulaWithSymmetry(
		suite.singleEisensteinEvenSumFormulaTerm,
		suite.wallpaperMultiplier,
		&wavepacket.Symmetry{
			P4g: true,
		},
	)

	checker.Assert(errEvenSum, IsNil)
	checker.Assert(squareFormula.Formula.WavePackets, HasLen, 2)
	checker.Assert(squareFormula.Formula.WavePackets[0].Terms, HasLen, 4)

	checker.Assert(real(squareFormula.Formula.WavePackets[1].Terms[0].Multiplier), utility.NumericallyCloseEnough{}, real(suite.wallpaperMultiplier), 1e-6)
	checker.Assert(imag(squareFormula.Formula.WavePackets[1].Terms[0].Multiplier), utility.NumericallyCloseEnough{}, imag(suite.wallpaperMultiplier), 1e-6)
	checker.Assert(squareFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, -1)
	checker.Assert(squareFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, 3)

	suite.SymmetryCheck(checker, squareFormula, wavepacket.Symmetry{
		P4:   true,
		P4m:   true,
		P4g:   true,
	})
}

func (suite *SquareCreatedWithDesiredSymmetry) TestP4mP4gOnlyCombinesForEvenSumTerms(checker *C) {
	_, errEven := wavepacket.NewSquareWallpaperFormulaWithSymmetry(
		suite.singleEisensteinEvenSumFormulaTerm,
		suite.wallpaperMultiplier,
		&wavepacket.Symmetry{
			P4: true,
			P4m: true,
			P4g: true,
		},
	)

	checker.Assert(errEven, IsNil)

	_, errIncompatible := wavepacket.NewSquareWallpaperFormulaWithSymmetry(
		suite.singleEisensteinOddSumFormulaTerm,
		suite.wallpaperMultiplier,
		&wavepacket.Symmetry{
			P4: true,
			P4m: true,
			P4g: true,
		},
	)

	checker.Assert(errIncompatible, ErrorMatches, "invalid desired symmetry")
}