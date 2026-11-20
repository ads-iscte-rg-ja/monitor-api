# monitor-api

Contrato de ingestão do Monitor: o serviço gRPC que o `monitor-remote` (cliente) usa para enviar
telemetria ao `monitor-central` (servidor). A decisão e as alternativas estão no
[ADR-0003](https://github.com/ads-iscte-rg-ja/monitor-platform/blob/main/docs/adr/0003-contrato-ingestao-grpc.md).

| Ficheiro | Conteúdo |
|---|---|
| `proto/monitor/ingest/v1/ingest.proto` | O contrato: serviço, mensagens, códigos de rejeição e regras de erro |
| `gen/go/monitor/ingest/v1/` | Código Go gerado (versionado, para os consumidores não precisarem do `buf`) |
| `gen/go/monitor/ingest/v1/limits.go` | Limites do contrato (`MaxBatchRecords`, `MaxMessageBytes`, `APIKeyMetadata`), escritos à mão |

## Usar

```go
import ingestv1 "github.com/ads-iscte-rg-ja/monitor-api/gen/go/monitor/ingest/v1"
```

```sh
go get github.com/ads-iscte-rg-ja/monitor-api@v1.1.0
```

## Versões

| Versão | Alteração |
|---|---|
| v1.0.0 | Serviço `IngestService`: estatísticas e erros em lote, códigos de rejeição ([ADR-0003](https://github.com/ads-iscte-rg-ja/monitor-platform/blob/main/docs/adr/0003-contrato-ingestao-grpc.md)) |
| v1.1.0 | `IngestStatsRequest.buffer` opcional: estado do *buffer* do agente ([ADR-0010](https://github.com/ads-iscte-rg-ja/monitor-platform/blob/main/docs/adr/0010-estado-do-buffer-e-health-ready.md)); aditivo |

## Alterar o contrato

1. Editar o `.proto`. Alterações compatíveis apenas: campos e valores novos com números novos, nunca
   reutilizar nem renumerar. O CI corre `buf breaking` contra o `main` e falha numa alteração incompatível.
2. Gerar o código e confirmar o lint:
   ```sh
   go install github.com/bufbuild/buf/cmd/buf@v1.75.0
   go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.36.12
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.6.2
   buf lint && buf generate
   ```
3. PR, *merge* e uma *tag* nova (`v1.x.y`). O central e o remote atualizam a versão no `go.mod`.

Uma alteração incompatível exige um pacote novo (`monitor.ingest.v2`), servido em paralelo com o v1
enquanto houver agentes antigos.

## Clientes noutras linguagens

O `.proto` é a especificação completa (cenário IO1): gerar o cliente com o `protoc`/`buf` da linguagem
e seguir as regras de autenticação e de erros descritas no início do ficheiro.
