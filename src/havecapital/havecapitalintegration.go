package havecapital

import "math"

func CalculateHaveCapital(feeCapital float64) float64 {

	if CheckHaveCapital300k1to50m(feeCapital) {
		freePrice := CalculateHaveCapital300k1to50m(feeCapital)

		if CheckFeePriceOver200k(freePrice) {
			return math.Floor(InitFreePrice200k())
		} else {
			return math.Floor(InitFreePriceEqualFreePrice(freePrice))
		}
	} else if CheckHaveCapitalOver50m(feeCapital) {

		return math.Floor(CaluculateHaveCapitalOver50m(feeCapital))

	} else if CheckHaveCapital1to300k(feeCapital) {

		freePrice := CalculateHaveCapital1to300k(feeCapital)
		if CheckFeePriceOver1000(freePrice) {
			return math.Floor(InitFreePrice1000())
		} else {
			return math.Floor(InitFreePriceEqualFreePrice(freePrice))
		}

	}
	return math.Floor(0.00)

}
