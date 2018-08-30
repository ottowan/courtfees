package arbitration

import "testing"

func Test_CheckCapitalBetween500m1to1000mPersonOverOne_Input_feeCapital_526845694_Should_be_true(t *testing.T) {
	feeCapital := 526845694.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween500m1to1000mPersonOverOne(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}
