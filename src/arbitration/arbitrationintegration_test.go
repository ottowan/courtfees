package arbitration

import "testing"

func Test_Arbitration_Input_feeCapital_0_Should_Be_6000(t *testing.T) {
	feeCapital := 0.00
	expectedfeePrice := 6000.00

	actualFeeCapital := CalculateArbitration(feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}
