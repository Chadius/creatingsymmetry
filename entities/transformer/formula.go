package transformer

import (
	"github.com/chadius/creatingsymmetry/entities/formula"
	"github.com/chadius/creatingsymmetry/entities/imageoutput"
	"github.com/chadius/creatingsymmetry/entities/mathutility"
	"image"
)

// FormulaTransformer turns one image stream into another using a oldformula
type FormulaTransformer struct {
}

// Transform converts the input image using the given oldformula.
func (f *FormulaTransformer) Transform(settings *Settings) *image.NRGBA {
	coordinateCollection := f.createCollectionBasedOnOutputImageSize(settings)
	f.scaleCoordinatesToViewport(settings, coordinateCollection)
	f.transformCoordinatesUsingFormula(settings, coordinateCollection)
	settings.CoordinateThreshold.FilterAndMarkMappedCoordinateCollection(coordinateCollection)
	settings.Eyedropper.ConvertCoordinatesToColors(coordinateCollection)
	return f.outputToImage(settings, coordinateCollection)
}

func (f *FormulaTransformer) createCollectionBasedOnOutputImageSize(settings *Settings) *imageoutput.CoordinateCollection {
	coordinates := []*imageoutput.MappedCoordinate{}
	for inputImageY := 0; inputImageY < settings.OutputHeight; inputImageY++ {
		for inputImageX := 0; inputImageX < settings.OutputWidth; inputImageX++ {
			coordinates = append(
				coordinates,
				imageoutput.NewMappedCoordinateUsingInputImageCoordinates(inputImageX, inputImageY),
			)
		}
	}

	return imageoutput.CoordinateCollectionBuilder().WithCoordinates(&coordinates).Build()
}

func (f *FormulaTransformer) scaleCoordinatesToViewport(settings *Settings, coordinateCollection *imageoutput.CoordinateCollection) {
	for _, coordinate := range *coordinateCollection.Coordinates() {
		patternViewportX := mathutility.ScaleValueBetweenTwoRanges(
			float64(coordinate.InputImageX()),
			float64(0),
			float64(settings.OutputWidth),
			settings.PatternViewportXMin,
			settings.PatternViewportXMax,
		)

		patternViewportY := mathutility.ScaleValueBetweenTwoRanges(
			float64(coordinate.InputImageY()),
			float64(0),
			float64(settings.OutputHeight),
			settings.PatternViewportYMin,
			settings.PatternViewportYMax,
		)
		coordinate.UpdatePatternViewportCoordinates(patternViewportX, patternViewportY)
	}
}

func (f *FormulaTransformer) transformCoordinatesUsingFormula(settings *Settings, coordinateCollection *imageoutput.CoordinateCollection) {
	if settings.Formula.Formula != nil {
		f.transformCoordinatesForArbitraryFormula(settings.Formula.Formula, coordinateCollection)
	}
}

func (f *FormulaTransformer) transformCoordinatesForArbitraryFormula(arbitraryFormula formula.Arbitrary, coordinateCollection *imageoutput.CoordinateCollection) {
	for _, coordinate := range *coordinateCollection.Coordinates() {
		complexCoordinate := complex(coordinate.PatternViewportX(), coordinate.PatternViewportY())
		transformedPoint := arbitraryFormula.Calculate(complexCoordinate)
		coordinate.UpdateTransformedCoordinates(real(transformedPoint), imag(transformedPoint))
	}
}

func (f *FormulaTransformer) outputToImage(settings *Settings, coordinateCollection *imageoutput.CoordinateCollection) *image.NRGBA {
	outputImage := image.NewNRGBA(image.Rect(0, 0, settings.OutputWidth, settings.OutputHeight))

	colorData := settings.Eyedropper.ConvertCoordinatesToColors(coordinateCollection)

	for index, colorToAdd := range *colorData {
		destinationPixelX := index % settings.OutputWidth
		destinationPixelY := index / settings.OutputWidth
		outputImage.Set(
			destinationPixelX,
			destinationPixelY,
			colorToAdd,
		)
	}
	return outputImage
}
