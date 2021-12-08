package transacao

import (
	"encoding/json"

	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/modelo/processo_transacao"
)

type ApresentacaoKafka struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	MensagemErro string `json:"mensagem_erro"`
}

func NovaTransacaoApresentacaoKafka() *ApresentacaoKafka {
	return &ApresentacaoKafka{}
}

// definindo o formato dos dados a serem enviados para o kafka (commit)
func (transacao *ApresentacaoKafka) Vincular(saida interface{}) error {
	transacao.ID = saida.(processo_transacao.SaidaTransacaoDTO).ID
	transacao.Status = saida.(processo_transacao.SaidaTransacaoDTO).Status
	transacao.MensagemErro = saida.(processo_transacao.SaidaTransacaoDTO).MensagemErro

	return nil
}

// formatando o formato em que o kafka deverá apresentar os dados (json)
func (transacao *ApresentacaoKafka) Exibir() ([]byte, error) {
	json_gerado, err := json.Marshal(transacao)

	if err != nil {
		return nil, err
	}

	return json_gerado, nil
}
