package havecapital

const feeCapitalConst = 50000000.00
const feeTypeConst = "มีทุนทรัพย์"
const failFeePriceConst = 0.00

func CalculateCapitalEqual50Milion(feeType string, feeCapital float64) float64 {

	const feePrice = 200000.00

	if feeType == feeTypeConst && feeCapital == feeCapitalConst {
		return feePrice
	}

	return failFeePriceConst

}

func CalculateCapitalLess50Milion(feeType string, feeCapital float64) float64 {

	const feePrice = 200000.00

	if feeType == feeTypeConst && feeCapital < feeCapitalConst {
		return (feeCapital * 2) / 100
	}

	return failFeePriceConst

}


func CalculateCapitalOver50Milion(feeType string, feeCapital float64) float64 {

	const feePrice = 200000.00

	if feeType == feeTypeConst && feeCapital > feeCapitalConst {
		return feePrice + (((feeCapital - feeCapitalConst) * 0.1) / 100)
	}

	return failFeePriceConst

}