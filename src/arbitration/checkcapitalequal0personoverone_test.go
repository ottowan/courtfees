package arbitration

import "testing"

func Test_CheckCapitalEqual0PersonOverOne_Input_feeCapital_0_Should_Be_true(t *testing.T) {
	feeCapital := 0.00
	expectedOutput := true
	actualOutput := CheckCapitalEqual0PersonOverOne(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}
