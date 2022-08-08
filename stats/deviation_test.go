package stats

import "testing"

func TestDeviation(t *testing.T) {
	rol, err := Rol(ORDER_BY_ASC, 3, 7, 6, 5, 4)
	if err != nil {
		t.Errorf("invalid rol: %s", err.Error())
	}
	expected := []float64{2, 1, 0, 1, 2}
	got := Deviation(rol)
	for i := range expected {
		if expected[i] != got[i] {
			t.Errorf("unexpected deviation result: expected=%v, got=%v", expected, got)
		}
	}
}
