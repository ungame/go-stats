package stats

import "testing"

func TestRolOrderedByAsc(t *testing.T) {
	imutable := []int{5, 4, 3, 2, 1}
	n := make([]int, len(imutable))
	copy(n, imutable)

	got, err := Rol(ORDER_BY_ASC, n...)
	if err != nil {
		t.Errorf("invalid rol: %s", err.Error())
	}

	for i := range imutable {
		if imutable[i] != n[i] {
			t.Errorf("rol should not source slice: expected=%v, got=%v", imutable, n)
		}
	}

	expected := []float64{1, 2, 3, 4, 5}
	for i := range expected {
		if expected[i] != got[i] {

		}
	}
}

func TestRolOrderedByDesc(t *testing.T) {
	imutable := []int{1, 2, 3, 4, 5}
	n := make([]int, len(imutable))
	copy(n, imutable)

	got, err := Rol(ORDER_BY_DESC, n...)
	if err != nil {
		t.Errorf("invalid rol: %s", err.Error())
	}

	for i := range imutable {
		if imutable[i] != n[i] {
			t.Errorf("rol should not change source slice: expected=%v, got=%v", imutable, n)
		}
	}

	expected := []float64{5, 4, 3, 2, 1}
	for i := range expected {
		if expected[i] != got[i] {

		}
	}
}
