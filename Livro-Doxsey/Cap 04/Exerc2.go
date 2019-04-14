// Exercicio 2, pag 46
// Mostra todos os multiplos de 3 ate 100

package main

import "fmt"

func main() {
	i := 1
	for i <= 100 {
		if i%3 == 0 {
			fmt.Println("O numero ", i, " é multiplo de 3")
		} else {
			fmt.Println("O numero ", i, " não é multiplo de 3")
		}
		i++
	}
}
