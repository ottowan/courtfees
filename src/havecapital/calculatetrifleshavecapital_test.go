package havecapital

import "testing"

func Test_CalculateTriflesHaveCapital_Input_feeCapital_250000_Should_be_5000(t *testing.T){
	expectedFeePrice := 5000.00
	feeCapital := 250000.00

	actualFeePrice := CalculateTriflesHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}

}