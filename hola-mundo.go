package main

import "fmt"

//crear un objeto
type Gorra struct {
	marca  string
	color  string
	precio float32
	plana  bool
}

func main() {
	//llenar un objecto
	//op1
	/*var gorra_negra = Gorra{
	marca:  "Nike",
	color:  "Negro",
	precio: 10000,
	plana:  false}*/
	//op2
	var gorra_negra = Gorra{"Nike", "Negro", 10000, false}

	fmt.Println(gorra_negra)
	var result float32 = operacion(10, 6, "/")
	fmt.Println(result)
	fmt.Println(devolverTexto())
	fmt.Println(calcularPrecio(7, 5000, "$"))
	printParamsList("param1", "param2", "param3", "param4", "param5")
	//array
	var peliculas [3]string
	peliculas[0] = "Titanic"
	peliculas[1] = "Goku"
	peliculas[2] = "Pantera"

	peliculas2 := [3]string{
		"Titanic", "Goku", "Pantera"}

	fmt.Println(peliculas2)

	var peliculas3 [3][2]string
	peliculas3[0][0] = "Up"
	peliculas3[0][1] = "Harry potter"

	peliculas3[1][0] = "dragonball Z"
	peliculas3[1][1] = "moppes"

	peliculas3[2][0] = "el paseo 20"
	peliculas3[2][1] = "la bruja"

	fmt.Println(peliculas3[2][1])

	peliculas4 := []string{
		"Titanic", "Goku", "Pantera", "la bruja", "moppes", "Up"}

	
	peliculas4 = append(peliculas4,"add Pictures")
	fmt.Println(peliculas4)	
	fmt.Println("Total de paliculas",  len(peliculas4))	
	fmt.Println("3 primeras",  peliculas4[0:3])	
	//end array

}

func operacion(n1 float32, n2 float32, op string) float32 {
	var result float32
	switch op {
	case "+":
		result = n1 + n2
	case "-":
		result = n1 - n2
	case "*":
		result = n1 * n2
	case "/":
		result = n1 / n2
	}
	return result
}

/**
Devolver varios resultados
*/
//op 1
/*
func devolverTexto() (string, int) {
	dato1 := "Esteban"
	dato2 := 31
	return dato1, dato2
}*/

//op 2
func devolverTexto() (dato1 string, dato2 int) {
	dato1 = "Esteban"
	dato2 = 31
	return
}

/*
closures = ejecutar una func, dentro de una func
*/

func calcularPrecio(cantidad float32, precio float32, moneda string) (msj string, totPag float32, signo string) {
	subTotal := func() float32 {
		return cantidad * precio
	}
	msj = "El precio a pagar es: "
	totPag = subTotal()
	signo = moneda
	return
}

/*
* Ejemplo de parametros de Lista
*
*params : Lista de strings
*
 */
func printParamsList(params ...string) {
	for _, param := range params {
		fmt.Println(param)
	}
	/*
	for i := 0; i < len(params); i++ {
		fmt.Println(params[i])
	}
	*/

}
