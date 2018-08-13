package havecapital

func CheckFeePriceOver200k(feePrice float64) float64 {

	const feePriceConst = 200000.00
	if feePrice > feePriceConst {
		return feePriceConst
	} else {
		return feePrice
	}
}
