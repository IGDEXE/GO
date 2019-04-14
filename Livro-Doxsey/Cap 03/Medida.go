// Exercicio 6, pagina 38
// Converter pes em metros

package main

import "fmt"

func main() {
	fmt.Printf("Informe a medida em pes: ")
	var pes float64
	fmt.Scanf("%f", &pes)

	metros := pes * 0.3048
	fmt.Println("A medida em metros é ", metros)
}
