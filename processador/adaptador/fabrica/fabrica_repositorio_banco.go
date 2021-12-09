package fabrica

import (
	"database/sql"

	adaptadorRepositorio "github.com/ArikBartzadok/ecossistema-gateway-pagamentos/adaptador/repositorio"
	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/dominio/repositorio"
)

type FabricaRepositorioBanco struct {
	Banco *sql.DB
}

func NovaFabricaRepositorioBanco(banco *sql.DB) *FabricaRepositorioBanco {
	return &FabricaRepositorioBanco{Banco: banco}
}

func (fabrica FabricaRepositorioBanco) CriarRepositorioTransacao() repositorio.RepositorioTransacao {
	return adaptadorRepositorio.NovoRepositorioTransacaoBanco(fabrica.Banco)
}
