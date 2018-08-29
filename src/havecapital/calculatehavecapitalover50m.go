package havecapital

func CaluculateHaveCapitalOver50m(feeCapital float64) float64 {
	return ((feeCapital - 50000000.00) * 0.001) + 200000.00
}
