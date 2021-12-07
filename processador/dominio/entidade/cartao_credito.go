package entidade

import (
	"errors"
	"regexp"
	"time"
)

type CartaoCredito struct {
	numero       string
	nome         string
	mesExpiracao int
	anoExpiracao int
	cvv          int
}

func NovoCartaoCredito(numero string, nome string, mesExpiracao int, anoExpiracao int, cvvExpiracao int) (*CartaoCredito, error) {
	cc := &CartaoCredito{
		numero:       numero,
		nome:         nome,
		mesExpiracao: mesExpiracao,
		anoExpiracao: anoExpiracao,
		cvv:          cvvExpiracao,
	}

	erro := cc.Valido()

	if erro != nil {
		return nil, erro // se válido, retorna nil para o cartão e o erro
	}
	return cc, nil // se válido, retorna o cartão e nil para o erro
}

// se válido, retorna nil, caso contrário, retorna o erro lançado
func (c *CartaoCredito) Valido() error {
	erro := c.validarNumero()
	if erro != nil {
		return erro
	}

	erro = c.validarMes()
	if erro != nil {
		return erro
	}

	erro = c.validarAno()
	if erro != nil {
		return erro
	}

	return nil // se todos passarem
}

// se válido, retorna nil, caso contrário, retorna o erro lançado
func (c *CartaoCredito) validarNumero() error {
	regexValidacao := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)

	if !regexValidacao.MatchString(c.numero) {
		return errors.New("Número do cartão de crédito inválido")
	}
	return nil
}

// se válido, retorna nil, caso contrário, retorna o erro lançado
func (c *CartaoCredito) validarMes() error {
	if c.mesExpiracao > 0 && c.mesExpiracao < 13 {
		return nil
	}
	return errors.New("Mês de expiração do cartão de crédito inválido")
}

// se válido, retorna nil, caso contrário, retorna o erro lançado
func (c *CartaoCredito) validarAno() error {
	if c.anoExpiracao >= time.Now().Year() {
		return nil
	}

	return errors.New("Ano de expiração do cartão de crédito inválido")
}
