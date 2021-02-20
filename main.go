package main

import (
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"log"
	"math"
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
	bounds := sourceImage.Bounds()
	destinationImage := image.NewNRGBA(image.Rect(0, 0, bounds.Size().X * 2, bounds.Size().Y * 2))
	destinationBounds := destinationImage.Bounds()

	for destinationY := destinationBounds.Min.Y ; destinationY < destinationBounds.Max.Y; destinationY++ {
		for destinationX := destinationBounds.Min.X ; destinationX < destinationBounds.Max.X; destinationX++ {
			r, g, b, a := calculateDestinationColor(destinationX, destinationY, sourceImage).RGBA()
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

func addFunctionPart (sourceVector complex128, powerN, powerM int, scale float64) complex128 {
	radius := math.Pow(real(sourceVector), float64(powerN - powerM)) * scale
	angle := cmplx.Exp(complex(0, imag(sourceVector) * float64(powerN - powerM)))
	return complex(radius, 0) * angle
}

func addFunctionPart2 (sourceVector complex128, powerN, powerM int, scale float64) complex128 {
	zRaised := cmplx.Pow(sourceVector, complex(float64(powerN - powerM), 0))
	return complex(real(zRaised) * scale, imag(zRaised) * scale)
}

func calculateDestinationColor(destinationX, destinationY int, sourceImage image.Image) color.Color {
	destinationVector := complex(float64(destinationX), float64(destinationY))
	destinationPoint := addFunctionPart2(destinationVector, 4, 2, 0.001)
	destinationPoint += addFunctionPart2(destinationVector, 2, 4, 1000)

	sourceX := real(destinationPoint)
	sourceY := imag(destinationPoint)

	bounds := sourceImage.Bounds()
	if (sourceX < 0 || sourceX > float64(bounds.Dx()) || sourceY < 0 || sourceY >= float64(bounds.Dy())) {
		return color.NRGBA{
			R: 0,
			G: 0,
			B: 0,
			A: 255,
		}
	}
	return sourceImage.At(int(sourceX), int(sourceY))
}