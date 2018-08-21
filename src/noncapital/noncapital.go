package noncapital

func CheckFeeTypeNonCapital(feeType string) bool {

	feeTypeResult := "ไม่มีทุนทรัพย์"

	if feeTypeResult == feeType {
		return true
	}

	return false
}

func InitFeePrice0() int {

	return 0
}

func CheckSectionType288_1(section string) bool {

	sectionReturn := "288(1)"
	if section == sectionReturn {
		return true
	}

	return false

}

