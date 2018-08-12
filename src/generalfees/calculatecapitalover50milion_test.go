package generalfees

import "testing"

func Test_CalculateCapitalOver50Milion_Input_HaveCapital_650Milion200k_Shoud_Be_800200(t *testing.T) {
	expectedFeePrice := 800200.00
	feeType := "มีทุนทรัพย์"
	feeCapital := 650200000.00

	actualFeePrice := CalculateCapitalOver50Milion(feeType, feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)
	}

}
