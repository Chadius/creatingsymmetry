package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Chadius/creating-symmetry/entities/command"
	"github.com/Chadius/creating-symmetry/entities/formula/frieze"
	"github.com/Chadius/creating-symmetry/entities/formula/rosette"
	"github.com/Chadius/creating-symmetry/entities/formula/wallpaper"
	"github.com/Chadius/creating-symmetry/entities/imageoutput"
	"github.com/Chadius/creating-symmetry/entities/mathutility"
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	filenameArguments := extractFilenameArguments()
	wallpaperCommand, err := loadFormulaFile(filenameArguments)
	destinationBounds, destinationCoordinates := getDestinationBoundary(filenameArguments)
	scaledCoordinates := scaleDestinationToPatternViewport(wallpaperCommand, destinationBounds, destinationCoordinates)
	transformedCoordinates := transformCoordinatesAndReport(wallpaperCommand, scaledCoordinates)
	outputImage := eyedropperFinalColorAndSaveToImage(wallpaperCommand, err, filenameArguments, destinationCoordinates, transformedCoordinates)
	outputToFile(filenameArguments.OutputFilename, outputImage)
}

// eyedropperFinalColorAndSaveToImage uses the CoordinateThreshold to select the colors in the output image.
//   It returns an image buffer.
func eyedropperFinalColorAndSaveToImage(wallpaperCommand *command.CreateSymmetryPattern, err error, filenameArguments *FilenameArguments, destinationCoordinates []complex128, transformedCoordinates []complex128) *image.NRGBA {
	//colorValueBoundMin := complex(wallpaperCommand.CoordinateThreshold.XMin, wallpaperCommand.CoordinateThreshold.YMin)
	//colorValueBoundMax := complex(wallpaperCommand.CoordinateThreshold.XMax, wallpaperCommand.CoordinateThreshold.YMax)
	//colorSourceImage := openSourceImage(err, filenameArguments)
	//outputImage := image.NewNRGBA(image.Rect(0, 0, filenameArguments.OutputWidth, filenameArguments.OutputHeight))
	//colorDestinationImage(outputImage, colorSourceImage, destinationCoordinates, transformedCoordinates, colorValueBoundMin, colorValueBoundMax)
	//return outputImage
	return helperForMapTransformedPointsToOutputImageBuffer(wallpaperCommand, filenameArguments, transformedCoordinates)
}

// transformCoordinatesAndReport applies the formula on the destination
//   returning a flat array of coordinates.
//   It also reports on the bounding box.
func transformCoordinatesAndReport(wallpaperCommand *command.CreateSymmetryPattern, scaledCoordinates []complex128) []complex128 {
	transformedCoordinates := transformCoordinatesForFormula(wallpaperCommand, scaledCoordinates)
	zMin, zMax := mathutility.GetBoundingBox(transformedCoordinates)
	println(zMin)
	println(zMax)
	return transformedCoordinates
}

func scaleDestinationToPatternViewport(wallpaperCommand *command.CreateSymmetryPattern, destinationBounds image.Rectangle, destinationCoordinates []complex128) []complex128 {
	patternViewportMin := complex(wallpaperCommand.PatternViewport.XMin, wallpaperCommand.PatternViewport.YMin)
	patternViewportMax := complex(wallpaperCommand.PatternViewport.XMax, wallpaperCommand.PatternViewport.YMax)
	scaledCoordinates := scaleDestinationPixels(
		destinationBounds,
		destinationCoordinates,
		patternViewportMin,
		patternViewportMax,
	)
	return scaledCoordinates
}

// getDestinationBoundary uses the Argument's OutputWidth and OutputHeight to return
//  a Rectangular buffer of those dimensions, and a flat array of the individual pixels.
func getDestinationBoundary(filenameArguments *FilenameArguments) (image.Rectangle, []complex128) {
	destinationBounds := image.Rect(0, 0, filenameArguments.OutputWidth, filenameArguments.OutputHeight)
	destinationCoordinates := flattenCoordinates(destinationBounds)
	return destinationBounds, destinationCoordinates
}

