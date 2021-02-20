package main

import (
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"log"
	"math/cmplx"
	"os"
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
	destinationImage := image.NewNRGBA(image.Rect(0, 0, 200, 200))
	//destinationImage := image.NewNRGBA(image.Rect(0, 0, 2160, 2160))
	destinationBounds := destinationImage.Bounds()

	for destinationY := destinationBounds.Min.Y ; destinationY < destinationBounds.Max.Y; destinationY++ {
		for destinationX := destinationBounds.Min.X ; destinationX < destinationBounds.Max.X; destinationX++ {
			interpolatedDestinationCoordinates := interpolateCoordinatesToRange(destinationX, destinationY, destinationImage)
			r, g, b, a := calculateDestinationColor(
				interpolatedDestinationCoordinates,
				sourceImage,
				).RGBA()
			destinationImage.Set(
				destinationX,
				destinationY,
				color.NRGBA{
					R: uint8(r>>8),
					G: uint8(g>>8),
					B: uint8(b>>8),
					A: uint8(a>>8),
				},
			)
		}
	}

	outFile, err := os.Create("exampleImage/newBrownie.png")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	png.Encode(outFile, destinationImage)
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

	sourceColorVector += addRaisedAndScaledVector(destinationVector, 5, 0, complex(1, 0))
	sourceColorVector += addRaisedAndScaledVector(destinationVector, 6, 1, complex(-0.5, -0))
	sourceColorVector += addRaisedAndScaledVector(destinationVector, 4, -6, complex(-1.5, 0))

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