package arbitration

func CalculateCommissionBetween50m1to100m(feeCapital float64) float64 {
	return 250000 + ((feeCapital - 50000000) * 0.001)
}
