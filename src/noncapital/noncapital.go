package noncapital

func CheckFeeTypeNonCapital(feeType string) bool {

	feeTypeResult := "ไม่มีทุนทรัพย์"

	if feeTypeResult == feeType {
		return true
	}

	return false
}
