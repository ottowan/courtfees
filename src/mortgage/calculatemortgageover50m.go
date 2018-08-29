package mortgage

func CalculateMortgageOver50m(feeCapital float64) float64{
	return ((feeCapital-50000000.00)*0.001)+100000.00
}