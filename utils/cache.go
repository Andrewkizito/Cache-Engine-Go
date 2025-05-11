package utils

type CacheEntry struct {
	ContentType string
	Data string
}

var cache = make(map[string]CacheEntry)

func GetData(key string) CacheEntry {
	return cache[key]
}

func SetData(key string, data CacheEntry) {
    cache[key] = data
}

func RemoveData(key string) {
    delete(cache, key)
}

func HasKey(key string) bool {
    _, exists := cache[key]
    return exists
}
