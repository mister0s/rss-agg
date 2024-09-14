package auth

import (
	"errors"
	"net/http"
	"strings"
)

// format should be Authorization: ApiKey {ApiKey}
func GetApiKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no auth header found")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed auth api key header")
	}

	return vals[1], nil
}
