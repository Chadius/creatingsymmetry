package transformer_test

import (
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

func (suite *FormulaTests) SetUpTest(checker *C) {
	sourceColors := image.NewNRGBA(image.Rect(0, 0, 2, 2))
	sourceColors.Set(0, 0, color.NRGBA{
		R: uint8(255),
		G: 0,
		B: 0,
		A: uint8(255),
	})
	sourceColors.Set(1, 0, color.NRGBA{
		R: 0,
		G: uint8(255),
		B: 0,
		A: uint8(255),
	})
	sourceColors.Set(0, 1, color.NRGBA{
		R: 0,
		G: 0,
		B: uint8(255),
		A: uint8(255),
	})
	sourceColors.Set(1, 1, color.NRGBA{
		R: 0,
		G: 0,
		B: 0,
		A: uint8(255),
	})
	suite.sourceImage = sourceColors.SubImage(image.Rect(0, 0, 2, 2))

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
