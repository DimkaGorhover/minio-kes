package kestest_test

import (
	"context"
	"flag"
	"os"
	"testing"

	"github.com/minio/kes/edge"
)

var azureConfigFile = flag.String("azure.config", "", "Path to a KES config file with Azure SecretsManager config")

func TestGatewayAzure(t *testing.T) {
	if *azureConfigFile == "" {
		t.Skip("Azure tests disabled. Use -azure.config=<config file with Azure SecretManager config> to enable them")
	}
	file, err := os.Open(*azureConfigFile)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	srvrConfig, err := edge.ReadServerConfigYAML(file)
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := testingContext(t)
	defer cancel()

	store, err := srvrConfig.KeyStore.Connect(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Metrics", func(t *testing.T) { testMetrics(ctx, store, t) })
	t.Run("APIs", func(t *testing.T) { testAPIs(ctx, store, t) })
	t.Run("CreateKey", func(t *testing.T) { testCreateKey(ctx, store, t) })
	t.Run("ImportKey", func(t *testing.T) { testImportKey(ctx, store, t) })
	t.Run("GenerateKey", func(t *testing.T) { testGenerateKey(ctx, store, t) })
	t.Run("EncryptKey", func(t *testing.T) { testEncryptKey(ctx, store, t) })
	t.Run("DecryptKey", func(t *testing.T) { testDecryptKey(ctx, store, t) })
	t.Run("DecryptKeyAll", func(t *testing.T) { testDecryptKeyAll(ctx, store, t) })
	t.Run("DescribePolicy", func(t *testing.T) { testDescribePolicy(ctx, store, t) })
	t.Run("GetPolicy", func(t *testing.T) { testGetPolicy(ctx, store, t) })
	t.Run("SelfDescribe", func(t *testing.T) { testSelfDescribe(ctx, store, t) })
}