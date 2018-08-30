package arbitration

import "testing"

func Test_CheckCapitalBetween50m1to100m_Input_feeCapital_750000000_Should_be_true(t *testing.T) {
	feeCapital := 75000000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween50m1to100m(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}
