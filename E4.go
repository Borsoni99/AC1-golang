//Elabore uma função e_bissexto() que receba um número correspondente a um determinado ano e em seguida informe se este ano é ou não bissexto.
//Um ano é bissexto se ele é múltiplo de quatro. No entanto anos múltiplos de 100 que não são múltiplos de 400 não são bissextos.
//Então 1995 não é bissexto, 2012 é bissexto, 1900 não é bissexto e 2000 é bissexto.

package main

import "fmt"

func main() {
	var ano int

	fmt.Println("informe um ano: ")
	fmt.Scan(&ano)
	if e_bissexto(ano) {
		fmt.Println("o ano ", ano, " é bissexto")
	} else {
		fmt.Println("o ano ", ano, " não é bissexto")
	}
}

func e_bissexto(ano int) bool {
	if ano%100 == 0 && ano%400 != 0 {
		return false
	} else if ano%4 == 0 {
		return true
	} else {
		return false
	}
}
