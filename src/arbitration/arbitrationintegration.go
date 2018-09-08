package arbitration

func CalculateArbitration(person int, feeCapital float64) float64 {

	if CheckPersonAmountOverOne(person) {
		return 55555.00
	} else {

		if CheckCapitalEqual0(feeCapital) {
			return CalculateCommissionEqual6000()
		} else if CheckCapitalBetween1to2m(feeCapital) {
			return CalculateCommissionEqual30k()
		}
	}
	return 0.00
}
