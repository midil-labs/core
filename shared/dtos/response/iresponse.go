package response

import (
	"github.com/midil-labs/core/shared/dtos"
)

type IResponse[T dtos.DTOInterface] interface {
	GetIdentifier() ResourceIdentifier
	GetAttributes() T
	GetRelationships() map[string]Relationship
	GetLinks() *Links
	GetMeta() map[string]any
}

func (r *Resource[T]) GetIdentifier() ResourceIdentifier {
	return r.ResourceIdentifier
}

func (r *Resource[T]) GetAttributes() T {
	return r.Attributes
}

func (r *Resource[T]) GetRelationships() map[string]Relationship {
	return r.Relationships
}

func (r *Resource[T]) GetLinks() *Links {
	return r.Links
}

func (r *Resource[T]) GetMeta() map[string]any {
	return r.Meta
}
