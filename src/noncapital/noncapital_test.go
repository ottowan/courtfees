package noncapital

import "testing"

func Test_CheckFeeTypeNonCapital_Input_feeType_noncapital_Should_Be_True(t *testing.T) {

	feeType := "ไม่มีทุนทรัพย์"
	expectedFeeType := true

	actualFeeType := CheckFeeTypeNonCapital(feeType)

	if expectedFeeType != actualFeeType {
		t.Errorf("Expected %v but got %v", expectedFeeType, actualFeeType)
	}
}
