package mortgage

func CalculateMortgage(feeCapital float64) float64 {

	if CheckFeeCapitalBetween300k1to50m(feeCapital) {
		feePrice := CalculateMortgageBetween300k1to50m(feeCapital)
		if CheckFeePriceOver100k(feePrice) {
			return InitFeePriceEqual100k()
		} else {
			return InitFeePriceEqualFeePrice(feePrice)
		}

	} else if CheckFeeCapitalOver50m(feeCapital) {
		return CalculateMortgageOver50m(feeCapital)
	}

	return 0.00
}
