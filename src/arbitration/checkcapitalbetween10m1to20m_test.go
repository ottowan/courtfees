package arbitration

import "testing"

func Test_CheckCapitalBetween10m1to20m_Input_feeCapital_15000000_Should_be_true(t *testing.T) {
	feeCapital := 15000000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween10m1to20m(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}
