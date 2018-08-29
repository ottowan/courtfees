package mortgage

import "testing"

func Test_InitFeePriceEqual1000_Should_Be_1000(t *testing.T) {

	expectedFeePrice := 1000.00

	actualFeePrice := InitFeePriceEqual1000()

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but got %v", expectedFeePrice, actualFeePrice)
	}

}
