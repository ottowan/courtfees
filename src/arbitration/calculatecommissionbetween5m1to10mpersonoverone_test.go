package arbitration

import "testing"

func Test_CalculateCommissionBetween5m1to10mPersonOverOne_Input_feeCapital_7500000_Should_be_160000(t *testing.T) {
	feeCapital := 7500000.00
	expectedCommission := 160000.00
	actualCommission := CalculateCommissionBetween5m1to10mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
