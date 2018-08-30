package arbitration

func CalculateCommissionBetween2m1to5mPersonOverOne(feeCapital float64) float64 {
	return 60000 + ((feeCapital - 2000000) * 0.02)
}
