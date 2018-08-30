package arbitration

func CalculateCommissionBetween2m1to5m(feeCapital float64) float64 {
	return 30000 + ((feeCapital - 2000000) * 0.01)
}
