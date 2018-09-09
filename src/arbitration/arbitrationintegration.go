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
		} else if CheckCapitalBetween20m1to35m(feeCapital) {
			return CalculateCommissionBetween20m1to35m(feeCapital)
		} else if CheckCapitalBetween35m1to50m(feeCapital) {
			return CalculateCommissionBetween35m1to50m(feeCapital)
		} else if CheckCapitalBetween50m1to100m(feeCapital) {
			return CalculateCommissionBetween50m1to100m(feeCapital)
		} else if CheckCapitalBetween100m1to500m(feeCapital) {
			return CalculateCommissionBetween100m1to500m(feeCapital)
		} else if CheckCapitalBetween500m1to1000m(feeCapital) {
			return CalculateCommissionBetween500m1to1000m(feeCapital)
		} else if CheckCapitalBetween1000m1to2000m(feeCapital) {
			return CalculateCommissionBetween1000m1to2000m(feeCapital)
		}

	}
	return 0.00
}
