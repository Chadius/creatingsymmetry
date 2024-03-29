package formula_test

import (
	"github.com/chadius/creatingsymmetry/entities/formula"
	"github.com/chadius/creatingsymmetry/entities/formula/coefficient"
	"github.com/chadius/creatingsymmetry/entities/utility"
	. "gopkg.in/check.v1"
	"math"
	"math/cmplx"
)

type WavePacketBuilderTest struct{}

var _ = Suite(&WavePacketBuilderTest{})

func (suite *WavePacketBuilderTest) TestCreateWavePackets(checker *C) {
	newWavePacket := formula.NewWavePacketBuilder().
		Multiplier(complex(2e-3, -5e7)).
		AddTerm(
			formula.NewTermBuilder().PowerN(1).PowerM(-2).Build(),
		).
		AddTerm(
			formula.NewTermBuilder().PowerN(-2).PowerM(1).Build(),
		).
		AddTerm(
			formula.NewTermBuilder().PowerN(1).PowerM(1).Build(),
		).
		Build()

	checker.Assert(real(newWavePacket.Multiplier()), utility.NumericallyCloseEnough{}, 2e-3, 1e-6)
	checker.Assert(imag(newWavePacket.Multiplier()), utility.NumericallyCloseEnough{}, -5e7, 1e-6)
	checker.Assert(newWavePacket.Terms(), HasLen, 3)
	checker.Assert(newWavePacket.Terms()[0].PowerN, Equals, 1)
	checker.Assert(newWavePacket.Terms()[0].PowerM, Equals, -2)
	checker.Assert(newWavePacket.Terms()[1].PowerN, Equals, -2)
	checker.Assert(newWavePacket.Terms()[1].PowerM, Equals, 1)
	checker.Assert(newWavePacket.Terms()[2].PowerN, Equals, 1)
	checker.Assert(newWavePacket.Terms()[2].PowerM, Equals, 1)
}

type WaveFormulaTests struct {
	hexLatticeVectors   []complex128
	hexagonalWavePacket *formula.WavePacket
}

var _ = Suite(&WaveFormulaTests{})

func (suite *WaveFormulaTests) SetUpTest(checker *C) {
	suite.hexLatticeVectors = []complex128{
		complex(1, 0),
		complex(-0.5, math.Sqrt(3.0)/2.0),
	}

	suite.hexagonalWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(1, 0)).
		AddTerm(
			formula.NewTermBuilder().PowerN(1).PowerM(-2).Build(),
		).
		AddTerm(
			formula.NewTermBuilder().PowerN(-2).PowerM(1).Build(),
		).
		AddTerm(
			formula.NewTermBuilder().PowerN(1).PowerM(1).Build(),
		).
		Build()
}

