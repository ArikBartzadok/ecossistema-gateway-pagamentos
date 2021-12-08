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
```