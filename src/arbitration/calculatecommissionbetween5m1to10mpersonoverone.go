package arbitration

func CalculateCommissionBetween5m1to10mPersonOverOne(feeCapital float64) float64 {
	return 120000 + ((feeCapital - 5000000) * 0.016)
}
