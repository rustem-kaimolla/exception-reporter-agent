package fingerprint

import (
	"crypto/sha1"
	"encoding/hex"
	"exception-reporter-agent/model"
	"fmt"
)

func Generate(payload model.ExceptionPayload) string {
	input := fmt.Sprintf("%s|%s|%d|%s",
		payload.Message,
		payload.File,
		payload.Line,
		payload.App,
	)

	hash := sha1.Sum([]byte(input))
	return hex.EncodeToString(hash[:])
}
