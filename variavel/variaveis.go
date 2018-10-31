//Nome do pacote da aplicação
package main

//Importando pacotes externos
import (
	//Pacote para trabalhar com formatação de textos
	"fmt"
)

//A função de inicialização do go, o nome da função tem que ser o mesmo que o do pacote
func main() {
	//Podemos definir o seu tipo
	//var nome int = "Paulo Bressan"
	//Ou deixar o Go gerenciar a tipagem das variaveis.
	// var nome = "Paulo Bressan"
	//O go infere na declaração e tipagem usando :=
	nome := "Paulo Bressan"
	//quando se trata de versão de variaveis, o go tipa a variavel com a maior versão, por exemplo float64
	versao := 1.1
	//A variavel depois da virgula vai ser concatenada com o texto
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa esta na versão", versao)
	//Usando reflect para saber o tipo de uma variavel
	// fmt.Println("O tipo da variavel é", reflect.TypeOf(idade))

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do programa")

	var comando int
	//Receber valores que o usuário digitou no terminal
	//Temos que definir o tipo ou modificador que vai ser recebido no primeiro parametro
	//O & é o caminho  variavel(Ponteiro)
	// fmt.Scanf("%d", &comando)
	//Scan simples utilizando o rcurso para definir o modficador de  acordo com a variavel
	fmt.Scan(&comando)
	fmt.Println("O comando escolhido foi", comando)
}
