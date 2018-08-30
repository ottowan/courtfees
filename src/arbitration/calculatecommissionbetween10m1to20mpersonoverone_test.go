package arbitration

import "testing"

func Test_CalculateCommissionBetween10m1to20mPersonOverOne_Input_feeCapital_15000000_Should_be_260000(t *testing.T) {
	feeCapital := 15000000.00
	expectedCommission := 260000.00
	actualCommission := CalculateCommissionBetween10m1to20mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
