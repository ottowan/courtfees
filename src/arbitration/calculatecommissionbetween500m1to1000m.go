package arbitration

func CalculateCommissionBetween500m1to1000m(feeCapital float64) float64 {
	return 500000 + ((feeCapital - 500000000) * 0.0004)
}
