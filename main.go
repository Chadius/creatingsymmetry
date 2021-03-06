package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"wallpaper/entities/formula"

	//"image/png"
	_ "image/png"
	"os"
	"wallpaper/entities/mathutility"
)

func main() {
	sampleSpaceMin := complex(-5e0, -5e0)
	sampleSpaceMax := complex(5e0, 5e0)
	outputWidth := 800
	outputHeight := 450
	//outputWidth := 3840
	//outputHeight := 2160
	colorSourceFilename := "exampleImage/brownie.png"
	outputFilename := "exampleImage/newBrownie.png"
	colorValueBoundMin := complex(-10e0, -3e1)
	colorValueBoundMax := complex(10e0, 3e1)

	reader, err := os.Open(colorSourceFilename)
	if err != nil {
	  log.Fatal(err)
	}
	defer reader.Close()

	colorSourceImage, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	destinationBounds := image.Rect(0,0, outputWidth, outputHeight)
	destinationCoordinates := flattenCoordinates(destinationBounds)

	scaledCoordinates := scaleDestinationPixels(
		destinationBounds,
		destinationCoordinates,
		sampleSpaceMin,
		sampleSpaceMax,
	)

	transformedCoordinates := transformCoordinates(scaledCoordinates)
	minz, maxz := mathutility.GetBoundingBox(transformedCoordinates)
	println(minz)
	println(maxz)

	// Consider how to give a preview image? What's the picture ration
	outputImage := image.NewNRGBA(image.Rect(0, 0, outputWidth, outputHeight))
	colorDestinationImage(outputImage, colorSourceImage, destinationCoordinates, transformedCoordinates, colorValueBoundMin, colorValueBoundMax)

	outputToFile(outputFilename, outputImage)
}

func outputToFile(outputFilename string, outputImage image.Image) {
	outputImageFile, err := os.Create(outputFilename)
	if err != nil {
		panic(err)
	}
	defer outputImageFile.Close()
	png.Encode(outputImageFile, outputImage)
}

func transformCoordinates(scaledCoordinates []complex128) []complex128 {
	friezeFormula := formula.FriezeFormula{
		Elements: []*formula.EulerFormulaElement{
			{
				Scale:                  complex(1e0, 0e2),
				PowerN:                 2,
				PowerM:                 -1,
				IgnoreComplexConjugate: false,
				CoefficientPairs: formula.LockedCoefficientPair{
					Multiplier: 1,
					OtherCoefficientRelationships: []formula.CoefficientRelationship{
						formula.MinusMMinusNMaybeFlipScale,
					},
				},
			},
		},
	}

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
	for _, complexCoordinate := range scaledCoordinates {
		transformedCoordinate := friezeFormula.Calculate(complexCoordinate)
		transformedCoordinates = append(transformedCoordinates, transformedCoordinate)
	}
	return transformedCoordinates
}

func flattenCoordinates(destinationBounds image.Rectangle) []complex128 {
	flattenedCoordinates := []complex128{}
	for destinationY := destinationBounds.Min.Y ; destinationY < destinationBounds.Max.Y; destinationY++ {
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
			sourceColorR,sourceColorG,sourceColorB,sourceColorA = 0,0,0,0
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
				R: uint8(sourceColorR>>8),
				G: uint8(sourceColorG>>8),
				B: uint8(sourceColorB>>8),
				A: uint8(sourceColorA>>8),
			},
		)
	}
}
