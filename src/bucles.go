package main

import "fmt"

func main() {
	//For condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Printf("\n")

	//For while
	counter := 0
	for counter < 10 {
		fmt.Println(counter)
		counter++
	}

	//For forever
	conunterForever := 0
	for {
		fmt.Println(conunterForever)
		conunterForever++
	}
}
