package havecapital

import "testing"

func Test_CheckFeePriceOver1000_Input_feePrice_1000000_Should_Be_20000(t *testing.T) {
	expectedResult := true
	feePrice := 5000.00

	actualResult := CheckFeePriceOver1000(feePrice)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got it %v", expectedResult, actualResult)
	}
}
