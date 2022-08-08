package stats

import "testing"

func TestVariance(t *testing.T) {
	rol, err := Rol(ORDER_BY_ASC, 3, 7, 6, 5, 4)
	if err != nil {
		t.Errorf("invalid rol: %s", err.Error())
	}
	expected := float64(2)
	got := Variance(rol)
	if expected != got {
		t.Errorf("unexpected variance result: expected=%v, got=%v", expected, got)
	}
}
