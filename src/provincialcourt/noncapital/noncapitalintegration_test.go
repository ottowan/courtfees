package noncapital

import "testing"

func Test_CalculateNonCapital_Input_feeType_noncapital_section_section288_1_feeCapital_0_Should_Be_0(t *testing.T) {
	section := "288(1)"
	feeCapital := 0.00

	expectedFeePrice := 0.00

	actualFeePrice := CalculateNonCapital(section, feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateNonCapital_Input_feeType_noncapital_section_section288_2_feeCapital_0_Should_Be_0(t *testing.T) {
	section := "288(2)"
	feeCapital := 0.00

	expectedFeePrice := 200.00

	actualFeePrice := CalculateNonCapital(section, feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}
func Test_CalculateNonCapital_Input_feeType_noncapital_section_section288_3_feeCapital_0_Should_Be_0(t *testing.T) {
	section := "288(3)"
	feeCapital := 0.00

	expectedFeePrice := 200.00

	actualFeePrice := CalculateNonCapital(section, feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}
func Test_CalculateNonCapital_Input_feeType_noncapital_section_blank_feeCapital_0_Should_Be_0(t *testing.T) {
	section := ""
	feeCapital := 200.00

	expectedFeePrice := 200.00

	actualFeePrice := CalculateNonCapital(section, feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}
