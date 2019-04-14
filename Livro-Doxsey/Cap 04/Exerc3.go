// Exercicio 3, pag 46
// Para multiplos de 3 mostra Fizz, para os de 5 Buzz

package main

import "fmt"

func main() {
	i := 1
	for i <= 100 {
		if (i%3 == 0) && (i%5 == 0) {
			fmt.Println("FizzBuzz")

		} else if i%5 == 0 {
			fmt.Println("Buzz")

		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
		i++
	}
}
