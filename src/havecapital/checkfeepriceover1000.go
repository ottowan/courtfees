package havecapital

func CheckFeePriceOver1000(feePrice float64) bool {
	assignFeePrice := 1000.00
	if feePrice > assignFeePrice {
		return true
	} else {
		return false
	}
}
