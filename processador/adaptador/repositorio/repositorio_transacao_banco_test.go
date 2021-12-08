package repositorio

import (
	"os"
	"testing"

	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/adaptador/repositorio/fixture"
	"github.com/ArikBartzadok/ecossistema-gateway-pagamentos/dominio/entidade"
	"github.com/stretchr/testify/assert"
)

func TesteRepositorioTransacao_Inserir(t *testing.T) {
	migrationsDiretorio := os.DirFS("fixture/sql")

	banco := fixture.Subir(migrationsDiretorio)
	defer fixture.Descer(banco, migrationsDiretorio)

	repositorio := NovoRepositorioTransacaoBanco(banco)
	err := repositorio.Inserir(
		"1",
		"1",
		10.99,
		entidade.APROVADO,
		"",
	)

	assert.Nil(t, err)
}
