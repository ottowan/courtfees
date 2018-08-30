package arbitration

import "testing"

func Test_CalculateCommissionBetween50m1to100m_Input_feeCapital_75000000_Should_be_275000(t *testing.T) {
	feeCapital := 75000000.00
	expectedCommission := 275000.00
	actualCommission := CalculateCommissionBetween50m1to100m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
