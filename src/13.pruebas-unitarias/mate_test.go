package main

import "testing"

func TestSumar(t *testing.T) {
	// result := Sumar(2, 3)
	// expect := 5

	// if result != expect {
	// 	t.Errorf("Error al sumar: se esperaba %d, pero se obtuvo %d", expect, result)
	// }

	casos := []struct {
		a, b, e int
	}{
		{3, 2, 5},
		{4, 5, 9},
		{456, 123, 579},
		{0, 0, 0},
		{-1, 1, 0},
	}

	for _, caso := range casos {
		r := Sumar(caso.a, caso.b)
		if r != caso.e {
			t.Errorf("Error al sumar: se esperaba %d, pero se obtuvo %d", caso.e, r)
		}
	}
}

func TestMayor(t *testing.T) {
	casos := []struct {
		a, b, e int
	}{
		{3, 2, 3},
		{4, 5, 5},
		{456, 123, 456},
		{0, 0, 0},
		{-1, 1, 1},
	}

	for _, caso := range casos {
		r := Mayor(caso.a, caso.b)
		if r != caso.e {
			t.Errorf("Error al comparar: se esperaba %d, pero se obtuvo %d", caso.e, r)
		}
	}
}

func TestFibonacci(t *testing.T) {
	casos := []struct {
		n, e int
	}{
		{0, 0},
		{1, 1},
		{7, 13},
		{40, 102334155},
	}

	for _, caso := range casos {
		r := Fibonacci(caso.n)
		if r != caso.e {
			t.Errorf("Error al calcular Fibonacci: se esperaba %d, pero se obtuvo %d", caso.e, r)
		}
	}

}
