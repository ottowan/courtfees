package arbitration

func CalculateCommissionBetween500m1to1000mPersonOverOne(feeCapital float64) float64 {
	return 1000000 + ((feeCapital - 500000000) * 0.0008)
}
