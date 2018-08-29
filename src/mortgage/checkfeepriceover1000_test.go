package mortgage

import "testing"

func Test_CheckFeePriceOver1000_Input_500_Should_Be_true(t *testing.T) {

	feePrice := 500.00
	expectedResult := false

	actualResult := CheckFeePriceOver1000(feePrice)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}

}
