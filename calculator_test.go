package calculator_test

import (
	"calculator"
	"math/rand"
	"testing"
)

func TestAdd2(t *testing.T) {
	testCases := []struct {
		desc string
		num  float64
	}{
		{
			desc: "Test Case Varying Function",
			num:  4,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

		})
	}
}

func TestAdd(t *testing.T) {

	testCases := []struct {
		desc string
		a, b float64
		want float64
	}{
		{
			desc: "2+2=4",
			a:    2,
			b:    2,
			want: 4,
		},
		{
			desc: "4+2=6",
			a:    4,
			b:    2,
			want: 6,
		},
		{
			desc: "123+(-1)=122",
			a:    123,
			b:    -1,
			want: 122,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := calculator.Add(tC.a, tC.b)
			if tC.want != got {
				t.Errorf("Error on %s: want %f, got %f", tC.desc, tC.want, got)
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

func TestDivide(t *testing.T) {
	testCases := []struct {
		desc        string
		a, b, want  float64
		errExpected bool
	}{
		{
			desc: "Divisão simples",
			a:    2,
			b:    2,
			want: 1,
		},
		{
			desc: "Divisão por 1",
			a:    2,
			b:    1,
			want: 2,
		},
		{
			desc:        "Divisão por zero",
			a:           2,
			b:           0,
			want:        0,
			errExpected: true,
		},
		{
			desc:        "Divisão negativa",
			a:           1,
			b:           -1,
			want:        -1,
			errExpected: false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := calculator.Divide(tC.a, tC.b)

			receivedErr := err != nil

			if receivedErr != tC.errExpected {
				t.Fatalf("Unexpected error: %s", err)
			}

			if tC.want != got {
				t.Errorf("want %f, got %f", tC.want, got)
			}

		})
	}
}

func TestAddRandom(t *testing.T) {
	t.Parallel()

	var randomTests int = 100
	for testLoop := 0; testLoop < randomTests; testLoop++ {
		var a float64 = rand.Float64()
		var b float64 = rand.Float64()
		want := a + b
		got := calculator.Add(a, b)
		if want != got {
			t.Errorf("want %f, got %f", want, got)
		}

	}

}

func TestIsPrimo(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc              string
		number            int
		want, errExpected bool
	}{
		{
			desc:   "2 Primo",
			number: 2,
			want:   true,
		},
		{
			desc:   "3 Primo",
			number: 3,
			want:   true,
		},
		{
			desc:   "4 Nao Primo",
			number: 4,
			want:   false,
		},
		{
			desc:   "5 Primo",
			number: 5,
			want:   true,
		},
		{
			desc:   "6 Nao Primo",
			number: 6,
			want:   false,
		},
		{
			desc:   "7 Primo",
			number: 7,
			want:   true,
		},

		{
			desc:   "8 Nao Primo",
			number: 8,
			want:   false,
		},

		{
			desc:   "9 Nao Primo",
			number: 9,
			want:   false,
		},

		{
			desc:   "10 Nao Primo",
			number: 10,
			want:   false,
		},

		{
			desc:   "11 Primo",
			number: 11,
			want:   true,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := calculator.IsPrimo(tC.number)
			receivedErr := err != nil
			if receivedErr != tC.errExpected {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tC.want != got {
				t.Errorf("Unexpected result for number %d", tC.number)
			}

		})
	}

}

func TestRemainder(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		desc       string
		a, b, want int64
	}{
		{
			desc: "Remainder of 2/2",
			a:    2,
			b:    2,
			want: 0,
		},
		{
			desc: "Remainder of 3/2",
			a:    3,
			b:    2,
			want: 1,
		},
		{
			desc: "Remainder of 4/2",
			a:    4,
			b:    2,
			want: 0,
		},
		{
			desc: "Remainder of 5/2",
			a:    5,
			b:    2,
			want: 1,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := calculator.Remainder(tC.a, tC.b)
			if tC.want != got {
				t.Errorf("Unexpected %s: want %d got %d", tC.desc, tC.want, got)
			}
		})
	}
}

func TestSqrt(t *testing.T) {
	testCases := []struct {
		desc         string
		number, want float64
		errExpected  bool
	}{
		{
			desc:   "Raiz de 1",
			number: 1,
			want:   1,
		},
		{
			desc:   "Raiz de 4",
			number: 4,
			want:   2,
		},
		{
			desc:   "Raiz de 9",
			number: 9,
			want:   3,
		},
		{
			desc:   "Raiz de 16",
			number: 16,
			want:   4,
		},
		{
			desc:        "Raiz Negativa de 4",
			number:      -4,
			want:        0,
			errExpected: true,
		},
		{
			desc:        "Raiz Negativa de 9",
			number:      -9,
			want:        0,
			errExpected: true,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got, err := calculator.Sqrt(tC.number)
			receivedErr := err != nil
			if receivedErr != tC.errExpected {
				t.Fatalf("Error received from sqrt: %s", err)
			}
			if tC.want != got {
				t.Errorf("Unexpected %s: want %f got %f", tC.desc, tC.want, got)
			}

		})
	}
}
