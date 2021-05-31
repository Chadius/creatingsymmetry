package wavepacket

import "wallpaper/entities/formula"

func addNewWavePacketsBasedOnSymmetry(term *formula.EisensteinFormulaTerm, desiredSymmetry *Symmetry, newWavePackets []*Formula) []*Formula {
	powerN := term.PowerN
	powerM := term.PowerM

	if desiredSymmetry.P31m || desiredSymmetry.P4m {
		newWavePackets = append(newWavePackets, &Formula{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerM: powerN,
					PowerN: powerM,
				},
			},
			Multiplier: term.Multiplier,
		})
	}
	if desiredSymmetry.P3m1 {
		newWavePackets = append(newWavePackets, &Formula{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerM: powerN * -1,
					PowerN: powerM * -1,
				},
			},
			Multiplier: term.Multiplier,
		})
	}
	if desiredSymmetry.P6 {
		newWavePackets = append(newWavePackets, &Formula{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerM: powerM * -1,
					PowerN: powerN * -1,
				},
			},
			Multiplier: term.Multiplier,
		})
	}
	if desiredSymmetry.P6m {
		newWavePackets = append(newWavePackets, &Formula{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerM: powerM * -1,
					PowerN: powerN * -1,
				},
			},
			Multiplier: term.Multiplier,
		})
		newWavePackets = append(newWavePackets, &Formula{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerM: powerN,
					PowerN: powerM,
				},
			},
			Multiplier: term.Multiplier,
		})
		newWavePackets = append(newWavePackets, &Formula{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerM: powerN * -1,
					PowerN: powerM * -1,
				},
			},
			Multiplier: term.Multiplier,
		})
	}

	powerSumIsEven := (powerN + powerM) % 2 == 0
	if desiredSymmetry.P4g {
		multiplierMaybeNegatedBasedOnSum := term.Multiplier
		if !powerSumIsEven {
			multiplierMaybeNegatedBasedOnSum *= -1
		}

		newWavePackets = append(newWavePackets, &Formula{
			Terms: []*formula.EisensteinFormulaTerm{
				{
					PowerM: powerN,
					PowerN: powerM,
				},
			},
			Multiplier: multiplierMaybeNegatedBasedOnSum,
		})
	}

	return newWavePackets
}
