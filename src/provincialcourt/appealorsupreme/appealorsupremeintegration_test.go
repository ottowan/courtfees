package appealorsupreme

import "testing"

func Test_CalculateAppealOrSupreme_Should_Be_200(t *testing.T) {
	expectedFeePrice := 200.00

	actualFeePrice := CalculateAppealOrSupreme()

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)
	}
}
