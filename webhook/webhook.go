package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"git-webhook/config"
	"io"
	"net/http"
)

func VerifySignature(r *http.Request, cfg *config.Config) (bool, []byte, error) {
	payload, err := io.ReadAll(r.Body)
	if err != nil {
		return false, nil, err
	}
	signature := r.Header.Get("X-Hub-Signature-256")
	mac := hmac.New(sha256.New, []byte(cfg.Secret))
	mac.Write(payload)
	expectedSignature := "sha256=" + hex.EncodeToString(mac.Sum(nil))
	return hmac.Equal([]byte(expectedSignature), []byte(signature)), payload, nil
}
