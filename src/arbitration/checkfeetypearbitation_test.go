package arbitration

import "testing"

func Test_CheckFeeTypeArbitation_Input_arbritation_Should_Be_true(t *testing.T) {

	feeType := "อนุญาโต"
	expectedResult := true

	actualTResult := CheckFeeTypeArbitation(feeType)

	if expectedResult != actualTResult {
		t.Errorf("Expected %v but i got %v", expectedResult, actualTResult)
	}
}
