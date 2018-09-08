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
