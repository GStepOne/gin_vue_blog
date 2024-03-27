package utils

func InList(key string, collection []string) bool {
	for _, k := range collection {
		if key == k {
			return true
		}
	}
	return false
}
