package Contato

type Contato struct {
	Nome		string
	Email		string
}

func (c *Contato) AlterarEmail(email string) {
	c.Email = email
}

