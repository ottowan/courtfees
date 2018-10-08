package appealorsupreme

import "testing"

func Test_CalculateCompensation_Should_Be_200(t *testing.T) {
	expectedFeePrice := 100.00

	actualFeePrice := CalculateCompensation()

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)
	}
}
