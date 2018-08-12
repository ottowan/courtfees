package generalfees

const feeCapitalConst = 50000000.00
const feeTypeConst = "มีทุนทรัพย์"

func CalculateCapitalOver50Milion(feeType string, feeCapital float64) float64 {

	const feePrice = 200000.00
	const failFeePrice = 0.00

	if feeType == feeTypeConst && feeCapital > feeCapitalConst {
		return feePrice + (((feeCapital - feeCapitalConst) * 0.1) / 100)
	}

	return failFeePrice

}
