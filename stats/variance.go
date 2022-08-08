package stats

import "math"

func Variance(rol []float64) float64 {
	deviation := Deviation(rol)
	variance := make([]float64, 0, len(deviation))
	for _, d := range deviation {
		variance = append(variance, math.Pow(d, 2))
	}
	return Mean(variance)
}
