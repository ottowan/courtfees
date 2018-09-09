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

func Test_Arbitration_Input_feeCapital_27500000_Should_Be_190000(t *testing.T) {
	person := 1
	feeCapital := 27500000.00
	expectedfeePrice := 190000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_35000000_Should_Be_220000(t *testing.T) {
	person := 1
	feeCapital := 35000000.00
	expectedfeePrice := 220000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_47500000_Should_Be_245000(t *testing.T) {
	person := 1
	feeCapital := 47500000.00
	expectedfeePrice := 245000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_50000000_Should_Be_250000(t *testing.T) {
	person := 1
	feeCapital := 50000000.00
	expectedfeePrice := 250000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_75000000_Should_Be_275000(t *testing.T) {
	person := 1
	feeCapital := 75000000.00
	expectedfeePrice := 275000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_100000000_Should_Be_300000(t *testing.T) {
	person := 1
	feeCapital := 100000000.00
	expectedfeePrice := 300000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_250000000_Should_Be_375000(t *testing.T) {
	person := 1
	feeCapital := 250000000.00
	expectedfeePrice := 375000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_500000000_Should_Be_500000(t *testing.T) {
	person := 1
	feeCapital := 500000000.00
	expectedfeePrice := 500000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_526845694_Should_Be_510738_28(t *testing.T) {
	person := 1
	feeCapital := 526845694.00
	expectedfeePrice := 510738.27760000003

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_1000000000_Should_Be_700000(t *testing.T) {
	person := 1
	feeCapital := 1000000000.00
	expectedfeePrice := 700000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_1500000000_Should_Be_850000(t *testing.T) {
	person := 1
	feeCapital := 1500000000.00
	expectedfeePrice := 850000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_2000000000_Should_Be_1000000(t *testing.T) {
	person := 1
	feeCapital := 2000000000.00
	expectedfeePrice := 1000000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}
