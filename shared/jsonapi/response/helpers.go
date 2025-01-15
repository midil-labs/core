package response

import (
	"fmt"
	"github.com/midil-labs/core/shared/jsonapi/common"
)

type Scope string

const (
    ScopeTop          Scope = "top"
    ScopeResource     Scope = "resource"
    ScopeRelationship Scope = "relationship"
    ScopeLink         Scope = "link"
)

func (s Scope) IsValid() bool {
	switch s {
	case ScopeTop, ScopeResource, ScopeRelationship, ScopeLink:
		return true
	default:
		return false
	}
}

var ErrInvalidScope = fmt.Errorf("invalid scope")

// Helper function to merge meta into a single Resource's Meta.
func mergeResourceMeta(r *Resource, meta map[string]any) {
	if r.Meta == nil {
		r.Meta = make(map[string]any)
	}
	for k, v := range meta {
		r.Meta[k] = v
	}
}

// Helper function to merge meta into all relationships in a single Resource.
func mergeRelationshipMeta(r *Resource, meta map[string]any) {
	if r.Relationships == nil {
		return
	}
	for relName, rel := range r.Relationships {
		if rel.Meta == nil {
			rel.Meta = make(map[string]any)
		}
		for k, v := range meta {
			rel.Meta[k] = v
		}
		r.Relationships[relName] = rel
	}
}

// Helper function to merge meta into a Resource's Links.
// If a Resource has no Links, we create an empty one.
func mergeLinksMeta(r *Resource, meta map[string]any) {
	if r.Links == nil {
		r.Links = &common.Links{}
	}
	if r.Links.Related.Meta == nil {
		r.Links.Related.Meta = make(map[string]any)
	}
	for k, v := range meta {
		r.Links.Related.Meta[k] = v
	}
}
