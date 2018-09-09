package arbitration

func CalculateCommissionBetween2m1to5m(feeCapital float64) float64 {
	return 30000 + ((feeCapital - 2000000) * 0.01)
}

func CalculateCommissionBetween2m1to5mPersonOverOne(feeCapital float64) float64 {
	return 60000 + ((feeCapital - 2000000) * 0.02)
}
func CalculateCommissionBetween5m1to10m(feeCapital float64) float64 {
	return 60000 + ((feeCapital - 5000000) * 0.008)
}

func CalculateCommissionBetween5m1to10mPersonOverOne(feeCapital float64) float64 {
	return 120000 + ((feeCapital - 5000000) * 0.016)
}

func CalculateCommissionBetween10m1to20m(feeCapital float64) float64 {
	return 100000 + ((feeCapital - 10000000) * 0.006)
}

func CalculateCommissionBetween10m1to20mPersonOverOne(feeCapital float64) float64 {
	return 200000 + ((feeCapital - 10000000) * 0.012)
}

func CalculateCommissionBetween20m1to35m(feeCapital float64) float64 {
	return 160000 + ((feeCapital - 20000000) * 0.004)
}

func CalculateCommissionBetween20m1to35mPersonOverOne(feeCapital float64) float64 {
	return 320000 + ((feeCapital - 20000000) * 0.008)
}

func CalculateCommissionBetween35m1to50m(feeCapital float64) float64 {
	return 220000 + ((feeCapital - 35000000) * 0.002)
}

func CalculateCommissionBetween35m1to50mPersonOverOne(feeCapital float64) float64 {
	return 440000 + ((feeCapital - 35000000) * 0.004)
}

func CalculateCommissionBetween50m1to100m(feeCapital float64) float64 {
	return 250000 + ((feeCapital - 50000000) * 0.001)
}

func CalculateCommissionBetween50m1to100mPersonOverOne(feeCapital float64) float64 {
	return 500000 + ((feeCapital - 50000000) * 0.002)
}

func CalculateCommissionBetween100m1to500m(feeCapital float64) float64 {
	return 300000 + ((feeCapital - 100000000) * 0.0005)
}

func CalculateCommissionBetween100m1to500mPersonOverOne(feeCapital float64) float64 {
	return 600000 + ((feeCapital - 100000000) * 0.001)
}

func CalculateCommissionBetween500m1to1000m(feeCapital float64) float64 {
	return 500000 + ((feeCapital - 500000000) * 0.0004)
}

func CalculateCommissionBetween500m1to1000mPersonOverOne(feeCapital float64) float64 {
	return 1000000 + ((feeCapital - 500000000) * 0.0008)
}

func CalculateCommissionBetween1000m1to2000m(feeCapital float64) float64 {
	return 700000 + ((feeCapital - 1000000000) * 0.0003)
}

func CalculateCommissionBetween1000m1to2000mPersonOverOne(feeCapital float64) float64 {
	return 1400000 + ((feeCapital - 1000000000) * 0.0006)
}

func CalculateCommissionEqual30k() float64 {
	return 30000.00
}

func CalculateCommissionEqual30kPersonOverOne() float64 {
	return 30000.00
}

func CalculateCommissionEqual6000() float64 {
	return 6000.00
}

func CalculateCommissionEqual60000PersonOverOne() float64 {
	return 60000.00
}

func CalculateCommissionOver2000m(feeCapital float64) float64 {
	return 1000000 + ((feeCapital - 2000000000) * 0.0002)
}

func CalculateCommissionOver2000mPersonOverOne(feeCapital float64) float64 {
	return 2000000 + ((feeCapital - 2000000000) * 0.0004)
}

func CheckCapitalBetween1to2m(feeCapital float64) bool {
	if feeCapital > 0 && feeCapital <= 2000000 {
		return true
	}

	return false
}

func CheckCapitalBetween1to2mPersonOverOne(feeCapital float64) bool {
	if feeCapital > 0 && feeCapital <= 2000000 {
		return true
	}

	return false
}

func CheckCapitalBetween2m1to5m(feeCapital float64) bool {
	if feeCapital > 2000000 && feeCapital <= 5000000 {
		return true
	}

	return false
}

func CheckCapitalBetween2m1to5mPersonOverOne(feeCapital float64) bool {
	if feeCapital > 2000000 && feeCapital <= 5000000 {
		return true
	}

	return false
}

