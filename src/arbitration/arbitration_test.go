package arbitration

import "testing"

func Test_CalculateCommissionBetween2m1to5m_Input_feeCapital_3500000_Should_be_45000(t *testing.T) {
	feeCapital := 3500000.00
	expectedCommission := 45000.00
	actualCommission := CalculateCommissionBetween2m1to5m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionBetween2m1to5mPersonOverOne_Input_feeCapital_3500000_Should_be_90000(t *testing.T) {
	feeCapital := 3500000.00
	expectedCommission := 90000.00
	actualCommission := CalculateCommissionBetween2m1to5mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionBetween5m1to10m_Input_feeCapital_7500000_Should_be_80000(t *testing.T) {
	feeCapital := 7500000.00
	expectedCommission := 80000.00
	actualCommission := CalculateCommissionBetween5m1to10m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionBetween5m1to10mPersonOverOne_Input_feeCapital_7500000_Should_be_160000(t *testing.T) {
	feeCapital := 7500000.00
	expectedCommission := 160000.00
	actualCommission := CalculateCommissionBetween5m1to10mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionBetween10m1to20m_Input_feeCapital_15000000_Should_be_130000(t *testing.T) {
	feeCapital := 15000000.00
	expectedCommission := 130000.00
	actualCommission := CalculateCommissionBetween10m1to20m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionBetween10m1to20mPersonOverOne_Input_feeCapital_15000000_Should_be_260000(t *testing.T) {
	feeCapital := 15000000.00
	expectedCommission := 260000.00
	actualCommission := CalculateCommissionBetween10m1to20mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionBetween20m1to35m_Input_feeCapital_27500000_Should_be_190000(t *testing.T) {
	feeCapital := 27500000.00
	expectedCommission := 190000.00
	actualCommission := CalculateCommissionBetween20m1to35m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionBetween20m1to35mPersonOverOne_Input_feeCapital_27500000_Should_be_380000(t *testing.T) {
	feeCapital := 27500000.00
	expectedCommission := 380000.00
	actualCommission := CalculateCommissionBetween20m1to35mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionBetween35m1to50m_Input_feeCapital_47500000_Should_be_245000(t *testing.T) {
	feeCapital := 47500000.00
	expectedCommission := 245000.00
	actualCommission := CalculateCommissionBetween35m1to50m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionBetween35m1to50mPersonOverOne_Input_feeCapital_47500000_Should_be_490000(t *testing.T) {
	feeCapital := 47500000.00
	expectedCommission := 490000.00
	actualCommission := CalculateCommissionBetween35m1to50mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionBetween50m1to100m_Input_feeCapital_75000000_Should_be_275000(t *testing.T) {
	feeCapital := 75000000.00
	expectedCommission := 275000.00
	actualCommission := CalculateCommissionBetween50m1to100m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionBetween50m1to100mPersonOverOne_Input_feeCapital_75000000_Should_be_550000(t *testing.T) {
	feeCapital := 75000000.00
	expectedCommission := 550000.00
	actualCommission := CalculateCommissionBetween50m1to100mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionBetween100m1to500m_Input_feeCapital_250000000_Should_be_375000(t *testing.T) {
	feeCapital := 250000000.00
	expectedCommission := 375000.00
	actualCommission := CalculateCommissionBetween100m1to500m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionBetween100m1to500mPersonOverOne_Input_feeCapital_250000000_Should_be_750000(t *testing.T) {
	feeCapital := 250000000.00
	expectedCommission := 750000.00
	actualCommission := CalculateCommissionBetween100m1to500mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionBetween500m1to1000m_Input_feeCapital_526845694_Should_be_510738dot28(t *testing.T) {
	feeCapital := 526845694.00
	expectedCommission := 510738.27760000003
	actualCommission := CalculateCommissionBetween500m1to1000m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionBetween500m1to1000mPersonOverOne_Input_feeCapital_526845694_Should_be_510738dot28(t *testing.T) {
	feeCapital := 526845694.00
	expectedCommission := 1021476.5552000001
	actualCommission := CalculateCommissionBetween500m1to1000mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionBetween1000m1to2000m_Input_feeCapital_1500m_Should_be_850k(t *testing.T) {
	feeCapital := 1500000000.00
	expectedCommission := 850000.00
	actualCommission := CalculateCommissionBetween1000m1to2000m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionBetween1000m1to2000mPersonOverOne_Input_feeCapital_1500m_Should_be_1m700k(t *testing.T) {
	feeCapital := 1500000000.00
	expectedCommission := 1700000.00
	actualCommission := CalculateCommissionBetween1000m1to2000mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionEqual30k_Should_be_30000(t *testing.T) {
	expectedCommission := 30000.00

	actualCommission := CalculateCommissionEqual30k()

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionEqual30kPersonOverOne_Should_be_30000(t *testing.T) {
	expectedCommission := 30000.00

	actualCommission := CalculateCommissionEqual30kPersonOverOne()

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionEqual6000_Should_be_6000(t *testing.T) {
	expectedCommission := 6000.00

	actualCommission := CalculateCommissionEqual6000()

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionEqual60000PersonOverOne_Should_be_60000(t *testing.T) {
	expectedCommission := 60000.00

	actualCommission := CalculateCommissionEqual60000PersonOverOne()

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionOver2000m_Input_feeCapital_10000m_Should_be_2m6k(t *testing.T) {
	feeCapital := 10000000000.00
	expectedCommission := 2600000.00
	actualCommission := CalculateCommissionOver2000m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CalculateCommissionOver2000mPersonOverOne_Input_feeCapital_10000m_Should_be_5m2k(t *testing.T) {
	feeCapital := 10000000000.00
	expectedCommission := 5200000.00
	actualCommission := CalculateCommissionOver2000mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}

func Test_CheckCapitalBetween1to2m_Input_feeCapital_2000000_Should_be_true(t *testing.T) {
	feeCapital := 2000000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween1to2m(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalBetween1to2mPersonOverOne_Input_feeCapital_2000000_Should_be_true(t *testing.T) {
	feeCapital := 2000000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween1to2mPersonOverOne(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalBetween2m1to5m_Input_feeCapital_3500000_Should_be_true(t *testing.T) {
	feeCapital := 3500000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween2m1to5m(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalBetween2m1to5mPersonOverOne_Input_feeCapital_3500000_Should_be_true(t *testing.T) {
	feeCapital := 3500000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween2m1to5mPersonOverOne(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalBetween5m1to10m_Input_feeCapital_7500000_Should_be_true(t *testing.T) {
	feeCapital := 7500000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween5m1to10m(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalBetween5m1to10mPersonOverOne_Input_feeCapital_7500000_Should_be_true(t *testing.T) {
	feeCapital := 7500000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween5m1to10mPersonOverOne(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalBetween10m1to20m_Input_feeCapital_15000000_Should_be_true(t *testing.T) {
	feeCapital := 15000000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween10m1to20m(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}
func Test_CheckCapitalBetween10m1to20mPersonOverOne_Input_feeCapital_15000000_Should_be_true(t *testing.T) {
	feeCapital := 15000000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween10m1to20mPersonOverOne(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalBetween20m1to35m_Input_27m500k_Shoud_Be_True(t *testing.T) {

	feeCapital := 27500000.00
	expectedResult := true

	actualResult := CheckCapitalBetween20m1to35m(feeCapital)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got it %v", expectedResult, actualResult)
	}

}

func Test_CheckCapitalBetween20m1to35mPersonOverOne_Input_25m_Shoud_Be_True(t *testing.T) {

	feeCapital := 25000000.00
	expectedResult := true

	actualResult := CheckCapitalBetween20m1to35mPersonOverOne(feeCapital)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got it %v", expectedResult, actualResult)
	}

}

func Test_CheckCapitalBetween35m1to50m_Input_47m500k_Shoud_Be_True(t *testing.T) {

	feeCapital := 47500000.00
	expectedResult := true

	actualResult := CheckCapitalBetween35m1to50m(feeCapital)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got it %v", expectedResult, actualResult)
	}

}

func Test_CheckCapitalBetween35m1to50mPersonOverOne_Input_47m500k_Shoud_Be_True(t *testing.T) {

	feeCapital := 47500000.00
	expectedResult := true

	actualResult := CheckCapitalBetween35m1to50mPersonOverOne(feeCapital)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but got it %v", expectedResult, actualResult)
	}

}

func Test_CheckCapitalBetween50m1to100m_Input_feeCapital_750000000_Should_be_true(t *testing.T) {
	feeCapital := 75000000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween50m1to100m(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalBetween50m1to100mPersonOverOne_Input_feeCapital_750000000_Should_be_true(t *testing.T) {
	feeCapital := 75000000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween50m1to100mPersonOverOne(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalBetween100m1to500m_Input_feeCapital_250000000_Should_be_true(t *testing.T) {
	feeCapital := 250000000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween100m1to500m(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalBetween100m1to500mPersonOverOne_Input_feeCapital_250000000_Should_be_true(t *testing.T) {
	feeCapital := 250000000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween100m1to500mPersonOverOne(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalBetween500m1to1000m_Input_feeCapital_526845694_Should_be_true(t *testing.T) {
	feeCapital := 526845694.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween500m1to1000m(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalBetween500m1to1000mPersonOverOne_Input_feeCapital_526845694_Should_be_true(t *testing.T) {
	feeCapital := 526845694.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween500m1to1000mPersonOverOne(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalBetween1000m1to2000m_Input_feeCapital_1500m_Should_be_true(t *testing.T) {
	feeCapital := 1500000000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween1000m1to2000m(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalBetween1000m1to2000mPersonOverOne_Input_feeCapital_1500m_Should_be_true(t *testing.T) {
	feeCapital := 1500000000.00
	expectedOutput := true
	actualOutput := CheckCapitalBetween1000m1to2000mPersonOverOne(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalEqual0_Input_feeCapital_0_Should_be_true(t *testing.T) {
	feeCapital := 0.00
	expectedOutput := true
	actualOutput := CheckCapitalEqual0(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalEqual0PersonOverOne_Input_feeCapital_0_Should_Be_true(t *testing.T) {
	feeCapital := 0.00
	expectedOutput := true
	actualOutput := CheckCapitalEqual0PersonOverOne(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalOver2000m_Input_feeCapital_10000m_Should_be_true(t *testing.T) {
	feeCapital := 10000000000.00
	expectedOutput := true
	actualOutput := CheckCapitalOver2000m(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckCapitalOver2000mPersonOverOne_Input_feeCapital_10000m_Should_be_true(t *testing.T) {
	feeCapital := 10000000000.00
	expectedOutput := true
	actualOutput := CheckCapitalOver2000mPersonOverOne(feeCapital)

	if expectedOutput != actualOutput {
		t.Errorf("Expected %v but i got %v", expectedOutput, actualOutput)
	}
}

func Test_CheckFeeTypeArbitation_Input_arbritation_Should_Be_true(t *testing.T) {

	feeType := "อนุญาโต"
	expectedResult := true

	actualTResult := CheckFeeTypeArbitation(feeType)

	if expectedResult != actualTResult {
		t.Errorf("Expected %v but i got %v", expectedResult, actualTResult)
	}
}

func Test_CheckPersonAmountOverOne_Input_1_Should_Be_false(t *testing.T) {
	amount := 1
	expectedResult := false

	actualResult := CheckPersonAmountOverOne(amount)

	if expectedResult != actualResult {
		t.Errorf("Expected %v but i got %v", expectedResult, actualResult)
	}
}
