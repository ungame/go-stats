package stats

func Mean(rol []float64) float64 {
	if IsInvalidRol(rol) {
		return 0
	}
	l := len(rol)
	return Sum(rol...) / float64(l)
}
