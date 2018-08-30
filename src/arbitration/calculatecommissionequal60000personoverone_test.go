package arbitration

import "testing"

func Test_CalculateCommissionEqual60000PersonOverOne_Should_be_60000(t *testing.T) {
	expectedCommission := 60000.00

	actualCommission := CalculateCommissionEqual60000PersonOverOne()

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
