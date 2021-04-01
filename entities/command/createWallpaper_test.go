package command_test

import (
	. "gopkg.in/check.v1"
	"testing"
	"wallpaper/entities/command"
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
sample_space:
  minx: 0
  miny: 0
  maxx: 1e-10
  maxy: 3e5
color_value_space:
  minx: -50
  miny: 9001
  maxx: -1e-1
  maxy: 2e10
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

	checker.Assert(wallpaperCommand.SampleSpace.MinX, Equals, 0.0)
	checker.Assert(wallpaperCommand.SampleSpace.MinY, Equals, 0.0)
	checker.Assert(wallpaperCommand.SampleSpace.MaxX, Equals, 1e-10)
	checker.Assert(wallpaperCommand.SampleSpace.MaxY, Equals, 3e5)

	checker.Assert(wallpaperCommand.ColorValueSpace.MinX, Equals, -50.0)
	checker.Assert(wallpaperCommand.ColorValueSpace.MinY, Equals, 9001.0)
	checker.Assert(wallpaperCommand.ColorValueSpace.MaxX, Equals, -1e-1)
	checker.Assert(wallpaperCommand.ColorValueSpace.MaxY, Equals, 2e10)

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
				"sample_space": {
				  "minx": 0,
				  "miny": 0,
				  "maxx": 1e-10,
				  "maxy": 3e5
				},
				"color_value_space": {
				  "minx": -50,
				  "miny": 9001,
				  "maxx": -1e-1,
				  "maxy": 2e10
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

	checker.Assert(wallpaperCommand.SampleSpace.MinX, Equals, 0.0)
	checker.Assert(wallpaperCommand.SampleSpace.MinY, Equals, 0.0)
	checker.Assert(wallpaperCommand.SampleSpace.MaxX, Equals, 1e-10)
	checker.Assert(wallpaperCommand.SampleSpace.MaxY, Equals, 3e5)

	checker.Assert(wallpaperCommand.ColorValueSpace.MinX, Equals, -50.0)
	checker.Assert(wallpaperCommand.ColorValueSpace.MinY, Equals, 9001.0)
	checker.Assert(wallpaperCommand.ColorValueSpace.MaxX, Equals, -1e-1)
	checker.Assert(wallpaperCommand.ColorValueSpace.MaxY, Equals, 2e10)

	checker.Assert(wallpaperCommand.FriezeFormula.Terms, HasLen, 2)
}
