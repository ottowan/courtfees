package havecapital

func CalculateHaveCapital(feeCapital float64) float64 {

	if CheckHaveCapital300k1to50m(feeCapital) {
		freePrice := CalculateHaveCapital300k1to50m(feeCapital)

		if CheckFeePriceOver200k(freePrice) {
			return InitFreePrice200k()
		} else {
			return InitFreePriceEqualFreePrice(freePrice)
		}
	} else if CheckHaveCapitalOver50m(feeCapital) {
		return CaluculateHaveCapitalOver50m(feeCapital)
	}
	return 0.00

}
