package havecapital

import "testing"

func Test_CalculateCapitalEqual50Milion_Input_HaveCapital_50Milion_Shoud_Be_200k(t *testing.T) {
	expectedFeePrice := 200000.00
	feeType := "มีทุนทรัพย์"
	feeCapital := 50000000.00

	actualFeePrice := CalculateCapitalEqual50Milion(feeType, feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)
	}

}

func Test_CalculateCapitalLess50Milion_Input_HaveCapital_45680000_Shoud_Be_200k(t *testing.T) {
	expectedFeePrice := 200000.00
	feeType := "มีทุนทรัพย์"
	feeCapital := 45680000.00

	actualFeePrice := CalculateCapitalLess50Milion(feeType, feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)
	}

}

func Test_CheckFeePriceOver200k_Input_913600_Should_Be_200k(t *testing.T) {
	expectedFeePrice := 200000.00
	feePrice := 913600.00

	actualFeePrice := CheckFeePriceOver200k(feePrice)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)
	}
}
func Test_CalculateCapitalOver50Milion_Input_HaveCapital_650Milion200k_Shoud_Be_800200(t *testing.T) {
	expectedFeePrice := 800200.00
	feeType := "มีทุนทรัพย์"
	feeCapital := 650200000.00

	actualFeePrice := CalculateCapitalOver50Milion(feeType, feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)
	}

}

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

