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

func ReadCachEntry(key string) (RawCacheEntry, error) {
	if key == "" {
		return RawCacheEntry{}, errors.New("key can not be empty")
	}

	hashedKey := GenerateHash(key)

	entry := GetData(hashedKey)
	decodedData, err := base64.StdEncoding.DecodeString(entry.Data)

	if err != nil {
		return RawCacheEntry{}, errors.New("failed to decode data")
	}

	return RawCacheEntry{
		ContentType: entry.ContentType,
		Raw:         decodedData,
	}, nil
}
