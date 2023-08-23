package operacoesarray

import (
	c "AC3/Contato"
	"fmt"
)

func AdicionaContato(contatoNovo c.Contato, contatos *[5]c.Contato){
	for i, contato := range contatos{
		if (contato == c.Contato{}) {
			contatos[i] = contatoNovo
			break
		}
	}
}

func ExcluiContato(contatos *[5]c.Contato){
	for i := len(contatos)-1; i >= 0; i-- {
		if (contatos[i] != c.Contato{}) {
			contatos[i] = c.Contato{}
			break
		}
	}
}

func ExibirContatos(contatos [5]c.Contato){
	for i, contato := range contatos{
		fmt.Println(i, "- nome: ", contato.Nome, "- email: ", contato.Email)
		}
	}

func EditarEmail(indice int, email string, contatos *[5]c.Contato) {
	contatos[indice].AlterarEmail(email) 
}