func CheckCapitalBetween5m1to10m(feeCapital float64) bool {
	if feeCapital > 5000000 && feeCapital <= 10000000 {
		return true
	}

	return false
}

func CheckCapitalBetween5m1to10mPersonOverOne(feeCapital float64) bool {
	if feeCapital > 5000000 && feeCapital <= 10000000 {
		return true
	}

	return false
}

func CheckCapitalBetween10m1to20m(feeCapital float64) bool {
	if feeCapital > 10000000 && feeCapital <= 20000000 {
		return true
	}

	return false
}

func CheckCapitalBetween10m1to20mPersonOverOne(feeCapital float64) bool {
	if feeCapital > 10000000 && feeCapital <= 20000000 {
		return true
	}

	return false
}

func CheckCapitalBetween20m1to35m(feeCapital float64) bool {

	capitalMin := 20000000.00
	capitalMaX := 35000000.00

	if capitalMin < feeCapital && feeCapital <= capitalMaX {
		return true
	}

	return false
}

func CheckCapitalBetween20m1to35mPersonOverOne(feeCapital float64) bool {

	capitalMin := 20000000.00
	capitalMaX := 35000000.00

	if capitalMin < feeCapital && feeCapital <= capitalMaX {
		return true
	}

	return false
}

func CheckCapitalBetween35m1to50m(feeCapital float64) bool {

	capitalMin := 35000000.00
	capitalMaX := 50000000.00

	if capitalMin < feeCapital && feeCapital <= capitalMaX {
		return true
	}

	return false
}

func CheckCapitalBetween35m1to50mPersonOverOne(feeCapital float64) bool {

	capitalMin := 35000000.00
	capitalMaX := 50000000.00

	if capitalMin < feeCapital && feeCapital <= capitalMaX {
		return true
	}

	return false
}

func CheckCapitalBetween50m1to100m(feeCapital float64) bool {

	capitalMin := 50000000.00
	capitalMaX := 100000000.00

	if capitalMin < feeCapital && feeCapital <= capitalMaX {
		return true
	}

	return false
}

func CheckCapitalBetween50m1to100mPersonOverOne(feeCapital float64) bool {

	capitalMin := 50000000.00
	capitalMaX := 100000000.00

	if capitalMin < feeCapital && feeCapital <= capitalMaX {
		return true
	}

	return false
}

func CheckCapitalBetween100m1to500m(feeCapital float64) bool {
	if feeCapital > 100000000 && feeCapital <= 500000000 {
		return true
	}

	return false
}

func CheckCapitalBetween100m1to500mPersonOverOne(feeCapital float64) bool {
	if feeCapital > 100000000 && feeCapital <= 500000000 {
		return true
	}

	return false
}

func CheckCapitalBetween500m1to1000m(feeCapital float64) bool {
	if feeCapital > 500000000 && feeCapital <= 10000000000 {
		return true
	}

	return false
}

func CheckCapitalBetween500m1to1000mPersonOverOne(feeCapital float64) bool {
	if feeCapital > 500000000 && feeCapital <= 10000000000 {
		return true
	}

	return false
}

func CheckCapitalBetween1000m1to2000m(feeCapital float64) bool {
	if feeCapital > 1000000000 && feeCapital <= 20000000000 {
		return true
	}

	return false
}

func CheckCapitalBetween1000m1to2000mPersonOverOne(feeCapital float64) bool {
	if feeCapital > 1000000000 && feeCapital <= 20000000000 {
		return true
	}

	return false
}

func CheckCapitalEqual0(feeCapital float64) bool {
	if feeCapital == 0 {
		return true
	}

	return false
}

func CheckCapitalEqual0PersonOverOne(feeCapital float64) bool {
	if feeCapital == 0 {
		return true
	}
	return false
}

func CheckCapitalOver2000m(feeCapital float64) bool {
	if feeCapital > 2000000000 {
		return true
	}

	return false
}

func CheckCapitalOver2000mPersonOverOne(feeCapital float64) bool {
	if feeCapital > 2000000000 {
		return true
	}

	return false
}

func CheckFeeTypeArbitation(feeType string) bool {
	if feeType == "อนุญาโต" {
		return true
	}

	return false
}

func CheckPersonAmountOverOne(amount int) bool {
	if amount > 1 {
		return true
	}

	return false
}
