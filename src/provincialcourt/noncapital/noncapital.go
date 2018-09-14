package noncapital

func CheckFeeTypeNonCapital(feeType string) bool {

	feeTypeResult := "ไม่มีทุนทรัพย์"

	if feeTypeResult == feeType {
		return true
	}

	return false
}

func InitFeePrice0() float64 {

	return 0.00
}

func InitFeePrice200() float64 {
	return 200.00
}

func CheckSectionType288_1(section string) bool {

	sectionReturn := "288(1)"
	if section == sectionReturn {
		return true
	}

	return false

}

func CheckSectionType227(section string) bool {

	sectionReturn := "227"
	if section == sectionReturn {
		return true
	}

	return false

}
func CheckSectionType288_2(section string) bool {

	sectionReturn := "288(2)"
	if section == sectionReturn {
		return true
	}

	return false

}
func CheckSectionType288_3(section string) bool {

	sectionReturn := "288(3)"
	if section == sectionReturn {
		return true
	}

	return false

}
