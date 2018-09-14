package noncapital
import (
	"math"
)
func CalculateNonCapital(section string, feeCapital float64) float64 {

	if CheckSectionType288_1(section) && feeCapital == 0.00 {
		return math.Floor(InitFeePrice0())
	} else if CheckSectionType288_2(section) && feeCapital == 0.00 {
		return math.Floor(InitFeePrice200())
	} else if CheckSectionType288_3(section) && feeCapital == 0.00 {
		return math.Floor(InitFeePrice200())
	} else if CheckSectionType227(section) && feeCapital == 0.00 {
		return math.Floor(InitFeePrice200())
	}

	return math.Floor(InitFeePrice200())
}
