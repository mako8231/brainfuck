package main

import (
	"brainfuck/modulos"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("insira o código")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	modulos.Interpretar(input)
}
