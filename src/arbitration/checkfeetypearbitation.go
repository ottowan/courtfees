package arbitration

func CheckFeeTypeArbitation(feeType string) bool {
	if feeType == "อนุญาโต" {
		return true
	}

	return false
}
