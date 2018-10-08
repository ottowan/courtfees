package appealorsupreme

import "testing"

func Test_AppealOrSupreme_Should_Be_100(t *testing.T) {
	expectedFeePrice := 100.00

	actualFeePrice := InitFeePrice100()

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got it %v", expectedFeePrice, actualFeePrice)
	}
}
