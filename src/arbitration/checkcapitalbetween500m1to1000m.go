package arbitration

func CheckCapitalBetween500m1to1000m(feeCapital float64) bool {
	if feeCapital > 500000000 && feeCapital <= 10000000000 {
		return true
	}

	return false
}
