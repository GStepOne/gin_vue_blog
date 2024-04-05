package utils

func InList(key string, collection []string) bool {
	for _, k := range collection {
		if key == k {
			return true
		}
	}
	return false
}

func Reverse[T any](s []T) []T {
	// 定义反转后的切片
	var reversed []T

	// 倒序遍历原始切片，并将元素添加到反转切片中
	for i := len(s) - 1; i >= 0; i-- {
		reversed = append(reversed, s[i])
	}

	return reversed
}
