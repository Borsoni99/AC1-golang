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
	if *n == M{
		fmt.Println("overflow")
	}else{
		for i := 0; i < *n; i++ {
			ind = buscaOrd(*v, *n, novoValor + i + 1)
			if ind != -1{
				break
			}
		}
		if ind == -1{
			v[*n] = novoValor
		}else{
			for i := *n - 1; i >= ind; i--{
				if v[i] == 0{
					continue
				}else{
					v[i + 1] = v[i]
				}
			}
			
			v[ind] = novoValor
		}
	}
}

func buscaOrd(v [M]int, n, x int) int {
	i := 0
	v[n] = x
	for v[i] < x {
		i++
	}
	if i == n + 1 || v[i] != x {
		return -1
	}
	return i
}
