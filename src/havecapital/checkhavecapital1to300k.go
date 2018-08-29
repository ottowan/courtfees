package havecapital

func CheckHaveCapital1to300k(feeCapital float64) bool {

	if feeCapital >= 1.00 && feeCapital <= 300000.00 {
		return true
	}

	return false
}
