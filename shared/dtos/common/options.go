package common

type Option[T any] func(*T)
type MetaOption = Option[NonStandardMeta]