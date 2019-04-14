// Recebe um numero e mostra o dobro
package main

import "fmt"

func main() {
	// Recebe um valor e armazena em uma variavel
	fmt.Printf("Informe um numero: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2
	fmt.Println(output)
}
