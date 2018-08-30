package arbitration

import "testing"

func Test_CalculateCommissionBetween500m1to1000m_Input_feeCapital_526845694_Should_be_510738dot28(t *testing.T) {
	feeCapital := 526845694.00
	expectedCommission := 510738.27760000003
	actualCommission := CalculateCommissionBetween500m1to1000m(feeCapital)

	if expectedCommission != actualCommission {
		t.Errorf("Expected %v but i got %v", expectedCommission, actualCommission)
	}
}
