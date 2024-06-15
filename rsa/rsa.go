package main

import (
	"fmt"
	"strconv"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func fastExponentiation(base, exponent, mod int) int {
	result := 1
	base = base % mod

	for exponent > 0 {
		// Si el exponente es impar, multiplica la base con el resultado
		if exponent%2 == 1 {
			result = (result * base) % mod
		}
		// Divide el exponente por 2
		exponent = exponent >> 1
		// Multiplica la base por sí misma
		base = (base * base) % mod
	}

	return result
}

func extendedGCD(a, b int) (int, int, int) {
	if b == 0 {
		return a, 1, 0
	}

	gcd, x1, y1 := extendedGCD(b, a%b)
	x := y1
	y := x1 - (a/b)*y1

	return gcd, x, y
}

func generarLlaves(p_a, q_a, e_a, p_b, q_b, e_b int) (int, int, int, int) {

	n_a := p_a * q_a
	n_b := p_b * q_b

	phi_a := (p_a - 1) * (q_a - 1)
	phi_b := (p_b - 1) * (q_b - 1)

	gcd_a, x_a, _ := extendedGCD(e_a, phi_a)
	gcd_b, x_b, _ := extendedGCD(e_b, phi_b)

	if x_a < 0 {
		x_a += phi_a
	}

	if x_b < 0 {
		x_b += phi_b
	}

	d_a := gcd_a * x_a
	d_b := gcd_b * x_b

	return d_a, n_a, d_b, n_b
}

func cifrar(n_b, e_b, mensaje int) int {
	texto_cifrado := fastExponentiation(mensaje, e_b, n_b)
	return texto_cifrado
}

func descifrar(d_b, n_b, texto_cifrado int) int {
	mensaje := fastExponentiation(texto_cifrado, d_b, n_b)
	return mensaje
}

func rsa(p_a, q_a, e_a, p_b, q_b, e_b, mensaje int) {
	d_a, _, d_b, n_b := generarLlaves(p_a, q_a, e_a, p_b, q_b, e_b)
	fmt.Println("d_a = ", d_a)
	fmt.Println("d_b = ", d_b)

	texto_cifrado := cifrar(n_b, e_b, mensaje)
	fmt.Println("texto Cifrado = ", texto_cifrado)

	texto_claro := descifrar(d_b, n_b, texto_cifrado)
	fmt.Println("texto claro = ", texto_claro)

	if texto_claro != mensaje {
		fmt.Println("Error en el cifrado o descifrado")
	}
}

func intToBinaryString(x int) string {
	return strconv.FormatInt(int64(x), 2)
}
func main() {
	//fmt.Println("Hola Mundo")

	//p_a, q_a, e_a, p_b, q_b, e_b, mensaje := 191, 223, 31, 191, 241, 17, 1234
	//rsa(p_a, q_a, e_a, p_b, q_b, e_b, mensaje)
	//resultado := generarLlaves(839, 947, 41, 761, 1019, 53)
	//fmt.Println("resultado: ", resultado)
	n, p := 17, 17
	//num := 27
	gcd_a, x_a, y := extendedGCD(n, p)
	if gcd_a != 1 {
		fmt.Println("no existe")
	}
	/*if x_a < 0 {
		x_a += p
	}*/
	//d_a := gcd_a * x_a
	fmt.Printf("Inverso %d = %d, %d, %d\n", n, x_a, y, gcd_a)
	//fmt.Println("num*inverso = ", d_a*num)
	//fmt.Println("division = ", (d_a*num)%p)
	/*
		for i := 1; i < 8; i++ {
			//p es el modulo
			n, p := i, 11
			//num := 27
			gcd_a, x_a, _ := extendedGCD(n, p)
			if gcd_a != 1 {
				fmt.Println("no existe")
			}
			if x_a < 0 {
				x_a += p
			}
			//d_a := gcd_a * x_a
			fmt.Printf("Inverso %d = %d\n", i, x_a)
			//fmt.Println("num*inverso = ", d_a*num)
			//fmt.Println("division = ", (d_a*num)%p)
		}
	*/
	//1:1, 10:101, 11:110, 100:111, 101:101, 110:11, 111:100
}
