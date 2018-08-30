package arbitration

import "testing"

func Test_CalculateCommissionBetween50m1to100mPersonOverOne_Input_feeCapital_75000000_Should_be_550000(t *testing.T) {
	feeCapital := 75000000.00
	expectedCommission := 550000.00
	actualCommission := CalculateCommissionBetween50m1to100mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
