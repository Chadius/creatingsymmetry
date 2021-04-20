package wave_test

import (
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/wave"
	"wallpaper/entities/utility"
)

type HexagonalWaveFormula struct {
	hexagonalWavePacket *wave.HexagonalWallpaperFormula
}

var _ = Suite(&HexagonalWaveFormula{})

func (suite *HexagonalWaveFormula) SetUpTest(checker *C) {
	suite.hexagonalWavePacket = &wave.HexagonalWallpaperFormula{
		WavePackets: []*wave.Formula{
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
	}
}

func (suite *HexagonalWaveFormula) TestHexagonalWallpaperImpliesAveragedLockedTerms(checker *C) {
	suite.hexagonalWavePacket.SetUp()
	calculation := suite.hexagonalWavePacket.Calculate(complex(math.Sqrt(3), -1 * math.Sqrt(3)))
	total := calculation.Total

	expectedAnswer := (cmplx.Exp(complex(0, 2 * math.Pi * (3 + math.Sqrt(3)))) +
		cmplx.Exp(complex(0, 2 * math.Pi * (-2 * math.Sqrt(3)))) +
		cmplx.Exp(complex(0, 2 * math.Pi * (-3 + math.Sqrt(3))))) / 3

	checker.Assert(real(total), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(total), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}

func (suite *HexagonalWaveFormula) TestUnmarshalFromJSON(checker *C) {
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
	hexFormula, err := wave.NewHexagonalWallpaperFormulaFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(real(hexFormula.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(hexFormula.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
	checker.Assert(hexFormula.WavePackets, HasLen, 1)
	checker.Assert(hexFormula.WavePackets[0].Terms[0].PowerM, Equals, -10)
}

func (suite *HexagonalWaveFormula) TestUnmarshalFromYAML(checker *C) {
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
	hexFormula, err := wave.NewHexagonalWallpaperFormulaFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)
	checker.Assert(real(hexFormula.Multiplier), utility.NumericallyCloseEnough{}, -1.0, 1e-6)
	checker.Assert(imag(hexFormula.Multiplier), utility.NumericallyCloseEnough{}, 2e-2, 1e-6)
	checker.Assert(hexFormula.WavePackets, HasLen, 1)
	checker.Assert(hexFormula.WavePackets[0].Terms[0].PowerM, Equals, -10)
}