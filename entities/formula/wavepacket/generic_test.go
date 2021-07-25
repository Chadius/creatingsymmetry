package wavepacket_test

import (
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/wavepacket"
	"wallpaper/entities/utility"
)

type GenericWallpaper struct {
	GenericFormula *wavepacket.GenericWallpaperFormula
}

var _ = Suite(&GenericWallpaper{})

func (suite *GenericWallpaper) SetUpTest (checker *C) {
	suite.GenericFormula = &wavepacket.GenericWallpaperFormula{
		Formula: &wavepacket.WallpaperFormula{
			WavePackets: []*wavepacket.WavePacket{
				{
					Terms: []*formula.EisensteinFormulaTerm{
						{
							PowerN: 3,
							PowerM: -4,
						},
					},
					Multiplier: complex(1, 0),
				},
			},
			Multiplier: complex(1, 0),
		},
		VectorWidth: 2,
		VectorHeight: -0.5,
	}
	suite.GenericFormula.SetUp()
}

func (suite *GenericWallpaper) TestSetupCreatesLatticeVectors (checker *C) {
	checker.Assert(real(suite.GenericFormula.Formula.Lattice.XLatticeVector), utility.NumericallyCloseEnough{}, 1, 1e-6)
	checker.Assert(imag(suite.GenericFormula.Formula.Lattice.XLatticeVector), utility.NumericallyCloseEnough{}, 0, 1e-6)

	checker.Assert(real(suite.GenericFormula.Formula.Lattice.YLatticeVector), utility.NumericallyCloseEnough{}, 2, 1e-6)
	checker.Assert(imag(suite.GenericFormula.Formula.Lattice.YLatticeVector), utility.NumericallyCloseEnough{}, -0.5, 1e-6)
}

func (suite *GenericWallpaper) TestRaiseErrorIfHeightIsZero (checker *C) {
	GenericFormulaWithNoHeight := &wavepacket.GenericWallpaperFormula{
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
		VectorWidth: -1,
		VectorHeight: 0,
	}
	err := GenericFormulaWithNoHeight.SetUp()

	checker.Assert(err, ErrorMatches, "vectors cannot be collinear: (.*,.*) and (.*,.*)")
}

func (suite *GenericWallpaper) TestSetupDoesNotAddLockedPairs (checker *C) {
	checker.Assert(suite.GenericFormula.Formula.WavePackets[0].Terms, HasLen, 1)
}

