package arbitration

import "testing"

func Test_CalculateCommissionOver2000mPersonOverOne_Input_feeCapital_10000m_Should_be_5m2k(t *testing.T) {
	feeCapital := 10000000000.00
	expectedCommission := 5200000.00
	actualCommission := CalculateCommissionOver2000mPersonOverOne(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
