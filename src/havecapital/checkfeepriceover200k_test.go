package havecapital

import "testing"

func Test_CheckFeePriceOver200k(t *testing.T){
	expectedFeePrice := 200000.00
	feePrice := 1000000.00

	actualFeePrice := CheckFeePriceOver200k(feePrice)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)
	}
}