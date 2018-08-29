package havecapital

import "testing"

func Test_CalculateHaveCapitalRate_Input_feeCapital_40000000_Should_Be_80000(t *testing.T) {
	expectedFeeprice := 800000.00
	feeCapital := 40000000.00

	actualFeeprice := CalculateHaveCapitalRate(feeCapital)

	if expectedFeeprice != actualFeeprice {
		t.Errorf("Expected %v but i got %v ", expectedFeeprice, actualFeeprice)
	}
}
