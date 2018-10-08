package noncapital

import "testing"

func Test_InitFeePrice0_Input_Should_Be_0(t *testing.T) {
	expectedFeePrice := 0.00

	actualFeePrice := InitFeePrice0()

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)

	}
}

func Test_InitFeePrice200_Input__Should_Be_200(t *testing.T) {
	expectedFeePrice := 200.00

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

func Test_CheckSectionType227_Input_section_227_Should_Be_TRUE(t *testing.T) {

	section := "227"
	expectedSection := true
	actualSection := CheckSectionType227(section)

	if expectedSection != actualSection {
		t.Errorf("Expected %v but got it %v", expectedSection, actualSection)
	}

}

func Test_CheckSectionType288_2_Input_section_288_2_Should_Be_TRUE(t *testing.T) {

	section := "288(2)"
	expectedSection := true
	actualSection := CheckSectionType288_2(section)

	if expectedSection != actualSection {
		t.Errorf("Expected %v but got it %v", expectedSection, actualSection)
	}

}

func Test_CheckSectionType288_3_Input_section_288_3_Should_Be_TRUE(t *testing.T) {

	section := "288(3)"
	expectedSection := true
	actualSection := CheckSectionType288_3(section)

	if expectedSection != actualSection {
		t.Errorf("Expected %v but got it %v", expectedSection, actualSection)
	}

}
