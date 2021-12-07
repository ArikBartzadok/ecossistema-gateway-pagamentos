package repositorio

type RepositorioTransacao interface {
	Inserir(id string, conta string, valor float64, status string, mensagemErro string) error
}
