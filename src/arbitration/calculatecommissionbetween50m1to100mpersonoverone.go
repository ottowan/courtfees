package arbitration

func CalculateCommissionBetween50m1to100mPersonOverOne(feeCapital float64) float64 {
	return 500000 + ((feeCapital - 50000000) * 0.002)
}
