package arbitration

func CalculateCommissionBetween20m1to35m(feeCapital float64) float64 {
	return 160000 + ((feeCapital - 20000000) * 0.004)
}
