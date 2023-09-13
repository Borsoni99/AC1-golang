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
	ind := -1
	if *n == M {
		fmt.Println("overflow")
	} else {
		aux := novoValor
		for true {
			aux++
			ind = buscaBin(*v, *n, aux)
			if ind != -1 || aux > v[*n-1] {
				break
			}
		}
		if ind == -1 {
			v[*n] = novoValor
		} else {
			for i := *n - 1; i >= ind; i-- {
				if v[i] == 0 {
					continue
				} else {
					v[i+1] = v[i]
				}
			}
			v[ind] = novoValor
		}
	}
	*n++
}

func buscaBin(v [M]int, n, x int) int {
	var meio int
	inf, sup := 0, n-1
	for inf <= sup {
		meio = int((inf + sup) / 2)
		if v[meio] == x {
			return meio
		}
		if v[meio] < x {
			inf = meio + 1
		} else {
			sup = meio - 1
		}
	}
	return -1
}
