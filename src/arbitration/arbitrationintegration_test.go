package arbitration

import "testing"

func Test_Arbitration_Input_person_1_feeCapital_0_Should_Be_6000(t *testing.T) {
	person := 1
	feeCapital := 0.00
	expectedfeePrice := 6000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_1000000_Should_Be_30000(t *testing.T) {
	person := 1
	feeCapital := 1000000.00
	expectedfeePrice := 30000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}
func Test_Arbitration_Input_person_1_feeCapital_2000000_Should_Be_30000(t *testing.T) {
	person := 1
	feeCapital := 2000000.00
	expectedfeePrice := 30000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_3500000_Should_Be_45000(t *testing.T) {
	person := 1
	feeCapital := 3500000.00
	expectedfeePrice := 45000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_5000000_Should_Be_60000(t *testing.T) {
	person := 1
	feeCapital := 5000000.00
	expectedfeePrice := 60000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_7500000_Should_Be_80000(t *testing.T) {
	person := 1
	feeCapital := 7500000.00
	expectedfeePrice := 80000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_10000000_Should_Be_100000(t *testing.T) {
	person := 1
	feeCapital := 10000000.00
	expectedfeePrice := 100000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_15000000_Should_Be_130000(t *testing.T) {
	person := 1
	feeCapital := 15000000.00
	expectedfeePrice := 130000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_20000000_Should_Be_160000(t *testing.T) {
	person := 1
	feeCapital := 20000000.00
	expectedfeePrice := 160000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_27500000_Should_Be_190000(t *testing.T) {
	person := 1
	feeCapital := 27500000.00
	expectedfeePrice := 190000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_35000000_Should_Be_220000(t *testing.T) {
	person := 1
	feeCapital := 35000000.00
	expectedfeePrice := 220000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_47500000_Should_Be_245000(t *testing.T) {
	person := 1
	feeCapital := 47500000.00
	expectedfeePrice := 245000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_50000000_Should_Be_250000(t *testing.T) {
	person := 1
	feeCapital := 50000000.00
	expectedfeePrice := 250000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_75000000_Should_Be_275000(t *testing.T) {
	person := 1
	feeCapital := 75000000.00
	expectedfeePrice := 275000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_100000000_Should_Be_300000(t *testing.T) {
	person := 1
	feeCapital := 100000000.00
	expectedfeePrice := 300000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_250000000_Should_Be_375000(t *testing.T) {
	person := 1
	feeCapital := 250000000.00
	expectedfeePrice := 375000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_500000000_Should_Be_500000(t *testing.T) {
	person := 1
	feeCapital := 500000000.00
	expectedfeePrice := 500000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_526845694_Should_Be_510738(t *testing.T) {
	person := 1
	feeCapital := 526845694.00
	expectedfeePrice := 510738.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_1000000000_Should_Be_700000(t *testing.T) {
	person := 1
	feeCapital := 1000000000.00
	expectedfeePrice := 700000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_1500000000_Should_Be_850000(t *testing.T) {
	person := 1
	feeCapital := 1500000000.00
	expectedfeePrice := 850000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_2000000000_Should_Be_1000000(t *testing.T) {
	person := 1
	feeCapital := 2000000000.00
	expectedfeePrice := 1000000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_1_feeCapital_10000000000_Should_Be_2600000(t *testing.T) {
	person := 1
	feeCapital := 10000000000.00
	expectedfeePrice := 2600000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_0_Should_Be_30000(t *testing.T) {
	person := 3
	feeCapital := 0.00
	expectedfeePrice := 30000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_1000000_Should_Be_60000(t *testing.T) {
	person := 3
	feeCapital := 1000000.00
	expectedfeePrice := 60000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}
func Test_Arbitration_Input_person_3_feeCapital_2000000_Should_Be_60000(t *testing.T) {
	person := 3
	feeCapital := 2000000.00
	expectedfeePrice := 60000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_3500000_Should_Be_90000(t *testing.T) {
	person := 3
	feeCapital := 3500000.00
	expectedfeePrice := 90000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_5000000_Should_Be_120000(t *testing.T) {
	person := 3
	feeCapital := 5000000.00
	expectedfeePrice := 120000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_7500000_Should_Be_160000(t *testing.T) {
	person := 3
	feeCapital := 7500000.00
	expectedfeePrice := 160000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_10000000_Should_Be_200000(t *testing.T) {
	person := 3
	feeCapital := 10000000.00
	expectedfeePrice := 200000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_15000000_Should_Be_260000(t *testing.T) {
	person := 3
	feeCapital := 15000000.00
	expectedfeePrice := 260000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_20000000_Should_Be_320000(t *testing.T) {
	person := 3
	feeCapital := 20000000.00
	expectedfeePrice := 320000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_27500000_Should_Be_380000(t *testing.T) {
	person := 3
	feeCapital := 27500000.00
	expectedfeePrice := 380000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_35000000_Should_Be_440000(t *testing.T) {
	person := 3
	feeCapital := 35000000.00
	expectedfeePrice := 440000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_47500000_Should_Be_490000(t *testing.T) {
	person := 3
	feeCapital := 47500000.00
	expectedfeePrice := 490000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_50000000_Should_Be_500000(t *testing.T) {
	person := 3
	feeCapital := 50000000.00
	expectedfeePrice := 500000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_75000000_Should_Be_550000(t *testing.T) {
	person := 3
	feeCapital := 75000000.00
	expectedfeePrice := 550000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_100000000_Should_Be_600000(t *testing.T) {
	person := 3
	feeCapital := 100000000.00
	expectedfeePrice := 600000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_250000000_Should_Be_750000(t *testing.T) {
	person := 3
	feeCapital := 250000000.00
	expectedfeePrice := 750000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_500000000_Should_Be_1000000(t *testing.T) {
	person := 3
	feeCapital := 500000000.00
	expectedfeePrice := 1000000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_526845694_Should_Be_1021476(t *testing.T) {
	person := 3
	feeCapital := 526845694.00
	expectedfeePrice := 1021476.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_1000000000_Should_Be_1400000(t *testing.T) {
	person := 3
	feeCapital := 1000000000.00
	expectedfeePrice := 1400000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_1500000000_Should_Be_1700000(t *testing.T) {
	person := 3
	feeCapital := 1500000000.00
	expectedfeePrice := 1700000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_2000000000_Should_Be_2000000(t *testing.T) {
	person := 3
	feeCapital := 2000000000.00
	expectedfeePrice := 2000000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_person_3_feeCapital_10000000000_Should_Be_5200000(t *testing.T) {
	person := 3
	feeCapital := 10000000000.00
	expectedfeePrice := 5200000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}
