package arbitration

func CheckCapitalBetween30m1to50mPersonOverOne(feeCapital float64) bool {

	capitalMin := 35000000.00
	capitalMaX := 50000000.00

	if capitalMin < feeCapital && feeCapital <= capitalMaX {
		return true
	}

	return false
}
