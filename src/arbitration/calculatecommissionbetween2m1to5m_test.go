package arbitration

import "testing"

func Test_CalculateCommissionBetween2m1to5m_Input_feeCapital_3500000_Should_be_45000(t *testing.T) {
	feeCapital := 3500000.00
	expectedCommission:= 45000.00
	actualCommission := CalculateCommissionBetween2m1to5m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
