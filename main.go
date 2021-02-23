package main

import (
	"image"
	"image/color"
	"image/png"
	"wallpaper/entities/formula"

	//"image/png"
	_ "image/png"
	"log"
	"os"
	"wallpaper/entities/mathutility"
)

func main() {
	reader, err := os.Open("exampleImage/brownie.png")
	if err != nil {
	    log.Fatal(err)
	}
	defer reader.Close()

	sourceImage, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}

	// Consider how to give a preview image? What's the picture ration
	destinationImage := image.NewNRGBA(image.Rect(0, 0, 800, 450))
	//destinationImage := image.NewNRGBA(image.Rect(0, 0, 3840, 2160))
	destinationBounds := destinationImage.Bounds()

	destinationCoordinates := flattenCoordinates(destinationBounds)

	size := float64(1)
	scaledCoordinates := scaleDestinationPixels(
		destinationBounds,
		destinationCoordinates,
		complex(-1.45 * size, -1 * size),
		complex(1.45 * size, size),
	)
	transformedCoordinates := transformCoordinates(scaledCoordinates)

	colorDestinationImage(destinationImage, sourceImage, destinationCoordinates, transformedCoordinates)
	outFile, err := os.Create("exampleImage/newBrownie.png")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	png.Encode(outFile, destinationImage)
}

func transformCoordinates(scaledCoordinates []complex128) []complex128 {
	form := formula.SymmetryFormula{
		PairedCoefficients : []*formula.CoefficientPairs{
			{
				Scale: complex(1, 0),
				PowerN: 5,
				PowerM: 0,
			},
			{
				Scale: complex(1, 0),
				PowerN: 0,
				PowerM: 5,
			},
			{
				Scale: complex(-0.85, 1),
				PowerN: 6,
				PowerM: 1,
			},
			{
				Scale: complex(-0.85, 1),
				PowerN: 1,
				PowerM: 6,
			},
			{
				Scale: complex(0, 1),
				PowerN: 4,
				PowerM: -6,
			},
			{
				Scale: complex(0, 1),
				PowerN: -6,
				PowerM: 4,
			},
		},
	}

	transformedCoordinates := []complex128{}
	for _, complexCoordinate := range scaledCoordinates {
		transformedCoordinate := form.Calculate(complexCoordinate)
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

func colorDestinationImage(destinationImage *image.NRGBA, sourceImage image.Image, destinationCoordinates []complex128, transformedCoordinates []complex128) {
	transformedBoundingBoxMin, transformedBoundingBoxMax := mathutility.GetBoundingBox(transformedCoordinates)
	sourceImageBounds := sourceImage.Bounds()
	for index, transformedCoordinate := range transformedCoordinates {
		sourceImagePixelX := int(mathutility.ScaleValueBetweenTwoRanges(
			float64(real(transformedCoordinate)),
			real(transformedBoundingBoxMin),
			real(transformedBoundingBoxMax),
			float64(sourceImageBounds.Min.X),
			float64(sourceImageBounds.Max.X),
		))
		sourceImagePixelY := int(mathutility.ScaleValueBetweenTwoRanges(
			float64(imag(transformedCoordinate)),
			imag(transformedBoundingBoxMin),
			imag(transformedBoundingBoxMax),
			float64(sourceImageBounds.Min.Y),
			float64(sourceImageBounds.Max.Y),
		))

		sourceColorR,sourceColorG,sourceColorB,sourceColorA := sourceImage.At(sourceImagePixelX, sourceImagePixelY).RGBA()

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
