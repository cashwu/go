package mymath

import "testing"

func TestAvg(t *testing.T) {

	if Avg([]float64{1,2}) != 1.5 {
		t.Error("expected 1.5, get", v)
	}

}
