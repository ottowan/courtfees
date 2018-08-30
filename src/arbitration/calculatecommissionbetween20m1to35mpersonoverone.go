package arbitration

func CalculateCommissionBetween20m1to35mPersonOverOne(feeCapital float64) float64 {
	return 320000 + ((feeCapital - 20000000) * 0.008)
}
