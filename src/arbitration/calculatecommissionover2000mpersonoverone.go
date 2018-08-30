package arbitration

func CalculateCommissionOver2000mPersonOverOne(feeCapital float64) float64 {
	return 2000000 + ((feeCapital - 2000000000) * 0.0004)
}
