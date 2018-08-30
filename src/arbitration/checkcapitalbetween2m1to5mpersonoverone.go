package arbitration

func CheckCapitalBetween2m1to5mPersonOverOne(feeCapital float64) bool {
	if feeCapital > 2000000 && feeCapital <= 5000000 {
		return true
	}

	return false
}
