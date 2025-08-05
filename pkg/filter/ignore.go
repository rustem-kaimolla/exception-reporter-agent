package filter

import (
	"exception-reporter-agent/model"
	"strings"
)

// Какие ошибки игнорирую то есть нет смысла их логировать
var ignoreMessages = []string{
	"Connection refused",
	"connection reset by peer",
	"connection timed out",
	"read timeout",
	"SSL certificate",
	"no such host",
	"Too many open files",
	"Redis connection",
	"elasticsearch cluster is down",
	"Rate limit exceeded",
	"circuit breaker is open",
	"timeout while waiting",
}

func ShouldIgnore(exc model.ExceptionPayload) bool {
	msg := strings.ToLower(exc.Message)

	for _, pattern := range ignoreMessages {
		if strings.Contains(msg, strings.ToLower(pattern)) {
			return true
		}
	}

	return false
}
