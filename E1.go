//Elabore uma função e_primo() que retorna se um número é primo ou não. Caso o número não seja primo, liste todos os números pelos quais ele é divisível.

package main

import "fmt"

func main(){
	var num int
	var eh_primo bool

	fmt.Println("informe um número: ")
	fmt.Scan(&num)

	eh_primo = e_primo(num)

	if eh_primo {
		fmt.Println("O número ",num," é primo!")
	} else{
		fmt.Println("O número ",num," não é primo!")
		fmt.Println(num," é divisivel por: ")
		numeros_divisiveis(num)
	}

}

func e_primo(num int) bool{
	e_primo := true
	for i := 2; i < num; i++ {
		if num % i == 0 {
			return !e_primo
		}
	}
	return e_primo
}

func numeros_divisiveis(num int){
	for i := 1; i <= num; i++ {
		if num % i == 0 {
			fmt.Print(i," ")
		}
	}
}