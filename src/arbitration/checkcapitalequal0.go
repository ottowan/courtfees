package arbitration

func CheckCapitalEqual0(feeCapital float64) bool {
	if feeCapital == 0 {
		return true
	}

	return false
}
