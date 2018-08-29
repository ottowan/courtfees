package havecapital

import "testing"

func Test_CheckFeeTypeHaveCapital_Input_havecapital_Shoud_Be_true(t *testing.T) {
	expectedFeeType := true
	feeType := "มีทุนทรัพย์"

	actualFeeType := CheckFeeTypeHaveCapital(feeType)

	if expectedFeeType != actualFeeType {
		t.Errorf("Expected %v but got it %v", expectedFeeType, actualFeeType)
	}
}
