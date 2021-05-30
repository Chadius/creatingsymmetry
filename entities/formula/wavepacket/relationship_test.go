package wavepacket_test

import (
	. "gopkg.in/check.v1"
	"wallpaper/entities/formula"
	"wallpaper/entities/formula/wavepacket"
)

type WaveSymmetryFormulaTests struct {
	baseOddWave *wavepacket.Formula
}

var _ = Suite(&WaveSymmetryFormulaTests{})

func (suite *WaveSymmetryFormulaTests) SetUpTest(checker *C) {
	suite.baseOddWave = &wavepacket.Formula{
		Terms: []*formula.EisensteinFormulaTerm{
			{
				PowerN:         1,
				PowerM:         -2,
			},
		},
		Multiplier: complex(1, 0),
	}
}

func (suite *WaveSymmetryFormulaTests) TestNeedTwoFormulasToDetectSymmetry(checker *C) {
	relationship := wavepacket.FindWaveRelationships([]*wavepacket.Formula{
		suite.baseOddWave,
	})
	checker.Assert(relationship.NoRelation, Equals, true)
}

func (suite *WaveSymmetryFormulaTests) TestPlusMPlusN(checker *C) {
	relationship := wavepacket.FindWaveRelationships([]*wavepacket.Formula{
		suite.baseOddWave,
		{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerN:         suite.baseOddWave.Terms[0].PowerM,
					PowerM:         suite.baseOddWave.Terms[0].PowerN,
				},
			},
			Multiplier: complex(1, 0),
		},
	})
	checker.Assert(relationship.PlusMPlusN, Equals, true)
}

func (suite *WaveSymmetryFormulaTests) TestMinusNMinusM(checker *C) {
	relationship := wavepacket.FindWaveRelationships([]*wavepacket.Formula{
		suite.baseOddWave,
		{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerN:         suite.baseOddWave.Terms[0].PowerN * -1,
					PowerM:         suite.baseOddWave.Terms[0].PowerM * -1,
				},
			},
			Multiplier: complex(1, 0),
		},
	})
	checker.Assert(relationship.MinusNMinusM, Equals, true)
}

func (suite *WaveSymmetryFormulaTests) TestMinusMMinusN(checker *C) {
	relationship := wavepacket.FindWaveRelationships([]*wavepacket.Formula{
		suite.baseOddWave,
		{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerN:         suite.baseOddWave.Terms[0].PowerM * -1,
					PowerM:         suite.baseOddWave.Terms[0].PowerN * -1,
				},
			},
			Multiplier: complex(1, 0),
		},
	})
	checker.Assert(relationship.MinusMMinusN, Equals, true)
}

func (suite *WaveSymmetryFormulaTests) TestMinusNMinusMPlusMPlusNMinusMMinusN(checker *C) {
	relationship := wavepacket.FindWaveRelationships([]*wavepacket.Formula{
		suite.baseOddWave,
		{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerN:         suite.baseOddWave.Terms[0].PowerN * -1,
					PowerM:         suite.baseOddWave.Terms[0].PowerM * -1,
				},
			},
			Multiplier: complex(1, 0),
		},
		{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerN:         suite.baseOddWave.Terms[0].PowerM * -1,
					PowerM:         suite.baseOddWave.Terms[0].PowerN * -1,
				},
			},
			Multiplier: complex(1, 0),
		},
		{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerN:         suite.baseOddWave.Terms[0].PowerM,
					PowerM:         suite.baseOddWave.Terms[0].PowerN,
				},
			},
			Multiplier: complex(1, 0),
		},
	})
	checker.Assert(relationship.MinusNMinusMPlusMPlusNMinusMMinusN, Equals, true)
}

func (suite *WaveSymmetryFormulaTests) TestMultipleSymmetries(checker *C) {
	relationship := wavepacket.FindWaveRelationships([]*wavepacket.Formula{
		{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerN:         1,
					PowerM:         -1,
				},
			},
			Multiplier: complex(1, 0),
		},
		{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerN:         -1,
					PowerM:         1,
				},
			},
			Multiplier: complex(1, 0),
		},
	})
	checker.Assert(relationship.PlusMPlusN, Equals, true)
	checker.Assert(relationship.MinusNMinusM, Equals, true)
}

func (suite *WaveSymmetryFormulaTests) TestMaybeNegateBasedOnSumPlusMPlusNWithOddSum(checker *C) {
	relationship := wavepacket.FindWaveRelationships([]*wavepacket.Formula{
		suite.baseOddWave,
		{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerN:         suite.baseOddWave.Terms[0].PowerM,
					PowerM:         suite.baseOddWave.Terms[0].PowerN,
				},
			},
			Multiplier: complex(-1, 0),
		},
	})
	checker.Assert(relationship.PlusMPlusN, Equals, false)
	checker.Assert(relationship.MaybeNegateBasedOnSumPlusMPlusN, Equals, true)
}

func (suite *WaveSymmetryFormulaTests) TestMaybeNegateBasedOnSumPlusMPlusNWithEvenSum(checker *C) {
	relationship := wavepacket.FindWaveRelationships([]*wavepacket.Formula{
		{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerN:         4,
					PowerM:         -2,
				},
			},
			Multiplier: complex(-1, 0),
		},
		{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerN:         -2,
					PowerM:         4,
				},
			},
			Multiplier: complex(1, 0),
		},
	})
	checker.Assert(relationship.PlusMPlusN, Equals, true)
	checker.Assert(relationship.MaybeNegateBasedOnSumPlusMPlusN, Equals, true)
}