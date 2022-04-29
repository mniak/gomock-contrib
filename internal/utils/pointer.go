package utils

func ToPointer[T any](i T) *T {
	return &i
}
