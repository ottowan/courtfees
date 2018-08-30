package arbitration

func CalculateCommissionBetween1000m1to2000m(feeCapital float64) float64 {
	return 700000 + ((feeCapital - 1000000000) * 0.0003)
}
