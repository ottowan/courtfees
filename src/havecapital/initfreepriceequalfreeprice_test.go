package havecapital

import "testing"

func Test_InitFreePriceEqualFreePrice(t *testing.T) {
	expectedFreePrice := 20000.00
	feePrice := 20000.00


	actualFeePrice := InitFreePriceEqualFreePrice(feePrice)

	if expectedFreePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFreePrice, actualFeePrice)
	}
}
