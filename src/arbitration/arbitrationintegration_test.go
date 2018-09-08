package arbitration

import "testing"

func Test_Arbitration_Input_feeCapital_0_Should_Be_6000(t *testing.T) {
	person := 1
	feeCapital := 0.00
	expectedfeePrice := 6000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_2000000_Should_Be_30000(t *testing.T) {
	person := 1
	feeCapital := 2000000.00
	expectedfeePrice := 30000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_3500000_Should_Be_45000(t *testing.T) {
	person := 1
	feeCapital := 3500000.00
	expectedfeePrice := 45000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_5000000_Should_Be_60000(t *testing.T) {
	person := 1
	feeCapital := 5000000.00
	expectedfeePrice := 60000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_7500000_Should_Be_80000(t *testing.T) {
	person := 1
	feeCapital := 7500000.00
	expectedfeePrice := 80000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_10000000_Should_Be_100000(t *testing.T) {
	person := 1
	feeCapital := 10000000.00
	expectedfeePrice := 100000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_15000000_Should_Be_130000(t *testing.T) {
	person := 1
	feeCapital := 15000000.00
	expectedfeePrice := 130000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_20000000_Should_Be_160000(t *testing.T) {
	person := 1
	feeCapital := 20000000.00
	expectedfeePrice := 160000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}
