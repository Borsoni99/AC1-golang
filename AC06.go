package main

import "fmt"

const M = 5

func main() {
	var lista [M]int
	var n = 0

	insereOrd(&lista, &n, 4)
	fmt.Println(lista)

	insereOrd(&lista, &n, 12)
	fmt.Println(lista)

	insereOrd(&lista, &n, 2)
	fmt.Println(lista)

	insereOrd(&lista, &n, 6)
	fmt.Println(lista)

	insereOrd(&lista, &n, 17)
	fmt.Println(lista)

	insereOrd(&lista, &n, 1)
	fmt.Println(lista)
}

func insereOrd(v *[M]int, n *int, novoValor int) {
	i := 0

	if *n == M {
		fmt.Println("Overflow")
		return
	}

	for ; i < *n; i++ {
		if v[i] > novoValor {
			break
		}
	}

	for j := *n - 1; j >= i; j-- {
		v[j+1] = v[j]
	}

	v[i] = novoValor
	*n++
}
