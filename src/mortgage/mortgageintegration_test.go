package mortgage

import "testing"

func Test_CalculateMortgage_Input_feeCapital_500000_Sholud_Be_5000(t *testing.T) {
	feeCapital := 500000.00
	expectedFeeCapital := 5000.00
	actualFeeCapital := CalculateMortgage(feeCapital)

	if expectedFeeCapital != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedFeeCapital, actualFeeCapital)
	}
}

func Test_CalculateMortgage_Input_feeCapital_20000000_Sholud_Be_100000(t *testing.T) {
	feeCapital := 20000000.00
	expectedFeeCapital := 100000.00
	actualFeeCapital := CalculateMortgage(feeCapital)

	if expectedFeeCapital != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedFeeCapital, actualFeeCapital)
	}
}

func Test_CalculateMortgage_Input_feeCapital_20000000_Sholud_Be_100000(t *testing.T) {
	feeCapital := 20000000.00
	expectedFeeCapital := 100000.00
	actualFeeCapital := CalculateMortgage(feeCapital)

	if expectedFeeCapital != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedFeeCapital, actualFeeCapital)
	}
}
