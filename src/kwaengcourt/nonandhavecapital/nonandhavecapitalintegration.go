package nonandhavecapital

import "math"

func CalculateNonAndHaveCapital(feeCapital float64) float64 {

	if CheckHaveCapital1to300k(feeCapital) {

		feePrice := CalculateHaveCapital1to300k(feeCapital)
		// if CheckFeePriceOver200(feePrice) {
		// 	if CheckFeePriceOver1000(feePrice) {
		// 		return math.Floor(InitFreePrice1000())
		// 	} else {
		// 		return math.Floor(InitFreePriceEqualFreePrice(feePrice))
		// 	}
		// } else {
		// 	return math.Floor(InitFreePrice200())
		// }

		if CheckFeePriceOver1000(feePrice) {
			return math.Floor(InitFreePrice1000())
		} else if CheckFeePriceOver200(feePrice) {
			return math.Floor(InitFreePriceEqualFreePrice(feePrice))
		} else {
			return math.Floor(InitFreePrice200())
		}

	}
	return math.Floor(0.00)

}
