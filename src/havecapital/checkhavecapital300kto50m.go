package havecapital

func CheckHaveCapital300kto50m(feeCapital float64) bool {

	if feeCapital >= 300000.00 && feeCapital <= 50000000.00 {
		return true
	}

	return false
}