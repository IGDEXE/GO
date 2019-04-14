// Go By Example
// Variables
// Ivo Dias

// Define o pacote principal
package main

// Importa o pacote FMT
import "fmt"

// Chama a funcao principal
func main() {
	// Utiliza "var" para declarar uma variavel
	var inicio = "Hii"
	fmt.Println(inicio)
	// Eh possivel declarar varias de uma vez
	var b, c int = 1, 2
	fmt.Println(b, c)
	// Go consegue identificar o tipo
	var d = true
	fmt.Println(d)
	// Quando nao definido um valor, ela tem como padrao 0
	var e int
	fmt.Println(e)
	// Uma forma simples de declarar uma variavel eh utilizar :=
	f := "short"
	fmt.Println(f)
}
