package noncapital

import "testing"
func Test_CheckSectionType288_1_Input_section_288_1_Should_Be_TRUE(t *testing.T){
	
	section := "288(1)"
	expectedSection := true
	actualSection := CheckSectionType288_1(section)

	if expectedSection != actualSection {
		t.Errorf("Expected %v but got it %v", expectedSection, actualSection)
	}


}

