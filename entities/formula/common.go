package formula

import (
	"errors"
	"fmt"
	"github.com/Chadius/creating-symmetry/entities/formula/coefficient"
	"math"
)

// ConvertToLatticeCoordinates changes the coordinate to match the axes defined by the latticeVectors.
func ConvertToLatticeCoordinates(cartesianPoint complex128, latticeVectors []complex128) complex128 {
	vector1 := latticeVectors[0]
	vector2 := latticeVectors[1]
	swapVectorsDuringCalculation := real(vector1) < 1e-6

	if swapVectorsDuringCalculation == true {
		vector1 = latticeVectors[1]
		vector2 = latticeVectors[0]
	}

	scalarForVector2Numerator := (real(vector1) * imag(cartesianPoint)) - (imag(vector1) * real(cartesianPoint))
	scalarForVector2Denominator := (real(vector1) * imag(vector2)) - (imag(vector1) * real(vector2))
	scalarForVector2 := scalarForVector2Numerator / scalarForVector2Denominator

	scalarForVector1Numerator := real(cartesianPoint) - (scalarForVector2 * real(vector2))
	scalarForVector1Denominator := real(vector1)
	scalarForVector1 := scalarForVector1Numerator / scalarForVector1Denominator

	if swapVectorsDuringCalculation {
		return complex(scalarForVector2, scalarForVector1)
	}

	return complex(scalarForVector1, scalarForVector2)
}

func vectorIsZero(vector complex128) bool {
	return real(vector) == 0 && imag(vector) == 0
}

// vectorsAreCollinear returns true if both vectors are perfectly lined up
func vectorsAreCollinear(vector1 complex128, vector2 complex128) bool {
	absoluteValueDotProduct := math.Abs((real(vector1) * real(vector2)) + (imag(vector1) * imag(vector2)))
	lengthOfVector1 := math.Sqrt((real(vector1) * real(vector1)) + (imag(vector1) * imag(vector1)))
	lengthOfVector2 := math.Sqrt((real(vector2) * real(vector2)) + (imag(vector2) * imag(vector2)))

	tolerance := 1e-8
	return math.Abs(absoluteValueDotProduct-(lengthOfVector1*lengthOfVector2)) < tolerance
}

// ValidateLatticeVectors returns an error if the lattice vectors are invalid.
func ValidateLatticeVectors(latticeVectors []complex128) error {
	if vectorIsZero(latticeVectors[0]) || vectorIsZero(latticeVectors[1]) {
		return errors.New(`lattice vectors cannot be (0,0)`)
	}
	if vectorsAreCollinear(latticeVectors[0], latticeVectors[1]) {
		return fmt.Errorf(
			`vectors cannot be collinear: (%f,%f) and \(%f,%f)`,
			real(latticeVectors[0]),
			imag(latticeVectors[0]),
			real(latticeVectors[1]),
			imag(latticeVectors[1]),
		)
	}
	return nil
}

// lockTermsBasedOnRelationship adds terms based on the first term of each WavePacket.
func lockTermsBasedOnRelationship(
	lockedRelationships []coefficient.Relationship,
	originalWavePackets []WavePacket) []WavePacket {
	newWavePackets := []WavePacket{}

	for _, wavePacket := range originalWavePackets {
		baseCoefficientPairing := coefficient.Pairing{
			PowerN: wavePacket.Terms()[0].PowerN,
			PowerM: wavePacket.Terms()[0].PowerM,
		}

		newWavePacket := NewWavePacketBuilder().
			Multiplier(wavePacket.Multiplier()).
			AddTerm(&wavePacket.Terms()[0])

		newPairings := baseCoefficientPairing.GenerateCoefficientSets(lockedRelationships)
		for _, newCoefficientPair := range newPairings {
			newTerm := NewTermBuilder().
				PowerN(newCoefficientPair.PowerN).
				PowerM(newCoefficientPair.PowerM).
				Multiplier(complex(1,0)).
				Build()

			newWavePacket.AddTerm(newTerm)
		}
		
		newWavePackets = append(newWavePackets, *newWavePacket.Build())
	}
	
	return newWavePackets
}

// CalculateCoordinateUsingWavePackets transforms the coordinate using the lattice and its wave packets.
func CalculateCoordinateUsingWavePackets(coordinate complex128, latticeVectors []complex128, wavePackets []WavePacket) complex128 {
	result := complex(0,0)

	zInLatticeCoordinates := ConvertToLatticeCoordinates(coordinate, latticeVectors)

	for _, wavePacket := range wavePackets {
		termContribution := wavePacket.Calculate(zInLatticeCoordinates)
		result += termContribution / complex(float64(len(wavePacket.Terms())), 0)
	}

	return result
}

