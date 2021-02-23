package calculator_test

import (
	"calculator"
	"testing"
)

func TestAddSubMul(t *testing.T) {
	t.Parallel()
	tcs := []struct{
		a, b float64
		want float64
		op func(float64, float64) float64
	}{
		{ a: 4, b: 2, want: 6, op: calculator.Add },
		{ a: -2, b: 2, want: 0, op: calculator.Add },
		{ a: -4, b: -2, want: -6, op: calculator.Add },
		{ a: -4, b: -2, want: -2, op: calculator.Subtract },
		{ a: 2, b: 2, want: 0, op: calculator.Subtract },
		{ a: -9, b: 99, want: -108, op: calculator.Subtract },
		{ a: -4, b: -2, want: 8, op: calculator.Multiply },
		{ a: 2, b: 2, want: 4, op: calculator.Multiply },
		{ a: -9, b: 9, want: -81, op: calculator.Multiply },
	}
	for _, tc := range tcs {
		got := tc.op(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

