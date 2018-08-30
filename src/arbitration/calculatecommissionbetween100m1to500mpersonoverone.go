package arbitration

func CalculateCommissionBetween100m1to500mPersonOverOne(feeCapital float64) float64 {
	return 600000 + ((feeCapital - 100000000) * 0.001)
}
