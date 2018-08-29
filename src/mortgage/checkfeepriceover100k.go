package mortgage

func CheckFeePriceOver100k(feePrice float64) bool{
	if feePrice > 100000.00 {
		return true
	}

	return false
}