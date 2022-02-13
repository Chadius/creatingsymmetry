package imageoutput_test

import (
	creatingsymmetryfakes "github.com/chadius/creatingsymmetry/creatingsymmetryfakes"
	"github.com/chadius/creatingsymmetry/entities/imageoutput"
	. "gopkg.in/check.v1"
	"image"
	"image/color"
	_ "image/jpeg"
	"math"
)

type RectangularEyedropperTests struct {
}

var _ = Suite(&RectangularEyedropperTests{})

func (suite *RectangularEyedropperTests) TestCreateEyedropperWithBoundaries(checker *C) {
	eyedropper := imageoutput.EyedropperBuilder().WithLeftSide(-50).WithRightSide(200).WithTopSide(-100).WithBottomSide(400).Build()
	checker.Assert(eyedropper.LeftSide(), Equals, -50)
	checker.Assert(eyedropper.RightSide(), Equals, 200)
	checker.Assert(eyedropper.TopSide(), Equals, -100)
	checker.Assert(eyedropper.BottomSide(), Equals, 400)
}

func (suite *RectangularEyedropperTests) TestCreateEyedropperWithSourceImage(checker *C) {
	sourceImage := generate2x2ImageWithRedGreenBlueBlackPixels()
	eyedropper := imageoutput.EyedropperBuilder().WithImage(sourceImage).Build()

	checker.Assert(eyedropper.Image(), Equals, sourceImage)
}

func (suite *RectangularEyedropperTests) TestEyedropperMapsCoordinatesAndSamplesSourceImage(checker *C) {
	sourceImage := generate2x2ImageWithRedGreenBlueBlackPixels()

	coordinates := []*imageoutput.MappedCoordinate{
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(1, 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, 1),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(1, 1),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(2, 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, 2),
	}
	collection := imageoutput.CoordinateCollectionBuilder().WithCoordinates(&coordinates).Build()
	(*collection.Coordinates())[0].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[1].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[2].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[3].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[4].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[5].MarkAsSatisfyingFilter()

	eyedropper := imageoutput.EyedropperBuilder().WithLeftSide(0).WithRightSide(3).WithTopSide(0).WithBottomSide(3).WithImage(sourceImage).Build()
	convertedColors := eyedropper.ConvertCoordinatesToColors(collection)

	assertPixelHasExpectedColor := func(colorIndex int, atX, atY int) {
		var sourceR, sourceG, sourceB, sourceA uint32
		var outputR, outputG, outputB, outputA uint32

		outputR, outputG, outputB, outputA = (*convertedColors)[colorIndex].RGBA()
		sourceR, sourceG, sourceB, sourceA = sourceImage.At(atX, atY).RGBA()
		checker.Assert(outputR, Equals, sourceR)
		checker.Assert(outputG, Equals, sourceG)
		checker.Assert(outputB, Equals, sourceB)
		checker.Assert(outputA, Equals, sourceA)
	}

	assertPixelHasExpectedColor(0, 0, 0)
	assertPixelHasExpectedColor(1, 1, 0)
	assertPixelHasExpectedColor(2, 0, 1)
	assertPixelHasExpectedColor(3, 1, 1)

	assertPixelHasNoAlpha := func(colorIndex int) {
		var outputA uint32
		_, _, _, outputA = (*convertedColors)[colorIndex].RGBA()
		checker.Assert(outputA, Equals, uint32(0))
	}

	assertPixelHasNoAlpha(4)
	assertPixelHasNoAlpha(5)
}

func generate2x2ImageWithRedGreenBlueBlackPixels() image.Image {
	return &creatingsymmetryfakes.FakeImage{
		AtStub: func(x int, y int) color.Color {
			if x == 0 && y == 0 {
				return color.NRGBA{
					R: uint8(255),
					G: 0,
					B: 0,
					A: uint8(255),
				}
			} else if x == 1 && y == 0 {
				return color.NRGBA{
					R: 0,
					G: uint8(255),
					B: 0,
					A: uint8(255),
				}
			} else if x == 0 && y == 1 {
				return color.NRGBA{
					R: 0,
					G: 0,
					B: uint8(255),
					A: uint8(255),
				}
			} else if x == 1 && y == 1 {
				return color.NRGBA{
					R: 0,
					G: 0,
					B: 0,
					A: uint8(255),
				}
			} else {
				return color.NRGBA{
					R: 0,
					G: 0,
					B: 0,
					A: 0,
				}
			}
		},
	}
}

func (suite *RectangularEyedropperTests) TestEyedropperDoesNotMapInvalidCoordinates(checker *C) {
	sourceImage := generate2x2ImageWithRedGreenBlueBlackPixels()

	coordinates := []*imageoutput.MappedCoordinate{
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(1, 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, 1),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(1, 1),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, 0),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(0, math.Inf(1)),
		imageoutput.NewMappedCoordinateUsingTransformedCoordinates(math.Inf(-1), 0),
	}
	collection := imageoutput.CoordinateCollectionBuilder().WithCoordinates(&coordinates).Build()
	(*collection.Coordinates())[0].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[1].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[2].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[3].MarkAsSatisfyingFilter()

	(*collection.Coordinates())[5].MarkAsSatisfyingFilter()
	(*collection.Coordinates())[6].MarkAsSatisfyingFilter()

	eyedropper := imageoutput.EyedropperBuilder().WithLeftSide(0).WithRightSide(2).WithTopSide(0).WithBottomSide(2).WithImage(sourceImage).Build()
	convertedColors := eyedropper.ConvertCoordinatesToColors(collection)

	assertPixelHasNoAlpha := func(colorIndex int) {
		var outputA uint32
		_, _, _, outputA = (*convertedColors)[colorIndex].RGBA()
		checker.Assert(outputA, Equals, uint32(0))
	}

	assertPixelHasNoAlpha(4)
	assertPixelHasNoAlpha(5)
	assertPixelHasNoAlpha(6)
}
