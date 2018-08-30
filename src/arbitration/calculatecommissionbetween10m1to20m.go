package arbitration

func CalculateCommissionBetween10m1to20m(feeCapital float64) float64 {
	return 100000 + ((feeCapital - 10000000) * 0.006)
}
