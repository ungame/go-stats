package stats

import (
	"testing"
)

func TestMean(t *testing.T) {
	rol, err := Rol(ORDER_BY_ASC, 5, 3, 1, 4, 2)
	if err != nil {
		t.Errorf("invalid rol: %s", err.Error())
	}
	expected := float64(3)
	got := Mean(rol)
	if expected != got {
		t.Errorf("unexpected mean result: expected=%v, got=%v", expected, got)
	}
}