// addNewWavePacketsBasedOnSymmetry creates new WavePackets based on the given term, multiplier and desired symmetry
func addNewWavePacketsBasedOnSymmetry(term Term, multiplier complex128, desiredSymmetry Symmetry, newWavePackets []WavePacket) []WavePacket {
	powerN := term.PowerN
	powerM := term.PowerM
	powerNIsEven := powerN%2 == 0
	powerSumIsEven := (powerN+powerM)%2 == 0

	multiplierMaybeNegatedBasedOnSum := multiplier
	if !powerSumIsEven {
		multiplierMaybeNegatedBasedOnSum *= -1
	}

	multiplierMaybeNegatedBasedOnPowerN := multiplier
	if !powerNIsEven {
		multiplierMaybeNegatedBasedOnPowerN *= -1
	}

	//if desiredSymmetry == P31m || desiredSymmetry == P4m || desiredSymmetry == Cm {
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplier).
	//		AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerM, powerN)).
	//		Build())
	//}
	//if desiredSymmetry == Pm {
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplier).
	//		AddTerm(NewTermWithMultiplierAndPowers(multiplier, powerN, powerM * -1)).
	//		Build())
	//}
	//if desiredSymmetry == Pg {
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplier).
	//		AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerN, powerM * -1)).
	//		Build())
	//}
	//if desiredSymmetry == Pmm {
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplier).
	//		AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerN * -1, powerM * -1)).
	//		Build())
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplier).
	//		AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerN * -1, powerM)).
	//		Build())
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplier).
	//		AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerN, powerM * -1)).
	//		Build())
	//}
	//if desiredSymmetry == Pmg {
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplier).
	//		AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerN * -1, powerM * -1)). // TODO double check that multiplier, may be multiplierMaybeNegatedBasedOnPowerN
	//		Build())
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplierMaybeNegatedBasedOnPowerN).
	//		AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerN * -1, powerM)).
	//		Build())
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplierMaybeNegatedBasedOnPowerN).
	//		AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerN, powerM * -1)).
	//		Build())
	//}
	//if desiredSymmetry == Pgg {
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplier).
	//		AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerN * -1, powerM * -1)). // TODO double check that multiplier, may be multiplierMaybeNegatedBasedOnSum
	//		Build())
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplierMaybeNegatedBasedOnSum).
	//		AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerN * -1, powerM)).
	//		Build())
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplierMaybeNegatedBasedOnSum).
	//		AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerN, powerM * -1)).
	//		Build())
	//}
	//if desiredSymmetry == P3m1 {
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplier).
	//		AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerM * -1, powerN * -1)).
	//		Build())
	//}
	//if desiredSymmetry == P6 {
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplier).
	//		AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerN * -1, powerM * -1)).
	//		Build())
	//}
	//if desiredSymmetry == P6m || desiredSymmetry == Cmm {
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplier).
	//		AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerN * -1, powerM * -1)).
	//		Build())
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplier).
	//		AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerM, powerN)).
	//		Build())
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplier).
	//		AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerM * -1, powerN * -1)).
	//		Build())
	//}
	//
	//if desiredSymmetry == P4g {
	//	newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
	//Multiplier(multiplierMaybeNegatedBasedOnSum).
	//		AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerM, powerN)).
	//		Build())
	//}
	if desiredSymmetry == P2 {
		newWavePackets = append(newWavePackets, *NewWavePacketBuilder().
			Multiplier(multiplier).
			AddTerm(NewTermWithMultiplierAndPowers(term.Multiplier, powerN * -1, powerM * -1)).
			Build())
	}

	return newWavePackets
}

// HasSymmetry returns true if the WavePackets involved form the desired symmetry.
func HasSymmetry(wavePackets []WavePacket, desiredSymmetry Symmetry, desiredSymmetryToCoefficients map[Symmetry][]coefficient.Relationship) bool {
	numberOfWavePackets := len(wavePackets)
	if numberOfWavePackets < 2 || numberOfWavePackets%2 == 1 {
		return false
	}

	coefficientsToFind := desiredSymmetryToCoefficients[desiredSymmetry]

	if coefficientsToFind == nil {
		return false
	}

	return CanWavePacketsBeGroupedAmongCoefficientRelationships(wavePackets, coefficientsToFind)
}

// CanWavePacketsBeGroupedAmongCoefficientRelationships returns true if the WavePackets involved satisfy the relationships.
func CanWavePacketsBeGroupedAmongCoefficientRelationships(wavePackets []WavePacket, desiredRelationships []coefficient.Relationship) bool {
	wavePacketsMatched := []bool{}
	for range wavePackets {
		wavePacketsMatched = append(wavePacketsMatched, false)
	}

	for indexA, wavePacketA := range wavePackets {
		relationshipWasFound := map[coefficient.Relationship]bool{}
		for _, r := range desiredRelationships {
			relationshipWasFound[r] = false
		}

		if wavePacketsMatched[indexA] == true {
			continue
		}

		for offsetB, wavePacketB := range wavePackets[indexA+1:] {
			relationshipsFound := GetWavePacketRelationship(
				wavePacketA,
				wavePacketB,
			)

			for _, relationshipToLookFor := range desiredRelationships {
				if ContainsRelationship(relationshipsFound, relationshipToLookFor) {
					wavePacketsMatched[indexA+offsetB+1] = true
					relationshipWasFound[relationshipToLookFor] = true
					break
				}
			}
		}

		for _, relationshipFound := range relationshipWasFound {
			if relationshipFound != true {
				return false
			}
		}
		wavePacketsMatched[indexA] = true
	}

	return true
}
