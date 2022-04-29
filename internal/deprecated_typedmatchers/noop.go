package typedmatchers

type noopMatcher[T any] struct{}

func Noop[T any]() *noopMatcher[T] {
	return &noopMatcher[T]{}
}

func (m *noopMatcher[T]) Matches(x T) bool {
	return true
}

func (m *noopMatcher[T]) String() string {
	return "it should never fail. if you are seeing this please contact github.com/mniak/gomock-contrib maintainers"
}
