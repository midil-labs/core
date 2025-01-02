package error

import "encoding/json"

type SerializationOption func(*json.Encoder)

type Option func(*ErrorObject)

func WithID(id string) Option {
	return func(e *ErrorObject) {
		e.ID = id
	}
}

func WithLinks(about, typeURL string) Option {
	return func(e *ErrorObject) {
		e.Links = &ErrorLinks{
			About: about,
			Type:  typeURL,
		}
	}
}

func WithSource(pointer, parameter, header string) Option {
	return func(e *ErrorObject) {
		e.Source = &ErrorSource{
			Pointer:   pointer,
			Parameter: parameter,
			Header:    header,
		}
	}
}

func WithMeta(meta map[string]interface{}) Option {
	return func(e *ErrorObject) {
		e.Meta = meta
	}
}

func WithLocalization(lang Language, title, detail string) Option {
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