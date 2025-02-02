package goutils

type Option[T any] func(*T)


type ErrableOption[T any] func(*T) []*error


func ApplyErrableOptions[T any](target *T, opts ...ErrableOption[T]) []*error {
	for _, opt := range opts {
		if err := opt(target); err != nil {
			return err
		}
	}
	return nil
}


func ApplyOptions[T any](target *T, opts ...Option[T]) {
	for _, opt := range opts {
		opt(target)
	}
}
