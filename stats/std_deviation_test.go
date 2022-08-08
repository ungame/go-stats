package stats

import "testing"

func TestStandardDeviation(t *testing.T) {
	rol, err := Rol(ORDER_BY_ASC, 3, 7, 6, 5, 4)
	if err != nil {
		t.Errorf("invalid rol: %s", err.Error())
	}
	expected := 1.4142135623730951
	got := StandardDeviation(rol)
	if expected != got {
		t.Errorf("unexpected standard deviation result: expected=%v, got=%v", expected, got)
	}
}
