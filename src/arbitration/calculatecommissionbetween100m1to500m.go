package arbitration

func CalculateCommissionBetween100m1to500m(feeCapital float64) float64 {
	return 300000 + ((feeCapital - 100000000) * 0.0005)
}
