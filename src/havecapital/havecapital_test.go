package havecapital

import "testing"

func Test_InitFreePrice1000_Should_Be_1000(t *testing.T) {
	expectedFreePrice := 1000.00

	actualFeePrice := InitFreePrice1000()

	if expectedFreePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFreePrice, actualFeePrice)
	}
}

func Test_InitFreePrice200k_Should_Be_200000(t *testing.T) {
	expectedFreePrice := 200000.00

	actualFeePrice := InitFreePrice200k()

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

func Test_CheckHaveCapitalOver50m_Input_feeCapital_526845694_Should_Be_true(t *testing.T) {
	expectedResult := true
	feeCapital := 526845694.00

	actualResult := CheckHaveCapitalOver50m(feeCapital)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but i got %v", expectedResult, actualResult)
	}

}

func Test_CheckHaveCapital300k1to50m_Input_feeCapital_40000000_Should_Be_true(t *testing.T) {
	expectedResult := true
	feeCapital := 40000000.00

	actualResult := CheckHaveCapital300k1to50m(feeCapital)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but i got %v", expectedResult, actualResult)
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

func Test_CheckFeeTypeHaveCapital_Input_feeType_havecapital_Shoud_Be_true(t *testing.T) {
	expectedFeeType := true
	feeType := "มีทุนทรัพย์"

	actualFeeType := CheckFeeTypeHaveCapital(feeType)

	if expectedFeeType != actualFeeType {
		t.Errorf("Expected %v but got it %v", expectedFeeType, actualFeeType)
	}
}

func Test_CheckFeePriceOver1000_Input_feePrice_1000000_Should_Be_20000(t *testing.T) {
	expectedResult := true
	feePrice := 5000.00

	actualResult := CheckFeePriceOver1000(feePrice)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got it %v", expectedResult, actualResult)
	}
}

func Test_CheckFeePriceOver200k_Input_feePrice_1000000_Should_Be_20000(t *testing.T) {
	expectedFeePrice := true
	feePrice := 1000000.00

	actualFeePrice := CheckFeePriceOver200k(feePrice)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateTriflesHaveCapital_Input_feeCapital_250000_Should_be_5000(t *testing.T) {
	expectedFeePrice := 5000.00
	feeCapital := 250000.00

	actualFeePrice := CalculateTriflesHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}

}

func Test_CalculateHaveCapitalRate_Input_feeCapital_40000000_Should_Be_80000(t *testing.T) {
	expectedFeeprice := 800000.00
	feeCapital := 40000000.00

	actualFeeprice := CalculateHaveCapitalRate(feeCapital)

	if expectedFeeprice != actualFeeprice {
		t.Errorf("Expected %v but i got %v ", expectedFeeprice, actualFeeprice)
	}
}

func Test_CaluculateHaveCapitalOver50m_Input_feeCapital_526845694_Should_Be_10536913dot88(t *testing.T){

	expectedFeePrice := 676845.694
	feeCapital := 526845694.00

	actualCapital := CaluculateHaveCapitalOver50m(feeCapital)

	if expectedFeePrice != actualCapital {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualCapital)
	}

}