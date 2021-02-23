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

func TestDivide(t *testing.T) {
	tcs := []struct{
		a, b float64
		want float64
		errExpected bool
	}{
		{ a: 4, b: 2, want: 2, errExpected: false },
		{ a: 4, b: 0, want: 999, errExpected: true },
		{ a: 2, b: 2, want: 1, errExpected: false },
		{ a: 100, b: 0, want: 999, errExpected: true },
	}
	for _, tc := range tcs {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := (err != nil)
		if tc.errExpected != errReceived {
			t.Fatalf("Divide(%f, %f): unexpected error status %v", tc.a, tc.b, err)
		}
		if !tc.errExpected && tc.want != got { // only for valid cases
			t.Errorf("Divide(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	/// ???
}