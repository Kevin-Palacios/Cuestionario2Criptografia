package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(numeros ...string) string {
	//var1, var2 := 0b0010, 0b0110
	suma := 0b0
	for _, numero := range numeros {
		binario, err := strconv.ParseInt(numero, 2, 64)
		if err != nil {
			fmt.Println("Error al convertir la cadena binaria a entero:", err)
			return ""
		}
		suma = suma ^ int(binario)
	}

	return strconv.FormatInt(int64(suma), 2)
}

// Divide dos polinomios binarios representados como enteros y devuelve el cociente y el residuo
func dividePolinomiosBinarios(dividendo, divisor int) (int, int) {
	// Inicializa el cociente y el residuo
	cociente := 0
	residuo := dividendo

	// Obtén el número de bits del divisor y del dividendo
	numBitsDivisor := bitLen(divisor)
	numBitsDividendo := bitLen(dividendo)

	// Realiza la división
	for i := numBitsDividendo - numBitsDivisor; i >= 0; i-- {
		if (residuo & (1 << (i + numBitsDivisor - 1))) != 0 {
			residuo ^= divisor << i
			cociente |= 1 << i
		}
	}

	return cociente, residuo
}

// Calcula la longitud en bits de un número entero
func bitLen(x int) int {
	length := 0
	for x > 0 {
		x >>= 1
		length++
	}
	return length
}

// Convierte un número entero a una cadena binaria
func intToBinaryString(x int) string {
	return strconv.FormatInt(int64(x), 2)
}

// Convierte un número entero a una cadena binaria
func BinaryStringToInt(x string) int64 {
	numero, err := strconv.ParseInt(x, 2, 64)
	if err != nil {
		fmt.Println("Error al convertir la cadena binaria a entero:", err)
		return 0
	}
	return numero
}

func multiplicar(n1, n2 string) string {
	//n1, n2 := "0010", "0111"
	matriz := make([]string, len(n2))
	for i, k := len(n2)-1, 0; i >= 0; i, k = i-1, k+1 {
		//for j := 0; j < len(n1); j++ {
		sumando := strings.Repeat("0", k+len(n1))
		if n2[i] == '1' {

			sumando = n1 + strings.Repeat("0", k)
		}
		matriz[k] = sumando
		//}
	}
	//fmt.Println(matriz)
	producto := sumar(matriz...)
	mod := "1011"
	//fmt.Println(producto)
	if modulo(producto, mod) {
		if len(producto) <= len(n1) {

			return producto
		}
	}
	_, residuo := dividePolinomiosBinarios(int(BinaryStringToInt(producto)), int(BinaryStringToInt(mod)))
	//fmt.Println("cociente=", intToBinaryString(int(cociente)))
	//fmt.Println("residuo=", intToBinaryString(int(residuo)))
	return intToBinaryString(residuo)
}

func modulo(elemento, modulo string) bool {
	//modulo = "111"
	//elemento = "1000101"
	elemento_binario, err := strconv.ParseInt(elemento, 2, 64)
	if err != nil {
		fmt.Println("Error al convertir la cadena binaria a entero:", err)
		return false
	}
	modulo_binario, err := strconv.ParseInt(modulo, 2, 64)
	if err != nil {
		fmt.Println("Error al convertir la cadena binaria a entero:", err)
		return false
	}

	if elemento_binario <= modulo_binario {
		return true
	}

	return false
}

func main() {
	//a=x(a la 2) y b=x(a la 2)+1 o 1+x(a la 2)
	fmt.Println("Hola Mundo")
	//x, y := 0b0010, 0b0110
	f := 3 - 1
	x, y := "0001", "0000"
	a, b := "100", "101"
	fmt.Println("Lado1=", 2<<f)
	puntos := make([][]string, 0)
	for i := 0; i < 2<<f; i++ {
		x = intToBinaryString(i)
		for j := 0; j < 2<<f; j++ {
			y = intToBinaryString(j)
			lado1 := sumar(multiplicar(y, y), multiplicar(x, y))
			lado2 := sumar((multiplicar(multiplicar(x, x), x)), multiplicar(a, multiplicar(x, x)), b)
			//fmt.Println("Lado1=", lado1)
			//fmt.Println("Lado2=", lado2)
			if lado1 == lado2 {
				//fmt.Println("Son iguales owo")
				puntos = append(puntos, []string{x, y})
			}
		}
	}
	fmt.Println(puntos)
	//fmt.Println("multiplicar(y, y)", multiplicar(y, y))
	//fmt.Println("multiplicar(x, x)", multiplicar(x, x))
	/*
		//E := "y^2+xy=x^3+ax^2+b"
		arreglo := make([]string, 0)
		arreglo = append(arreglo, x, y)
		funcion := sumar(arreglo...)
		multiplicar(x, x)
		fmt.Println(funcion)
	*/
}
