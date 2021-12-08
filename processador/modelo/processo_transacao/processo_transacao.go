package processo_transacao

import (
	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/dominio/entidade"
	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/dominio/repositorio"
)

type ProcessoTransacao struct {
	repositorio repositorio.RepositorioTransacao
}

func NovoProcessoTransacao(repositorio repositorio.RepositorioTransacao) *ProcessoTransacao {
	return &ProcessoTransacao{
		repositorio: repositorio,
	}
}

func (p *ProcessoTransacao) Executar(entrada EntradaTransacaoDTO) (SaidaTransacaoDTO, error) {
	transacao := entidade.NovaTransacao()

	transacao.ID = entrada.ID
	transacao.ContaID = entrada.ContaID
	transacao.Valor = entrada.Valor

	_, ccInvalido := entidade.NovoCartaoCredito(
		entrada.NumeroCartaoCredito,
		entrada.NomeCartaoCredito,
		entrada.MesExpiracaoCartaoCredito,
		entrada.AnoExpiracaoCartaoCredito,
		entrada.CvvCartaoCredito,
	)

	if ccInvalido != nil {
		err := p.repositorio.Inserir(
			transacao.ID,
			transacao.ContaID,
			transacao.Valor,
			entidade.REJEITADO,
			ccInvalido.Error(),
		)

		if err != nil {
			return SaidaTransacaoDTO{}, err
		}

		saida := SaidaTransacaoDTO{
			ID:           transacao.ID,
			Status:       entidade.REJEITADO,
			MensagemErro: ccInvalido.Error(),
		}

		return saida, nil
	}

	return SaidaTransacaoDTO{}, nil

}
