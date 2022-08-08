package stats

func Median(rol []float64) float64 {
	if IsInvalidRol(rol) {
		return 0
	}
	l := len(rol)
	middle := l / 2
	if l%2 == 1 {
		return rol[middle]
	}
	return Mean([]float64{rol[middle-1], rol[middle]})
}
