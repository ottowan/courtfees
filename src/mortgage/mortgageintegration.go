package mortgage

import "math"

func CalculateMortgage(feeCapital float64) float64 {

	if CheckFeeCapitalBetween300k1to50m(feeCapital) {
		feePrice := CalculateMortgageBetween300k1to50m(feeCapital)
		if CheckFeePriceOver100k(feePrice) {
			return math.Floor(InitFeePriceEqual100k())
		} else {
			return math.Floor(InitFeePriceEqualFeePrice(feePrice))
		}

	} else if CheckFeeCapitalOver50m(feeCapital) {

		return math.Floor(CalculateMortgageOver50m(feeCapital))

	} else if CheckFeeCapitalBetween1To300k(feeCapital) {

		feePrice := CalculateMortgageBetween1to300k(feeCapital)
		if CheckFeePriceOver1000(feePrice) {
			return math.Floor(InitFeePriceEqual1000())
		} else {
			return math.Floor(InitFeePriceEqualFeePrice(feePrice))
		}
	}

	return math.Floor(0.00)
}
