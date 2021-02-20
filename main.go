package main

import (
	"image"
	"image/color"
	"image/png"
	_ "image/png"
	"log"
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
	destinationImage := image.NewRGBA(image.Rect(0, 0, bounds.Size().X, bounds.Size().Y))

	for sourceY := bounds.Min.Y ; sourceY < bounds.Max.Y; sourceY++ {
		for sourceX := bounds.Min.X ; sourceX < bounds.Max.X; sourceX++ {
			r, g, b, a := sourceImage.At(sourceX, sourceY).RGBA()
			destinationImage.Set(
				sourceX,
				sourceY,
				color.NRGBA{
					R: uint8(g>>8),
					G: uint8(b>>8),
					B: uint8(r>>8),
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