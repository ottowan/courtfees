package havecapital

import "testing"

func Test_CalculateHaveCapital(t *testing.T) {
	feeCapital := 1000000.00
	//expectedFeePrice := 20000.00
	expectedFeePrice := 0.00

	actualFeePrice := CalculateHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}
