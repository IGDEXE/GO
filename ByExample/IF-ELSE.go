// Go By Example
// IF-ELSE
// Ivo Dias

// Define o pacote principal
package main

// Importa o pacote FMT
import "fmt"

// Chama a funcao principal
func main() {
	// Verifica o resto
	if 7%2 == 0 {
		fmt.Println("7 eh par")
	} else {
		fmt.Println("7 eh impar")
	}
	// Pode utilizar sem o ELSE
	if 8%4 == 0 {
		fmt.Println("8 eh divisivel por 4")
	}
	// Podemos declarar variaveis antes das condicoes
	if num := 9; num < 0 {
		fmt.Println(num, "eh negativo")
	} else if num < 10 {
		fmt.Println(num, "tem 1 digito")
	} else {
		fmt.Println(num, "tem varios digitos")
	}
}

// Nao precisa colocar ()
