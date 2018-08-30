package arbitration

import "testing"

func Test_CalculateCommissionEqual6000_Should_be_6000(t *testing.T) {
	expectedCommission := 6000.00

	actualCommission := CalculateCommissionEqual6000()

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
