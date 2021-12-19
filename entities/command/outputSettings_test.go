package command_test

import (
	"github.com/Chadius/creating-symmetry/entities/command"
	. "gopkg.in/check.v1"
)

type CreateOutputSettings struct {
}

var _ = Suite(&CreateOutputSettings{})

func (suite *CreateOutputSettings) TestBuildOutputSettings(checker *C) {
	settings := command.NewOutputSettingsBuilder().OutputWidth(100).OutputHeight(20).Build()

	checker.Assert(settings.OutputWidth(), Equals, 100)
	checker.Assert(settings.OutputHeight(), Equals, 20)
}

func (suite *CreateOutputSettings) TestBuildOutputIgnoresInvalidOutputWidthHeight(checker *C) {
	settings := command.NewOutputSettingsBuilder().OutputWidth(-100).OutputHeight(-20).Build()

	checker.Assert(settings.OutputWidth(), Equals, 0)
	checker.Assert(settings.OutputHeight(), Equals, 0)
}

func (suite *CreateOutputSettings) TestCreateFromYAMLWithBuilder(checker *C) {
	yamlByteStream := []byte(`
output_width: 150
output_height: 40
`)
	settings := command.NewOutputSettingsBuilder().WithYAML(yamlByteStream).Build()

	checker.Assert(settings.OutputWidth(), Equals, 150)
	checker.Assert(settings.OutputHeight(), Equals, 40)
}

func (suite *CreateOutputSettings) TestCreateFromJSONWithBuilder(checker *C) {
	jsonByteStream := []byte(`{
"output_width": 150,
"output_height": 40
}`)
	settings := command.NewOutputSettingsBuilder().WithJSON(jsonByteStream).Build()

	checker.Assert(settings.OutputWidth(), Equals, 150)
	checker.Assert(settings.OutputHeight(), Equals, 40)
}
