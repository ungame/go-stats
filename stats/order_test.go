package stats

import (
	"sort"
	"testing"
)

func TestOrderByAsc(t *testing.T) {
	var (
		n        = []float64{5, 4, 3, 2, 1}
		expected = []float64{1, 2, 3, 4, 5}
		orderBy  = OrderBy(ORDER_BY_ASC, n)
	)

	sort.Slice(n, orderBy)
	for i := range expected {
		if expected[i] != n[i] {
			t.Errorf("unexpected slice ordered by asc: expected=%v, got=%v", expected, n)
		}
	}
}

func TestOrderByDesc(t *testing.T) {
	var (
		n        = []float64{1, 2, 3, 4, 5}
		expected = []float64{5, 4, 3, 2, 1}
		orderBy  = OrderBy(ORDER_BY_DESC, n)
	)

	sort.Slice(n, orderBy)
	for i := range expected {
		if expected[i] != n[i] {
			t.Errorf("unexpected slice ordered by desc: expected=%v, got=%v", expected, n)
		}
	}
}
