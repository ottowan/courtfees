package arbitration

import "testing"

func Test_CheckCapitalOver2000m_Input_feeCapital_10000m_Should_be_true(t *testing.T) {
	feeCapital := 10000000000.00
	expectedOutput := true
	actualOutput := CheckCapitalOver2000m(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}
