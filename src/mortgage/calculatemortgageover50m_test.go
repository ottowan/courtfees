package mortgage

import "testing"

func Test_CalculateMortgageOver50m_Input_feeCapital_100000000_Should_Be_150000(t *testing.T) {

	feeCapital := 100000000.00
	expectedFeePrice := 150000.00

	actualFeePrice := CalculateMortgageOver50m(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got %v", expectedFeePrice, actualFeePrice)
	}

}
