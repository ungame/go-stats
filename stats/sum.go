package stats

func Sum(n ...float64) float64 {
	l := len(n)
	if l == 0 {
		return 0
	}
	if l == 1 {
		return n[0]
	}
	return n[0] + Sum(n[1:]...)
}
