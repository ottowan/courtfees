package arbitration

import "math"

func CalculateArbitration(person int, feeCapital float64) float64 {

	if CheckPersonAmountOverOne(person) {
		if CheckCapitalEqual0PersonOverOne(feeCapital) {
			return math.Floor(CalculateCommissionEqual30kPersonOverOne())
		} else if CheckCapitalBetween1to2mPersonOverOne(feeCapital) {
			return math.Floor(CalculateCommissionEqual60000PersonOverOne())
		} else if CheckCapitalBetween2m1to5mPersonOverOne(feeCapital) {
			return math.Floor(CalculateCommissionBetween2m1to5mPersonOverOne(feeCapital))
		} else if CheckCapitalBetween5m1to10mPersonOverOne(feeCapital) {
			return math.Floor(CalculateCommissionBetween5m1to10mPersonOverOne(feeCapital))
		} else if CheckCapitalBetween10m1to20mPersonOverOne(feeCapital) {
			return math.Floor(CalculateCommissionBetween10m1to20mPersonOverOne(feeCapital))
		} else if CheckCapitalBetween20m1to35mPersonOverOne(feeCapital) {
			return math.Floor(CalculateCommissionBetween20m1to35mPersonOverOne(feeCapital))
		} else if CheckCapitalBetween35m1to50mPersonOverOne(feeCapital) {
			return math.Floor(CalculateCommissionBetween35m1to50mPersonOverOne(feeCapital))
		} else if CheckCapitalBetween50m1to100mPersonOverOne(feeCapital) {
			return math.Floor(CalculateCommissionBetween50m1to100mPersonOverOne(feeCapital))
		} else if CheckCapitalBetween100m1to500mPersonOverOne(feeCapital) {
			return math.Floor(CalculateCommissionBetween100m1to500mPersonOverOne(feeCapital))
		} else if CheckCapitalBetween500m1to1000mPersonOverOne(feeCapital) {
			return math.Floor(CalculateCommissionBetween500m1to1000mPersonOverOne(feeCapital))
		} else if CheckCapitalBetween1000m1to2000mPersonOverOne(feeCapital) {
			return math.Floor(CalculateCommissionBetween1000m1to2000mPersonOverOne(feeCapital))
		} else if CheckCapitalOver2000mPersonOverOne(feeCapital) {
			return math.Floor(CalculateCommissionOver2000mPersonOverOne(feeCapital))
		}

	} else {

		if CheckCapitalEqual0(feeCapital) {
			return math.Floor(CalculateCommissionEqual6000())
		} else if CheckCapitalBetween1to2m(feeCapital) {
			return math.Floor(CalculateCommissionEqual30k())
		} else if CheckCapitalBetween2m1to5m(feeCapital) {
			return math.Floor(CalculateCommissionBetween2m1to5m(feeCapital))
		} else if CheckCapitalBetween5m1to10m(feeCapital) {
			return math.Floor(CalculateCommissionBetween5m1to10m(feeCapital))
		} else if CheckCapitalBetween10m1to20m(feeCapital) {
			return math.Floor(CalculateCommissionBetween10m1to20m(feeCapital))
		} else if CheckCapitalBetween20m1to35m(feeCapital) {
			return math.Floor(CalculateCommissionBetween20m1to35m(feeCapital))
		} else if CheckCapitalBetween35m1to50m(feeCapital) {
			return math.Floor(CalculateCommissionBetween35m1to50m(feeCapital))
		} else if CheckCapitalBetween50m1to100m(feeCapital) {
			return math.Floor(CalculateCommissionBetween50m1to100m(feeCapital))
		} else if CheckCapitalBetween100m1to500m(feeCapital) {
			return math.Floor(CalculateCommissionBetween100m1to500m(feeCapital))
		} else if CheckCapitalBetween500m1to1000m(feeCapital) {
			return math.Floor(CalculateCommissionBetween500m1to1000m(feeCapital))
		} else if CheckCapitalBetween1000m1to2000m(feeCapital) {
			return math.Floor(CalculateCommissionBetween1000m1to2000m(feeCapital))
		} else if CheckCapitalOver2000m(feeCapital) {
			return math.Floor(CalculateCommissionOver2000m(feeCapital))
		}

	}
	return math.Floor(0.00)
}
