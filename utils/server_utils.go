package utils

import (
	"encoding/base64"
	"errors"
)

type RawCacheEntry struct {
	ContentType string
	Raw         []byte
}

func AddCacheEntry(key string, data RawCacheEntry) (string, error) {
	if key == "" || len(data.Raw) == 0 {
		return "", errors.New("key and data can not be empty")
	}

	hashedKey := GenerateHash(key)
	encodedData := base64.StdEncoding.EncodeToString(data.Raw)

	exists := HasKey(hashedKey)
	if exists {
		return "", errors.New("key is already taken, use update action or delete entry")
	}

	newCacheEntry := CacheEntry{
		ContentType: data.ContentType,
		Data:        encodedData,
	}

	SetData(hashedKey, newCacheEntry)
	return key, nil
}

func ReadCachEntry() {
	fmt.Println("Reading from cache")
}
