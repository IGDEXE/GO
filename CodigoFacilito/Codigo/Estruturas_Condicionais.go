// Estruturas Condicionais
// Ivo Dias

// Pacote principal
package main

// Importa a biblioteca
import "fmt"

/*
	Operadores:
		== igual
		!= diferente
		<  menor
		>  maior
		>= maior ou igual
		<= menor ou igual
		&& e
		|| ou
*/

// Funcao principal
func main() {
	// Declara as variaveis
	x := 10
	y := 5
	// Verifica qual numero eh maior
	if x > y {
		fmt.Printf("%d é maior que %d \n",x,y)
	} else if x < y {
		fmt.Printf("%d é maior que %d \n",y,x)
	} else {
		fmt.Printf("%d é igual a %d \n",x,y)
	}
}