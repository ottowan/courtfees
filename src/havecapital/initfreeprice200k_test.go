package havecapital

import "testing"

func Test_InitFreePrice200k_Should_Be_200000(t *testing.T) {
	expectedFreePrice := 200000.00

	actualFeePrice := InitFreePrice200k()

	if expectedFreePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFreePrice, actualFeePrice)
	}
}
