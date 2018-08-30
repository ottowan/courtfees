package arbitration

func CheckCapitalEqual0PersonOverOne(feeCapital float64) bool{
	if feeCapital == 0 {
		return true
	}
	return false
}