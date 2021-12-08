```sh
# subindo containers
docker-compose up -d
```

```sh
# executando container
docker exec -it processador_go bash
```

```sh
# iniciando uma aplicação go
go mod init github.com/ArikBartzadok/ecossistema-gateway-pagamentos
```

```sh
# executando testes da aplicação
go test ./...
# ou
go test -count=1 ./...

# limpando cache dos testes
go clean -testcache
```

```sh
# gerando mocks para testes com o repositório
mockgen -destination=dominio/repositorio/mock/mock.go -source=dominio\/repositorio\/repositorio.go
```