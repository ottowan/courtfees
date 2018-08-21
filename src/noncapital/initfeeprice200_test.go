package noncapital

import "testing"

func Test_InitFeePrice200_Input__Should_Be_200(t *testing.T) {
	expectedFeePrice := 200

	actualFeePrice := InitFeePrice200()

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it  %v", expectedFeePrice, actualFeePrice)
	}
}
