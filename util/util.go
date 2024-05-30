package util

import (
	"encoding/json"
	"os"
)

func GetDatabaseURL() string {
	return os.Getenv("DATABASE_URL")
}

func GetDatabaseName() string {
	return os.Getenv("DATABASE_NAME")
}

func GetResponseBodyKeys(bodyBytes []byte) []string {
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
