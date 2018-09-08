package havecapital

import "testing"

func Test_CalculateHaveCapital_feeCapital_1000000_Should_Be_20000(t *testing.T) {
	feeCapital := 1000000.00
	expectedFeePrice := 20000.00

	actualFeePrice := CalculateHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateHaveCapital_feeCapital_40000000_Should_Be_200000(t *testing.T) {
	feeCapital := 40000000.00
	expectedFeePrice := 200000.00

	actualFeePrice := CalculateHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateHaveCapital_feeCapital_50000000_Should_Be_200000(t *testing.T) {
	feeCapital := 50000000.00
	expectedFeePrice := 200000.00

	actualFeePrice := CalculateHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateHaveCapital_feeCapital_300001_Should_Be_6000_02(t *testing.T) {
	feeCapital := 300001.00
	expectedFeePrice := 6000.02

	actualFeePrice := CalculateHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateHaveCapital_feeCapital_50000001_Should_Be_200000_001(t *testing.T) {
	feeCapital := 50000001.00
	expectedFeePrice := 200000.001

	actualFeePrice := CalculateHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateHaveCapital_feeCapital_526845694_Should_Be_676845_694(t *testing.T) {
	feeCapital := 526845694.00
	expectedFeePrice := 676845.694

	actualFeePrice := CalculateHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateHaveCapital_feeCapital_250000_Should_Be_1000(t *testing.T) {
	feeCapital := 250000.00
	expectedFeePrice := 1000.00

	actualFeePrice := CalculateHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateHaveCapital_feeCapital_20000_Should_Be_400(t *testing.T) {
	feeCapital := 20000.00
	expectedFeePrice := 400.00

	actualFeePrice := CalculateHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}
