package arbitration

import "testing"

func Test_Arbitration_Input_feeCapital_0_Should_Be_6000(t *testing.T) {
	person := 1
	feeCapital := 0.00
	expectedfeePrice := 6000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}

func Test_Arbitration_Input_feeCapital_2000000_Should_Be_30000(t *testing.T) {
	person := 1
	feeCapital := 2000000.00
	expectedfeePrice := 30000.00

	actualFeeCapital := CalculateArbitration(person, feeCapital)

	if expectedfeePrice != actualFeeCapital {
		t.Errorf("Expected %v but i got %v", expectedfeePrice, actualFeeCapital)
	}
}
