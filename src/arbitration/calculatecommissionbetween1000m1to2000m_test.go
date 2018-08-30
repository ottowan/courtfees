package arbitration

import "testing"

func Test_CalculateCommissionBetween1000m1to2000m_Input_feeCapital_1500m_Should_be_850k(t *testing.T) {
	feeCapital := 1500000000.00
	expectedCommission := 850000.00
	actualCommission := CalculateCommissionBetween1000m1to2000m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
