package processo_transacao

import (
	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/adaptador/broker"
	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/dominio/entidade"
	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/dominio/repositorio"
)

type ProcessoTransacao struct {
	repositorio repositorio.RepositorioTransacao
	Produtor    broker.InterfaceProdutor
	Topico      string
}

func NovoProcessoTransacao(repositorio repositorio.RepositorioTransacao, interfaceProdutor broker.InterfaceProdutor, topico string) *ProcessoTransacao {
	return &ProcessoTransacao{
		repositorio: repositorio,
		Produtor:    interfaceProdutor,
		Topico:      topico,
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

	err = p.publicar(saida, []byte(transacao.ID))

	if err != nil {
		return SaidaTransacaoDTO{}, err
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

	err = p.publicar(saida, []byte(transacao.ID))

	if err != nil {
		return SaidaTransacaoDTO{}, err
	}

	return saida, nil
}

func (p *ProcessoTransacao) publicar(saida SaidaTransacaoDTO, chave []byte) error {
	err := p.Produtor.Publicar(saida, chave, p.Topico)

	if err != nil {
		return err
	}

	return nil
}
