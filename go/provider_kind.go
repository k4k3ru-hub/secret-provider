//
// provider_kind.go
//
package secretprovider

import (
    "fmt"

    "github.com/k4k3ru-hub/secret-provider/go/env_aes_gcm"
)

type ProviderKind string

const (
    ProviderKindEnvAESGCM ProviderKind = envaesgcm.ProviderKind
    ProviderKindGCPKMS    ProviderKind = "gcp-kms"
    ProviderKindAWSKMS    ProviderKind = "aws-kms"
)



//
// Check whether provider kind is valid.
//
// Version:
//   - 2026-05-25: Added.
//
func (k ProviderKind) IsValid() bool {
    switch k {
    case ProviderKindEnvAESGCM, ProviderKindGCPKMS, ProviderKindAWSKMS:
        return true
    default:
        return false
    }
}


//
// Validate provider kind.
//
// Version:
//   - 2026-05-25: Added.
//
func (k ProviderKind) Validate() error {
    s := string(k)
    if len(s) > 32 {
        return fmt.Errorf("invalid parameter: provider_kind=%q", "too long")
    }
    if !k.IsValid() {
        return fmt.Errorf("invalid parameter: provider_kind=%q", s)
    }
    return nil
}


//
// Convert provider kind to string.
//
// Version:
//   - 2026-05-25: Added.
//
func (k ProviderKind) String() string {
    return string(k)
}
