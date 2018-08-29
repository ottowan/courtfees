package havecapital

import "testing"

func Test_InitFreePrice200k(t *testing.T) {
	expectedFreePrice := 200000.00
	//feePrice := 800000.00


	actualFeePrice := InitFreePrice200k()

	if expectedFreePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFreePrice, actualFeePrice)
	}
}
