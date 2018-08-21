package mortgage

import "testing"

func Test_checkFeeCapitalOver50m_Input_feeCapital_100000000_Should_Be_True(t *testing.T) {

	feeCapital := 100000000.00
	expectedFeeCapital := true

	actualFeeCapital := checkFeeCapitalOver50m(feeCapital)

	if expectedFeeCapital != actualFeeCapital {
		t.Errorf("Expected %v but got %v", expectedFeeCapital, actualFeeCapital)
	}

}
