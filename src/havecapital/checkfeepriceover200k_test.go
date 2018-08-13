package havecapital

import "testing"

func Test_CheckFeePriceOver200k_Input_913600_Should_Be_200k(t *testing.T) {
	expectedFeePrice := 200000.00
	feePrice := 913600.00

	actualFeePrice := CheckFeePriceOver200k(feePrice)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)
	}
}
