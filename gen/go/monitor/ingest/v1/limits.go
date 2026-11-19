package ingestv1

// Limites e nomes que fazem parte do contrato (ver o comentário inicial de ingest.proto).
// Escrito à mão: o buf não gera nem apaga este ficheiro.
const (
	// MaxBatchRecords é o número máximo de registos por pedido.
	MaxBatchRecords = 100
	// MaxMessageBytes é o tamanho máximo de uma mensagem de pedido, codificada.
	MaxMessageBytes = 1 << 20
	// APIKeyMetadata é o metadado gRPC que leva a chave do host.
	APIKeyMetadata = "x-api-key"
)
