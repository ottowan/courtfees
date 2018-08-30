package arbitration

import "testing"

func Test_CalculateCommissionBetween20m1to35m_Input_feeCapital_27500000_Should_be_190000(t *testing.T) {
	feeCapital := 27500000.00
	expectedCommission := 190000.00
	actualCommission := CalculateCommissionBetween20m1to35m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
