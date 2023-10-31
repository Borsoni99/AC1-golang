///////////////////////////////
// Thiago Borsoni, Daniel Lloyd
///////////////////////////////
package main

import "fmt"

type No struct {
	Chave  int
	Esq, Dir *No
}

type Arvore struct {
	Raiz *No
}

func buscaArvore(valor int, no *No) int {
	if no == nil {
		return 0
	}

	if valor == no.Chave {
		return 1
	}

	if valor < no.Chave {
		if no.Esq == nil {
			return 0
		}
		return buscaArvore(valor, no.Esq)
	}

	if no.Dir == nil {
		return 0
	}
	return buscaArvore(valor, no.Dir)
}


func insere(valor int, no *No) *No {
	if no == nil {
		novoNo := &No{Chave: valor, Esq: nil, Dir: nil}
		return novoNo
	}

	if valor < no.Chave {
		no.Esq = insere(valor, no.Esq)
	} else if valor > no.Chave {
		no.Dir = insere(valor, no.Dir)
	} else {
		fmt.Println("Inserção inválida: valor já existe na árvore.")
	}

	return no
}

func imprimeArvore(a *Arvore) {
	if a.Raiz != nil { imprimeSimetrico(a.Raiz) }
}

func imprimeSimetrico(n *No) {
	if n.Esq != nil { imprimeSimetrico(n.Esq) }
	fmt.Println(n.Chave)
	if n.Dir != nil { imprimeSimetrico(n.Dir) }
}

func main() {
	arv := &Arvore{}
	arv.Raiz = insere(12, arv.Raiz)
	arv.Raiz = insere(7, arv.Raiz)
	arv.Raiz = insere(8, arv.Raiz)
	arv.Raiz = insere(9, arv.Raiz)
	arv.Raiz = insere(11, arv.Raiz)
	arv.Raiz = insere(6, arv.Raiz)
	arv.Raiz = insere(2, arv.Raiz)
	arv.Raiz = insere(3, arv.Raiz)
	arv.Raiz = insere(20, arv.Raiz)

	valor := 11
	resultado := buscaArvore(valor, arv.Raiz)

		switch resultado {
		case 0:
			fmt.Printf("Valor %d não encontrado na árvore.\n", valor)
		case 1:
			fmt.Printf("Valor %d encontrado.\n", valor)
		}

	imprimeArvore(arv)

}


