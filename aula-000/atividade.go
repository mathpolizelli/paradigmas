package main

import "fmt"

func main() {
	var input int

	fmt.Println("Digite um número para descobrir sua tabuada")
	fmt.Scanln(&input)
	
	for i:=1; i < 11; i++ {
	    fmt.Println(input, " x ", i, " = ", input*i)
	}
}