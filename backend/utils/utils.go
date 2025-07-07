package utils

func Ptr[T any](v T) *T {
	return &v
}

func StringEqual(a, b *string) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	return *a == *b
}
