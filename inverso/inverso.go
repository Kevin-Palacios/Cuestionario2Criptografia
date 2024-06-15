package main

import (
	"fmt"
)

// Degree devuelve el grado de un polinomio binario representado como un entero
func Degree(p int) int {
	degree := -1
	for p > 0 {
		p >>= 1
		degree++
	}
	return degree
}

// Modulus realiza la operación módulo de un polinomio binario
func Modulus(p, q int) int {
	for Degree(p) >= Degree(q) {
		p ^= q << (Degree(p) - Degree(q))
	}
	return p
}

// EuclidesExtendido para polinomios binarios
func EuclidesExtendido(a, b int) (gcd, s, t int) {
	if b == 0 {
		return a, 1, 0
	}

	// Inicialización de los coeficientes para a y b
	s0, s1 := 1, 0
	t0, t1 := 0, 1

	for b != 0 {
		quotient := 0
		for Degree(a) >= Degree(b) {
			shift := Degree(a) - Degree(b)
			quotient ^= 1 << shift
			a ^= b << shift
		}
		a, b = b, a
		s0, s1 = s1, s0^multiplyPoly(s1, quotient)
		t0, t1 = t1, t0^multiplyPoly(t1, quotient)
	}

	return a, s0, t0
}

// multiplyPoly multiplica dos polinomios binarios
func multiplyPoly(p, q int) int {
	result := 0
	for q > 0 {
		if q&1 == 1 {
			result ^= p // Si el bit menos significativo de q es 1, añade p a result
		}
		p <<= 1 // Multiplica p por x
		q >>= 1 // Divide q por x
	}
	return result
}

func main() {
	// Ejemplo de uso del algoritmo de Euclides extendido para polinomios binarios
	a := 0b10011 // x^3 + x^2 + 1
	b := 0b1110  // x^2 + 1

	gcd, s, t := EuclidesExtendido(a, b)

	fmt.Printf("GCD: %b\n", gcd)
	fmt.Printf("s: %b\n", s)
	fmt.Printf("t: %b\n", t)

	for i := 1; i < 8; i++ {
		//p es el modulo
		n, p := i, 0b1011
		//num := 27
		gcd_a, x_a, y := EuclidesExtendido(n, p)
		if gcd_a != 1 {
			fmt.Println("no existe")
		}
		/*if x_a < 0 {
			x_a += p
		}*/
		//d_a := gcd_a * x_a
		fmt.Printf("Inverso %d = %d, %d\n", i, x_a, y)
		//fmt.Printf("Inversoy %d = %d\n", i, y)
		//fmt.Println("num*inverso = ", d_a*num)
		//fmt.Println("division = ", (d_a*num)%p)
	}
	//1:1, 10:101, 11:110, 100:111, 101:101, 110:11, 111:100
}
