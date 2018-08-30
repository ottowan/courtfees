package arbitration

func CheckCapitalOver2000mPersonOverOne(feeCapital float64) bool {
	if feeCapital > 2000000000 {
		return true
	}

	return false
}
