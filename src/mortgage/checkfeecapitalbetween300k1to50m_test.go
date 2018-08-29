package mortgage

import "testing"

func Test_checkFeeCapitalBetween300kto50m_Input_feeCapital_500000_Should_Be_True(t *testing.T) {

	feeCapital := 500000.00
	expectedFeeCapital := true

	actualFeeCapital := checkFeeCapitalBetween300k1to50m(feeCapital)

	if expectedFeeCapital != actualFeeCapital {
		t.Errorf("Expected %v but got %v", expectedFeeCapital, actualFeeCapital)
	}

}
