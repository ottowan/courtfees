package arbitration

import "testing"

func Test_CalculateCommissionBetween1000m1to2000mPersonOverOne_Input_feeCapital_1500m_Should_be_1m700k(t *testing.T) {
	feeCapital := 1500000000.00
	expectedCommission := 1700000.00
	actualCommission := CalculateCommissionBetween1000m1to2000mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
