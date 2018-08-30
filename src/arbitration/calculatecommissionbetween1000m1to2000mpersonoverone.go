package arbitration

func CalculateCommissionBetween1000m1to2000mPersonOverOne(feeCapital float64) float64 {
	return 1400000 + ((feeCapital - 1000000000) * 0.0006)
}
