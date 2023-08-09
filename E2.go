//Implemente uma função fibo() que recebe um número e calcule o “n-ésimo” elemento da série de Fibonacci.
//A série de Fibonacci é dada de tal forma que um número em uma dada posição é igual à soma dos dois números anteriores.
//Considere que a série é formada pela sequência 1, 1, 2, 3, …

package main

import "fmt"

func main(){
	var num int

	fmt.Println("informe um número: ")
	fmt.Scan(&num)

	fmt.Println(calcular_nesimo_elemento_fibo(num))

}

func calcular_nesimo_elemento_fibo(num int) int{
	prox_num := 1
	num_anterior := 0

	for i := 1; i < num; i++ {
		prox_num_temp := prox_num
		prox_num = prox_num + num_anterior
		num_anterior = prox_num_temp
	}
	return prox_num
}