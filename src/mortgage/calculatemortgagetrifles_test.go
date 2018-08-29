package mortgage

import "testing"

func Test_CalculateMortgageTrifles_Input_feeCapital_500000_Should_Be_500(t *testing.T) {

	feeCapital := 500000.00
	expectedFeePrice := 500.00

	actualFeePrice := CalculateMortgageTrifles(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got %v", expectedFeePrice, actualFeePrice)
	}

}
