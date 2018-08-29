package havecapital

import "testing"

func Test_CheckHaveCapital1to300k_Input_feeCapital_250000_Should_Be_true(t *testing.T) {
	expectedResult := true
	feeCapital := 250000.00

	actualResult := CheckHaveCapital1to300k(feeCapital)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but i got %v", expectedResult, actualResult)
	}
}
