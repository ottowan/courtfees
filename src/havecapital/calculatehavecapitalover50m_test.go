package havecapital

import "testing"

func Test_CaluculateHaveCapitalOver50m_Input_feeCapital_526845694_Should_Be_10536913dot88(t *testing.T){

	expectedFeePrice := 676845.694
	feeCapital := 526845694.00

	actualCapital := CaluculateHaveCapitalOver50m(feeCapital)

	if expectedFeePrice != actualCapital {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualCapital)
	}

}