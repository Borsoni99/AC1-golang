package main

import (
	"fmt"
)

func hanoi(n int, origem string, destino string, auxiliar string) int {
	if n == 1 {
		fmt.Printf("Mover disco 1 de %s para %s\n", origem, destino)
		return 1
	}

	movimentos := hanoi(n-1, origem, auxiliar, destino)
	fmt.Printf("Mover disco %d de %s para %s\n", n, origem, destino)
	movimentos++
	movimentos += hanoi(n-1, auxiliar, destino, origem)

	return movimentos
}

func main() {
	numDiscos := 3
	origem, auxiliar, destino := "A", "B", "C"
	movimentos := hanoi(numDiscos, origem, destino, auxiliar)
	fmt.Println("Número de movimentos necessários: \n", movimentos)
}
