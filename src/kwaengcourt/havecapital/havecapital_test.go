package havecapital

import "testing"

func Test_CheckFeePriceOver1000_Input_feePrice_5000_Should_Be_true(t *testing.T) {
	expectedResult := true
	feePrice := 5000.00

	actualResult := CheckFeePriceOver1000(feePrice)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got it %v", expectedResult, actualResult)
	}
}

func Test_InitFreePrice1000_Should_Be_1000(t *testing.T) {
	expectedFreePrice := 1000.00

	actualFeePrice := InitFreePrice1000()

	if expectedFreePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFreePrice, actualFeePrice)
	}
}

func Test_InitFreePriceEqualFreePrice_Input_feePrice_20000_Should_be_20000(t *testing.T) {
	expectedFreePrice := 20000.00
	feePrice := 20000.00

	actualFeePrice := InitFreePriceEqualFreePrice(feePrice)

	if expectedFreePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFreePrice, actualFeePrice)
	}
}

func Test_CheckHaveCapital1to300k_Input_feeCapital_250000_Should_Be_true(t *testing.T) {
	expectedResult := true
	feeCapital := 250000.00

	actualResult := CheckHaveCapital1to300k(feeCapital)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but i got %v", expectedResult, actualResult)
	}
}

func Test_CalculateTriflesHaveCapital_Input_feeCapital_250000_Should_be_5000(t *testing.T) {
	expectedFeePrice := 5000.00
	feeCapital := 250000.00

	actualFeePrice := CalculateHaveCapital1to300k(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}

}
