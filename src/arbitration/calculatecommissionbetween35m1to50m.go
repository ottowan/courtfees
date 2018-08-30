package arbitration

func CalculateCommissionBetween35m1to50m(feeCapital float64) float64 {
	return 220000 + ((feeCapital - 35000000) * 0.002)
}
