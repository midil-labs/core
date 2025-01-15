package error

import (
	"encoding/json"
	"github.com/midil-labs/core/shared/utils/goutils"
)

type SerializationOption func(*json.Encoder)
type Option goutils.Option[ErrorObject]

func WithID(id string) Option {
	return func(e *ErrorObject) {
		e.ID = id
	}
}

func WithLinks[T ErrorObject](about, typeURL string) Option {
	return func(e *ErrorObject) {
		e.Links = &ErrorLinks{
			About: about,
			Type:  typeURL,
		}
	}
}

func WithSource[T ErrorObject](pointer, parameter, header string) Option {
	return func(e *ErrorObject) {
		e.Source = &ErrorSource{
			Pointer:   pointer,
			Parameter: parameter,
			Header:    header,
		}
	}
}

func WithMeta[T ErrorObject](meta map[string]interface{}) Option {
	return func(e *ErrorObject) {
		e.Meta = meta
	}
}

func WithLocalization[T ErrorObject](lang Language, title, detail string) Option {
	return func(e *ErrorObject) {
		if e.localizer == nil {
			e.localizer = make(map[Language]Localization)
		}
		e.localizer[lang] = Localization{
			Title:  title,
			Detail: detail,
		}
	}
}
