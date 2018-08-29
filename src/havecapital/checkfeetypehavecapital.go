package havecapital

func CheckFeeTypeHaveCapital(feeType string) bool {

	if feeType == "มีทุนทรัพย์" {
		return true
	}

	return false
}
