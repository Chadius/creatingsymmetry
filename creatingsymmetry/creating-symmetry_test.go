package creatingsymmetry_test

import (
	"bytes"
	"github.com/Chadius/creating-symmetry/creatingsymmetry"
	. "gopkg.in/check.v1"
	"image"
	"image/color"
	"image/png"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type ReadInputStreamsSuite struct{}

var _ = Suite(&ReadInputStreamsSuite{})

func (suite *ReadInputStreamsSuite) TestAcceptStream(checker *C) {
	formulaData := []byte(`pattern_viewport:
  x_min: 0
  y_min: 0
  x_max: 10
  y_max: 10
formula:
  type: identity
`)
	formulaDataByteStream := bytes.NewBuffer(formulaData)

	outputSettingsData := []byte(`
output_width: 2
output_height: 1
`)
	outputSettingsDataByteStream := bytes.NewBuffer(outputSettingsData)

	sourceColors := image.NewNRGBA(image.Rect(0, 0, 1, 1))
	sourceColors.Set(0, 0, color.NRGBA{
		R: uint8(255),
		G: 0,
		B: 0,
		A: uint8(255),
	})
	sourceImage := sourceColors.SubImage(image.Rect(0, 0, 1, 1))
	inputImageDataByteStream := new(bytes.Buffer)
	png.Encode(inputImageDataByteStream, sourceImage)

	var output bytes.Buffer

	transformer := creatingsymmetry.FileTransformer{}
	err := transformer.ApplyFormulaToTransformImage(inputImageDataByteStream, formulaDataByteStream, outputSettingsDataByteStream, &output)

	checker.Assert(err, IsNil)

	outputImageByteReader := bytes.NewReader(output.Bytes())
	outputImage, decodeError := png.Decode(outputImageByteReader)

	checker.Assert(decodeError, IsNil)
	checker.Assert(outputImage.Bounds().Max.X, Equals, 2)
	checker.Assert(outputImage.Bounds().Max.Y, Equals, 1)

	r0, g0, b0, a0 := outputImage.At(0, 0).RGBA()
	checker.Assert(r0, Equals, uint32(0xffff))
	checker.Assert(g0, Equals, uint32(0))
	checker.Assert(b0, Equals, uint32(0))
	checker.Assert(a0, Equals, uint32(0xffff))
}
