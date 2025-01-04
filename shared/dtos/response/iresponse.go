package response

import (
	"github.com/midil-labs/core/shared/dtos"
	error "github.com/midil-labs/core/shared/dtos/error"

)

// IResponse defines the behavior for a generic JSON:API response.
type IResponse[T dtos.DTOInterface] interface {
    // GetData returns the 'data' portion of the JSON:API response.
    GetData() Data[T]

    // SetData sets the 'data' portion of the JSON:API response
    // and returns the interface for chaining (optional).
    SetData(Data[T]) IResponse[T]

    // GetErrors returns the 'errors' array for this JSON:API response.
    GetErrors() []error.ErrorObject

    // SetErrors sets the 'errors' array for this JSON:API response,
    // returning the interface for chaining (optional).
    SetErrors([]error.ErrorObject) IResponse[T]

    // GetMeta returns a map of non-standard meta-information.
    GetMeta() map[string]any

    // SetMeta sets the non-standard meta-information for the response,
    // returning the interface for chaining (optional).
    SetMeta(map[string]any) IResponse[T]

    // GetLinks returns a map of top-level links for the response.
    GetLinks() map[string]string

    // SetLinks sets a map of top-level links for the response,
    // returning the interface for chaining (optional).
    SetLinks(map[string]string) IResponse[T]

    // GetIncluded returns a list of resources included alongside primary data.
    GetIncluded() []Resource[dtos.DTOInterface]

    // SetIncluded sets the list of included resources,
    // returning the interface for chaining (optional).
    SetIncluded([]Resource[dtos.DTOInterface]) IResponse[T]
}