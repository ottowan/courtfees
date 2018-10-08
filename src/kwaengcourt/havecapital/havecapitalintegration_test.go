package havecapital

import "testing"

func Test_CalculateHaveCapital_feeCapital_300000_Should_Be_1000(t *testing.T) {
	feeCapital := 300000.00
	expectedFeePrice := 1000.00

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

func Test_CalculateHaveCapital_feeCapital_300001_Should_Be_0(t *testing.T) {
	feeCapital := 300001.00
	expectedFeePrice := 0.00

	actualFeePrice := CalculateHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}
