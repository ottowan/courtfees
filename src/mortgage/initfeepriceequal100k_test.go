package mortgage

import "testing"

func Test_InitFeePriceEqual100k_Should_Be_100000(t *testing.T) {

	expectedFeePrice := 100000.00

	actualFeePrice := InitFeePriceEqual100k()

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got %v", expectedFeePrice, actualFeePrice)
	}

}
