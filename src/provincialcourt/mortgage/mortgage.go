package mortgage

func CalculateMortgageBetween1to50m(feeCapital float64) float64 {
	return feeCapital * 0.01
}

func CalculateMortgageOver50m(feeCapital float64) float64 {
	return ((feeCapital - 50000000.00) * 0.001) + 100000.00
}

func CheckFeeCapitalBetween1to50m(feeCapital float64) bool {

	if feeCapital > 0.00 && feeCapital <= 50000000.00 {
		return true
	} else {
		return false
	}
}

func CheckFeeCapitalOver50m(feeCapital float64) bool {

	checkFeeCapital := 50000000.00

	if feeCapital > checkFeeCapital {
		return true
	}
	return false
}

func CheckFeePriceOver100k(feePrice float64) bool {
	if feePrice > 100000.00 {
		return true
	}

	return false
}

func InitFeePriceEqual100k() float64 {
	return 100000.00
}

func InitFeePriceEqualFeePrice(feePrice float64) float64 {
	return feePrice
}
