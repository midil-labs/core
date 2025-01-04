package common


func ApplyOptions[T any](target *T, opts ...Option[T]) {
	for _, opt := range opts {
		opt(target)
	}
}