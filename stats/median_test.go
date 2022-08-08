package stats

import "testing"

func TestMedianOddLength(t *testing.T) {
	rol, err := Rol(ORDER_BY_ASC, 3, 1, 5, 2, 4)
	if err != nil {
		t.Errorf("invalid rol: %s", err.Error())
	}
	expected := float64(3)
	got := Median(rol)
	if expected != got {
		t.Errorf("unexpected median result: expected=%v, got=%v", expected, got)
	}
}

func TestMedianEvenLength(t *testing.T) {
	rol, err := Rol(ORDER_BY_ASC, 3, 1, 5, 3, 2, 4)
	if err != nil {
		t.Errorf("invalid rol: %s", err.Error())
	}
	expected := float64(3)
	got := Median(rol)
	if expected != got {
		t.Errorf("unexpected median result: expected=%v, got=%v", expected, got)
	}
}
