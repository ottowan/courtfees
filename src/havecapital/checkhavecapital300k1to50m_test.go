package havecapital

import "testing"

func Test_CheckHaveCapital300k1to50m_Input_feeCapital_40000000_Should_Be_true(t *testing.T) {
	expectedResult := true
	feeCapital := 40000000.00

	actualResult := CheckHaveCapital300k1to50m(feeCapital)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but i got %v", expectedResult, actualResult)
	}
}
