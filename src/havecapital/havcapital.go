package havecapital

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