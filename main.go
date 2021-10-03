package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/Chadius/creating-symmetry/entities/command"
	"github.com/Chadius/creating-symmetry/entities/formula/frieze"
	"github.com/Chadius/creating-symmetry/entities/formula/rosette"
	"github.com/Chadius/creating-symmetry/entities/formula/wallpaper"
	"github.com/Chadius/creating-symmetry/entities/mathutility"
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	formulaFilename := "data/formula.yml"
	flag.StringVar(&formulaFilename, "f", "data/formula.yml", "The filename of the formula file. Defaults to data/formula.yml")
	flag.Parse()

	createWallpaperYAML, err := ioutil.ReadFile(formulaFilename)
	if err != nil {
		log.Fatal(err)
	}
	wallpaperCommand, err := command.NewCreateWallpaperCommandFromYAML(createWallpaperYAML)
	if err != nil {
		log.Fatal(err)
	}
	sampleSpaceMin := complex(wallpaperCommand.PatternViewport.XMin, wallpaperCommand.PatternViewport.YMin)
	sampleSpaceMax := complex(wallpaperCommand.PatternViewport.XMax, wallpaperCommand.PatternViewport.YMax)
	outputWidth := wallpaperCommand.OutputImageSize.Width
	outputHeight := wallpaperCommand.OutputImageSize.Height
	colorSourceFilename := wallpaperCommand.SampleSourceFilename
	outputFilename := wallpaperCommand.OutputFilename
	colorValueBoundMin := complex(wallpaperCommand.EyedropperBoundary.XMin, wallpaperCommand.EyedropperBoundary.YMin)
	colorValueBoundMax := complex(wallpaperCommand.EyedropperBoundary.XMax, wallpaperCommand.EyedropperBoundary.YMax)

	reader, err := os.Open(colorSourceFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	colorSourceImage, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	destinationBounds := image.Rect(0, 0, outputWidth, outputHeight)
	destinationCoordinates := flattenCoordinates(destinationBounds)

	scaledCoordinates := scaleDestinationPixels(
		destinationBounds,
		destinationCoordinates,
		sampleSpaceMin,
		sampleSpaceMax,
	)

	transformedCoordinates := transformCoordinatesForFormula(wallpaperCommand, scaledCoordinates)
	minz, maxz := mathutility.GetBoundingBox(transformedCoordinates)
	println(minz)
	println(maxz)

	// Consider how to give a preview image? What's the picture ration
	outputImage := image.NewNRGBA(image.Rect(0, 0, outputWidth, outputHeight))
	colorDestinationImage(outputImage, colorSourceImage, destinationCoordinates, transformedCoordinates, colorValueBoundMin, colorValueBoundMax)

	outputToFile(outputFilename, outputImage)
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

func flattenCoordinates(destinationBounds image.Rectangle) []complex128 {
	flattenedCoordinates := []complex128{}
	for destinationY := destinationBounds.Min.Y; destinationY < destinationBounds.Max.Y; destinationY++ {
		for destinationX := destinationBounds.Min.X; destinationX < destinationBounds.Max.X; destinationX++ {
			flattenedCoordinates = append(flattenedCoordinates, complex(float64(destinationX), float64(destinationY)))
		}
	}
	return flattenedCoordinates
}

func scaleDestinationPixels(destinationBounds image.Rectangle, destinationCoordinates []complex128, viewPortMin complex128, viewPortMax complex128) []complex128 {
	scaledCoordinates := []complex128{}
	for _, destinationCoordinate := range destinationCoordinates {
		destinationScaledX := mathutility.ScaleValueBetweenTwoRanges(
			float64(real(destinationCoordinate)),
			float64(destinationBounds.Min.X),
			float64(destinationBounds.Max.X),
			real(viewPortMin),
			real(viewPortMax),
		)
		destinationScaledY := mathutility.ScaleValueBetweenTwoRanges(
			float64(imag(destinationCoordinate)),
			float64(destinationBounds.Min.Y),
			float64(destinationBounds.Max.Y),
			imag(viewPortMin),
			imag(viewPortMax),
		)
		scaledCoordinates = append(scaledCoordinates, complex(destinationScaledX, destinationScaledY))
	}
	return scaledCoordinates
}

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
