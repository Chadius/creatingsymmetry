package creatingsymmetry

import (
	"github.com/Chadius/creating-symmetry/entities/command"
	"github.com/Chadius/creating-symmetry/entities/imageoutput"
	"github.com/Chadius/creating-symmetry/entities/transformer"
	"image"
	"image/png"
	"io"
	"io/ioutil"
)

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . TransformerStrategy
// TransformerStrategy shapes the expected messages and the expected responses when running the rules.
type TransformerStrategy interface {
	ApplyFormulaToTransformImage(inputImageDataByteStream, formulaDataByteStream, outputSettingsDataByteStream io.Reader, output io.Writer) error
}

type FileTransformer struct{}

func (f *FileTransformer) ApplyFormulaToTransformImage(inputImageDataByteStream, formulaDataByteStream, outputSettingsDataByteStream io.Reader, output io.Writer) error {
	wallpaperCommand, wallpaperErr := readWallpaperCommand(formulaDataByteStream)
	if wallpaperErr != nil {
		return wallpaperErr
	}
	sourceImage, sourceImageErr := readSourceImage(inputImageDataByteStream)
	if sourceImageErr != nil {
		return sourceImageErr
	}
	outputSettings, outputSettingsErr := readOutputSettings(outputSettingsDataByteStream)
	if outputSettingsErr != nil {
		return outputSettingsErr
	}
	outputImage := transformImage(sourceImage, wallpaperCommand, outputSettings)
	png.Encode(output, outputImage)
	return nil
}

func readSourceImage(input io.Reader) (image.Image, error) {
	colorSourceImage, _, err := image.Decode(input)
	if err != nil {
		return nil, err
	}
	return colorSourceImage, nil
}

func readWallpaperCommand(input io.Reader) (*command.CreateSymmetryPattern, error) {
	createWallpaperYAML, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, err
	}
	wallpaperCommand, err := command.NewCreateWallpaperCommandFromYAML(createWallpaperYAML)
	if err != nil {
		return nil, err
	}
	return wallpaperCommand, nil
}

func readOutputSettings(input io.Reader) (*command.OutputSettings, error) {
	outputSettingsYAML, err := ioutil.ReadAll(input)
	if err != nil {
		return nil, err
	}

	outputSettings := command.NewOutputSettingsBuilder().WithYAML(outputSettingsYAML).Build()
	return outputSettings, nil
}

func transformImage(sourceImage image.Image, wallpaperCommand *command.CreateSymmetryPattern, outputSettings *command.OutputSettings) *image.NRGBA {
	coordinateThreshold := imageoutput.CoordinateFilterBuilder().
		WithMinimumX(wallpaperCommand.CoordinateThreshold.XMin).
		WithMaximumX(wallpaperCommand.CoordinateThreshold.XMax).
		WithMinimumY(wallpaperCommand.CoordinateThreshold.YMin).
		WithMaximumY(wallpaperCommand.CoordinateThreshold.YMax).
		Build()

	var eyedropper *imageoutput.RectangularEyedropper
	if wallpaperCommand.Eyedropper != nil {
		eyedropper = imageoutput.EyedropperBuilder().
			WithLeftSide(wallpaperCommand.Eyedropper.LeftSide).
			WithRightSide(wallpaperCommand.Eyedropper.RightSide).
			WithTopSide(wallpaperCommand.Eyedropper.TopSide).
			WithBottomSide(wallpaperCommand.Eyedropper.BottomSide).
			WithImage(sourceImage).
			Build()
	} else {
		eyedropper = imageoutput.EyedropperBuilder().
			WithLeftSide(sourceImage.Bounds().Min.X).
			WithRightSide(sourceImage.Bounds().Max.X).
			WithTopSide(sourceImage.Bounds().Min.Y).
			WithBottomSide(sourceImage.Bounds().Max.Y).
			WithImage(sourceImage).
			Build()
	}

	transformerEntity := transformer.FormulaTransformer{}

	outputImage := transformerEntity.Transform(&transformer.Settings{
		PatternViewportXMin: wallpaperCommand.PatternViewport.XMin,
		PatternViewportXMax: wallpaperCommand.PatternViewport.XMax,
		PatternViewportYMin: wallpaperCommand.PatternViewport.YMin,
		PatternViewportYMax: wallpaperCommand.PatternViewport.YMax,
		InputImage:          sourceImage,
		Formula:             wallpaperCommand,
		CoordinateThreshold: coordinateThreshold,
		Eyedropper:          eyedropper,
		OutputWidth:         outputSettings.OutputWidth(),
		OutputHeight:        outputSettings.OutputHeight(),
	})
	return outputImage
}
