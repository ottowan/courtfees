package arbitration

func CheckCapitalBetween5m1to10m(feeCapital float64) bool {
	if feeCapital > 5000000 && feeCapital <= 10000000 {
		return true
	}

	return false
}