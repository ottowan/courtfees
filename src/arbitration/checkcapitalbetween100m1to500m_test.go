package arbitration

import "testing"

func Test_CheckCapitalBetween100m1to500m_Input_feeCapital_250000000_Should_be_true(t *testing.T) {
	feeCapital := 250000000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween100m1to500m(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}
