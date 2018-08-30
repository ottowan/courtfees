package arbitration

import "testing"

func Test_CalculateCommissionOver2000m_Input_feeCapital_10000m_Should_be_2m6k(t *testing.T) {
	feeCapital := 10000000000.00
	expectedCommission := 2600000.00
	actualCommission := CalculateCommissionOver2000m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
