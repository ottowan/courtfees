package noncapital

func CalculateNonCapital(section string, feeCapital float64) float64 {

	if CheckSectionType288_1(section) && feeCapital == 0.00 {
		return InitFeePrice0()
	} else {
		return InitFeePrice200()
	}
}