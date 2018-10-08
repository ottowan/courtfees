package mortgage

import "testing"

func Test_CheckFeeCapitalBetween1To300k_Input_feeCapital_50000_Should_Be_True(t *testing.T) {

	feeCapital := 50000.00
	expectedResult := true

	actualResult := CheckFeeCapitalBetween1To300k(feeCapital)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}

}

func Test_CheckFeePriceOver1000_Input_500_Should_Be_true(t *testing.T) {

	feePrice := 500.00
	expectedResult := false

	actualResult := CheckFeePriceOver1000(feePrice)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}

}

func Test_InitFeePriceEqual1000_Should_Be_1000(t *testing.T) {

	expectedFeePrice := 1000.00

	actualFeePrice := InitFeePriceEqual1000()

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_InitFeePriceEqualFeePrice_Input_500_Should_Be_500(t *testing.T) {
	feePrice := 500.00
	expectedFeePrice := 500.00

	actualFeePrice := InitFeePriceEqualFeePrice(feePrice)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got %v", expectedFeePrice, actualFeePrice)
	}
}
