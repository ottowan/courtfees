package arbitration

import "testing"

func Test_CheckCapitalBetween30m1to50m_Input_47m500k_Shoud_Be_True(t *testing.T) {

	feeCapital := 47500000.00
	expectedResult := true

	actualResult := CheckCapitalBetween30m1to50m(feeCapital)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got it %v", expectedResult, actualResult)
	}

}
