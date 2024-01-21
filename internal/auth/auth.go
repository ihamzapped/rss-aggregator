package auth

import (
	"errors"
	"net/http"
	"strings"
)

/*
Extract api key from auth headers
Example:
Authorization: ApiKey {an apikey of len 64}
*/
func GetApiKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")

	if val == "" {
		return "", errors.New("no auth header found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 || vals[0] != "ApiKey" || len(vals[1]) != 64 {
		return "", errors.New("malformed auth header")
	}

	return vals[1], nil
}
