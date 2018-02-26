package base64

import (
	"encoding/base64"
)

// ToB64 encodes string to base64
func Encode(data *string) *string {
	b64 := base64.StdEncoding.EncodeToString([]byte(*data))
	return &b64
}
