package command_test

import (
	"github.com/Chadius/creating-symmetry/entities/command"
	"github.com/Chadius/creating-symmetry/entities/formula/wallpaper"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type CreateWallpaperCommandSuite struct {
}

var _ = Suite(&CreateWallpaperCommandSuite{})

func (suite *CreateWallpaperCommandSuite) SetUpTest(checker *C) {
}

func (suite *CreateWallpaperCommandSuite) TestCreateFromYAML(checker *C) {
	yamlByteStream := []byte(`sample_source_filename: input.png
output_filename: output.png
output_size:
  width: 800
  height: 600
pattern_viewport:
  x_min: 0
  y_min: 0
  x_max: 1e-10
  y_max: 3e5
eyedropper_boundary:
  x_min: -50
  y_min: 9001
  x_max: -1e-1
  y_max: 2e10
rosette_formula:
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
	checker.Assert(wallpaperCommand.SampleSourceFilename, Equals, "input.png")

	checker.Assert(wallpaperCommand.OutputFilename, Equals, "output.png")
	checker.Assert(wallpaperCommand.OutputImageSize.Width, Equals, 800)
	checker.Assert(wallpaperCommand.OutputImageSize.Height, Equals, 600)

	checker.Assert(wallpaperCommand.PatternViewport.XMin, Equals, 0.0)
	checker.Assert(wallpaperCommand.PatternViewport.YMin, Equals, 0.0)
	checker.Assert(wallpaperCommand.PatternViewport.XMax, Equals, 1e-10)
	checker.Assert(wallpaperCommand.PatternViewport.YMax, Equals, 3e5)

	checker.Assert(wallpaperCommand.EyedropperBoundary.XMin, Equals, -50.0)
	checker.Assert(wallpaperCommand.EyedropperBoundary.YMin, Equals, 9001.0)
	checker.Assert(wallpaperCommand.EyedropperBoundary.XMax, Equals, -1e-1)
	checker.Assert(wallpaperCommand.EyedropperBoundary.YMax, Equals, 2e10)

	checker.Assert(wallpaperCommand.RosetteFormula.Terms, HasLen, 2)
}

func (suite *CreateWallpaperCommandSuite) TestCreateFromJSON(checker *C) {
	jsonByteStream := []byte(`{
				"sample_source_filename": "input.png",
				"output_filename": "output.png",
				"output_size": {
				  "width": 800,
				  "height": 600
				},
				"pattern_viewport": {
				  "x_min": 0,
				  "y_min": 0,
				  "x_max": 1e-10,
				  "y_max": 3e5
				},
				"eyedropper_boundary": {
				  "x_min": -50,
				  "y_min": 9001,
				  "x_max": -1e-1,
				  "y_max": 2e10
				},
				"frieze_formula": {
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
	checker.Assert(wallpaperCommand.SampleSourceFilename, Equals, "input.png")

	checker.Assert(wallpaperCommand.OutputFilename, Equals, "output.png")
	checker.Assert(wallpaperCommand.OutputImageSize.Width, Equals, 800)
	checker.Assert(wallpaperCommand.OutputImageSize.Height, Equals, 600)

	checker.Assert(wallpaperCommand.PatternViewport.XMin, Equals, 0.0)
	checker.Assert(wallpaperCommand.PatternViewport.YMin, Equals, 0.0)
	checker.Assert(wallpaperCommand.PatternViewport.XMax, Equals, 1e-10)
	checker.Assert(wallpaperCommand.PatternViewport.YMax, Equals, 3e5)

	checker.Assert(wallpaperCommand.EyedropperBoundary.XMin, Equals, -50.0)
	checker.Assert(wallpaperCommand.EyedropperBoundary.YMin, Equals, 9001.0)
	checker.Assert(wallpaperCommand.EyedropperBoundary.XMax, Equals, -1e-1)
	checker.Assert(wallpaperCommand.EyedropperBoundary.YMax, Equals, 2e10)

	checker.Assert(wallpaperCommand.FriezeFormula.Terms, HasLen, 2)
}

func (suite *CreateWallpaperCommandSuite) TestMarshalWallpaperFormula(checker *C) {
	yamlByteStream := []byte(`sample_source_filename: input.png
output_filename: output.png
output_size:
  width: 800
  height: 600
pattern_viewport:
  x_min: 0
  y_min: 0
  x_max: 1e-10
  y_max: 3e5
eyedropper_boundary:
  x_min: -50
  y_min: 9001
  x_max: -1e-1
  y_max: 2e10
lattice_pattern:
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
	wallpaperCommand, err := command.NewCreateWallpaperCommandFromYAML(yamlByteStream)
	checker.Assert(err, IsNil)

	checker.Assert(wallpaperCommand.LatticePattern.LatticeType, Equals, wallpaper.Generic)
	checker.Assert(wallpaperCommand.LatticePattern.DesiredSymmetry, Equals, wallpaper.P2)
	checker.Assert(wallpaperCommand.LatticePattern.Multiplier, Equals, complex(-1.0, 2e-2))

	checker.Assert(wallpaperCommand.LatticePattern.LatticeSize.Width, Equals, 0.8)
	checker.Assert(wallpaperCommand.LatticePattern.LatticeSize.Height, Equals, 0.3)

	checker.Assert(wallpaperCommand.LatticePattern.WavePackets, HasLen, 1)

}