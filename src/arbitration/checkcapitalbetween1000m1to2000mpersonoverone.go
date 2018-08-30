package arbitration

func CheckCapitalBetween1000m1to2000mPersonOverOne(feeCapital float64) bool {
	if feeCapital > 1000000000 && feeCapital <= 20000000000 {
		return true
	}

	return false
}
