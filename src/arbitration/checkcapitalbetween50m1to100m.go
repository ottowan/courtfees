package arbitration

func CheckCapitalBetween50m1to100m(feeCapital float64) bool {

	capitalMin := 50000000.00
	capitalMaX := 100000000.00

	if capitalMin < feeCapital && feeCapital <= capitalMaX {
		return true
	}

	return false
}
