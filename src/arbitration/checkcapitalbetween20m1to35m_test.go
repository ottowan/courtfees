package arbitration

import "testing"

func Test_CheckCapitalBetween20m1to35m_Input_27m500k_Shoud_Be_True(t *testing.T) {

	feeCapital := 27500000.00
	expectedResult := true

	actualResult := CheckCapitalBetween20m1to35m(feeCapital)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got it %v", expectedResult, actualResult)
	}

}
