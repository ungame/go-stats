package stats

import "math"

func StandardDeviation(rol []float64) float64 {
	return math.Sqrt(Variance(rol))
}
