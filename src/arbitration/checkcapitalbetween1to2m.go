package arbitration

func CheckCapitalBetween1to2m(feeCapital float64) bool {
	if feeCapital > 0 && feeCapital <= 2000000 {
		return true
	}

	return false
}