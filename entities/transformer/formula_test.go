package transformer_test

import (
	"github.com/Chadius/creating-symmetry/entities/command"
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	"github.com/Chadius/creating-symmetry/entities/formula/exponential"
	"github.com/Chadius/creating-symmetry/entities/formula/rosette"
	"github.com/Chadius/creating-symmetry/entities/imageoutput"
	transformerEntity "github.com/Chadius/creating-symmetry/entities/transformer"
	. "gopkg.in/check.v1"
	"image"
	"image/color"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type FormulaTests struct {
	sourceImage             image.Image
	mockCoordinateThreshold MockCoordinateThreshold
	mockEyedropper          MockEyedropper
	commandShouldBeAFormula *command.CreateSymmetryPattern
}

type MockCoordinateThreshold struct {
	MockWasCalled bool
}

func (m *MockCoordinateThreshold) FilterAndMarkMappedCoordinateCollection(collection *imageoutput.CoordinateCollection) {
	m.MockWasCalled = true
}

type MockEyedropper struct {
	MockWasCalled bool
}

func (m *MockEyedropper) ConvertCoordinatesToColors(collection *imageoutput.CoordinateCollection) *[]color.Color {
	m.MockWasCalled = true
	return &[]color.Color{}
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

	rosetteFormula := rosette.Formula{
		Terms: []*exponential.RosetteFriezeTerm{
			{
				Multiplier:             complex(3, 0),
				PowerN:                 1,
				PowerM:                 0,
				IgnoreComplexConjugate: false,
				CoefficientRelationships: []coefficient.Relationship{
					coefficient.PlusMPlusN,
				},
			},
		},
	}

	suite.commandShouldBeAFormula = &command.CreateSymmetryPattern{
		PatternViewport:     command.ComplexNumberCorners{},
		CoordinateThreshold: command.ComplexNumberCorners{},
		Eyedropper:          nil,
		RosetteFormula:      &rosetteFormula,
		FriezeFormula:       nil,
		LatticePattern:      nil,
	}

	suite.mockCoordinateThreshold = MockCoordinateThreshold{}
	suite.mockEyedropper = MockEyedropper{}
}

func (suite *FormulaTests) TestTransformerCallsThresholdAndEyedropper(checker *C) {
	transformer := transformerEntity.FormulaTransformer{}

	transformer.Transform(&transformerEntity.Settings{
		PatternViewportXMin: 0,
		PatternViewportXMax: 1,
		PatternViewportYMin: 0,
		PatternViewportYMax: 1,
		InputImage:          suite.sourceImage,
		Formula:             suite.commandShouldBeAFormula,
		CoordinateThreshold: &suite.mockCoordinateThreshold,
		Eyedropper:          &suite.mockEyedropper,
		OutputWidth:         1,
		OutputHeight:        1,
	})

	checker.Assert(suite.mockCoordinateThreshold.MockWasCalled, Equals, true)
	checker.Assert(suite.mockEyedropper.MockWasCalled, Equals, true)
}

func (suite *FormulaTests) TestTransformerOutputsToImageOfGivenSize(checker *C) {
	transformer := transformerEntity.FormulaTransformer{}

	outputImage := transformer.Transform(&transformerEntity.Settings{
		PatternViewportXMin: 0,
		PatternViewportXMax: 1,
		PatternViewportYMin: 0,
		PatternViewportYMax: 1,
		InputImage:          suite.sourceImage,
		Formula:             suite.commandShouldBeAFormula,
		CoordinateThreshold: &suite.mockCoordinateThreshold,
		Eyedropper:          &suite.mockEyedropper,
		OutputWidth:         3,
		OutputHeight:        1,
	})

	checker.Assert(outputImage.Bounds().Max.X, Equals, 3)
	checker.Assert(outputImage.Bounds().Max.Y, Equals, 1)
}
