package mortgage

import "testing"

func Test_checkFeeTypeMortgage_Input_feeType_mortgage_Slould_Be_True(t *testing.T) {

	feeType := "บังคับจำนอง"
	expectedFeeType := true

	actualFeeType := checkFeeTypeMortgage(feeType)

	if expectedFeeType != actualFeeType {
		t.Errorf("Expected %v but got %v", expectedFeeType, actualFeeType)
	}
}
