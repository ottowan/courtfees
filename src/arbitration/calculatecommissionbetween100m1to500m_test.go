package arbitration

import "testing"

func Test_CalculateCommissionBetween100m1to500m_Input_feeCapital_250000000_Should_be_375000(t *testing.T) {
	feeCapital := 250000000.00
	expectedCommission := 375000.00
	actualCommission := CalculateCommissionBetween100m1to500m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
