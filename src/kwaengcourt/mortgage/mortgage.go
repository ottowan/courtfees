package mortgage

func CalculateMortgageBetween1to300k(feeCapital float64) float64 {
	return feeCapital * 0.01
}

func CheckFeeCapitalBetween1To300k(feeCapital float64) bool {

	if feeCapital > 0 && feeCapital <= 300000.00 {
		return true
	}

	return false
}

func CheckFeePriceOver1000(feePrice float64) bool {
	if feePrice > 1000.00 {
		return true
	}
	return false
}

func InitFeePriceEqual1000() float64 {
	return 1000.00
}

func InitFeePriceEqualFeePrice(feePrice float64) float64 {
	return feePrice
}
