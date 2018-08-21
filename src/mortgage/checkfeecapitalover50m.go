package mortgage

func checkFeeCapitalOver50m(feeCapital float64) bool {

	checkFeeCapital := 50000000.00

	if feeCapital > checkFeeCapital {
		return true
	}
	return false
}
