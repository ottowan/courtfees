package arbitration

import "testing"

func Test_CalculateCommissionBetween35m1to50mPersonOverOne_Input_feeCapital_47500000_Should_be_490000(t *testing.T) {
	feeCapital := 47500000.00
	expectedCommission := 490000.00
	actualCommission := CalculateCommissionBetween35m1to50mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
