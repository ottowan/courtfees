package arbitration

func CalculateCommissionOver2000m(feeCapital float64) float64 {
	return 1000000 + ((feeCapital - 2000000000) * 0.0002)
}
