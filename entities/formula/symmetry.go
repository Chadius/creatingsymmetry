package formula

//// addNewWavePacketsBasedOnSymmetry creates new WavePackets based on the given term, multiplier and desired symmetry
//func addNewWavePacketsBasedOnSymmetry(term *eisenstien.EisensteinFormulaTerm, multiplier complex128, desiredSymmetry Symmetry, newWavePackets []*WavePacket) []*WavePacket {
//	powerN := term.PowerN
//	powerM := term.PowerM
//	powerNIsEven := powerN%2 == 0
//	powerSumIsEven := (powerN+powerM)%2 == 0
//
//	multiplierMaybeNegatedBasedOnSum := multiplier
//	if !powerSumIsEven {
//		multiplierMaybeNegatedBasedOnSum *= -1
//	}
//
//	multiplierMaybeNegatedBasedOnPowerN := multiplier
//	if !powerNIsEven {
//		multiplierMaybeNegatedBasedOnPowerN *= -1
//	}
//
//	if desiredSymmetry == P31m || desiredSymmetry == P4m || desiredSymmetry == Cm {
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerN: powerM,
//					PowerM: powerN,
//				},
//			},
//			multiplier: multiplier,
//		})
//	}
//	if desiredSymmetry == Pm {
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerN: powerN,
//					PowerM: powerM * -1,
//				},
//			},
//			multiplier: multiplier,
//		})
//	}
//	if desiredSymmetry == Pg {
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerN: powerN,
//					PowerM: powerM * -1,
//				},
//			},
//			multiplier: multiplierMaybeNegatedBasedOnPowerN,
//		})
//	}
//	if desiredSymmetry == Pmm {
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerN: powerN * -1,
//					PowerM: powerM * -1,
//				},
//			},
//			multiplier: multiplier,
//		})
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerN: powerN * -1,
//					PowerM: powerM,
//				},
//			},
//			multiplier: multiplier,
//		})
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerN: powerN,
//					PowerM: powerM * -1,
//				},
//			},
//			multiplier: multiplier,
//		})
//	}
//	if desiredSymmetry == Pmg {
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerN: powerN * -1,
//					PowerM: powerM * -1,
//				},
//			},
//			multiplier: multiplier,
//		})
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerN: powerN * -1,
//					PowerM: powerM,
//				},
//			},
//			multiplier: multiplierMaybeNegatedBasedOnPowerN,
//		})
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerN: powerN,
//					PowerM: powerM * -1,
//				},
//			},
//			multiplier: multiplierMaybeNegatedBasedOnPowerN,
//		})
//	}
//	if desiredSymmetry == Pgg {
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerN: powerN * -1,
//					PowerM: powerM * -1,
//				},
//			},
//			multiplier: multiplier,
//		})
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerN: powerN * -1,
//					PowerM: powerM,
//				},
//			},
//			multiplier: multiplierMaybeNegatedBasedOnSum,
//		})
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerN: powerN,
//					PowerM: powerM * -1,
//				},
//			},
//			multiplier: multiplierMaybeNegatedBasedOnSum,
//		})
//	}
//
//	if desiredSymmetry == P3m1 {
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerM: powerN * -1,
//					PowerN: powerM * -1,
//				},
//			},
//			multiplier: multiplier,
//		})
//	}
//	if desiredSymmetry == P6 {
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerM: powerM * -1,
//					PowerN: powerN * -1,
//				},
//			},
//			multiplier: multiplier,
//		})
//	}
//	if desiredSymmetry == P6m || desiredSymmetry == Cmm {
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerN: powerN * -1,
//					PowerM: powerM * -1,
//				},
//			},
//			multiplier: multiplier,
//		})
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerN: powerM,
//					PowerM: powerN,
//				},
//			},
//			multiplier: multiplier,
//		})
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerN: powerM * -1,
//					PowerM: powerN * -1,
//				},
//			},
//			multiplier: multiplier,
//		})
//	}
//
//	if desiredSymmetry == P4g {
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerM: powerN,
//					PowerN: powerM,
//				},
//			},
//			multiplier: multiplierMaybeNegatedBasedOnSum,
//		})
//	}
//	if desiredSymmetry == P2 {
//		newWavePackets = append(newWavePackets, &WavePacket{
//			terms: []*eisenstien.EisensteinFormulaTerm{
//				{
//					PowerN: powerN * -1,
//					PowerM: powerM * -1,
//				},
//			},
//			multiplier: multiplier,
//		})
//	}
//
//	return newWavePackets
//}

// Symmetry encodes all possible symmetries for wallpaper patterns.
type Symmetry string

// All possible symmetries for wallpaper patterns, based on crystallography.
const (
	P1   Symmetry = "p1"
	P2   Symmetry = "p2"
	P3   Symmetry = "p3"
	P3m1 Symmetry = "p3m1"
	P31m Symmetry = "p31m"
	P6   Symmetry = "p6"
	P6m  Symmetry = "p6m"
	P4   Symmetry = "p4"
	P4m  Symmetry = "p4m"
	P4g  Symmetry = "p4g"
	Cm   Symmetry = "cm"
	Cmm  Symmetry = "cmm"
	Pm   Symmetry = "pm"
	Pg   Symmetry = "pg"
	Pgg  Symmetry = "pgg"
	Pmm  Symmetry = "pmm"
	Pmg  Symmetry = "pmg"
)
