package havecapital

import "testing"

func Test_CalculateWritRate_Input_5_maelaos_chiangkham_phayao_Should_Be_500(t *testing.T) {
	expectedWritRate := 500.00
	moo := 2
	subDistrict := "แม่ลาว"
	district := "เชียงคำ"
	province := "พะเยา"

	actualWritRate := CalculateWritRate(moo, subDistrict, district, province)

	if expectedWritRate != actualWritRate {
		t.Errorf("Expected %v but got it %v", expectedWritRate, actualWritRate)
	}
}
