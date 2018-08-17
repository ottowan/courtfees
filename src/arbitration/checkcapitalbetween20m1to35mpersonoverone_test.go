package arbitration

import "testing"

func Test_CheckCapitalBetween20m1to35mPersonOverOne_Input_25m_Shoud_Be_True(t *testing.T){

	feeCapital := 25000000.00
	expectedResult := true

	actualResult := CheckCapitalBetween20m1to35mPersonOverOne(feeCapital)

	
	if expectedResult != actualResult {
		t.Errorf("Expected %v but got it %v", expectedResult, actualResult)
	}


}