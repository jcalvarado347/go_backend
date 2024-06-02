package main

import "fmt"

func main() {

	//Area Cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area Cuadrado", areaCuadrado)

	x := 10
	y := 50

	//suma
	result := x + y
	fmt.Println("Suma", result)

	//resta
	result = y - x
	fmt.Println("Resta", result)

	//Multiplicacion
	result = x * y
	fmt.Println("Multiplicación", result)

	//Division
	result = y / x
	fmt.Println("Division", result)

	//Modulo
	result = y % x
	fmt.Println("Modulo", result)

	//Incremental
	x++
	fmt.Println("Incrementeal", x)

	//Decremento
	x--
	fmt.Println("Decremental", x)

}