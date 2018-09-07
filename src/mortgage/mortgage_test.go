package mortgage

import "testing"

func Test_CalculateMortgageBetween300k1to50m_Input_feeCapital_500000_Should_Be_5000(t *testing.T) {

	feeCapital := 500000.00
	expectedFeePrice := 5000.00

	actualFeePrice := CalculateMortgageBetween300k1to50m(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got %v", expectedFeePrice, actualFeePrice)
	}

}

func Test_CalculateMortgageOver50m_Input_feeCapital_100000000_Should_Be_150000(t *testing.T) {

	feeCapital := 100000000.00
	expectedFeePrice := 150000.00

	actualFeePrice := CalculateMortgageOver50m(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got %v", expectedFeePrice, actualFeePrice)
	}

}

func Test_CalculateMortgageTrifles_Input_feeCapital_500000_Should_Be_500(t *testing.T) {

	feeCapital := 500000.00
	expectedFeePrice := 500.00

	actualFeePrice := CalculateMortgageTrifles(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got %v", expectedFeePrice, actualFeePrice)
	}

}

func Test_CheckFeeCapitalBetween1To300k_Input_feeCapital_50000_Should_Be_True(t *testing.T) {

	feeCapital := 50000.00
	expectedResult := true

	actualResult := CheckFeeCapitalBetween1To300k(feeCapital)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}

}

func Test_checkFeeCapitalBetween300kto50m_Input_feeCapital_500000_Should_Be_True(t *testing.T) {

	feeCapital := 500000.00
	expectedFeeCapital := true

	actualFeeCapital := checkFeeCapitalBetween300k1to50m(feeCapital)

	if expectedFeeCapital != actualFeeCapital {
		t.Errorf("Expected %v but got %v", expectedFeeCapital, actualFeeCapital)
	}

}

func Test_checkFeeCapitalOver50m_Input_feeCapital_100000000_Should_Be_True(t *testing.T) {

	feeCapital := 100000000.00
	expectedFeeCapital := true

	actualFeeCapital := checkFeeCapitalOver50m(feeCapital)

	if expectedFeeCapital != actualFeeCapital {
		t.Errorf("Expected %v but got %v", expectedFeeCapital, actualFeeCapital)
	}

}

func Test_CheckFeePriceOver100k_Input_5000_Should_Be_true(t *testing.T) {

	feePrice := 5000.00
	expectedResult := false

	actualResult := CheckFeePriceOver100k(feePrice)

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

func Test_checkFeeTypeMortgage_Input_feeType_mortgage_Slould_Be_True(t *testing.T) {

	feeType := "บังคับจำนอง"
	expectedFeeType := true

	actualFeeType := checkFeeTypeMortgage(feeType)

	if expectedFeeType != actualFeeType {
		t.Errorf("Expected %v but got %v", expectedFeeType, actualFeeType)
	}
}

func Test_InitFeePriceEqual100k_Should_Be_100000(t *testing.T) {

	expectedFeePrice := 100000.00

	actualFeePrice := InitFeePriceEqual100k()

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got %v", expectedFeePrice, actualFeePrice)
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
