package havecapital

import "math"

func CalculateHaveCapital(feeCapital float64) float64 {

	if CheckHaveCapital1to300k(feeCapital) {

		freePrice := CalculateHaveCapital1to300k(feeCapital)
		if CheckFeePriceOver1000(freePrice) {
			return math.Floor(InitFreePrice1000())
		} else {
			return math.Floor(InitFreePriceEqualFreePrice(freePrice))
		}

	}
	return math.Floor(0.00)

}
