package mortgage

import "testing"

func Test_CheckFeePriceOver100k_Input_5000_Should_Be_true(t *testing.T) {

	feePrice := 5000.00
	expectedResult := false

	actualResult := CheckFeePriceOver100k(feePrice)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}

}
