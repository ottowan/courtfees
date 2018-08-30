package arbitration

func CalculateCommissionBetween5m1to10m(feeCapital float64) float64 {
	return 60000 + ((feeCapital - 5000000) * 0.008)
}
