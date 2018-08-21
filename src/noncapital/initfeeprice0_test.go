package noncapital

import "testing"

func Test_InitFeePrice0_Input_Should_Be_0(t *testing.T) {
	expectedFeePrice := 0

	actualFeePrice := InitFeePrice0()

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)

	}
}