func (suite *WaveFormulaTests) TestWaveFormulaCombinesEisensteinTerms(checker *C) {
	zInLatticeCoordinates := formula.ConvertToLatticeCoordinates(
		complex(math.Sqrt(3), -1*math.Sqrt(3)),
		suite.hexLatticeVectors,
	)

	calculation := suite.hexagonalWavePacket.Calculate(zInLatticeCoordinates)

	expectedAnswer := cmplx.Exp(complex(0, 2*math.Pi*(3+math.Sqrt(3)))) +
		cmplx.Exp(complex(0, 2*math.Pi*(-2*math.Sqrt(3)))) +
		cmplx.Exp(complex(0, 2*math.Pi*(-3+math.Sqrt(3))))

	checker.Assert(real(calculation), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(calculation), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}

func (suite *WaveFormulaTests) TestWaveFormulaUsesMultiplier(checker *C) {
	hexagonalWavePacketWithNewMultiplier := formula.NewWavePacketBuilder().
		Multiplier(complex(1/3.0, 0)).
		AddTerm(&suite.hexagonalWavePacket.Terms()[0]).
		AddTerm(&suite.hexagonalWavePacket.Terms()[1]).
		AddTerm(&suite.hexagonalWavePacket.Terms()[2]).
		Build()

	zInLatticeCoordinates := formula.ConvertToLatticeCoordinates(
		complex(math.Sqrt(3), -1*math.Sqrt(3)),
		suite.hexLatticeVectors,
	)

	calculation := hexagonalWavePacketWithNewMultiplier.Calculate(zInLatticeCoordinates)

	expectedAnswer := (cmplx.Exp(complex(0, 2*math.Pi*(3+math.Sqrt(3)))) +
		cmplx.Exp(complex(0, 2*math.Pi*(-2*math.Sqrt(3)))) +
		cmplx.Exp(complex(0, 2*math.Pi*(-3+math.Sqrt(3))))) / 3

	checker.Assert(real(calculation), utility.NumericallyCloseEnough{}, real(expectedAnswer), 1e-6)
	checker.Assert(imag(calculation), utility.NumericallyCloseEnough{}, imag(expectedAnswer), 1e-6)
}

type WavePacketRelationshipTest struct {
	aPlusNPlusMOddWavePacket         *formula.WavePacket
	aPlusMMinusNOddWavePacket        *formula.WavePacket
	aPlusMPlusNOddWavePacket         *formula.WavePacket
	aMinusNMinusMOddWavePacket       *formula.WavePacket
	aMinusMMinusNOddWavePacket       *formula.WavePacket
	aMinusMPlusNOddWavePacket        *formula.WavePacket
	aPlusMPlusNOddNegatedWavePacket  *formula.WavePacket
	aMinusSumNAndMPlusNOddWavePacket *formula.WavePacket

	aPlusNPlusMEvenWavePacket          *formula.WavePacket
	aPlusMPlusNEvenWavePacket          *formula.WavePacket
	aPlusMPlusNEvenNegatedWavePacket   *formula.WavePacket
	aMinusMMinusNOddNegatedWavePacket  *formula.WavePacket
	aMinusMMinusNEvenWavePacket        *formula.WavePacket
	aMinusMMinusNEvenNegatedWavePacket *formula.WavePacket
	aPlusMMinusSumNAndMOddWavePacket   *formula.WavePacket
	aMinusSumNAndMPlusNWavePacket      *formula.WavePacket
}

var _ = Suite(&WavePacketRelationshipTest{})

func (suite *WavePacketRelationshipTest) SetUpTest(checker *C) {
	suite.aPlusNPlusMOddWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(2, 1)).
		AddTerm(formula.NewTermBuilder().PowerN(-1).PowerM(4).Build()).
		Build()
	suite.aPlusMPlusNOddWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(2, 1)).
		AddTerm(formula.NewTermBuilder().PowerN(4).PowerM(-1).Build()).
		Build()
	suite.aMinusNMinusMOddWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(2, 1)).
		AddTerm(formula.NewTermBuilder().PowerN(1).PowerM(-4).Build()).
		Build()
	suite.aMinusMMinusNOddWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(2, 1)).
		AddTerm(formula.NewTermBuilder().PowerN(-4).PowerM(1).Build()).
		Build()
	suite.aPlusMMinusNOddWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(2, 1)).
		AddTerm(formula.NewTermBuilder().PowerN(4).PowerM(1).Build()).
		Build()
	suite.aMinusMPlusNOddWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(2, 1)).
		AddTerm(formula.NewTermBuilder().PowerN(-4).PowerM(-1).Build()).
		Build()
	suite.aPlusMPlusNOddNegatedWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(-2, -1)).
		AddTerm(formula.NewTermBuilder().PowerN(4).PowerM(-1).Build()).
		Build()
	suite.aMinusMMinusNOddNegatedWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(-2, -1)).
		AddTerm(formula.NewTermBuilder().PowerN(-4).PowerM(1).Build()).
		Build()
	suite.aPlusMMinusSumNAndMOddWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(2, 1)).
		AddTerm(formula.NewTermBuilder().PowerN(4).PowerM(-3).Build()).
		Build()
	suite.aMinusSumNAndMPlusNWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(2, 1)).
		AddTerm(formula.NewTermBuilder().PowerN(-3).PowerM(-1).Build()).
		Build()
	suite.aMinusSumNAndMPlusNOddWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(2, 1)).
		AddTerm(formula.NewTermBuilder().PowerN(-3).PowerM(-1).Build()).
		Build()
	suite.aPlusNPlusMEvenWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(2, 1)).
		AddTerm(formula.NewTermBuilder().PowerN(-6).PowerM(2).Build()).
		Build()
	suite.aPlusMPlusNEvenWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(2, 1)).
		AddTerm(formula.NewTermBuilder().PowerN(2).PowerM(-6).Build()).
		Build()
	suite.aPlusMPlusNEvenNegatedWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(-2, -1)).
		AddTerm(formula.NewTermBuilder().PowerN(2).PowerM(-6).Build()).
		Build()
	suite.aMinusMMinusNEvenWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(2, 1)).
		AddTerm(formula.NewTermBuilder().PowerN(-2).PowerM(6).Build()).
		Build()
	suite.aMinusMMinusNEvenNegatedWavePacket = formula.NewWavePacketBuilder().
		Multiplier(complex(-2, -1)).
		AddTerm(formula.NewTermBuilder().PowerN(-2).PowerM(6).Build()).
		Build()
}

