package main

import "fmt"

func main() {

	//Defer, siempre termina ultimo en codificacion
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	//continue y break
	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//break, termina donde enuentre valor
		if i == 8{
			fmt.Println("Break")
			break
		}


	}
}
