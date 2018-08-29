package havecapital

func CheckHaveCapitalOver50m(feeCapital float64) bool {
	if feeCapital > 50000000 {
		return true
	}

	return false
}
