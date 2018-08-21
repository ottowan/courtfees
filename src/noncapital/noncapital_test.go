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

func Test_InitFeePrice0_Input_Should_Be_0(t *testing.T) {
	expectedFeePrice := 0

	actualFeePrice := InitFeePrice0()

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)

	}
}
