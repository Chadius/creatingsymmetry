package transformer

import (
	"github.com/Chadius/creating-symmetry/entities/formula/frieze"
	"github.com/Chadius/creating-symmetry/entities/formula/rosette"
	"github.com/Chadius/creating-symmetry/entities/formula/wallpaper"
	"github.com/Chadius/creating-symmetry/entities/imageoutput"
	"github.com/Chadius/creating-symmetry/entities/mathutility"
	"image"
)

// FormulaTransformer turns one image stream into another using a formula
type FormulaTransformer struct {

}

// Transform converts the input image using the given formula.
func (f *FormulaTransformer) Transform(settings *Settings) *image.NRGBA {
	coordinateCollection := f.createCollectionBasedOnInputImageSize(settings)
	f.scaleCoordinatesToViewport(settings, coordinateCollection)
	f.transformCoordinatesUsingFormula(settings, coordinateCollection)
	settings.CoordinateThreshold.FilterAndMarkMappedCoordinateCollection(coordinateCollection)
	settings.Eyedropper.ConvertCoordinatesToColors(coordinateCollection)
	return f.outputToImage(settings, coordinateCollection)
}

func (f *FormulaTransformer) createCollectionBasedOnInputImageSize(settings *Settings) *imageoutput.CoordinateCollection {
	coordinates := []*imageoutput.MappedCoordinate{}
	index := 0
	for inputImageY := settings.InputImage.Bounds().Min.Y; inputImageY < settings.InputImage.Bounds().Max.Y; inputImageY++ {
		for inputImageX := settings.InputImage.Bounds().Min.X; inputImageX < settings.InputImage.Bounds().Max.X; inputImageX++ {

			if index >= 190 && index < 210 {
				println(index, inputImageX, inputImageY)
			}
			index = index + 1
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


		// DELETE
		//if index == 1000 {
			//200
			//200
			//println(settings.OutputWidth)
			//println(settings.OutputHeight)

			//-8.000000e-001
			//+8.000000e-001
			//-8.000000e-001
			//+8.000000e-001
			//println(settings.PatternViewportXMin)
			//println(settings.PatternViewportXMax)
			//println(settings.PatternViewportYMin)
			//println(settings.PatternViewportYMax)

			// 10 ? Why is this 10?
			// 5
			//println(coordinate.InputImageX())
			//println(coordinate.InputImageY())

			//-7.200000e-001
			//-7.600000e-001
			//println(patternViewportX)
			//println(patternViewportY)
		//}

		coordinate.UpdatePatternViewportCoordinates(patternViewportX, patternViewportY)
	}
}

func (f *FormulaTransformer) transformCoordinatesUsingFormula(settings *Settings, coordinateCollection *imageoutput.CoordinateCollection) {
	if settings.Formula.FriezeFormula != nil {
		f.transformCoordinatesForFriezeFormula(settings.Formula.FriezeFormula, coordinateCollection)
		return
	}
	if settings.Formula.RosetteFormula != nil {
		f.transformCoordinatesForRosetteFormula(settings.Formula.RosetteFormula, coordinateCollection)
		return
	}
	if settings.Formula.LatticePattern != nil {
		f.transformCoordinatesForLatticePattern(settings.Formula.LatticePattern, coordinateCollection)
		return
	}
}

func (f *FormulaTransformer) transformCoordinatesForFriezeFormula(friezeFormula *frieze.Formula, coordinateCollection *imageoutput.CoordinateCollection) {
	for _, coordinate := range *coordinateCollection.Coordinates() {
		complexCoordinate := complex(coordinate.PatternViewportX(), coordinate.PatternViewportY())
		friezePatternResults := friezeFormula.Calculate(complexCoordinate)
		coordinate.UpdateTransformedCoordinates(real(friezePatternResults.Total), imag(friezePatternResults.Total))
	}
}

func (f *FormulaTransformer) transformCoordinatesForRosetteFormula(rosetteFormula *rosette.Formula, coordinateCollection *imageoutput.CoordinateCollection) {
	for _, coordinate := range *coordinateCollection.Coordinates() {
		complexCoordinate := complex(coordinate.PatternViewportX(), coordinate.PatternViewportY())
		rosettePatternResults := rosetteFormula.Calculate(complexCoordinate)

		// DELETE
		//if index == 1000 {
			//-7.200000e-001
			//-7.600000e-001
			//(-6.231386e-001+2.063578e+000i)
		//	println(coordinate.PatternViewportX())
		//	println(coordinate.PatternViewportY())
		//	println(rosettePatternResults.Total)
		//}
		coordinate.UpdateTransformedCoordinates(real(rosettePatternResults.Total), imag(rosettePatternResults.Total))
	}
}

func (f *FormulaTransformer) transformCoordinatesForLatticePattern(latticePattern *wallpaper.Formula, coordinateCollection *imageoutput.CoordinateCollection) {
	latticePattern.Setup()

	for _, coordinate := range *coordinateCollection.Coordinates() {
		complexCoordinate := complex(coordinate.PatternViewportX(), coordinate.PatternViewportY())
		latticePatternResults := latticePattern.Calculate(complexCoordinate)
		coordinate.UpdateTransformedCoordinates(real(latticePatternResults.Total), imag(latticePatternResults.Total))
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