package nonandhavecapital

import "testing"

func Test_CalculateNonAndHaveCapital_feeCapital_1_Should_Be_200(t *testing.T) {
	feeCapital := 1.00
	expectedFeePrice := 200.00

	actualFeePrice := CalculateNonAndHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateNonAndHaveCapital_feeCapital_9999_Should_Be_200(t *testing.T) {
	feeCapital := 9999.00
	expectedFeePrice := 200.00

	actualFeePrice := CalculateNonAndHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}
func Test_CalculateNonAndHaveCapital_feeCapital_20000_Should_Be_400(t *testing.T) {
	feeCapital := 20000.00
	expectedFeePrice := 400.00

	actualFeePrice := CalculateNonAndHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateNonAndHaveCapital_feeCapital_250000_Should_Be_5000(t *testing.T) {
	feeCapital := 250000.00
	expectedFeePrice := 5000.00

	actualFeePrice := CalculateNonAndHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateNonAndHaveCapital_feeCapital_300001_Should_Be_6000(t *testing.T) {
	feeCapital := 300001.00
	expectedFeePrice := 6000.00

	actualFeePrice := CalculateNonAndHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateNonAndHaveCapital_feeCapital_1000000_Should_Be_20000(t *testing.T) {
	feeCapital := 1000000.00
	expectedFeePrice := 20000.00

	actualFeePrice := CalculateNonAndHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateNonAndHaveCapital_feeCapital_40000000_Should_Be_200000(t *testing.T) {
	feeCapital := 40000000.00
	expectedFeePrice := 200000.00

	actualFeePrice := CalculateNonAndHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateNonAndHaveCapital_feeCapital_50000000_Should_Be_200000(t *testing.T) {
	feeCapital := 50000000.00
	expectedFeePrice := 200000.00

	actualFeePrice := CalculateNonAndHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateNonAndHaveCapital_feeCapital_50000001_Should_Be_200000(t *testing.T) {
	feeCapital := 50000001.00
	expectedFeePrice := 200000.00

	actualFeePrice := CalculateNonAndHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateNonAndHaveCapital_feeCapital_526845694_Should_Be_676845(t *testing.T) {
	feeCapital := 526845694.00
	expectedFeePrice := 676845.00

	actualFeePrice := CalculateNonAndHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateNonAndHaveCapital_feeCapital_0_Should_Be_0(t *testing.T) {
	feeCapital := 0.00
	expectedFeePrice := 0.00

	actualFeePrice := CalculateNonAndHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}

func Test_CalculateNonAndHaveCapital_feeCapital_300000_Should_Be_0(t *testing.T) {
	feeCapital := 300000.00
	expectedFeePrice := 6000.00

	actualFeePrice := CalculateNonAndHaveCapital(feeCapital)

	if expectedFeePrice != actualFeePrice {
		t.Errorf("Expected %v but i got %v", expectedFeePrice, actualFeePrice)
	}
}
