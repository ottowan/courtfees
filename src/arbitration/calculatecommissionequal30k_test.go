package arbitration

import "testing"

func Test_CalculateCommissionEqual30k_Should_be_30000(t *testing.T) {
	expectedCommission := 30000.00

	actualCommission := CalculateCommissionEqual30k()

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
