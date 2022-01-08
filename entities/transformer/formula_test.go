package transformer_test

import (
	creatingsymmetryfakes "github.com/Chadius/creating-symmetry/creating-symmetryfakes"
	"github.com/Chadius/creating-symmetry/entities/command"
	"github.com/Chadius/creating-symmetry/entities/formula"
	"github.com/Chadius/creating-symmetry/entities/imageoutput/imageoutputfakes"
	transformerEntity "github.com/Chadius/creating-symmetry/entities/transformer"
	. "gopkg.in/check.v1"
	"image"
	"image/color"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type FormulaTests struct {
	sourceImage             image.Image
	commandShouldBeAFormula *command.CreateSymmetryPattern
}

var _ = Suite(&FormulaTests{})

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

func (suite *FormulaTests) SetUpTest(checker *C) {
	suite.sourceImage = generate2x2ImageWithRedGreenBlueBlackPixels()

	rosetteFormula, _ := formula.NewBuilder().
		Rosette().
		AddTerm(
			formula.NewTermBuilder().
				PowerN(2).
				PowerM(1).
				Build(),
		).
		Build()

	suite.commandShouldBeAFormula = &command.CreateSymmetryPattern{
		PatternViewport:     command.ComplexNumberCorners{},
		CoordinateThreshold: command.ComplexNumberCorners{},
		Eyedropper:          nil,
		Formula:             rosetteFormula,
	}
}

func (suite *FormulaTests) TestTransformerCallsThresholdAndEyedropper(checker *C) {
	mockCoordinateThreshold := imageoutputfakes.FakeCoordinateThreshold{}
	mockEyedropper := imageoutputfakes.FakeEyedropper{}
	mockEyedropper.ConvertCoordinatesToColorsReturns(&[]color.Color{})
	transformer := transformerEntity.FormulaTransformer{}

	transformer.Transform(&transformerEntity.Settings{
		PatternViewportXMin: 0,
		PatternViewportXMax: 1,
		PatternViewportYMin: 0,
		PatternViewportYMax: 1,
		InputImage:          suite.sourceImage,
		Formula:             suite.commandShouldBeAFormula,
		CoordinateThreshold: &mockCoordinateThreshold,
		Eyedropper:          &mockEyedropper,
		OutputWidth:         1,
		OutputHeight:        1,
	})

	checker.Assert(mockCoordinateThreshold.FilterAndMarkMappedCoordinateCollectionCallCount(), Equals, 1)
	checker.Assert(mockEyedropper.ConvertCoordinatesToColorsCallCount(), Not(Equals), 0)
}

func (suite *FormulaTests) TestTransformerOutputsToImageOfGivenSize(checker *C) {
	mockCoordinateThreshold := imageoutputfakes.FakeCoordinateThreshold{}
	mockEyedropper := imageoutputfakes.FakeEyedropper{}
	mockEyedropper.ConvertCoordinatesToColorsReturns(&[]color.Color{})
	transformer := transformerEntity.FormulaTransformer{}

	outputImage := transformer.Transform(&transformerEntity.Settings{
		PatternViewportXMin: 0,
		PatternViewportXMax: 1,
		PatternViewportYMin: 0,
		PatternViewportYMax: 1,
		InputImage:          suite.sourceImage,
		Formula:             suite.commandShouldBeAFormula,
		CoordinateThreshold: &mockCoordinateThreshold,
		Eyedropper:          &mockEyedropper,
		OutputWidth:         3,
		OutputHeight:        1,
	})

	checker.Assert(outputImage.Bounds().Max.X, Equals, 3)
	checker.Assert(outputImage.Bounds().Max.Y, Equals, 1)
}
