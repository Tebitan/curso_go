package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//getParamsConsole()
	isParOImpar()
}

/**
*Captura los parametros de consola
*Comando : go run programa.go Pedro 5
*
*param nombre [0]
*param numeroEdad [1]
 */
func getParamsConsole() {
	nombre := os.Args[1]
	edadString := os.Args[2]
	// strconv.Atoi -> covert string to int
	numeroEdad, err := strconv.Atoi(edadString)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola " + nombre + " Bienvenido al programa de GO")
		if numeroEdad <= 17 {
			fmt.Println("Eres menor de edad")
		} else if numeroEdad <= 59 {
			fmt.Println("Eres mayor de edad")
		} else {
			fmt.Println("Eres un adulto mayor")
		}
	}
}

func isParOImpar() {
	numeroString := os.Args[1]
	fmt.Println("Numero ingresado : " + numeroString)
	// strconv.Atoi -> covert string to int
	numero, err := strconv.Atoi(numeroString)
	if err != nil {
		fmt.Println(err)
	} else {
		if numero%2 == 0 {
			fmt.Println("Es par")
		} else {
			fmt.Println("Es Impar")
		}
	}
}
