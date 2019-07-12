// Ler e imprimir dados
// Ivo Dias

// Pacote principal
package main

// Importa a biblioteca
import( 
	"fmt"
	"bufio"
	"os"
)

// Funcao principal
func main() {
	// C style
	var idade int // Declara uma variavel inteira
	fmt.Println("Informe sua idade: ") // Escreve uma pergunta na tela
	fmt.Scanf("%v\n",&idade) // Salva o valor escrito na tela
	fmt.Printf("Sua idade eh %d\n",idade) // Mostra o valor

	// GO style
	reader := bufio.NewReader(os.Stdin) // Cria um leitor
	fmt.Println("Informe o seu nome: ") // Escreve a pergunta
	nome,err := reader.ReadString('\n') // Usa o leitor para salvar o valor
	if err != nil {
		fmt.Println(err) // Em caso de erro, mostra o erro
	} else {
		fmt.Println("Ola "+nome) // Se nao tiver, escreve o nome
	}

}