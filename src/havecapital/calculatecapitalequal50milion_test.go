package havecapital

import "testing"

func Test_CalculateCapitalEqual50Milion_Input_HaveCapital_50Milion_Shoud_Be_200k(t *testing.T) {
	expectedFeePrice := 200000.00
	feeType := "มีทุนทรัพย์"
	feeCapital := 50000000.00

	actualFeePrice := CalculateCapitalEqual50Milion(feeType, feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)
	}

}
