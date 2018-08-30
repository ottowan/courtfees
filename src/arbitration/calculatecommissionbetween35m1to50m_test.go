package arbitration

import "testing"

func Test_CalculateCommissionBetween35m1to50m_Input_feeCapital_47500000_Should_be_245000(t *testing.T) {
	feeCapital := 47500000.00
	expectedCommission := 245000.00
	actualCommission := CalculateCommissionBetween35m1to50m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
