package arbitration

import "testing"

func Test_CheckCapitalBetween1to2mPersonOverOne_Input_feeCapital_2000000_Should_be_true(t *testing.T) {
	feeCapital := 2000000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween1to2mPersonOverOne(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}
