package arbitration

import "testing"

func Test_CalculateCommissionBetween500m1to1000mPersonOverOne_Input_feeCapital_526845694_Should_be_510738dot28(t *testing.T) {
	feeCapital := 526845694.00
	expectedCommission := 1021476.5552000001
	actualCommission := CalculateCommissionBetween500m1to1000mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
