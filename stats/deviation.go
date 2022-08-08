package stats

import "math"

func Deviation(rol []float64) []float64 {
	if IsInvalidRol(rol) {
		return nil
	}
	mean := Mean(rol)
	deviation := make([]float64, 0, len(rol))
	for i := range rol {
		deviation = append(deviation, math.Abs(rol[i]-mean))
	}
	return deviation
}
