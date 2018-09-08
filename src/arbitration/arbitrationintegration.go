package arbitration

func CalculateArbitration(feeCapital float64) float64 {

	if CheckCapitalEqual0(feeCapital){
		return CalculateCommissionEqual6000()
	}

	return 0.00
}
