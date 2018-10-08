package nonandhavecapital

import "math"

func CalculateNonAndHaveCapital(feeCapital float64) float64 {

	if CheckHaveCapital1to50m(feeCapital) {
		feePrice := CalculateHaveCapital1to50m(feeCapital)

		if CheckFeePriceOver200k(feePrice) {
			return math.Floor(InitFreePrice200k())
		} else if CheckFeePriceOver200(feePrice) {
			return math.Floor(InitFreePriceEqualFreePrice(feePrice))
		} else {
			return math.Floor(InitFreePrice200())
		}
	} else if CheckHaveCapitalOver50m(feeCapital) {

		return math.Floor(CaluculateHaveCapitalOver50m(feeCapital))

	}

	return math.Floor(0.00)

}
