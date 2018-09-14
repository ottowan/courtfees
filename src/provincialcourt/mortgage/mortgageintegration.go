package mortgage

import "math"

func CalculateMortgage(feeCapital float64) float64 {

	if CheckFeeCapitalBetween1to50m(feeCapital) {
		feePrice := CalculateMortgageBetween1to50m(feeCapital)
		if CheckFeePriceOver100k(feePrice) {
			return math.Floor(InitFeePriceEqual100k())
		} else {
			return math.Floor(InitFeePriceEqualFeePrice(feePrice))
		}

	} else if CheckFeeCapitalOver50m(feeCapital) {

		return math.Floor(CalculateMortgageOver50m(feeCapital))

	}
	return math.Floor(0.00)
}
