// Instalar BitDefender
// Ivo Dias

// Pacote principal
package main

// Importa a biblioteca
import (
	"fmt"
	"os/exec"
	"os"
	"bufio"
)

// Funcao principal
func main() {
	// Caminho do executavel
	wizard := "\\"
	wizard += "\\"
	wizard += "srsvm030.seniorsolution.com.br\\Scripts\\SRE\\BitDefenderDC.bat"
	// Faz o procedimento
	fmt.Println("Iniciando instalacao do BitDefender") // Escreve na tela
	fmt.Println("Isso pode demorar alguns minutos, favor aguardar") // Escreve na tela
	exec.Command("cmd","/C",wizard).Run()
	// Encerra
	fmt.Println("Procedimento concluido") // Escreve na tela
	fmt.Println("Mais detalhes em: C:\\INFRA\\BitDefender\\LOG") // Escreve na tela
	fmt.Println("Aperte alguma tecla para encerrar") // Escreve na tela
	bufio.NewReader(os.Stdin).ReadBytes('\n') // Aguarda uma entrada de dados para fechar a tela
}