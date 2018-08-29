package mortgage

import "testing"

func Test_CheckFeeCapitalBetween1To300k_Input_feeCapital_50000_Should_Be_True(t *testing.T) {

	feeCapital := 50000.00
	expectedResult := true

	actualResult := CheckFeeCapitalBetween1To300k(feeCapital)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}

}
