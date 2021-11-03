package main

import (
	"flag"
	"github.com/Chadius/creating-symmetry/entities/command"
	"github.com/Chadius/creating-symmetry/entities/imageoutput"
	transformerEntity "github.com/Chadius/creating-symmetry/entities/transformer"
	"image"
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
	wallpaperCommand := loadFormulaFile(filenameArguments)

	transformImage(filenameArguments, wallpaperCommand)
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

func checkSourceArgument(sourceImageFilename string) {
	if sourceImageFilename == "" {
		log.Fatal("missing source filename")
	}
}

func loadFormulaFile(filenameArguments *FilenameArguments) *command.CreateSymmetryPattern {
	createWallpaperYAML, err := ioutil.ReadFile(filenameArguments.FormulaFilename)
	if err != nil {
		log.Fatal(err)
	}
	wallpaperCommand, err := command.NewCreateWallpaperCommandFromYAML(createWallpaperYAML)
	if err != nil {
		log.Fatal(err)
	}
	return wallpaperCommand
}

func transformImage(filenameArguments *FilenameArguments, wallpaperCommand *command.CreateSymmetryPattern) {
	sourceImage := openSourceImage(filenameArguments)
	coordinateThreshold := imageoutput.CoordinateFilterBuilder().
		WithMinimumX(wallpaperCommand.CoordinateThreshold.XMin).
		WithMaximumX(wallpaperCommand.CoordinateThreshold.XMax).
		WithMinimumY(wallpaperCommand.CoordinateThreshold.YMin).
		WithMaximumY(wallpaperCommand.CoordinateThreshold.YMax).
		Build()

	var eyedropper *imageoutput.RectangularEyedropper
	if wallpaperCommand.Eyedropper != nil {
		eyedropper = imageoutput.EyedropperBuilder().
			WithLeftSide(wallpaperCommand.Eyedropper.LeftSide).
			WithRightSide(wallpaperCommand.Eyedropper.RightSide).
			WithTopSide(wallpaperCommand.Eyedropper.TopSide).
			WithBottomSide(wallpaperCommand.Eyedropper.BottomSide).
			WithImage(&sourceImage).
			Build()
	} else {
		eyedropper = imageoutput.EyedropperBuilder().
			WithLeftSide(sourceImage.Bounds().Min.X).
			WithRightSide(sourceImage.Bounds().Max.X).
			WithTopSide(sourceImage.Bounds().Min.Y).
			WithBottomSide(sourceImage.Bounds().Max.Y).
			WithImage(&sourceImage).
			Build()
	}

	transformer := transformerEntity.FormulaTransformer{}

	outputImage := transformer.Transform(&transformerEntity.Settings{
		PatternViewportXMin: wallpaperCommand.PatternViewport.XMin,
		PatternViewportXMax: wallpaperCommand.PatternViewport.XMax,
		PatternViewportYMin: wallpaperCommand.PatternViewport.YMin,
		PatternViewportYMax: wallpaperCommand.PatternViewport.YMax,
		InputImage:          sourceImage,
		Formula:             wallpaperCommand,
		CoordinateThreshold: coordinateThreshold,
		Eyedropper:          eyedropper,
		OutputWidth:         filenameArguments.OutputWidth,
		OutputHeight:        filenameArguments.OutputHeight,
	})
	outputToFile(filenameArguments.OutputFilename, outputImage)
}

func openSourceImage(filenameArguments *FilenameArguments) image.Image {
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

