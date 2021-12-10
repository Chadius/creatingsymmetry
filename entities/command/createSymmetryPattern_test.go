package command_test

import (
	"github.com/Chadius/creating-symmetry/entities/command"
	"github.com/Chadius/creating-symmetry/entities/formula"
	. "gopkg.in/check.v1"
	"reflect"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type CreateWallpaperCommandSuite struct {
}

var _ = Suite(&CreateWallpaperCommandSuite{})

func (suite *CreateWallpaperCommandSuite) SetUpTest(checker *C) {
}

func (suite *CreateWallpaperCommandSuite) TestCreateFromYAMLWithBuilder(checker *C) {
	yamlByteStream := []byte(`
pattern_viewport:
  x_min: 0
  y_min: 0
  x_max: 1e-10
  y_max: 3e5
coordinate_threshold:
  x_min: -50
  y_min: 9001
  x_max: -1e-1
  y_max: 2e10
eyedropper:
  left: 0
  right: 20
  top: -10
  bottom: 300
formula:
  type: rosette
  terms:
    -
      multiplier:
        real: -1.0
        imaginary: 2e-2
      power_n: 3
      power_m: 0
      coefficient_pairs:
        multiplier: 1
        relationships:
        - -M-N
        - "+M+NF"
    -
      multiplier:
        real: 1e-10
        imaginary: 0
      power_n: 1
      power_m: 1
      coefficient_pairs:
        multiplier: 1
        relationships:
        - -M-NF
`)
	wallpaperCommand, err := command.NewCreateWallpaperCommandFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)

	checker.Assert(reflect.TypeOf(wallpaperCommand.Formula).String(), Equals, "*formula.Rosette")
	checker.Assert(wallpaperCommand.Formula.FormulaLevelTerms(), HasLen, 2)
}

func (suite *CreateWallpaperCommandSuite) TestCreateFromJSONWithBuilder(checker *C) {
	jsonByteStream := []byte(`{
		"pattern_viewport": {
		  "x_min": 0,
		  "y_min": 0,
		  "x_max": 1e-10,
		  "y_max": 3e5
		},
		"coordinate_threshold": {
		  "x_min": -50,
		  "y_min": 9001,
		  "x_max": -1e-1,
		  "y_max": 2e10
		},
		"formula": {
			"type": "frieze",
			"terms": [
				{
					"multiplier": {
						"real": -1.0,
						"imaginary": 2e-2
					},
					"power_n": 3,
					"power_m": 0,
					"coefficient_pairs": {
					  "multiplier": 1,
					  "relationships": ["-M-N", "+M+NF"]
					}
				},
				{
					"multiplier": {
						"real": 1e-10,
						"imaginary": 0
					},
					"power_n": 1,
					"power_m": 1,
					"coefficient_pairs": {
					  "multiplier": 1,
					  "relationships": ["-M-NF"]
					}
				}
			]
		}
	}`)
	wallpaperCommand, err := command.NewCreateWallpaperCommandFromJSON(jsonByteStream)
	checker.Assert(err, IsNil)

	checker.Assert(wallpaperCommand.PatternViewport.XMin, Equals, 0.0)
	checker.Assert(wallpaperCommand.PatternViewport.YMin, Equals, 0.0)
	checker.Assert(wallpaperCommand.PatternViewport.XMax, Equals, 1e-10)
	checker.Assert(wallpaperCommand.PatternViewport.YMax, Equals, 3e5)

	checker.Assert(wallpaperCommand.CoordinateThreshold.XMin, Equals, -50.0)
	checker.Assert(wallpaperCommand.CoordinateThreshold.YMin, Equals, 9001.0)
	checker.Assert(wallpaperCommand.CoordinateThreshold.XMax, Equals, -1e-1)
	checker.Assert(wallpaperCommand.CoordinateThreshold.YMax, Equals, 2e10)

	checker.Assert(wallpaperCommand.Eyedropper, IsNil)

	checker.Assert(reflect.TypeOf(wallpaperCommand.Formula).String(), Equals, "*formula.Frieze")
	checker.Assert(wallpaperCommand.Formula.FormulaLevelTerms(), HasLen, 2)
}

func (suite *CreateWallpaperCommandSuite) TestMarshalWallpaperFormulaWithBuilder(checker *C) {
	yamlByteStream := []byte(`
pattern_viewport:
  x_min: 0
  y_min: 0
  x_max: 1e-10
  y_max: 3e5
coordinate_threshold:
  x_min: -50
  y_min: 9001
  x_max: -1e-1
  y_max: 2e10
formula:
  type: generic
  lattice_width: 0.8 
  lattice_height: 0.3
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
	wallpaperCommand, err := command.NewCreateWallpaperCommandFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)

	checker.Assert(reflect.TypeOf(wallpaperCommand.Formula).String(), Equals, "*formula.Generic")
	checker.Assert(wallpaperCommand.Formula.LatticeVectors()[1], Equals, complex(0.8, 0.3))
	p2SymmetryFound := false
	for _, symmetry := range wallpaperCommand.Formula.SymmetriesFound() {
		if symmetry == formula.P2 {
			p2SymmetryFound = true
		}
	}
	checker.Assert(p2SymmetryFound, Equals, true)
}
