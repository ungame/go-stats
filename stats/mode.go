package stats

import "sort"

func Mode(rol []float64) []float64 {
	if IsInvalidRol(rol) {
		return nil
	}

	repetitions := make(map[float64]uint16, len(rol))

	for _, val := range rol {
		if r, ok := repetitions[val]; ok {
			r++
			repetitions[val] = r
		} else {
			repetitions[val] = 0
		}
	}

	return getMode(repetitions)
}

func getMode(repetitions map[float64]uint16) []float64 {
	var m uint16
	for _, r := range repetitions {
		if r > m {
			m = r
		}
	}
	if m == 0 {
		return nil
	}
	mode := make([]float64, 0, 2)
	for v, r := range repetitions {
		if r == m {
			mode = append(mode, v)
		}
	}
	if len(mode) > 1 {
		sort.Slice(mode, OrderBy(ORDER_BY_ASC, mode))
	}
	return mode
}
