package havecapital

func InitFreePrice1000() float64 {
	return 1000.00
}

func InitFreePriceEqualFreePrice(feePrice float64) float64 {
	return feePrice
}

func CheckHaveCapital1to300k(feeCapital float64) bool {

	if feeCapital > 1000 && feeCapital <= 300000.00 {
		return true
	}

	return false
}

func CheckFeePriceOver1000(feePrice float64) bool {
	assignFeePrice := 1000.00
	if feePrice > assignFeePrice {
		return true
	} else {
		return false
	}
}

func CalculateHaveCapital1to300k(feeCapital float64) float64 {
	return feeCapital * 0.02
}
