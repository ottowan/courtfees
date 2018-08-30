package arbitration

func CalculateCommissionBetween10m1to20mPersonOverOne(feeCapital float64) float64 {
	return 200000 + ((feeCapital - 10000000) * 0.012)
}
