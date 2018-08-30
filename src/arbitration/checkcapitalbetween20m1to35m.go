package arbitration

func CheckCapitalBetween20m1to35m(feeCapital float64) bool {

	capitalMin := 20000000.00
	capitalMaX := 35000000.00

	if capitalMin < feeCapital && feeCapital <= capitalMaX {
		return true
	}

	return false
}
