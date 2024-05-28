package util

import (
	"encoding/json"
	"io"
)

func GetResponseBodyKeys(body io.ReadCloser) []string {
	bodyBytes, _ := io.ReadAll(body)
	jsonMap := make(map[string]json.RawMessage)
	err := json.Unmarshal(bodyBytes, &jsonMap)

	keys := make([]string, 0)
	if err != nil {
		return keys
	}

	for key := range jsonMap {
		keys = append(keys, key)
	}

	return keys
}
