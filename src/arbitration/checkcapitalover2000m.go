package arbitration

func CheckCapitalOver2000m(feeCapital float64) bool {
	if feeCapital > 2000000000 {
		return true
	}

	return false
}
