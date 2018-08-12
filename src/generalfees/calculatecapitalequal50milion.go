package generalfees

func CalculateCapitalEqual50Milion(feeType string, feeCapital float64) float64 {

	const feeTypeConst = "มีทุนทรัพย์"
	const feeCapitalConst = 50000000.00
	const feePrice = 200000.00
	const failFeePrice = 0.00

	if feeType == feeTypeConst && feeCapital == feeCapitalConst {
		return feePrice
	}

	return failFeePrice

}
