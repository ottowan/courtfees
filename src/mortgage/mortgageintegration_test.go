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

func Test_CalculateMortgage_Input_feeCapital_100000000_Sholud_Be_150000(t *testing.T) {
	feeCapital := 100000000.00
	expectedFeeCapital := 150000.00
	actualFeeCapital := CalculateMortgage(feeCapital)

	if expectedFeeCapital != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedFeeCapital, actualFeeCapital)
	}
}

func Test_CalculateMortgage_Input_feeCapital_150000_Sholud_Be_1000(t *testing.T) {
	feeCapital := 150000.00
	expectedFeeCapital := 1000.00
	actualFeeCapital := CalculateMortgage(feeCapital)

	if expectedFeeCapital != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedFeeCapital, actualFeeCapital)
	}
}

func Test_CalculateMortgage_Input_feeCapital_50000_Sholud_Be_500(t *testing.T) {
	feeCapital := 50000.00
	expectedFeeCapital := 500.00
	actualFeeCapital := CalculateMortgage(feeCapital)

	if expectedFeeCapital != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedFeeCapital, actualFeeCapital)
	}
}

func Test_CalculateMortgage_Input_feeCapital_300001_Sholud_Be_3000_01(t *testing.T) {
	feeCapital := 300001.00
	expectedFeeCapital := 3000.01
	actualFeeCapital := CalculateMortgage(feeCapital)

	if expectedFeeCapital != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedFeeCapital, actualFeeCapital)
	}
}

func Test_CalculateMortgage_Input_feeCapital_50000000_Sholud_Be_100000(t *testing.T) {
	feeCapital := 50000000.00
	expectedFeeCapital := 100000.00
	actualFeeCapital := CalculateMortgage(feeCapital)

	if expectedFeeCapital != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedFeeCapital, actualFeeCapital)
	}
}
