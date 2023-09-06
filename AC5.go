package main

import "fmt"

func main() {
	array := []int{1,2,3,4,5,7,9,10,23,34,45,67,78,89,91,99,100}
	alvo := 70

	resultado := busca(array, alvo)
	fmt.Println(resultado)

}

func busca(array []int, alvo int) []int{
	j := len(array) - 1
	for i := 0; i < len(array); {
		if (array[i] + array[j]) > alvo{
			j--
		}else if (array[i] + array[j]) < alvo{
			i++
		}else{
			return []int{array[i], array[j]}
		}
	}
	return []int{-1, -1}
}
