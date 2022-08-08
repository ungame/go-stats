package stats

import (
	"testing"
)

func TestMode(t *testing.T) {
	rol, err := Rol(ORDER_BY_ASC, 5, 3, 2, 4, 1, 3)
	if err != nil {
		t.Errorf("invalid rol: %s", err.Error())
	}
	expected := []float64{3}
	got := Mode(rol)
	for i := range expected {
		if expected[i] != got[i] {
			t.Errorf("unexpected mode result: expected=%v, got=%v", expected, got)
		}
	}
}

func TestModeBimodal(t *testing.T) {
	rol, err := Rol(ORDER_BY_ASC, 5, 3, 2, 4, 4, 1, 3)
	if err != nil {
		t.Errorf("invalid rol: %s", err.Error())
	}
	expected := []float64{3, 4}
	got := Mode(rol)
	for i := range expected {
		if expected[i] != got[i] {
			t.Errorf("unexpected mode result: expected=%v, got=%v", expected, got)
		}
	}
}
