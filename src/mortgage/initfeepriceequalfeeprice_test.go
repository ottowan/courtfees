package mortgage

import "testing"

func Test_InitFeePriceEqualFeePrice_Input_500_Should_Be_500(t *testing.T) {
	feePrice := 500.00
	expectedFeePrice := 500.00

	actualFeePrice := InitFeePriceEqualFeePrice(feePrice)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got %v", expectedFeePrice, actualFeePrice)
	}

}
