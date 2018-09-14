package havecapital

import "math"

func CalculateHaveCapital(feeCapital float64) float64 {

	if CheckHaveCapital1to50m(feeCapital) {
		freePrice := CalculateHaveCapital1to50m(feeCapital)

		if CheckFeePriceOver200k(freePrice) {
			return math.Floor(InitFreePrice200k())
		} else {
			return math.Floor(InitFreePriceEqualFreePrice(freePrice))
		}
	} else if CheckHaveCapitalOver50m(feeCapital) {

		return math.Floor(CaluculateHaveCapitalOver50m(feeCapital))

	}

	return math.Floor(0.00)

}
