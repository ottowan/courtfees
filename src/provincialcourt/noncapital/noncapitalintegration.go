package noncapital

import (
	"math"
)

func CalculateNonCapital(section string, feeCapital float64) float64 {

	if CheckSectionType288_1(section) && CheckFeeCapitalEqual0(feeCapital) {
		return math.Floor(InitFeePrice0())
	} else if CheckSectionType288_2(section) && CheckFeeCapitalEqual0(feeCapital) {
		return math.Floor(InitFeePrice200())
	} else if CheckSectionType288_3(section) && CheckFeeCapitalEqual0(feeCapital) {
		return math.Floor(InitFeePrice200())
	} else if CheckSectionType227(section) && CheckFeeCapitalEqual0(feeCapital) {
		return math.Floor(InitFeePrice200())
	} else if CheckFeeCapitalEqual0(feeCapital) {
		return math.Floor(InitFeePrice200())
	}

	return math.Floor(0.00)
}
