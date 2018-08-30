package arbitration

import "testing"

func Test_CheckPersonAmountOverOne_Input_1_Should_Be_false(t *testing.T){
		amount := 1
		expectedResult := false

		actualResult := CheckPersonAmountOverOne(amount)

		if expectedResult != actualResult {
			t.Errorf("Expected %v but i got %v", expectedResult, actualResult)
		}
}