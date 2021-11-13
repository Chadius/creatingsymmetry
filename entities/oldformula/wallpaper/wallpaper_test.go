package wallpaper_test

import (
	"github.com/Chadius/creating-symmetry/entities/oldformula/eisenstien"
	"github.com/Chadius/creating-symmetry/entities/oldformula/wallpaper"
	. "gopkg.in/check.v1"
	"math"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type WallpaperMarshalTest struct{}

var _ = Suite(&WallpaperMarshalTest{})

func (suite *WallpaperMarshalTest) SetUpTest(checker *C) {}

func (suite *WallpaperMarshalTest) TestCreateFormulaFromYAML(checker *C) {
	yamlByteStream := []byte(`
lattice_type: generic
lattice_size:
  width: 0.8 
  height: 0.3
multiplier:
  real: -1.0
  imaginary: 2e-2
wave_packets:
-
  multiplier:
    real: 2
    imaginary: 3
  terms:
  -
    power_n: 12
    power_m: -10
  -
    power_n: -5
    power_m: 3
desired_symmetry: p2
`)
	formula, err := wallpaper.NewFormulaFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)

	checker.Assert(formula.LatticeType, Equals, wallpaper.Generic)
	checker.Assert(formula.DesiredSymmetry, Equals, wallpaper.P2)
	checker.Assert(formula.Multiplier, Equals, complex(-1.0, 2e-2))

	checker.Assert(formula.LatticeSize.Width, Equals, 0.8)
	checker.Assert(formula.LatticeSize.Height, Equals, 0.3)

	checker.Assert(formula.WavePackets, HasLen, 1)
	checker.Assert(formula.WavePackets[0].Multiplier, Equals, complex(2, 3))
	checker.Assert(formula.WavePackets[0].Terms, HasLen, 2)
	checker.Assert(formula.WavePackets[0].Terms[0].PowerN, Equals, 12)
	checker.Assert(formula.WavePackets[0].Terms[0].PowerM, Equals, -10)
	checker.Assert(formula.WavePackets[0].Terms[1].PowerN, Equals, -5)
	checker.Assert(formula.WavePackets[0].Terms[1].PowerM, Equals, 3)
}

func (suite *WallpaperMarshalTest) TestCreateFormulaFromJSON(checker *C) {
	jsonByteStream := []byte(`{
	"lattice_type": "hexagonal",
	"desired_symmetry": "p3m1",
	"desired_symmetry": "p3m1",
	"multiplier": {
		"real": 1.0,
		"imaginary": 20
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
`)
	formula, err := wallpaper.NewFormulaFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)

	checker.Assert(formula.LatticeType, Equals, wallpaper.Hexagonal)
	checker.Assert(formula.DesiredSymmetry, Equals, wallpaper.P3m1)
	checker.Assert(formula.Multiplier, Equals, complex(1.0, 20))

	checker.Assert(formula.WavePackets, HasLen, 1)
	checker.Assert(formula.WavePackets[0].Multiplier, Equals, complex(1, 0))
	checker.Assert(formula.WavePackets[0].Terms, HasLen, 1)
	checker.Assert(formula.WavePackets[0].Terms[0].PowerN, Equals, 1)
	checker.Assert(formula.WavePackets[0].Terms[0].PowerM, Equals, -2)
}

type MakeNewFormulaBasedOnLatticeShape struct{}

var _ = Suite(&MakeNewFormulaBasedOnLatticeShape{})

func (suite *MakeNewFormulaBasedOnLatticeShape) SetUpTest(checker *C) {}

func (suite *MakeNewFormulaBasedOnLatticeShape) TestMakeGenericFormula(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Generic,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0.5,
			Height: 2.3,
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			{
				Multiplier: complex(1, 0),
				Terms: []*eisenstien.EisensteinFormulaTerm{
					{
						PowerN: 1,
						PowerM: -4,
					},
				},
			},
		},
		DesiredSymmetry: wallpaper.P1,
	}

	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.Lattice.XLatticeVector, Equals, complex(1, 0))
	checker.Assert(newFormula.Lattice.YLatticeVector, Equals, complex(0.5, 2.3))

	checker.Assert(newFormula.WavePackets, HasLen, 1)
	checker.Assert(newFormula.WavePackets[0].Terms, HasLen, 1)

	checker.Assert(newFormula.WavePackets[0].Terms[0].PowerN, Equals, 1)
	checker.Assert(newFormula.WavePackets[0].Terms[0].PowerM, Equals, -4)
}

func (suite *MakeNewFormulaBasedOnLatticeShape) TestMakeHexagonalFormula(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Hexagonal,
		LatticeSize: nil,
		Lattice:     nil,
		Multiplier:  complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			{
				Multiplier: complex(1, 0),
				Terms: []*eisenstien.EisensteinFormulaTerm{
					{
						PowerN: 1,
						PowerM: -4,
					},
				},
			},
		},
		DesiredSymmetry: "p3",
	}

	err := newFormula.Setup()
	checker.Assert(err, IsNil)

	checker.Assert(newFormula.Lattice.XLatticeVector, Equals, complex(1, 0))
	checker.Assert(newFormula.Lattice.YLatticeVector, Equals, complex(-0.5, math.Sqrt(3.0)/2.0))

	checker.Assert(newFormula.WavePackets, HasLen, 1)
	checker.Assert(newFormula.WavePackets[0].Terms, HasLen, 3)

	checker.Assert(newFormula.WavePackets[0].Terms[1].PowerN, Equals, -4)
	checker.Assert(newFormula.WavePackets[0].Terms[1].PowerM, Equals, 3)

	checker.Assert(newFormula.WavePackets[0].Terms[2].PowerN, Equals, 3)
	checker.Assert(newFormula.WavePackets[0].Terms[2].PowerM, Equals, 1)
}

func (suite *MakeNewFormulaBasedOnLatticeShape) TestSetupThrowsAnErrorIfVectorsAreZero(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Generic,
		LatticeSize: &wallpaper.Dimensions{
			Width:  0,
			Height: 0,
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			{
				Multiplier: complex(1, 0),
				Terms: []*eisenstien.EisensteinFormulaTerm{
					{
						PowerN: 1,
						PowerM: -4,
					},
				},
			},
		},
		DesiredSymmetry: "p1",
	}

	err := newFormula.Setup()
	checker.Assert(err, ErrorMatches, "lattice vectors cannot be \\(0,0\\)")
}

func (suite *MakeNewFormulaBasedOnLatticeShape) TestSetupThrowsAnErrorIfVectorsAreCollinear(checker *C) {
	newFormula := wallpaper.Formula{
		LatticeType: wallpaper.Generic,
		LatticeSize: &wallpaper.Dimensions{
			Width:  -10,
			Height: 0,
		},
		Lattice:    nil,
		Multiplier: complex(2, 0),
		WavePackets: []*wallpaper.WavePacket{
			{
				Multiplier: complex(1, 0),
				Terms: []*eisenstien.EisensteinFormulaTerm{
					{
						PowerN: 1,
						PowerM: -4,
					},
				},
			},
		},
		DesiredSymmetry: "p1",
	}

	err := newFormula.Setup()
	checker.Assert(err, ErrorMatches, "vectors cannot be collinear: (.*,.*) and (.*,.*)")
}

// (Start making tests for Hex and Generic wallpapers)
// (Like Symmetry checks)
