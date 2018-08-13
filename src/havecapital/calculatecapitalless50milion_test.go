package havecapital

import "testing"

func Test_CalculateCapitalLess50Milion_Input_HaveCapital_45680000_Shoud_Be_913600(t *testing.T) {
	expectedFeePrice := 913600.00
	feeType := "มีทุนทรัพย์"
	feeCapital := 45680000.00

	actualFeePrice := CalculateCapitalLess50Milion(feeType, feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)
	}

}
