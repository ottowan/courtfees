package mortgage

import "testing"

func Test_CalculateMortgageBetween300k1to50m_Input_feeCapital_500000_Should_Be_5000(t *testing.T) {

	feeCapital := 500000.00
	expectedFeePrice := 5000.00

	actualFeePrice := CalculateMortgageBetween300k1to50m(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got %v", expectedFeePrice, actualFeePrice)
	}

}
