package havecapital

import "testing"

func Test_InitFreePrice1000_Should_Be_1000(t *testing.T) {
	expectedFreePrice := 1000.00

	actualFeePrice := InitFreePrice1000()

	if expectedFreePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFreePrice, actualFeePrice)
	}
}
