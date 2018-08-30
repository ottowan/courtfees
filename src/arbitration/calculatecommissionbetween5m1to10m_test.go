package arbitration

import "testing"

func Test_CalculateCommissionBetween5m1to10m_Input_feeCapital_7500000_Should_be_80000(t *testing.T) {
	feeCapital := 7500000.00
	expectedCommission := 80000.00
	actualCommission := CalculateCommissionBetween5m1to10m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
