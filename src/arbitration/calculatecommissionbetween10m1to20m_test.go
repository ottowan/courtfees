package arbitration

import "testing"

func Test_CalculateCommissionBetween10m1to20m_Input_feeCapital_15000000_Should_be_130000(t *testing.T) {
	feeCapital := 15000000.00
	expectedCommission := 130000.00
	actualCommission := CalculateCommissionBetween10m1to20m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
