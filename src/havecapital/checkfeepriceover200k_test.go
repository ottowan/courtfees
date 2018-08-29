package havecapital

import "testing"

func Test_CheckFeePriceOver200k_Input_feePrice_1000000_Should_Be_20000(t *testing.T) {
	expectedFeePrice := true
	feePrice := 1000000.00

	actualFeePrice := CheckFeePriceOver200k(feePrice)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)
	}
}