func (suite *WavePacketRelationshipTest) TestPlusNPlusM(checker *C) {
	relationshipsFound := formula.GetWavePacketRelationship(
		*suite.aPlusNPlusMOddWavePacket,
		*suite.aPlusNPlusMOddWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(relationshipsFound[0], Equals, coefficient.PlusNPlusM)
}

func (suite *WavePacketRelationshipTest) TestMinusNMinusM(checker *C) {
	relationshipsFound := formula.GetWavePacketRelationship(
		*suite.aPlusNPlusMOddWavePacket,
		*suite.aMinusNMinusMOddWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(relationshipsFound[0], Equals, coefficient.MinusNMinusM)
}

func (suite *WavePacketRelationshipTest) TestMinusMMinusN(checker *C) {
	relationshipsFound := formula.GetWavePacketRelationship(
		*suite.aPlusNPlusMOddWavePacket,
		*suite.aMinusMMinusNOddWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(relationshipsFound[0], Equals, coefficient.MinusMMinusN)
}

func (suite *WavePacketRelationshipTest) TestPlusMPlusNMaybeFlipScale(checker *C) {
	relationshipsFound := formula.GetWavePacketRelationship(
		*suite.aPlusNPlusMOddWavePacket,
		*suite.aPlusMPlusNOddNegatedWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(formula.ContainsRelationship(relationshipsFound, coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum), Equals, true)

	relationshipsFound = formula.GetWavePacketRelationship(
		*suite.aPlusNPlusMEvenWavePacket,
		*suite.aPlusMPlusNEvenNegatedWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 0)

	relationshipsFound = formula.GetWavePacketRelationship(
		*suite.aPlusNPlusMEvenWavePacket,
		*suite.aPlusMPlusNEvenWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 2)
	checker.Assert(formula.ContainsRelationship(relationshipsFound, coefficient.PlusMPlusNNegateMultiplierIfOddPowerSum), Equals, true)
	checker.Assert(formula.ContainsRelationship(relationshipsFound, coefficient.PlusMPlusN), Equals, true)
}

func (suite *WavePacketRelationshipTest) TestMinusMMinusNMaybeFlipScale(checker *C) {
	relationshipsFound := formula.GetWavePacketRelationship(
		*suite.aPlusNPlusMOddWavePacket,
		*suite.aMinusMMinusNOddNegatedWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(relationshipsFound[0], Equals, coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum)

	relationshipsFound = formula.GetWavePacketRelationship(
		*suite.aPlusNPlusMEvenWavePacket,
		*suite.aMinusMMinusNEvenWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 2)
	checker.Assert(formula.ContainsRelationship(relationshipsFound, coefficient.MinusMMinusNNegateMultiplierIfOddPowerSum), Equals, true)
	checker.Assert(formula.ContainsRelationship(relationshipsFound, coefficient.MinusMMinusN), Equals, true)

	relationshipsFound = formula.GetWavePacketRelationship(
		*suite.aPlusNPlusMEvenWavePacket,
		*suite.aMinusMMinusNEvenNegatedWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 0)
}

func (suite *WavePacketRelationshipTest) TestPlusMMinusSumNAndM(checker *C) {
	relationshipsFound := formula.GetWavePacketRelationship(
		*suite.aPlusNPlusMOddWavePacket,
		*suite.aPlusMMinusSumNAndMOddWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(relationshipsFound[0], Equals, coefficient.PlusMMinusSumNAndM)
}

func (suite *WavePacketRelationshipTest) TestMinusSumNAndMPlusN(checker *C) {
	relationshipsFound := formula.GetWavePacketRelationship(
		*suite.aPlusNPlusMOddWavePacket,
		*suite.aMinusSumNAndMPlusNOddWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(relationshipsFound[0], Equals, coefficient.MinusSumNAndMPlusN)
}

func (suite *WavePacketRelationshipTest) TestPlusMMinusN(checker *C) {
	relationshipsFound := formula.GetWavePacketRelationship(
		*suite.aPlusNPlusMOddWavePacket,
		*suite.aPlusMMinusNOddWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(relationshipsFound[0], Equals, coefficient.PlusMMinusN)
}

func (suite *WavePacketRelationshipTest) TestMinusMPlusN(checker *C) {
	relationshipsFound := formula.GetWavePacketRelationship(
		*suite.aPlusNPlusMOddWavePacket,
		*suite.aMinusMPlusNOddWavePacket,
	)

	checker.Assert(relationshipsFound, HasLen, 1)
	checker.Assert(relationshipsFound[0], Equals, coefficient.MinusMPlusN)
}

type WavePacketDataStreamBuilder struct{}

var _ = Suite(&WavePacketDataStreamBuilder{})

func (suite *WavePacketDataStreamBuilder) TestCreateWavePacketUsingYAML(checker *C) {
	yamlByteStream := []byte(`
multiplier:
  real: 2e9
  imaginary: 1e-3
terms:
  -
    power_n: 3
    power_m: 19
`)
	packet := formula.NewWavePacketBuilder().UsingYAMLData(yamlByteStream).Build()
	checker.Assert(real(packet.Multiplier()), utility.NumericallyCloseEnough{}, 2e9, 1e-6)
	checker.Assert(imag(packet.Multiplier()), utility.NumericallyCloseEnough{}, 1e-3, 1e-6)

	checker.Assert(packet.Terms(), HasLen, 1)
	term := packet.Terms()[0]
	checker.Assert(term.PowerN, Equals, 3)
	checker.Assert(term.PowerM, Equals, 19)
}

func (suite *WavePacketDataStreamBuilder) TestCreateWavePacketUseDefaultMultiplier(checker *C) {
	yamlByteStream := []byte(``)
	packet := formula.NewWavePacketBuilder().UsingYAMLData(yamlByteStream).Build()
	checker.Assert(packet.Multiplier(), Equals, complex(1, 0))
}

func (suite *WavePacketDataStreamBuilder) TestCreateWavePacketUsingJSON(checker *C) {
	jsonByteStream := []byte(`{
"multiplier": {
	"real": 5e2,
	"imaginary": 3e-1
},
"terms": [
	{
		"power_n": 3,
		"power_m": 19
	},
	{
		"power_n": -11,
		"power_m": -7
	}
]
}`)
	packet := formula.NewWavePacketBuilder().UsingJSONData(jsonByteStream).Build()
	checker.Assert(real(packet.Multiplier()), utility.NumericallyCloseEnough{}, 5e2, 1e-6)
	checker.Assert(imag(packet.Multiplier()), utility.NumericallyCloseEnough{}, 3e-1, 1e-6)

	checker.Assert(packet.Terms(), HasLen, 2)
	term0 := packet.Terms()[0]
	checker.Assert(term0.PowerN, Equals, 3)
	checker.Assert(term0.PowerM, Equals, 19)
	term1 := packet.Terms()[1]
	checker.Assert(term1.PowerN, Equals, -11)
	checker.Assert(term1.PowerM, Equals, -7)
}
