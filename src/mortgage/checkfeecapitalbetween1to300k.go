package mortgage

func CheckFeeCapitalBetween1To300k(feeCapital float64) bool {

	if feeCapital > 0 && feeCapital <= 300000.00 {
		return true
	}

	return false
}
