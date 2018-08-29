package mortgage

func CheckFeePriceOver1000(feePrice float64) bool {
	if feePrice > 1000.00 {
		return true
	}
	return false
}
