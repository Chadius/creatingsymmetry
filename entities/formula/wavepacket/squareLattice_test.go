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
