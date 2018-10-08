package nonandhavecapital

func CheckFeePriceOver200(feeCapital float64) bool {
	if feeCapital >= 200 {
		return true
	}

	return false
}

func InitFreePrice200() float64 {
	return 200.00
}

func InitFreePrice200k() float64 {
	return 200000.00
}

func InitFreePriceEqualFreePrice(feePrice float64) float64 {
	return feePrice
}

func CheckHaveCapitalOver50m(feeCapital float64) bool {
	if feeCapital > 50000000 {
		return true
	}

	return false
}

func CheckHaveCapital1to50m(feeCapital float64) bool {

	if feeCapital > 0.00 && feeCapital <= 50000000.00 {
		return true
	}

	return false
}

func CheckFeePriceOver200k(feePrice float64) bool {
	assignFeePrice := 200000.00
	if feePrice >= assignFeePrice {
		return true
	} else {
		return false
	}
}

func CalculateHaveCapital1to50m(feeCapital float64) float64 {
	return feeCapital * 0.02
}

func CaluculateHaveCapitalOver50m(feeCapital float64) float64 {
	return ((feeCapital - 50000000.00) * 0.001) + 200000.00
}

func InitFreePrice1000() float64 {
	return 1000.00
}

// func InitFreePriceEqualFreePrice(feePrice float64) float64 {
// 	return feePrice
// }

func CheckHaveCapital1to300k(feeCapital float64) bool {

	if feeCapital >= 1 && feeCapital <= 300000.00 {
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
