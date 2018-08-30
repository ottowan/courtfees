package arbitration

import "testing"

func Test_CheckCapitalBetween2m1to5m_Input_feeCapital_3500000_Should_be_true(t *testing.T) {
	feeCapital := 3500000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween2m1to5m(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}
