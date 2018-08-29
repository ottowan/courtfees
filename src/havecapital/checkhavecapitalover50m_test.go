package havecapital

import "testing"

func Test_CheckHaveCapitalOver50m_Input_feeCapital_526845694_Should_Be_true(t *testing.T){
	expectedResult := true
	feeCapital := 526845694.00

	actualResult := CheckHaveCapitalOver50m(feeCapital)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but i got %v", expectedResult, actualResult)
	}

	
}