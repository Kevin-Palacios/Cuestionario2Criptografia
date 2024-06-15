package main

import (
	"fmt"
	"math"
	"strconv"
)

func intToBinaryString(x int) string {
	return strconv.FormatInt(int64(x), 2)
}

func intToHexaString(x int) string {
	return strconv.FormatInt(int64(x), 16)
}

func HexaStringToInt(x string) int64 {
	numero, err := strconv.ParseInt(x, 16, 64)
	if err != nil {
		fmt.Println("Error al convertir la cadena binaria a entero:", err)
		return 0
	}
	return numero
}

func multiplicar(n1, n2 string) string {
	decimal, err := strconv.ParseInt(n2, 16, 64)
	if err != nil {
		fmt.Println("Error al convertir hexadecimal a decimal1:", err)
		return ""
	}
	// Convertir el número decimal a binario
	cadena := strconv.FormatInt(decimal, 2)
	sumandos := make([]string, 0)
	potencias := make([]string, 0)
	xtimeActual := n1
	xorActual := "00"
	for i, k := len(cadena)-1, 0; i >= 0; i, k = i-1, k+1 {
		potenciaF := int(math.Pow(2, float64(k+1)))
		potenciaString := intToHexaString(potenciaF)
		xtimeAnterior := xtimeActual
		if cadena[i] == '1' {
			sumandos = append(sumandos, xtimeActual)
			potencia := int(math.Pow(2, float64(k)))
			potencias = append(potencias, intToHexaString(potencia))
		}
		xtimeActual = xtime(xtimeActual)
		fmt.Printf("{%s} • {%s} = xtime(%s) = {%s}\n", n1, potenciaString, xtimeAnterior, xtimeActual)
	}
	fmt.Printf("{%s} • {%s} = {%s} • (", n1, n2, n1)
	for i, potencia := range potencias {
		if i == len(potencias)-1 {
			fmt.Printf("{%s})\n", potencia)
			continue
		}
		fmt.Printf("{%s} ⊕ ", potencia)
	}
	fmt.Printf("{%s} • {%s} = ", n1, n2)
	for i, sumando := range sumandos {
		if i == len(potencias)-1 {
			fmt.Printf("{%s})\n", sumando)
			continue
		}
		fmt.Printf("{%s} ⊕ ", sumando)
	}
	for _, v := range sumandos {
		aux := HexaStringToInt(xorActual)

		aux2 := HexaStringToInt(v)

		xorActual = strconv.FormatInt(aux^aux2, 16)
		if len(xorActual) == 1 {
			xorActual = "0" + xorActual
		}
	}
	fmt.Printf("{%s} • {%s} = {%s}\n", n1, n2, xorActual)
	return xorActual
}

func xtime(numero string) string {
	// Convertir el número hexadecimal a decimal
	decimal, err := strconv.ParseInt(numero, 16, 64)
	if err != nil {
		fmt.Println("Error al convertir hexadecimal a decimal4:", err)
		return ""
	}
	// Convertir el número decimal a binario
	corrimiento := decimal << 1
	binario := strconv.FormatInt(corrimiento, 2)
	//fmt.Println(binario)
	polinomio := 283
	resultado := corrimiento
	if len(binario) >= 9 {
		for i := 0; i < len(binario)-8; i++ {
			if binario[i] == '1' {

				resultado = int64(corrimiento) ^ int64(polinomio)
				break
			}
		}
	}
	hexadecimal := strconv.FormatInt(int64(resultado), 16)
	if len(hexadecimal) == 1 {
		hexadecimal = "0" + hexadecimal
	}
	return hexadecimal
}

func main() {
	//resultado := multiplicar("6f", "16")
	resultado := multiplicar("57", "13")
	fmt.Printf("resultado: {%s}", resultado)
}
