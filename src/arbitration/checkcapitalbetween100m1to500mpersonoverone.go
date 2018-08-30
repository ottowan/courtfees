package arbitration

func CheckCapitalBetween100m1to500mPersonOverOne(feeCapital float64) bool {
	if feeCapital > 100000000 && feeCapital <= 500000000 {
		return true
	}

	return false
}
