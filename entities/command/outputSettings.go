package command

import (
	"encoding/json"
	"github.com/Chadius/creating-symmetry/entities/utility"
	"gopkg.in/yaml.v2"
)

// OutputSettingsBuilder is used to create output settings objects.
type OutputSettingsBuilder struct {
	outputWidth  int
	outputHeight int
}

// NewOutputSettingsBuilder returns a new object used to Build Formula objects.
func NewOutputSettingsBuilder() *OutputSettingsBuilder {
	return &OutputSettingsBuilder{
		outputWidth:  0,
		outputHeight: 0,
	}
}

// OutputWidth sets the outputWidth.
func (b *OutputSettingsBuilder) OutputWidth(width int) *OutputSettingsBuilder {
	if width <= 0 {
		return b
	}
	b.outputWidth = width
	return b
}

// OutputHeight sets the outputHeight.
func (b *OutputSettingsBuilder) OutputHeight(height int) *OutputSettingsBuilder {
	if height <= 0 {
		return b
	}
	b.outputHeight = height
	return b
}

// OutputSettingsBuilderMarshal can be marshaled and converted to a OutputSettingsBuilder
type OutputSettingsBuilderMarshal struct {
	OutputWidth  int `json:"output_width" yaml:"output_width"`
	OutputHeight int `json:"output_height" yaml:"output_height"`
}

// WithYAML consumes the yaml byte stream to fill settings.
func (b *OutputSettingsBuilder) WithYAML(byteStream []byte) *OutputSettingsBuilder {
	marshalSettings := newOutputSettingsBuilderMarshalFromByteStream(byteStream, yaml.Unmarshal)
	return b.setOutputSettingsUsingMarshal(marshalSettings)
}

// WithJSON consumes the yaml byte stream to fill settings.
func (b *OutputSettingsBuilder) WithJSON(byteStream []byte) *OutputSettingsBuilder {
	marshalSettings := newOutputSettingsBuilderMarshalFromByteStream(byteStream, json.Unmarshal)
	return b.setOutputSettingsUsingMarshal(marshalSettings)
}

func newOutputSettingsBuilderMarshalFromByteStream(data []byte, unmarshal utility.UnmarshalFunc) *OutputSettingsBuilderMarshal {
	var unmarshalError error
	var marshalSettings OutputSettingsBuilderMarshal
	unmarshalError = unmarshal(data, &marshalSettings)

	if unmarshalError != nil {
		return nil
	}
	return &marshalSettings
}

func (b *OutputSettingsBuilder) setOutputSettingsUsingMarshal(marshalSettings *OutputSettingsBuilderMarshal) *OutputSettingsBuilder {
	b.OutputHeight(marshalSettings.OutputHeight)
	b.OutputWidth(marshalSettings.OutputWidth)
	return b
}

// Build creates OutputSettings using the builder settings.
func (b *OutputSettingsBuilder) Build() *OutputSettings {
	return &OutputSettings{
		outputWidth:  b.outputWidth,
		outputHeight: b.outputHeight,
	}
}

// OutputSettings stores how to render the transformation.
type OutputSettings struct {
	outputWidth  int
	outputHeight int
}

// OutputWidth gets the output outputWidth.
func (o *OutputSettings) OutputWidth() int {
	return o.outputWidth
}

// OutputHeight is a getter.
func (o *OutputSettings) OutputHeight() int {
	return o.outputHeight
}
