package mortgage

import "math"

func CalculateMortgage(feeCapital float64) float64 {

	if CheckFeeCapitalBetween1To300k(feeCapital) {

		feePrice := CalculateMortgageBetween1to300k(feeCapital)
		if CheckFeePriceOver1000(feePrice) {
			return math.Floor(InitFeePriceEqual1000())
		} else {
			return math.Floor(InitFeePriceEqualFeePrice(feePrice))
		}
	}

	return math.Floor(0.00)
}
