package arbitration

func CheckCapitalBetween10m1to20mPersonOverOne(feeCapital float64) bool {
	if feeCapital > 10000000 && feeCapital <= 20000000 {
		return true
	}

	return false
}
