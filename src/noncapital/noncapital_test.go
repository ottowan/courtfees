package noncapital

import "testing"

func Test_CheckFeeTypeNonCapital_Input_feeType_noncapital_Should_Be_True(t *testing.T) {

	feeType := "ไม่มีทุนทรัพย์"
	expectedFeeType := true

	actualFeeType := CheckFeeTypeNonCapital(feeType)

	if expectedFeeType != actualFeeType {
		t.Errorf("Expected %v but got %v", expectedFeeType, actualFeeType)
	}
}

func Test_InitFeePrice0_Input_Should_Be_0(t *testing.T) {
	expectedFeePrice := 0

	actualFeePrice := InitFeePrice0()

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)

	}
}

func Test_InitFeePrice200_Input__Should_Be_200(t *testing.T) {
	expectedFeePrice := 200

	actualFeePrice := InitFeePrice200()

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it  %v", expectedFeePrice, actualFeePrice)
	}
}
func Test_CheckSectionType288_1_Input_section_288_1_Should_Be_TRUE(t *testing.T) {

	section := "288(1)"
	expectedSection := true
	actualSection := CheckSectionType288_1(section)

	if expectedSection != actualSection {
		t.Errorf("Expected %v but got it %v", expectedSection, actualSection)
	}

}