func loadFormulaFile(filenameArguments *FilenameArguments) (*command.CreateSymmetryPattern, error) {
	createWallpaperYAML, err := ioutil.ReadFile(filenameArguments.FormulaFilename)
	if err != nil {
		log.Fatal(err)
	}
	wallpaperCommand, err := command.NewCreateWallpaperCommandFromYAML(createWallpaperYAML)
	if err != nil {
		log.Fatal(err)
	}
	return wallpaperCommand, err
}

func openSourceImage(err error, filenameArguments *FilenameArguments) image.Image {
	reader, err := os.Open(filenameArguments.SourceImageFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	colorSourceImage, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	return colorSourceImage
}

func outputToFile(outputFilename string, outputImage image.Image) {
	err := os.MkdirAll(filepath.Dir(outputFilename), 0777)
	if err != nil {
		panic(err)
	}
	outputImageFile, err := os.Create(outputFilename)
	if err != nil {
		panic(err)
	}
	defer outputImageFile.Close()
	png.Encode(outputImageFile, outputImage)
}

func transformCoordinatesForFormula(command *command.CreateSymmetryPattern, scaledCoordinates []complex128) []complex128 {
	if command.FriezeFormula != nil {
		return transformCoordinatesForFriezeFormula(command.FriezeFormula, scaledCoordinates)
	}
	if command.RosetteFormula != nil {
		return transformCoordinatesForRosetteFormula(command.RosetteFormula, scaledCoordinates)
	}
	if command.LatticePattern != nil {
		return transformCoordinatesForLatticePattern(command.LatticePattern, scaledCoordinates)
	}
	log.Fatal(errors.New("no formula found"))
	return []complex128{}
}

func transformCoordinatesForFriezeFormula(friezeFormula *frieze.Formula, scaledCoordinates []complex128) []complex128 {
	symmetryAnalysis := friezeFormula.AnalyzeForSymmetry()
	if symmetryAnalysis.P111 {
		println("Has these symmetries: p111")
	}
	if symmetryAnalysis.P211 {
		println("  P211")
	}
	if symmetryAnalysis.P1m1 {
		println("  P1m1")
	}
	if symmetryAnalysis.P11g {
		println("  P11g")
	}
	if symmetryAnalysis.P11m {
		println("  P11m")
	}
	if symmetryAnalysis.P2mm {
		println("  P2mm")
	}
	if symmetryAnalysis.P2mg {
		println("  P2mg")
	}

	transformedCoordinates := []complex128{}
	resultsByTerm := [][]complex128{}
	for range friezeFormula.Terms {
		resultsByTerm = append(resultsByTerm, []complex128{})
	}

	for _, complexCoordinate := range scaledCoordinates {
		friezeResults := friezeFormula.Calculate(complexCoordinate)
		for index, formulaResult := range friezeResults.ContributionByTerm {
			resultsByTerm[index] = append(resultsByTerm[index], formulaResult)
		}

		transformedCoordinate := friezeResults.Total
		transformedCoordinates = append(transformedCoordinates, transformedCoordinate)
	}

	println("Min/Max ranges, by Term")
	for index, results := range resultsByTerm {
		minz, maxz := mathutility.GetBoundingBox(results)
		fmt.Printf("%d: %e - %e\n", index, minz, maxz)
	}
	return transformedCoordinates
}

func transformCoordinatesForRosetteFormula(rosetteFormula *rosette.Formula, scaledCoordinates []complex128) []complex128 {
	transformedCoordinates := []complex128{}
	resultsByTerm := [][]complex128{}
	for range rosetteFormula.Terms {
		resultsByTerm = append(resultsByTerm, []complex128{})
	}

	for _, complexCoordinate := range scaledCoordinates {
		rosetteResults := rosetteFormula.Calculate(complexCoordinate)
		for index, formulaResult := range rosetteResults.ContributionByTerm {
			resultsByTerm[index] = append(resultsByTerm[index], formulaResult)
		}

		transformedCoordinate := rosetteResults.Total
		transformedCoordinates = append(transformedCoordinates, transformedCoordinate)
	}

	println("Min/Max ranges, by Term")
	for index, results := range resultsByTerm {
		minz, maxz := mathutility.GetBoundingBox(results)
		fmt.Printf("%d: %e - %e\n", index, minz, maxz)
	}
	return transformedCoordinates
}

func transformCoordinatesForLatticePattern(latticePattern *wallpaper.Formula, scaledCoordinates []complex128) []complex128 {
	latticePattern.Setup()

	transformedCoordinates := []complex128{}
	resultsByTerm := [][]complex128{}
	for range latticePattern.WavePackets {
		resultsByTerm = append(resultsByTerm, []complex128{})
	}

	for _, complexCoordinate := range scaledCoordinates {
		latticePatternResults := latticePattern.Calculate(complexCoordinate)
		for index, formulaResult := range latticePatternResults.ContributionByTerm {
			resultsByTerm[index] = append(resultsByTerm[index], formulaResult)
		}

		transformedCoordinate := latticePatternResults.Total
		transformedCoordinates = append(transformedCoordinates, transformedCoordinate)
	}

	println("Min/Max ranges, by Term")
	for index, results := range resultsByTerm {
		minz, maxz := mathutility.GetBoundingBox(results)
		fmt.Printf("%d: %e - %e\n", index, minz, maxz)
	}

	println("Symmetries found:")
	if latticePattern.HasSymmetry(wallpaper.P1) {
		println("  p1")
	}
	if latticePattern.HasSymmetry(wallpaper.P2) {
		println("  p2")
	}
	if latticePattern.HasSymmetry(wallpaper.P31m) {
		println("  p31m")
	}
	if latticePattern.HasSymmetry(wallpaper.P3m1) {
		println("  p3m1")
	}
	if latticePattern.HasSymmetry(wallpaper.P6) {
		println("  p6")
	}
	if latticePattern.HasSymmetry(wallpaper.P6m) {
		println("  p6m")
	}
	if latticePattern.HasSymmetry(wallpaper.P3) {
		println("  p3")
	}
	if latticePattern.HasSymmetry(wallpaper.P4) {
		println("  p4")
	}
	if latticePattern.HasSymmetry(wallpaper.P4m) {
		println("  p4m")
	}
	if latticePattern.HasSymmetry(wallpaper.P4g) {
		println("  p4g")
	}
	if latticePattern.HasSymmetry(wallpaper.Cm) {
		println("  cm")
	}
	if latticePattern.HasSymmetry(wallpaper.Cmm) {
		println("  cmm")
	}
	if latticePattern.HasSymmetry(wallpaper.Pm) {
		println("  pm")
	}
	if latticePattern.HasSymmetry(wallpaper.Pg) {
		println("  pg")
	}
	if latticePattern.HasSymmetry(wallpaper.Pmm) {
		println("  pmm")
	}
	if latticePattern.HasSymmetry(wallpaper.Pmg) {
		println("  pmg")
	}
	if latticePattern.HasSymmetry(wallpaper.Pgg) {
		println("  pgg")
	}
	return transformedCoordinates
}

// flattenCoordinates returns an array of coordinates (using complex128 for each pair.)
//   Each item in the array ranges from (0,0) to (destinationBounds.Min.X,destinationBounds.Min.Y).
func flattenCoordinates(destinationBounds image.Rectangle) []complex128 {
	flattenedCoordinates := []complex128{}
	for destinationY := destinationBounds.Min.Y; destinationY < destinationBounds.Max.Y; destinationY++ {
		for destinationX := destinationBounds.Min.X; destinationX < destinationBounds.Max.X; destinationX++ {
			flattenedCoordinates = append(flattenedCoordinates, complex(float64(destinationX), float64(destinationY)))
		}
	}
	return flattenedCoordinates
}

// scaleDestinationPixels returns a list of pairs of float64. Each pair is mapped from the viewPort to the destinationBoundary.
func scaleDestinationPixels(destinationBounds image.Rectangle, destinationCoordinates []complex128, viewPortMin complex128, viewPortMax complex128) []complex128 {
	scaledCoordinates := []complex128{}
	for _, destinationCoordinate := range destinationCoordinates {
		destinationScaledX := mathutility.ScaleValueBetweenTwoRanges(
			real(destinationCoordinate),
			float64(destinationBounds.Min.X),
			float64(destinationBounds.Max.X),
			real(viewPortMin),
			real(viewPortMax),
		)
		destinationScaledY := mathutility.ScaleValueBetweenTwoRanges(
			imag(destinationCoordinate),
			float64(destinationBounds.Min.Y),
			float64(destinationBounds.Max.Y),
			imag(viewPortMin),
			imag(viewPortMax),
		)
		scaledCoordinates = append(scaledCoordinates, complex(destinationScaledX, destinationScaledY))
	}
	return scaledCoordinates
}

// colorDestinationImage maps each transformed coordinate to a color
//   inside the Eyedropper Boundary.
func colorDestinationImage(
	destinationImage *image.NRGBA,
	sourceImage image.Image,
	destinationCoordinates []complex128,
	transformedCoordinates []complex128,
	colorValueBoundMin complex128,
	colorValueBoundMax complex128,
) {
	sourceImageBounds := sourceImage.Bounds()
	for index, transformedCoordinate := range transformedCoordinates {
		var sourceColorR, sourceColorG, sourceColorB, sourceColorA uint32

		if real(transformedCoordinate) < real(colorValueBoundMin) ||
			imag(transformedCoordinate) < imag(colorValueBoundMin) ||
			real(transformedCoordinate) > real(colorValueBoundMax) ||
			imag(transformedCoordinate) > imag(colorValueBoundMax) {
			sourceColorR, sourceColorG, sourceColorB, sourceColorA = 0, 0, 0, 0
		} else {
			sourceImagePixelX := int(mathutility.ScaleValueBetweenTwoRanges(
				float64(real(transformedCoordinate)),
				real(colorValueBoundMin),
				real(colorValueBoundMax),
				float64(sourceImageBounds.Min.X),
				float64(sourceImageBounds.Max.X),
			))
			sourceImagePixelY := int(mathutility.ScaleValueBetweenTwoRanges(
				float64(imag(transformedCoordinate)),
				imag(colorValueBoundMin),
				imag(colorValueBoundMax),
				float64(sourceImageBounds.Min.Y),
				float64(sourceImageBounds.Max.Y),
			))
			sourceColorR, sourceColorG, sourceColorB, sourceColorA = sourceImage.At(sourceImagePixelX, sourceImagePixelY).RGBA()
		}

		destinationPixelX := int(real(destinationCoordinates[index]))
		destinationPixelY := int(imag(destinationCoordinates[index]))

		destinationImage.Set(
			destinationPixelX,
			destinationPixelY,
			color.NRGBA{
				R: uint8(sourceColorR >> 8),
				G: uint8(sourceColorG >> 8),
				B: uint8(sourceColorB >> 8),
				A: uint8(sourceColorA >> 8),
			},
		)
	}
}

func checkSourceArgument(sourceImageFilename string) {
	if sourceImageFilename == "" {
		log.Fatal("missing source filename")
	}
}

func checkOutputArgument(outputFilename, outputDimensions string) (int, int) {
	if outputFilename == "" {
		log.Fatal("missing output filename")
	}

	outputWidth, widthErr := strconv.Atoi(strings.Split(outputDimensions, "x")[0])
	if widthErr != nil {
		log.Fatal(widthErr)
	}

	outputHeight, heightErr := strconv.Atoi(strings.Split(outputDimensions, "x")[1])
	if heightErr != nil {
		log.Fatal(heightErr)
	}

	return outputWidth, outputHeight
}

// FilenameArguments assume the user provides filenames to create a pattern.
type FilenameArguments struct {
	FormulaFilename     string
	OutputFilename      string
	OutputHeight        int
	OutputWidth         int
	SourceImageFilename string
}

func extractFilenameArguments() *FilenameArguments {
	var sourceImageFilename, outputFilename, outputDimensions string
	formulaFilename := "data/formula.yml"
	flag.StringVar(&formulaFilename, "f", "data/formula.yml", "See -formula")
	flag.StringVar(&formulaFilename, "formula", "data/formula.yml", "The filename of the formula file. Defaults to data/formula.yml")

	flag.StringVar(&sourceImageFilename, "in", "", "See -source. Required.")
	flag.StringVar(&sourceImageFilename, "source", "", "Source filename. Required.")

	flag.StringVar(&outputFilename, "out", "", "Output filename. Required.")
	outputDimensions = "200x200"
	flag.StringVar(&outputDimensions, "size", "200x200", "Output size in pixels, separated with an x. Default to 200x200.")
	flag.Parse()

	checkSourceArgument(sourceImageFilename)
	outputWidth, outputHeight := checkOutputArgument(outputFilename, outputDimensions)

	return &FilenameArguments{
		FormulaFilename:     formulaFilename,
		OutputFilename:      outputFilename,
		OutputHeight:        outputHeight,
		OutputWidth:         outputWidth,
		SourceImageFilename: sourceImageFilename,
	}
}


func helperForMapTransformedPointsToOutputImageBuffer(command *command.CreateSymmetryPattern, arguments *FilenameArguments, transformedCoordinates []complex128) *image.NRGBA {
	var err error
	colorSourceImage := openSourceImage(err, arguments)

	// TODO Rename CoordinateThreshold
	filter := imageoutput.CoordinateFilterFactory().
		WithMinimumX(command.CoordinateThreshold.XMin).
		WithMaximumX(command.CoordinateThreshold.XMax).
		WithMinimumY(command.CoordinateThreshold.YMin).
		WithMaximumY(command.CoordinateThreshold.YMax).
		Build()

	// TODO Read Eyedropper Sides from a formula file
	eyedropper := imageoutput.EyedropperFactory().
		WithLeftSide(colorSourceImage.Bounds().Min.X).
		WithRightSide(colorSourceImage.Bounds().Max.X).
		WithTopSide(colorSourceImage.Bounds().Min.Y).
		WithBottomSide(colorSourceImage.Bounds().Max.Y).
		WithImage(&colorSourceImage).
		Build()

	transformedCoordinateCollection := imageoutput.CoordinateCollectionFactory().
		WithComplexNumbers(&transformedCoordinates).
		Build()

	// TODO: Add a pivotal tracker story to make a new object to represent an image file.

	return MapTransformedPointsToOutputImageBuffer(eyedropper, transformedCoordinateCollection, arguments, filter)
}

// MapTransformedPointsToOutputImageBuffer Uses the transformed points, source image and eyedropper to return an output image buffer.
func MapTransformedPointsToOutputImageBuffer(eyedropper *imageoutput.Eyedropper, transformedCoordinates *imageoutput.CoordinateCollection, arguments *FilenameArguments, filter *imageoutput.CoordinateFilter) *image.NRGBA{
	filter.FilterAndMarkMappedCoordinateCollection(transformedCoordinates)

	colorData := eyedropper.ConvertCoordinatesToColors(transformedCoordinates)
	outputImage := image.NewNRGBA(image.Rect(0, 0, arguments.OutputWidth, arguments.OutputHeight))

	for index, colorToAdd := range *colorData {
		destinationPixelX := index % arguments.OutputWidth
		destinationPixelY := index / arguments.OutputWidth
		outputImage.Set(
			destinationPixelX,
			destinationPixelY,
			colorToAdd,
		)
	}
	return outputImage
}