package fabrica

import "github.com/ArikBartzadok/ecossistema-gateway-pagamentos/dominio/repositorio"

type RepositorioFabrica interface {
	CriarRepositorioTransacao() repositorio.RepositorioTransacao
}
