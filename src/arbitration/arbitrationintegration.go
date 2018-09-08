package arbitration

func CalculateArbitration(person int, feeCapital float64) float64 {

	if CheckPersonAmountOverOne(person) {
		return 55555.00
	} else {

		if CheckCapitalEqual0(feeCapital) {
			return CalculateCommissionEqual6000()
		} else if CheckCapitalBetween1to2m(feeCapital) {
			return CalculateCommissionEqual30k()
		} else if CheckCapitalBetween2m1to5m(feeCapital) {
			return CalculateCommissionBetween2m1to5m(feeCapital)
		} else if CheckCapitalBetween5m1to10m(feeCapital) {
			return CalculateCommissionBetween5m1to10m(feeCapital)
		} else if CheckCapitalBetween10m1to20m(feeCapital) {
			return CalculateCommissionBetween10m1to20m(feeCapital)
		}
	}
	return 0.00
}
