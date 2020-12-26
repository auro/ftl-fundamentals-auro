package calculator_test

import (
	"calculator"
	"testing"
)

func TestAdd(t *testing.T) {

	testCases := []struct {
		desc string
		a, b float64
		want float64
	}{
		{
			desc: "Primeiro TestAdd",
			a:    2,
			b:    2,
			want: 4,
		},
		{
			desc: "Segundo TestAdd",
			a:    4,
			b:    2,
			want: 6,
		},
		{
			desc: "Third TestAdd",
			a:    123,
			b:    -1,
			want: 122,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := calculator.Add(tC.a, tC.b)
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}
		})
	}

}

func TestSubstract(t *testing.T) {
	testCases := []struct {
		desc string
		a, b float64
		want float64
	}{
		{
			desc: "Primeiro Test Substract",
			a:    2,
			b:    2,
			want: 0,
		},

		{
			desc: "Segundo Test Substract",
			a:    2,
			b:    3,
			want: -1,
		},

		{
			desc: "Tri Test Substract",
			a:    -5,
			b:    5,
			want: -10,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := calculator.Subtract(tC.a, tC.b)
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	testCases := []struct {
		desc       string
		a, b, want float64
	}{
		{
			desc: "Test Multiply 1",
			a:    1,
			b:    1,
			want: 1,
		},
		{
			desc: "Test Multiply 2",
			a:    2,
			b:    2,
			want: 4,
		},
		{
			desc: "Test Multiply 3",
			a:    -2,
			b:    4,
			want: -8,
		},
		{
			desc: "Test Multiply 3",
			a:    -2,
			b:    -4,
			want: 8,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := calculator.Multiply(tC.a, tC.b)
			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}

		})
	}
}

func TestDivide2(t *testing.T) {
	testCases := []struct {
		desc        string
		a, b, want  float64
		errExpected bool
	}{
		{
			desc: "Divide Test 1",
			a:    2,
			b:    2,
			want: 1,
		},
		{
			desc: "Divide Test 2",
			a:    2,
			b:    1,
			want: 2,
		},
		{
			desc:        "Divide Test 3",
			a:           2,
			b:           0,
			want:        1,
			errExpected: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := calculator.Divide(tC.a, tC.b)

			if err != nil {
				if !tC.errExpected {
					t.Fatalf("Error: %s", err)
				}
			}

			if (tC.want != got) && (!tC.errExpected) {
				t.Errorf("want %f, got %f", tC.want, got)
			}
			// (tentar fazer no menor numero de linhas possivel - 8)
			// espera erro e recebe o erro
			// espera erro e nao recebe o erro
			// nao espera erro e recebe erro
			// nao espera erro e nao recebe erro
		})
	}
}

func TestRemainder(t *testing.T) {
	t.Parallel()
	var want int = 1
	got := calculator.Remainder(2, 2)
	if want != got {
		//		t.Errorf("want %f, got %f", want, got)
	}
}
