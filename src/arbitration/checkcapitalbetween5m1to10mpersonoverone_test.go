package arbitration

import "testing"

func Test_CheckCapitalBetween5m1to10mPersonOverOne_Input_feeCapital_7500000_Should_be_true(t *testing.T) {
	feeCapital := 7500000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween5m1to10mPersonOverOne(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}
