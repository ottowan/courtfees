package mortgage

import "testing"

func Test_CalculateMortgageBetween1to50m_Input_feeCapital_500000_Should_Be_5000(t *testing.T) {

	feeCapital := 500000.00
	expectedFeePrice := 5000.00

	actualFeePrice := CalculateMortgageBetween1to50m(feeCapital)

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

func Test_CheckFeeCapitalBetween1To50m_Input_feeCapital_50000_Should_Be_True(t *testing.T) {

	feeCapital := 50000.00
	expectedResult := true

	actualResult := CheckFeeCapitalBetween1to50m(feeCapital)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got %v", expectedResult, actualResult)
	}

}

func Test_CheckFeeCapitalBetween300kto50m_Input_feeCapital_500000_Should_Be_True(t *testing.T) {

	feeCapital := 500000.00
	expectedFeeCapital := true

	actualFeeCapital := CheckFeeCapitalBetween1to50m(feeCapital)

	if expectedFeeCapital != actualFeeCapital {
		t.Errorf("Expected %v but got %v", expectedFeeCapital, actualFeeCapital)
	}

}

func Test_CheckFeeCapitalOver50m_Input_feeCapital_100000000_Should_Be_True(t *testing.T) {

	feeCapital := 100000000.00
	expectedFeeCapital := true

	actualFeeCapital := CheckFeeCapitalOver50m(feeCapital)

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

func Test_CheckFeeTypeMortgage_Input_feeType_mortgage_Slould_Be_True(t *testing.T) {

	feeType := "บังคับจำนอง"
	expectedFeeType := true

	actualFeeType := CheckFeeTypeMortgage(feeType)

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

func Test_InitFeePriceEqualFeePrice_Input_500_Should_Be_500(t *testing.T) {
	feePrice := 500.00
	expectedFeePrice := 500.00

	actualFeePrice := InitFeePriceEqualFeePrice(feePrice)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got %v", expectedFeePrice, actualFeePrice)
	}
}
