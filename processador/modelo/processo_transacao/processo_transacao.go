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

	cc, ccInvalido := entidade.NovoCartaoCredito(
		entrada.NumeroCartaoCredito,
		entrada.NomeCartaoCredito,
		entrada.MesExpiracaoCartaoCredito,
		entrada.AnoExpiracaoCartaoCredito,
		entrada.CvvCartaoCredito,
	)

	if ccInvalido != nil {
		return p.transacaoRejeitada(transacao, ccInvalido)
	}

	transacao.AdicionarCartaoCredito(*cc)
	transacaoInvalida := transacao.Valida()

	if transacaoInvalida != nil {
		return p.transacaoRejeitada(transacao, transacaoInvalida)
	}

	return p.transacaoAprovada(entrada, transacao)
}

func (p *ProcessoTransacao) transacaoAprovada(entrada EntradaTransacaoDTO, transacao *entidade.Transacao) (SaidaTransacaoDTO, error) {
	err := p.repositorio.Inserir(
		transacao.ID,
		transacao.ContaID,
		transacao.Valor,
		entidade.APROVADO,
		"",
	)

	if err != nil {
		return SaidaTransacaoDTO{}, err
	}

	saida := SaidaTransacaoDTO{
		ID:           transacao.ID,
		Status:       entidade.APROVADO,
		MensagemErro: "",
	}

	return saida, nil
}

func (p *ProcessoTransacao) transacaoRejeitada(transacao *entidade.Transacao, transacaoInvalida error) (SaidaTransacaoDTO, error) {
	err := p.repositorio.Inserir(
		transacao.ID,
		transacao.ContaID,
		transacao.Valor,
		entidade.REJEITADO,
		transacaoInvalida.Error(),
	)

	if err != nil {
		return SaidaTransacaoDTO{}, err
	}

	saida := SaidaTransacaoDTO{
		ID:           transacao.ID,
		Status:       entidade.REJEITADO,
		MensagemErro: transacaoInvalida.Error(),
	}

	return saida, nil
}
