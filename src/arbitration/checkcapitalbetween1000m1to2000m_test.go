package arbitration

import "testing"

func Test_CheckCapitalBetween1000m1to2000m_Input_feeCapital_1500m_Should_be_true(t *testing.T) {
	feeCapital := 1500000000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween1000m1to2000m(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}
