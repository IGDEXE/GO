// Exercicio 5, pag 38
// Converter temperatura

package main

import "fmt"

func main() {
	fmt.Printf("Informe a temperatura em Fahrenheit: ")
	var Fah float64
	fmt.Scanf("%f", &Fah)

	celsius := (Fah - 32) * 5 / 9
	fmt.Println("A temperatura em Celsius é", celsius)
}
