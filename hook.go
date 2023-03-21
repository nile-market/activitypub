package activitypub

package activitypub

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "database/sql"
    "encoding/json"
    "encoding/pem"
    "errors"
    "log"
    "net/http"
    "strings"
    "time"

    "github.com/pocketbase/pocketbase"
)

type ActivityPubHook struct{}

func (h *ActivityPubHook) HandleEvent(ctx context.Context, event pocketbase.Event) {
    if event.Type == "user-registered" {
        user := event.Data.(User)
        privateKey, publicKey, err := generateKeys()
        if err != nil {
            log.Println(err)
            return
        }
        err = ctx.DB.SetUserKeys(user.Username, privateKey, publicKey)
        if err != nil {
            log.Println(err)
        }
    }
}

func generateKeys() ([]byte, []byte, error) {
    privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
    if err != nil {
        return nil, nil, err
    }

    privateKeyBytes := x509.MarshalPKCS1PrivateKey(privateKey)
    publicKeyBytes, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
    if err != nil {
        return nil, nil, err
    }

    return privateKeyBytes, publicKeyBytes, nil
}
