package arbitration

func CalculateArbitration(person int, feeCapital float64) float64 {

	if CheckPersonAmountOverOne(person) {
		if CheckCapitalEqual0PersonOverOne(feeCapital) {
			return CalculateCommissionEqual30kPersonOverOne()
		} else if CheckCapitalBetween1to2mPersonOverOne(feeCapital) {
			return CalculateCommissionEqual60000PersonOverOne()
		} else if CheckCapitalBetween2m1to5mPersonOverOne(feeCapital) {
			return CalculateCommissionBetween2m1to5mPersonOverOne(feeCapital)
		} else if CheckCapitalBetween5m1to10mPersonOverOne(feeCapital) {
			return CalculateCommissionBetween5m1to10mPersonOverOne(feeCapital)
		} else if CheckCapitalBetween10m1to20mPersonOverOne(feeCapital) {
			return CalculateCommissionBetween10m1to20mPersonOverOne(feeCapital)
		} else if CheckCapitalBetween20m1to35mPersonOverOne(feeCapital) {
			return CalculateCommissionBetween20m1to35mPersonOverOne(feeCapital)
		} else if CheckCapitalBetween35m1to50mPersonOverOne(feeCapital) {
			return CalculateCommissionBetween35m1to50mPersonOverOne(feeCapital)
		} else if CheckCapitalBetween50m1to100mPersonOverOne(feeCapital) {
			return CalculateCommissionBetween50m1to100mPersonOverOne(feeCapital)
		} else if CheckCapitalBetween100m1to500mPersonOverOne(feeCapital) {
			return CalculateCommissionBetween100m1to500mPersonOverOne(feeCapital)
		} else if CheckCapitalBetween500m1to1000mPersonOverOne(feeCapital) {
			return CalculateCommissionBetween500m1to1000mPersonOverOne(feeCapital)
		} else if CheckCapitalBetween1000m1to2000mPersonOverOne(feeCapital) {
			return CalculateCommissionBetween1000m1to2000mPersonOverOne(feeCapital)
		} else if CheckCapitalOver2000mPersonOverOne(feeCapital) {
			return CalculateCommissionOver2000mPersonOverOne(feeCapital)
		}

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
		} else if CheckCapitalOver2000m(feeCapital) {
			return CalculateCommissionOver2000m(feeCapital)
		}

	}
	return 0.00
}
