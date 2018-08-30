package arbitration

import "testing"

func Test_CalculateCommissionBetween20m1to35mPersonOverOne_Input_feeCapital_27500000_Should_be_380000(t *testing.T) {
	feeCapital := 27500000.00
	expectedCommission := 380000.00
	actualCommission := CalculateCommissionBetween20m1to35mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
