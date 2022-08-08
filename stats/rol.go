package stats

import (
	"fmt"
	"sort"
)

func Rol[N Number](o Order, n ...N) ([]float64, error) {
	l := len(n)
	if isInvalidRol(l) {
		return nil, fmt.Errorf("rol should have two or more numbers, but received: %v", n)
	}
	r := make([]float64, l)
	for i := 0; i < l; i++ {
		r[i] = float64(n[i])
	}
	sort.Slice(r, OrderBy(o, r))
	return r, nil
}

func IsInvalidRol(rol []float64) bool {
	return len(rol) < 2
}

func isInvalidRol(length int) bool {
	return length < 2
}
