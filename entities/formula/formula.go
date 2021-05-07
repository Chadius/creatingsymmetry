package formula

// CalculationResultForFormula shows the results of a calculation
type CalculationResultForFormula struct {
	Total				complex128
	ContributionByTerm	[]complex128
}