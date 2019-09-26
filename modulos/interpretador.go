package modulos

import "fmt"

//tamanho máximo da memória
const (
	tam = 900000
)

//memória da linguagem
var memoria [tam]byte

//ponteiro
var ptr int

//Interpretar vai aceitar o texto em brainfuck como parâmetro
func Interpretar(s string) {
	//Brainfuck possui oito instruções
	//Para ler cada instrução, deveremos colocar a leitura da string dentro da estrutura de repetição
	for i := 0; i < len(s); i++ {
		var c = 0
		switch s[i] {
		//move o ponteiro da memória para a direita
		case '>':
			//se a memória estiver cheia
			if ptr == tam-1 {
				ptr = 0 //ponteiro volta para a posição zero
			} else {
				ptr++
			}
			break
		//Move o ponteiro da memória para a esquerda
		case '<':
			//Se o ponteiro for zero
			if ptr == 0 {
				//ponteiro será retornado para o extremo da direita
				ptr = tam - 1
			} else {
				ptr--
			}
			break
		// '+' incrementa o valor da célula da memória sobre o ponteiro
		case '+':
			memoria[ptr]++
			break
		// '-' decrementa esse valor
		case '-':
			memoria[ptr]--
			break
		//. printa o valor referente ao endereço na memória
		case '.':
			fmt.Printf("%s", string((memoria[ptr])))
			break
		//, entra um caracterer e insere na célula
		case ',':
			memoria[ptr] = byte([]byte(s)[0])
			break
		/*[pula de endereço caso o ponteiro for menor que zero ] */
		case '[':
			if memoria[ptr] == 0 {
				i++
				for c > 0 || s[i] != ']' {
					if s[i] == '[' {
						c++
					} else if s[i] == ']' {
						c--
					}
					i++
				}
			}
			break
		//Pula de votla se o ponteiro não for zero
		case ']':
			if memoria[ptr] != 0 {
				i--
				for c > 0 || s[i] != '[' {
					if s[i] == ']' {
						c++
					} else if s[i] == '[' {
						c--
					}
					i--
				}

			}
			break
		}

	}
}
