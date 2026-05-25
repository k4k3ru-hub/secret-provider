//
// provider.go
//
package secretprovider

import (
    "context"

    "github.com/k4k3ru-hub/secret-provider/go/env_aes_gcm"
)

type EnvAESGCMConfig = envaesgcm.Config


type Provider interface {
    ProviderKind() string
    KeyVersion() string
    Encrypt(ctx context.Context, plainText string) (string, error)
    Decrypt(ctx context.Context, cipherText string, keyVersion string) (string, error)
}


//
// Create AES GCM provider.
//
// Version:
//   - 2026-05-25: Added.
//
func NewEnvAESGCMProvider(config EnvAESGCMConfig) (Provider, error) {
    return envaesgcm.NewProvider(config)
}