func (suite *GenericWallpaper) TestCalculationOfPoints (checker *C) {
	calculation := suite.GenericFormula.Calculate(complex(1.5, 10))
	total := calculation.Total

	expectedAnswer := cmplx.Exp(complex(0, math.Pi ))

	checker.Assert(real(total), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(total), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}

func (suite *GenericWallpaper) TestUnmarshalFromJSON(checker *C) {
	jsonByteStream := []byte(`{
				"vector_width": 0.8,
				"vector_height": 0.3,
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
	GenericFormula, err := wavepacket.NewGenericWallpaperFormulaFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(GenericFormula.VectorWidth, utility.NumericallyCloseEnough{}, 0.8, 1e-6)
	checker.Assert(GenericFormula.VectorHeight, utility.NumericallyCloseEnough{}, 0.3, 1e-6)
	checker.Assert(real(GenericFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(GenericFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
	checker.Assert(GenericFormula.Formula.WavePackets, HasLen, 1)
	checker.Assert(GenericFormula.Formula.WavePackets[0].Terms[0].PowerN, Equals, 12)
	checker.Assert(GenericFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -10)
}

func (suite *GenericWallpaper) TestUnmarshalFromYAML(checker *C) {
	yamlByteStream := []byte(`
vector_width: 0.8
vector_height: 0.3
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
	GenericFormula, err := wavepacket.NewGenericWallpaperFormulaFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(GenericFormula.VectorWidth, utility.NumericallyCloseEnough{}, 0.8, 1e-6)
	checker.Assert(GenericFormula.VectorHeight, utility.NumericallyCloseEnough{}, 0.3, 1e-6)
	checker.Assert(real(GenericFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(GenericFormula.Formula.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
	checker.Assert(GenericFormula.Formula.WavePackets, HasLen, 1)
	checker.Assert(GenericFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -10)
}

type GenericWaveSymmetry struct {
	baseWavePacket *wavepacket.WavePacket
}

var _ = Suite(&GenericWaveSymmetry{})

func (suite *GenericWaveSymmetry) SetUpTest(checker *C) {
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

func (suite *GenericWaveSymmetry) TestOnlyP1SymmetryFound(checker *C) {
	GenericFormula := wavepacket.GenericWallpaperFormula{
		Formula: &wavepacket.WallpaperFormula{
			WavePackets: []*wavepacket.WavePacket{
				suite.baseWavePacket,
			},
			Multiplier: complex(1, 0),
		},
	}

	GenericFormula.SetUp()
	checker.Assert(GenericFormula.HasSymmetry(wavepacket.P1), Equals, true)
	checker.Assert(GenericFormula.HasSymmetry(wavepacket.P2), Equals, false)
}

func (suite *GenericWaveSymmetry) TestP2(checker *C) {
	GenericFormula := wavepacket.GenericWallpaperFormula{
		Formula: &wavepacket.WallpaperFormula{
			WavePackets: []*wavepacket.WavePacket{
				suite.baseWavePacket,
				{
					Terms: []*formula.EisensteinFormulaTerm{
						{
							PowerN:         suite.baseWavePacket.Terms[0].PowerN * -1,
							PowerM:         suite.baseWavePacket.Terms[0].PowerM * -1,
						},
					},
					Multiplier: suite.baseWavePacket.Multiplier,
				},
			},
			Multiplier: complex(1, 0),
		},
	}

	GenericFormula.SetUp()
	checker.Assert(GenericFormula.HasSymmetry(wavepacket.P1), Equals, true)
	checker.Assert(GenericFormula.HasSymmetry(wavepacket.P2), Equals, true)
}

type GenericCreatedWithDesiredSymmetry struct {
	eisensteinTerm []*formula.EisensteinFormulaTerm
	wallpaperMultiplier complex128
	VectorWidth float64
	VectorHeight float64
}

var _ = Suite(&GenericCreatedWithDesiredSymmetry{})

func (suite *GenericCreatedWithDesiredSymmetry) SetUpTest (checker *C) {
	suite.eisensteinTerm = []*formula.EisensteinFormulaTerm{
		{
			PowerN:         8,
			PowerM:         -3,
		},
	}

	suite.wallpaperMultiplier = complex(1, 0)
	suite.VectorWidth = 2.0
	suite.VectorHeight = 1.5
}

func (suite *GenericCreatedWithDesiredSymmetry) TestCreateWallpaperWithP1(checker *C) {
	GenericFormula, err := wavepacket.NewGenericWallpaperFormulaWithSymmetry(
		suite.eisensteinTerm,
		suite.wallpaperMultiplier,
		suite.VectorWidth,
		suite.VectorHeight,
		wavepacket.P1,
	)

	checker.Assert(err, IsNil)

	checker.Assert(GenericFormula.VectorWidth, Equals, suite.VectorWidth)
	checker.Assert(GenericFormula.VectorHeight, Equals, suite.VectorHeight)

	checker.Assert(GenericFormula.Formula.WavePackets, HasLen, 1)
	checker.Assert(GenericFormula.Formula.WavePackets[0].Terms, HasLen, 1)

	checker.Assert(GenericFormula.HasSymmetry(wavepacket.P1), Equals, true)
	checker.Assert(GenericFormula.HasSymmetry(wavepacket.P2), Equals, false)
}

func (suite *GenericCreatedWithDesiredSymmetry) TestCreateWallpaperWithP2(checker *C) {
	GenericFormula, err := wavepacket.NewGenericWallpaperFormulaWithSymmetry(
		suite.eisensteinTerm,
		suite.wallpaperMultiplier,
		suite.VectorWidth,
		suite.VectorHeight,
		wavepacket.P2,
	)

	checker.Assert(err, IsNil)

	checker.Assert(GenericFormula.VectorWidth, Equals, suite.VectorWidth)
	checker.Assert(GenericFormula.VectorHeight, Equals, suite.VectorHeight)

	checker.Assert(GenericFormula.Formula.WavePackets, HasLen, 2)
	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms, HasLen, 1)
	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, suite.eisensteinTerm[0].PowerN * -1)
	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, suite.eisensteinTerm[0].PowerM * -1)

	checker.Assert(GenericFormula.HasSymmetry(wavepacket.P1), Equals, true)
	checker.Assert(GenericFormula.HasSymmetry(wavepacket.P2), Equals, true)
}

func (suite *GenericCreatedWithDesiredSymmetry) TestCreateDesiredSymmetryWithYAML(checker *C) {
	yamlByteStream := []byte(`
vector_width: 0.6
vector_height: 0.3
formula:
  desired_symmetry: p2
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
          power_m: -9
`)

	GenericFormula, err := wavepacket.NewGenericWallpaperFormulaFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)

	checker.Assert(GenericFormula.VectorWidth, Equals, 0.6)
	checker.Assert(GenericFormula.VectorHeight, Equals, 0.3)

	checker.Assert(GenericFormula.Formula.WavePackets, HasLen, 2)
	checker.Assert(GenericFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -9)

	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, -12)
	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, 9)

	checker.Assert(GenericFormula.HasSymmetry(wavepacket.P1), Equals, true)
	checker.Assert(GenericFormula.HasSymmetry(wavepacket.P2), Equals, true)
}

func (suite *GenericCreatedWithDesiredSymmetry) TestCreateDesiredSymmetryWithJSON(checker *C) {
	jsonByteStream := []byte(`{
"vector_width": 0.6,
"vector_height": 0.3,
				"formula": {
					"desired_symmetry": "p2",
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
	GenericFormula, err := wavepacket.NewGenericWallpaperFormulaFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)

	checker.Assert(GenericFormula.VectorWidth, Equals, 0.6)
	checker.Assert(GenericFormula.VectorHeight, Equals, 0.3)

	checker.Assert(GenericFormula.Formula.WavePackets, HasLen, 2)
	checker.Assert(GenericFormula.Formula.WavePackets[0].Terms[0].PowerM, Equals, -2)

	checker.Assert(GenericFormula.Formula.WavePackets[1].Multiplier, Equals, complex(1, 0))
	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerN, Equals, -1)
	checker.Assert(GenericFormula.Formula.WavePackets[1].Terms[0].PowerM, Equals, 2)

	checker.Assert(GenericFormula.HasSymmetry(wavepacket.P1), Equals, true)
	checker.Assert(GenericFormula.HasSymmetry(wavepacket.P2), Equals, true)
}