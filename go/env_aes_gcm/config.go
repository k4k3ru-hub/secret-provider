//
// config.go
//
package envaesgcm

import (
    "fmt"
)

type Config struct {
    KeyVersion          string
    EncryptionKeyBase64 string
}


func (c Config) Validate() error {
    if c.KeyVersion == "" {
        return fmt.Errorf("missing required parameter: key_version=%q", "empty")
    }
    if c.EncryptionKeyBase64 == "" {
        return fmt.Errorf("missing required parameter: encryption_key_base64=%q", "empty")
    }
    return nil
}
