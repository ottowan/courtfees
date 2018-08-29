package havecapital

func CheckFeePriceOver200k(feePrice float64) float64 {
	assignFeePrice := 200000.00
	if feePrice >= assignFeePrice {
		return assignFeePrice
	} else {
		return feePrice
	}
}
