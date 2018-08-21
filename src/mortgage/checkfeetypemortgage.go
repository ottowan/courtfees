package mortgage

func checkFeeTypeMortgage(feeType string) bool {

	feeTypeResult := "บังคับจำนอง"

	if feeTypeResult == feeType {

		return true
	}
	return false

}
