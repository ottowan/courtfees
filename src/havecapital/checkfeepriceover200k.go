package havecapital

func CheckFeePriceOver200k(feePrice float64) bool {
	assignFeePrice := 200000.00
	if feePrice >= assignFeePrice {
		return true
	} else {
		return false
	}
}
