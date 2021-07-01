package wavepacket_test

import (
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/wavepacket"
	"wallpaper/entities/utility"
)

type RhombicWallpaper struct {
	rhombicFormula *wavepacket.RhombicWallpaperFormula
}

var _ = Suite(&RhombicWallpaper{})

func (suite *RhombicWallpaper) SetUpTest (checker *C) {
	suite.rhombicFormula = &wavepacket.RhombicWallpaperFormula{
		Formula: &wavepacket.WallpaperFormula{
			WavePackets: []*wavepacket.WavePacket{
				{
					Terms: []*formula.EisensteinFormulaTerm{
						{
							PowerN: 1,
							PowerM: -2,
						},
					},
					Multiplier: complex(1, 0),
				},
			},
			Multiplier: complex(1, 0),
		},
		LatticeHeight: 1.0,
	}
	suite.rhombicFormula.SetUp()
}

func (suite *RhombicWallpaper) TestSetupCreatesLatticeVectors (checker *C) {
	checker.Assert(real(suite.rhombicFormula.Formula.Lattice.XLatticeVector), utility.NumericallyCloseEnough{}, 0.5, 1e-6)
	checker.Assert(imag(suite.rhombicFormula.Formula.Lattice.XLatticeVector), utility.NumericallyCloseEnough{}, 1.0, 1e-6)

	checker.Assert(real(suite.rhombicFormula.Formula.Lattice.YLatticeVector), utility.NumericallyCloseEnough{}, 0.5, 1e-6)
	checker.Assert(imag(suite.rhombicFormula.Formula.Lattice.YLatticeVector), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
}

func (suite *RhombicWallpaper) TestRaiseErrorIfHeightIsZero (checker *C) {
	rhombicFormulaWithNoHeight := &wavepacket.RhombicWallpaperFormula{
		Formula: &wavepacket.WallpaperFormula{
			WavePackets: []*wavepacket.WavePacket{
				{
					Terms: []*formula.EisensteinFormulaTerm{
						{
							PowerN: 1,
							PowerM: -2,
						},
					},
					Multiplier: complex(1, 0),
				},
			},
			Multiplier: complex(1, 0),
		},
		LatticeHeight: 0,
	}
	err := rhombicFormulaWithNoHeight.SetUp()

	checker.Assert(err, ErrorMatches, "vectors cannot be collinear: (.*,.*) and (.*,.*)")
}

func (suite *RhombicWallpaper) TestSetupLocksPairs (checker *C) {
	checker.Assert(suite.rhombicFormula.Formula.WavePackets[0].Terms, HasLen, 2)
	checker.Assert(suite.rhombicFormula.Formula.WavePackets[0].Terms[1].PowerN, Equals, suite.rhombicFormula.Formula.WavePackets[0].Terms[0].PowerM)
	checker.Assert(suite.rhombicFormula.Formula.WavePackets[0].Terms[1].PowerM, Equals, suite.rhombicFormula.Formula.WavePackets[0].Terms[0].PowerN)
}

func (suite *RhombicWallpaper) TestCalculationOfPoints (checker *C) {
	calculation := suite.rhombicFormula.Calculate(complex(0.75, -0.25))
	total := calculation.Total

	expectedAnswer := (
		cmplx.Exp(complex(0, math.Pi * -9 / 4)) +
		cmplx.Exp(complex(0, math.Pi * -3 / 4))) / 2

	checker.Assert(real(total), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(total), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}

func (suite *RhombicWallpaper) TestUnmarshalFromJSON(checker *C) {
	jsonByteStream := []byte(`{
				"lattice_height": 0.3,
				"formula": {
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
				}
			}`)
	rhombicFormula, err := wavepacket.NewRhombicWallpaperFormulaFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(rhombicFormula.LatticeHeight, utility.NumericallyCloseEnough{}, 0.3, 1e-6)
	checker.Assert(real(rhombicFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(rhombicFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
	checker.Assert(rhombicFormula.Formula.WavePackets, HasLen, 1)
	checker.Assert(rhombicFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -10)
}

func (suite *RhombicWallpaper) TestUnmarshalFromYAML(checker *C) {
	yamlByteStream := []byte(`
lattice_height: 0.3
formula:
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
	rhombicFormula, err := wavepacket.NewRhombicWallpaperFormulaFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(rhombicFormula.LatticeHeight, utility.NumericallyCloseEnough{}, 0.3, 1e-6)
	checker.Assert(real(rhombicFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(rhombicFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
	checker.Assert(rhombicFormula.Formula.WavePackets, HasLen, 1)
	checker.Assert(rhombicFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -10)
}

type RhombicWaveSymmetry struct {
	baseWavePacket *wavepacket.WavePacket
}

var _ = Suite(&RhombicWaveSymmetry{})

func (suite *RhombicWaveSymmetry) SetUpTest(checker *C) {
	suite.baseWavePacket = &wavepacket.WavePacket{
		Terms:[]*formula.EisensteinFormulaTerm{
			{
				PowerN:         8,
				PowerM:         -3,
			},
		},
		Multiplier: complex(1, 0),
	}
}

func (suite *RhombicWaveSymmetry) TestNoSymmetryFound(checker *C) {
	rhombicFormula := wavepacket.RhombicWallpaperFormula{
		Formula: &wavepacket.WallpaperFormula{
			WavePackets: []*wavepacket.WavePacket{
				suite.baseWavePacket,
			},
			Multiplier: complex(1, 0),
		},
	}

	rhombicFormula.SetUp()
	checker.Assert(rhombicFormula.HasSymmetry(wavepacket.Cm), Equals, false)
	checker.Assert(rhombicFormula.HasSymmetry(wavepacket.Cmm), Equals, false)
}

func (suite *RhombicWaveSymmetry) TestCm(checker *C) {
	rhombicFormula := wavepacket.RhombicWallpaperFormula{
		Formula: &wavepacket.WallpaperFormula{
			WavePackets: []*wavepacket.WavePacket{
				suite.baseWavePacket,
				{
					Terms: []*formula.EisensteinFormulaTerm{
						{
							PowerN:         suite.baseWavePacket.Terms[0].PowerM,
							PowerM:         suite.baseWavePacket.Terms[0].PowerN,
						},
					},
				},
			},
			Multiplier: complex(1, 0),
		},
	}

	rhombicFormula.SetUp()
	checker.Assert(rhombicFormula.HasSymmetry(wavepacket.Cm), Equals, true)
	checker.Assert(rhombicFormula.HasSymmetry(wavepacket.Cmm), Equals, false)
}

func (suite *RhombicWaveSymmetry) TestCmm(checker *C) {
	rhombicFormula := wavepacket.RhombicWallpaperFormula{
		Formula: &wavepacket.WallpaperFormula{
			WavePackets: []*wavepacket.WavePacket{
				suite.baseWavePacket,
				{
					Terms: []*formula.EisensteinFormulaTerm{
						{
							PowerN:         suite.baseWavePacket.Terms[0].PowerM,
							PowerM:         suite.baseWavePacket.Terms[0].PowerN,
						},
					},
				},
				{
					Terms: []*formula.EisensteinFormulaTerm{
						{
							PowerN:         suite.baseWavePacket.Terms[0].PowerM * -1,
							PowerM:         suite.baseWavePacket.Terms[0].PowerN * -1,
						},
					},
				},
				{
					Terms: []*formula.EisensteinFormulaTerm{
						{
							PowerN:         suite.baseWavePacket.Terms[0].PowerN * -1,
							PowerM:         suite.baseWavePacket.Terms[0].PowerM * -1,
						},
					},
				},
			},
			Multiplier: complex(1, 0),
		},
	}

	rhombicFormula.SetUp()
	checker.Assert(rhombicFormula.HasSymmetry(wavepacket.Cm), Equals, true)
	checker.Assert(rhombicFormula.HasSymmetry(wavepacket.Cmm), Equals, true)
}

type RhombicCreatedWithDesiredSymmetry struct {
	singleEisensteinFormulaTerm []*formula.EisensteinFormulaTerm
	wallpaperMultiplier complex128
	LatticeHeight float64
}

var _ = Suite(&RhombicCreatedWithDesiredSymmetry{})

func (suite *RhombicCreatedWithDesiredSymmetry) SetUpTest (checker *C) {
	suite.singleEisensteinFormulaTerm = []*formula.EisensteinFormulaTerm{
		{
			PowerN: 1,
			PowerM: -2,
			Multiplier: complex(1, 0),
		},
	}
	suite.wallpaperMultiplier = complex(1, 0)
	suite.LatticeHeight = 1.0
}

func (suite *RhombicCreatedWithDesiredSymmetry) TestCreateWallpaperWithCm(checker *C) {
	rhombicFormula, err := wavepacket.NewRhombicWallpaperFormulaWithSymmetry(
		suite.singleEisensteinFormulaTerm,
		suite.wallpaperMultiplier,
		suite.LatticeHeight,
		wavepacket.Cm,
	)

	checker.Assert(err, IsNil)
	checker.Assert(rhombicFormula.Formula.WavePackets, HasLen, 2)
	checker.Assert(rhombicFormula.Formula.WavePackets[0].Terms, HasLen, 2)

	checker.Assert(rhombicFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, suite.singleEisensteinFormulaTerm[0].PowerM)
	checker.Assert(rhombicFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, suite.singleEisensteinFormulaTerm[0].PowerN)

	checker.Assert(rhombicFormula.HasSymmetry(wavepacket.Cm), Equals, true)
	checker.Assert(rhombicFormula.HasSymmetry(wavepacket.Cmm), Equals, false)
}

func (suite *RhombicCreatedWithDesiredSymmetry) TestCreateWallpaperWithCmm(checker *C) {
	rhombicFormula, err := wavepacket.NewRhombicWallpaperFormulaWithSymmetry(
		suite.singleEisensteinFormulaTerm,
		suite.wallpaperMultiplier,
		suite.LatticeHeight,
		wavepacket.Cmm,
	)

	checker.Assert(err, IsNil)
	checker.Assert(rhombicFormula.Formula.WavePackets, HasLen, 4)
	checker.Assert(rhombicFormula.Formula.WavePackets[0].Terms, HasLen, 2)

	checker.Assert(rhombicFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, suite.singleEisensteinFormulaTerm[0].PowerN * -1)
	checker.Assert(rhombicFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, suite.singleEisensteinFormulaTerm[0].PowerM * -1)
	checker.Assert(rhombicFormula.Formula.WavePackets[2].Terms[0].PowerN, Equals, suite.singleEisensteinFormulaTerm[0].PowerM)
	checker.Assert(rhombicFormula.Formula.WavePackets[2].Terms[0].PowerM, Equals, suite.singleEisensteinFormulaTerm[0].PowerN)
	checker.Assert(rhombicFormula.Formula.WavePackets[3].Terms[0].PowerN, Equals, suite.singleEisensteinFormulaTerm[0].PowerM * -1)
	checker.Assert(rhombicFormula.Formula.WavePackets[3].Terms[0].PowerM, Equals, suite.singleEisensteinFormulaTerm[0].PowerN * -1)

	checker.Assert(rhombicFormula.HasSymmetry(wavepacket.Cm), Equals, true)
	checker.Assert(rhombicFormula.HasSymmetry(wavepacket.Cmm), Equals, true)
}

func (suite *RhombicCreatedWithDesiredSymmetry) TestCreateDesiredSymmetryWithYAML(checker *C) {
	yamlByteStream := []byte(`
lattice_height: 0.3
formula:
  desired_symmetry: cm
  multiplier:
    real: 1.0
    imaginary: 0
  wave_packets:
    - 
      multiplier:
        real: 1.0
        imaginary: 0
      terms:
        -
          power_n: 1
          power_m: -2
`)
	rhombicFormula, err := wavepacket.NewRhombicWallpaperFormulaFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)

	checker.Assert(rhombicFormula.Formula.WavePackets, HasLen, 2)
	checker.Assert(rhombicFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -2)

	checker.Assert(rhombicFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, -2)
	checker.Assert(rhombicFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, 1)

	checker.Assert(rhombicFormula.HasSymmetry(wavepacket.Cm), Equals, true)
	checker.Assert(rhombicFormula.HasSymmetry(wavepacket.Cmm), Equals, false)
}

func (suite *RhombicCreatedWithDesiredSymmetry) TestCreateDesiredSymmetryWithJSON(checker *C) {
	jsonByteStream := []byte(`{
				"lattice_height": 0.3,
				"formula": {
					"desired_symmetry": "cmm",
					"multiplier": {
						"real": 1.0,
						"imaginary": 0
					},
					"wave_packets": [
						{
							"multiplier": {
								"real": 1.0,
								"imaginary": 0
							},
							"terms": [
								{
									"power_n": 1,
									"power_m": -2
								}
							]
						}
					]
				}
			}`)
	rhombicFormula, err := wavepacket.NewRhombicWallpaperFormulaFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)

	checker.Assert(rhombicFormula.Formula.WavePackets, HasLen, 4)
	checker.Assert(rhombicFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -2)

	checker.Assert(rhombicFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, -1)
	checker.Assert(rhombicFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, 2)
	checker.Assert(rhombicFormula.Formula.WavePackets[2].Terms[0].PowerN, Equals, -2)
	checker.Assert(rhombicFormula.Formula.WavePackets[2].Terms[0].PowerM, Equals, 1)
	checker.Assert(rhombicFormula.Formula.WavePackets[3].Terms[0].PowerN, Equals, 2)
	checker.Assert(rhombicFormula.Formula.WavePackets[3].Terms[0].PowerM, Equals, -1)

	checker.Assert(rhombicFormula.HasSymmetry(wavepacket.Cm), Equals, true)
	checker.Assert(rhombicFormula.HasSymmetry(wavepacket.Cmm), Equals, true)
}