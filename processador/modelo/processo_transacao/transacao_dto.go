package processo_transacao

// DTO -> objeto que serve apenas para tranferir dados entre as camadas, em um padrão

// Modelo dos dados entregues pelo kafka
type EntradaTransacaoDTO struct {
	ID                        string  `json:"id"`
	ContaID                   string  `json:"conta_id"`
	NumeroCartaoCredito       string  `json:"numero_cartao_credito"`
	NomeCartaoCredito         string  `json:"nome_cartao_credito"`
	MesExpiracaoCartaoCredito int     `json:"mes_expiracao_cartao_credito"`
	AnoExpiracaoCartaoCredito int     `json:"ano_expiracao_cartao_credito"`
	CvvCartaoCredito          int     `json:"cvv_cartao_credito"`
	Valor                     float64 `json:"valor"`
}

// Modelo dos dados retornados ao kafka
type SaidaTransacaoDTO struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	MensagemErro string `json:"mensagem_erro"`
}
