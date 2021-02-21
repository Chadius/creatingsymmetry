package main

import (
	"image"
	"image/color"
	"image/png"

	//"image/png"
	_ "image/png"
	"log"
	"math/cmplx"
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
	//destinationImage := image.NewNRGBA(image.Rect(0, 0, 400, 400))
	destinationImage := image.NewNRGBA(image.Rect(0, 0, 2160, 2160))
	destinationBounds := destinationImage.Bounds()

	destinationCoordinates := flattenCoordinates(destinationBounds)

	size := float64(9000)
	scaledCoordinates := scaleDestinationPixels(
		destinationBounds,
		destinationCoordinates,
		complex(-1 * size, -1 * size),
		complex(size, size),
		//complex(-100, -100),
		//complex(100, 100),
		//complex(-100000, -100000),
		//complex(100000, 100000),
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
	transformedCoordinates := []complex128{}
	for _, complexCoordinate := range scaledCoordinates {
		transformedCoordinate := complex(0, 0)
		transformedCoordinate += addRaisedAndScaledVector(complexCoordinate, 5, 0, complex(1, 0))
		transformedCoordinate += addRaisedAndScaledVector(complexCoordinate, 6, 1, complex(-9e-4, -1e-0))
		transformedCoordinate += addRaisedAndScaledVector(complexCoordinate, 4, -6, complex(-9e-30, 0))
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

func interpolateCoordinatesToRange(x, y int, picture image.Image) complex128 {
	xPercent := float64(x) / float64(picture.Bounds().Dx()) - 0.5
	yPercent := float64(y) / float64(picture.Bounds().Dy()) - 0.5
	return complex(xPercent, yPercent)
}

func interpolateRangeToCoordinates(point complex128, picture image.Image) image.Point {
	xCoordinate := int(float64(picture.Bounds().Dx()) * (real(point) + 0.5))
	yCoordinate := int(float64(picture.Bounds().Dy()) * (imag(point) + 0.5))
	return image.Point{X: xCoordinate, Y:yCoordinate,}
}

func addRaisedAndScaledVector(sourceVector complex128, powerN, powerM int, scale complex128) complex128 {
	zRaised := cmplx.Pow(sourceVector, complex(float64(powerN - powerM), 0))
	zRaised += cmplx.Pow(sourceVector, complex(float64(powerM - powerN), 0))
	return zRaised * scale
}

func calculateDestinationColor(destinationVector complex128, sourceImage image.Image) color.Color {
	sourceColorVector := complex(0,0)

	sourceColorVector += addRaisedAndScaledVector(destinationVector, 5, 0, complex(0.000001, 0))
	//sourceColorVector += addRaisedAndScaledVector(destinationVector, 6, 1, complex(-0.5, -0))
	//sourceColorVector += addRaisedAndScaledVector(destinationVector, 4, -6, complex(-1.5, 0))

	sourcePoint := interpolateRangeToCoordinates(
		sourceColorVector,
		sourceImage)

	sourceX := sourcePoint.X
	sourceY := sourcePoint.Y

	bounds := sourceImage.Bounds()

	if (sourceX < 0 || sourceX > bounds.Dx() || sourceY < 0 || sourceY >= bounds.Dy()) {
		return color.NRGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 255,
		}
	}
	return sourceImage.At(int(sourceX), int(sourceY))
}