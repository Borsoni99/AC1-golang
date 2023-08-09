//Implementa uma função que receba um número e exiba o dia correspondente da semana (1-Domingo, 2- Segunda, etc.),
//se digitar outro valor deve aparecer valor inválido.

package main

import "fmt"

func main(){
	var num int

	fmt.Println("informe um número entre 1 e 7: ")
	fmt.Scan(&num)

	for !entrada_valida(num){
		fmt.Println("número inválido!")
		fmt.Println("informe um número entre 1 e 7: ")
		fmt.Scan(&num)
	}
	dia_da_semana(num)
}

func dia_da_semana(num int){
	switch num {
	case 1:
		println("Domingo")
	case 2:
		println("Segunda")
	case 3:
		println("Terça")
	case 4:
		println("Quarta")
	case 5:
		println("Quinta")
	case 6:
		println("Sexta")
	case 7:
		println("Sábado")

	}
}


func entrada_valida(num int) bool{
	if num < 1 || num > 7{
		return false
	} else{
		return true
	}
}