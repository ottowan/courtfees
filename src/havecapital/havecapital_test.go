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

func Test_CalculateCapitalLess50Milion_Input_HaveCapital_45680000_Shoud_Be_913600(t *testing.T) {
	expectedFeePrice := 913600.00
	feeType := "มีทุนทรัพย์"
	feeCapital := 45680000.00

	actualFeePrice := CalculateCapitalLess50Milion(feeType, feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)
	}

}

func Test_CalculateCapitalOver50Milion_Input_HaveCapital_650Milion200k_Shoud_Be_800200(t *testing.T) {
	expectedFeePrice := 800200.00
	feeType := "มีทุนทรัพย์"
	feeCapital := 650200000.00

	actualFeePrice := CalculateCapitalOver50Milion(feeType, feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)
	}

}