package arbitration

func CalculateCommissionBetween35m1to50mPersonOverOne(feeCapital float64) float64 {
	return 440000 + ((feeCapital - 35000000) * 0.004)
}
